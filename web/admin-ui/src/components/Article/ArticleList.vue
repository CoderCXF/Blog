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
      <el-table :data="ArticleList.slice((currentPage-1)*pageSize,currentPage*pageSize)">
        <el-table-column
          prop="ID"
          label="ID"
          width="50px"
          align="center">
        </el-table-column>
        <el-table-column
          prop="cid"
          label="分类"
          width="50px"
          align="center">
        </el-table-column>
        <el-table-column
          prop="title"
          label="文章标题"
          width="250px"
          align="center">
        </el-table-column>
        <el-table-column
          prop="desc"
          label="文章描述"
          width="250px"
          align="center">
        </el-table-column>
        <el-table-column
          prop="img"
          label="缩略图"
          width="200px"
          align="center">
          <template slot-scope="scope">
            <!-- 这里插入图片需要用到slot机制，为src绑定参数 -->
            <img v-bind:src="scope.row.img" style="width:70px;height:70px;" alt="暂无图片">
          </template>
        </el-table-column>
        <el-table-column
          label="操作"
          width="180px"
          align="center">
            <template slot-scope="scope">
              <el-button
                size="mini"
                icon="el-icon-edit"
                @click="handleEdit(scope.$index, scope.row.ID, scope.row.Articlename, scope.row.role)">编辑
              </el-button>
              <el-button
                size="mini"
                type="danger"
                icon="el-icon-delete"
                @click="handleDelete(scope.$index, scope.row.ID)">删除
              </el-button>
            </template>
        </el-table-column>
      </el-table>
      <!-- 分页操作 -->
      <el-pagination align='center' 
        @size-change="handleSizeChange" 
        @current-change="handleCurrentChange"
        :current-page="currentPage" 
        :page-sizes="[1,5,10,20]" 
        :page-size="pageSize"
        layout="total, sizes, prev, pager, next, jumper" 
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
      title:'', // 搜索框内容
      img:'',  // 缩略图
      currentPage: 1, // 当前页码
      total: 10, // 总条数
      pageSize: 10, // 每页的数据条数
    }
  },
  // 声明周期函数
  created(){
    this.getArticleList()
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
      console.log(res)
      this.ArticleList = res.data
      this.total = res.total
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
      const confirmRes = await this.$confirm('此操作将永久删除该文件, 是否继续?', '提示', {
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
  }
}
</script>

<style scoped>
.container,.el-card{
  height: 100%;
}
</style>