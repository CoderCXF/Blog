<template>
  <div class="container">
    <el-card>
      <!-- el-row表明这是一行，包含在其中的内容都写在一行中。
            Element UI 规定24rows && 24cols   
       -->
      <el-row :gutter="20">
        <el-col :span="12">
        </el-col>
        <el-col :span="4">
          <el-button type="primary" @click="addCateClick">新增分类</el-button>
        </el-col>
      </el-row>
      <!-- table表格 -->
      <!-- 这里使用了自定义列表，并且使用scope可以获取表中数据-->
      <el-table :data="CateList.slice((currentPage-1)*pageSize,currentPage*pageSize)">
        <el-table-column
          prop="id"
          label="ID"
          width="180px"
          align="center">
        </el-table-column>
        <el-table-column
          prop="name"
          label="用户名"
          width="180px"
          align="center">
        </el-table-column>
        <el-table-column
          label="操作"
          width="270px"
          align="center">
            <template slot-scope="scope">
              <el-button
                size="mini"
                icon="el-icon-edit"
                @click="handleEdit(scope.$index, scope.row.id, scope.row.name, scope.row.role)">编辑
              </el-button>
              <el-button
                size="mini"
                type="danger"
                icon="el-icon-delete"
                @click="handleDelete(scope.$index, scope.row.id)">删除
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
        :total="CateList.length">
      </el-pagination>
    </el-card>

    <!-- 新增用户区域 dialog-->
    <!-- dialog的关闭按钮需要绑定一个close事件，不绑定的话点击关闭按钮没有反应 -->
    <el-dialog title="添加分类" :visible="dialogFormVisible" show-close @close="onClose">
      <el-form :model="newCateInfo" ref="addCateFormRef" :rules="addCateFormRules">
        <el-form-item label="分类名" prop="name">
          <el-input v-model="newCateInfo.name" placeholder="请输入用户名"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancelAddCate">取 消</el-button>
        <el-button type="primary" @click="confirmAddCate">确 定</el-button>
      </div>
    </el-dialog>

    <!-- 编辑用户区域 开始-->
        <el-dialog title="编辑分类" :visible="dialogEditFormVisible" show-close @close="onCloseEdit">
          <el-form :model="CateInfo" ref="editCateFormRef" :rules="editCateFormRules">
            <el-form-item label="分类名" prop="name">
              <el-input v-model="CateInfo.name" :placeholder="CateInfo.name"></el-input>
            </el-form-item>
          </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancelEditCate">取 消</el-button>
        <el-button type="primary" @click="confirmEditCate">确 定</el-button>
      </div>
    </el-dialog>
    <!-- 编辑用户 结束 -->
  </div>
</template>

<script>
export default {
  data(){
    return {
      CateList:[],
      name:'',
      currentPage: 1, // 当前页码
      total: 10, // 总条数
      pageSize: 10, // 每页的数据条数

      // 新增用户dialog 数据
      dialogFormVisible: false,
      newCateInfo: {
        name: '',
      },
      role: '1',
      addCateFormRules: {
        name: [
          { required: true, message: '请输入分类名', trigger: 'blur' }
        ]
      },
      // 编辑分类区域数据
      CateInfo: {
        name: ''
      },
      dialogEditFormVisible: false,
      // 将要修改的用户的ID，通过table的scope获得
      toEditID: 0,
      editCateFormRules: {
        name: [
          { required: true, message: '请输入分类名', trigger: 'blur' },
        ]
      },
    }
  },
  // 声明周期函数
  created(){
    this.getCateList()
  },
  methods:{
    async getCateList(){
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
        this.CateList = res.data
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
    // 新增用户:
    addCateClick(){
      this.dialogFormVisible = true
    },
    onClose(){
      this.dialogFormVisible = false
      this.$refs.addCateFormRef.resetFields()
    },
    // 新增用户:dialog取消
    cancelAddCate(){
      this.dialogFormVisible = false
      // 要使得resetFields生效，还必须定义rules-prop
      this.$refs.addCateFormRef.resetFields()
    },
    // 新增用户:dialog确认提交按钮
    confirmAddCate(){
      this.newCateInfo.role = parseInt(this.role)
      this.$refs.addCateFormRef.validate(async (valid)=>{
        if(!valid){
          return this.$message.error('参数格式不确定')
        }
        const {data:res} = await this.$http.post('category/add',{
          name: this.newCateInfo.name
        })
        if(res.status != 200) return this.$message.error(err.msg)
        this.$message.success('添加用户成功')
        this.dialogFormVisible = false
        this.getCateList()
      })
    },

    // 编辑用户:按钮事件
    handleEdit(index, id, name, role){
      this.dialogEditFormVisible = true
      this.toEditID = id
      this.CateInfo.name = name
    },
    // 编辑用户d:ialog关闭
    onCloseEdit() {
      this.dialogEditFormVisible = false
      this.$refs.editCateFormRef.resetFields()
    },
    // 编辑用户：dialog取消
    cancelEditCate() {
      this.onCloseEdit()
    },
    // 编辑用户：dialog确定
    confirmEditCate() {
      this.$refs.editCateFormRef.validate(async (valid)=>{
        if(!valid){
          return this.$message.error('参数格式不确定')
        }
        const {data:res} = await this.$http.put(`category/${this.toEditID}`,{
          name: this.CateInfo.name
        })
        if(res.status != 200) return this.$message.error(err.msg)
        this.$message.success('分类修改成功')
        this.dialogEditFormVisible = false
        this.getCateList()
      })
    },
    // 删除分类
    async handleDelete(index, id){
      // TODO:
      const confirmRes = await this.$confirm('此操作将永久删除该文件, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
        }).catch(err => err)
      if ('cancel' === confirmRes) {// 用户点击了取消
        return this.$message.success('取消删除')
      }
      if ('confirm' === confirmRes) { // 用户点击了确认
        const { data : res } = await this.$http.delete(`category/${id}`)
        if (res.status !== 200) {
          return this.$message.error(res.msg)
        }
        this.$message.success('删除成功')
        //刷新用户列表
        this.getCateList()
      }
    }
  }
}
</script>