<template>
  <div class="container">
    <el-card>
      <!-- el-row表明这是一行，包含在其中的内容都写在一行中。
            Element UI 规定24rows && 24cols   
       -->
      <el-row :gutter="20">
        <el-col :span="12">
          <!-- clearable属性比较有趣，删除inpout框内容会自动的调整 -->
          <el-input placeholder="输入文章查找" v-model="title" clearable @change="getArticleList">
            <el-button slot="append" icon="el-icon-search" @click="getArticleList"></el-button>
          </el-input>
        </el-col>
        <el-col :span="4">
          <el-button type="primary" @click="addArticleClick">新增</el-button>
        </el-col>
      </el-row>
      <!-- table表格 -->
      <!-- 这里使用了自定义列表，并且使用scope可以获取表中数据-->
      <el-table 
      :data="ArticleList.slice((currentPage-1)*pageSize,currentPage*pageSize)"
      :cell-style="CellStyle"
      >
        <el-table-column
          prop="title"
          label="文章标题"
          min-width="20%"
          align="left">
        </el-table-column>
        <el-table-column
          label="分类"
          min-width="20%"
          align="left">
          <template slot-scope="scope">
            <el-tag type="success" color="#fff2e8" size="small" style="color: #fa541c;">{{formatCid(scope.row)}}</el-tag>
          </template>
        </el-table-column>

        <el-table-column
          label="标签"
          min-width="20%"
          align="left">
          <template slot-scope="scope">
            <!-- TODO: -->
            <el-tag type="success" 
              color="rgb(207, 211, 215)" 
              size="small" 
              style="color: #484a4b;"
              v-for="(value, name) in formatTag(scope.row)"
              :key="name"
              >
              {{value}}
            </el-tag>
          </template>
        </el-table-column>
          
        <el-table-column
          prop="desc"
          label="文章描述"
          min-width="20%"
          align="left">
        </el-table-column>
        <el-table-column
          prop="img"
          label="缩略图"
          min-width="20%"
          align="left">
          <template slot-scope="scope">
            <!-- 这里插入图片需要用到slot机制，为src绑定参数 -->
            <img v-bind:src="scope.row.img" style="width:70px;height:70px;" alt="暂无图片">
          </template>
        </el-table-column>
        <el-table-column
          label="操作"
          min-width="30%"
          align="left">
            <template slot-scope="scope">
              <el-button
                class="rowActionBtn"
                type="text"
                size="mini"
                icon="el-icon-edit"
                @click="handleEdit(scope.$index, scope.row.ID, scope.row.Articlename, scope.row.role)">编辑
              </el-button>
              <span style="color:#e8e8e8; padding: 0 10px">|</span>
              <el-button
                class="rowActionBtn"
                type="text"
                size="mini"
                icon="el-icon-delete"
                @click="handleDelete(scope.$index, scope.row.ID)">删除
              </el-button>
            </template>
        </el-table-column>
      </el-table>
      <!-- 分页操作 -->
      <el-pagination align='right' 
        background
        @size-change="handleSizeChange" 
        @current-change="handleCurrentChange"
        :current-page="currentPage" 
        :page-sizes="[5,10,20]" 
        :page-size="pageSize"
        layout="prev, pager, next, sizes" 
        :total="ArticleList.length">
      </el-pagination>
    </el-card>

  </div>
</template>

<script>
export default {
  data(){
    return {
      ArticleList:[],
      cateList: [],
      tagList:[],
      title:'', // 搜索框内容
      img:'',  // 缩略图
      currentPage: 1, // 当前页码
      total: 10, // 总条数
      pageSize: 10, // 每页的数据条数
      cateName:'',
      tid: '',
      tagName:{}, // 当前文章的所有标签
    }
  },
  // 声明周期函数
  created(){
    this.getArticleList()
    this.getAllCates()
    this.getTagList()
  },
  methods:{
    async getArticleList(){
      const {data:res} = await this.$http.get('article', { 
        params:{
           title: this.title,
           pagesize:this.pageSize,
           pagenum:this.currentPage 
        }
      })
      if (res.status != 200){
        return this.$message.error(res.message)
      }
      this.ArticleList = res.data
      this.total = res.total
    },
    // 获取分类列表
    async getAllCates(){
      const {data:res} = await this.$http.get('category', { 
        params:{
           pagesize:this.pageSize,
           pagenum:this.currentPage 
           } 
        })
        // console.log(res)
        // console.log(res.status)
        if (res.status != 200){
          return this.$message.error(res.message)
        }
        this.cateList = res.data
    },
    // 获取标签列表
    async getTagList(){
      const {data:res} = await this.$http.get('tags')
      if(res.status !== 200){
        return this.$$message.error('获取标签列表失败')
      }
      this.tagList = res.data
    },
    //每页条数改变时触发 选择一页显示多少行
    handleSizeChange(val) {
        console.log(`每页 ${val} 条`);
        this.currentPage = 1; // 当改变'条/页'的时候总是从第一页开始显示
        this.pageSize = val; // 回调参数：每页条数
    },
    //当前页改变时触发 跳转其他页
    handleCurrentChange(val) {
        console.log(`当前页: ${val}`);
        this.currentPage = val; // 回调参数：当前页
    },
    // 新增文章:
    addArticleClick(){
       this.$router.push('/artadd')
    },

    // 编辑文章:按钮事件
    handleEdit(index, id,Articlename, role) {
      this.$router.push(`/artadd/${id}`)
    },

    // 删除文章
    async handleDelete(index, id) {
      // TODO:
      const confirmRes = await this.$confirm('此操作将永久删除该文章, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
        }).catch(err => err)
      if ('cancel' === confirmRes) {//用户点击了取消
        return this.$message.success('取消删除')
      }
      if ('confirm' === confirmRes) { //用户点击了确认
        const { data : res } = await this.$http.delete(`article/${id}`)
        if (res.status !== 200) {
          return this.$message.error(res.msg)
        }
        this.$message.success('删除成功')
        //刷新用户列表
        this.getArticleList()
      }
    },
    // 获取某一行文章的cid =>根据cid获取cid内容 
    formatCid(row, column){
      let cid = row.cid
      for(let i = 0; i < this.cateList.length; i++) {
        if(cid == this.cateList[i].id){
          this.cateName =  this.cateList[i].name
          return this.cateName
        }
      }
    },
    // 获取某一行文章的tid =>根据tagid获取tag字符串内容
    formatTag(row) {
      let id = row.ID
      let tid = row.tid
      let tidArr = tid.split(",")
      let temp = []
      // console.log(tidArr)
      for(let j = 0; j < tidArr.length; j++) {
        let ttid = parseInt(tidArr[j])  // 1
        for(let i = 0; i < this.tagList.length; i++) {
          if(ttid === this.tagList[i].ID){
            temp.push(this.tagList[i].name)
          }
        }
      }
      // this.tagName[id] = temp
      return temp.length > 5 ? temp.slice(0, 3):temp
      // console.log(this.tagName)
    },
    CellStyle(row,column,rowIndex,columnIndex) {
      if(row.columnIndex === 1){
        return 'color:red'
      }
    },   
  }
}
</script>

<style scoped>
.container,.el-card{
  /* height: 100%; */
}

.rowActionBtn{
  color:red
}


</style>