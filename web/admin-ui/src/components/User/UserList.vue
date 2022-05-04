<template>
  <div class="container">
    <el-card>
      <!-- el-row表明这是一行，包含在其中的内容都写在一行中。
            Element UI 规定24rows && 24cols   
       -->
      <el-row :gutter="20">
        <el-col :span="12">
          <!-- clearable属性比较有趣，删除inpout框内容会自动的调整 -->
          <el-input placeholder="输入用户名查找" v-model="username" clearable @change="getUserList">
            <el-button slot="append" icon="el-icon-search" @click="getUserList"></el-button>
          </el-input>
        </el-col>
        <el-col :span="4">
          <el-button type="primary" @click="addUserClick">新增</el-button>
        </el-col>
      </el-row>
      <!-- table表格 -->
      <!-- 这里使用了自定义列表，并且使用scope可以获取表中数据-->
      <el-table :data="userList.slice((currentPage-1)*pageSize,currentPage*pageSize)">
        <el-table-column
          prop="ID"
          label="ID"
          width="180px"
          align="center">
        </el-table-column>
        <el-table-column
          prop="username"
          label="用户名"
          width="180px"
          align="center">
        </el-table-column>
        <el-table-column
          label="角色"
          width="180px"
          align="center">
          <template slot-scope="scope">
          <i class="el-icon-time"></i>
          <!-- 这里将角色ID转换为文字 -->
          <span style="margin-left: 10px">{{ scope.row.role == 1?'管理员': scope.row.role == 2? '订阅用户':'无效用户' }}</span>
      </template>
        </el-table-column>
        <el-table-column
          label="操作"
          width="270px"
          align="center">
            <template slot-scope="scope">
              <el-button
                size="mini"
                icon="el-icon-edit"
                @click="handleEdit(scope.$index, scope.row.ID, scope.row.username, scope.row.role)">编辑
              </el-button>
              <el-button
                size="mini"
                type="danger"
                icon="el-icon-delete"
                @click="handleDelete(scope.$index, scope.row.ID)">删除
              </el-button>
              <el-button
                size="mini"
                type="info"
                icon="el-icon-warning-outline"
                @click="onResetPwd">重置
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
        :total="userList.length">
      </el-pagination>
    </el-card>

    <!-- 新增用户区域 dialog-->
    <!-- dialog的关闭按钮需要绑定一个close事件，不绑定的话点击关闭按钮没有反应 -->
    <el-dialog title="添加用户" :visible="dialogFormVisible" show-close @close="onClose">
      <el-form :model="newUserInfo" ref="addUserFormRef" :rules="addUserFormRules">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="newUserInfo.username" placeholder="请输入用户名"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input type="password" v-model="newUserInfo.password" placeholder="请输入密码" show-password></el-input>
        </el-form-item>
        <el-form-item label="确认密码" prop="checkpass">
          <el-input type="password" v-model="newUserInfo.checkpass" placeholder="确认密码" show-password suffix-icon="el-icon-circle-check"></el-input>
        </el-form-item>
        <el-form-item label="角色" prop="role">
          <br>
          <el-select v-model="role" placeholder="请选择角色" style="width:100%">
            <!-- value就是返回的值 value = role -->
            <el-option label="管理员" value="1"></el-option>
            <el-option label="订阅用户" value="2"></el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancelAddUser">取 消</el-button>
        <el-button type="primary" @click="confirmAddUser">确 定</el-button>
      </div>
    </el-dialog>

    <!-- 编辑用户区域 开始-->
        <el-dialog title="编辑用户" :visible="dialogEditFormVisible" show-close @close="onCloseEdit">
      <el-form :model="userInfo" ref="editUserFormRef" :rules="editUserFormRules">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="userInfo.username" :placeholder="userInfo.username"></el-input>
        </el-form-item>
        <el-form-item label="角色" prop="role">
          <br>
          <el-select v-model="role" placeholder="请选择用户角色" style="width:100%">
            <!-- value就是返回的值 value = role -->
            <el-option label="管理员" value="1"></el-option>
            <el-option label="订阅用户" value="2"></el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancelEditUser">取 消</el-button>
        <el-button type="primary" @click="confirmEditUser">确 定</el-button>
      </div>
    </el-dialog>
    <!-- 编辑用户 结束 -->
  </div>
</template>

<script>
export default {
  data(){
    return {
      userList:[],
      username:'',
      currentPage: 1, // 当前页码
      total: 10, // 总条数
      pageSize: 10, // 每页的数据条数

      // 新增用户dialog 数据
      dialogFormVisible: false,
      newUserInfo: {
        username: '',
        password: '',
        checkpass: '',
        role: 2
      },
      role: '1',
      addUserFormRules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 4, max: 12, message: '长度在 4 到 12 个字符', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 6, max: 12, message: '长度在 6 到 12 个字符', trigger: 'blur' }
        ],
        checkpass: [
           //简单校验密码
           { validator: (rule, value, callback)=>{
              if (value === '') {
                callback(new Error('请再次输入密码'));
              } else if (value !== this.userInfo.password) {
                callback(new Error('两次输入密码不一致!'));
              } else {
                callback();
              }
            }, trigger: 'blur' 
           }
        ],
      },
      // 编辑用户区域数据
      userInfo: {
        username: '',
        password: '',
        checkpass: '',
        role: 2
      },
      dialogEditFormVisible: false,
      // 将要修改的用户的ID，通过table的scope获得
      toEditID: 0,
      editUserFormRules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 4, max: 12, message: '长度在 4 到 12 个字符', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 6, max: 12, message: '长度在 6 到 12 个字符', trigger: 'blur' }
        ],
        checkpass: [
           //简单校验密码
           { validator: (rule, value, callback)=>{
              if (value === '') {
                callback(new Error('请再次输入密码'));
              } else if (value !== this.userInfo.password) {
                callback(new Error('两次输入密码不一致!'));
              } else {
                callback();
              }
            }, trigger: 'blur' 
           }
        ],
      },
    }
  },
  // 声明周期函数
  created(){
    this.getUserList()
  },
  methods:{
    async getUserList(){
      const {data:res} = await this.$http.get('users', { 
        params:{
           username: this.username,
           pagesize:this.pageSize,
           pagenum:this.currentPage 
           } 
        })
      // console.log(this.username)
      // console.log(res.status)
      if (res.status != 200){
        return this.$message.error(res.message)
      }
      this.userList = res.data
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
    addUserClick(){
      this.dialogFormVisible = true
    },
    onClose(){
      this.dialogFormVisible = false
      this.$refs.addUserFormRef.resetFields()
    },
    // 新增用户:dialog取消
    cancelAddUser(){
      this.dialogFormVisible = false
      // 要使得resetFields生效，还必须定义rules-prop
      this.$refs.addUserFormRef.resetFields()
    },
    // 新增用户:dialog确认提交按钮
    confirmAddUser(){
      this.newUserInfo.role = parseInt(this.role)
      this.$refs.addUserFormRef.validate(async (valid)=>{
        if(!valid){
          return this.$message.error('参数格式不确定')
        }
        const {data:res} = await this.$http.post('user/add',{
          username: this.newUserInfo.username,
          password: this.newUserInfo.password,
          role: this.newUserInfo.role
        })
        if(res.status != 200) return this.$message.error(err.msg)
        this.$message.success('添加用户成功')
        this.dialogFormVisible = false
        this.getUserList()
      })
    },

    // 编辑用户:按钮事件
    handleEdit(index, id,username, role){
      this.dialogEditFormVisible = true
      this.toEditID = id
      this.userInfo.username = username
      this.userInfo.role = role
      this.role = role.toString()
    },
    // 编辑用户d:ialog关闭
    onCloseEdit() {
      this.dialogEditFormVisible = false
      this.$refs.editUserFormRef.resetFields()
    },
    // 编辑用户：dialog取消
    cancelEditUser() {
      this.onCloseEdit()
    },
    // 编辑用户：dialog确定
    confirmEditUser() {
      this.userInfo.role = parseInt(this.role)
      this.$refs.editUserFormRef.validate(async (valid)=>{
        if(!valid){
          return this.$message.error('参数格式不确定')
        }
        const {data:res} = await this.$http.put(`user/${this.toEditID}`,{
          username: this.userInfo.username,
          role: this.userInfo.role
        })
        if(res.status != 200) return this.$message.error(err.msg)
        this.$message.success('用户修改成功')
        this.dialogEditFormVisible = false
        this.getUserList()
      })
    },
    // 删除用户
    async handleDelete(index, id){
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
        const { data : res } = await this.$http.delete(`/user/${id}`)
        if (res.status !== 200) {
          return this.$message.error(res.msg)
        }
        this.$message.success('删除成功')
        //刷新用户列表
        this.getUserList()
    }
    },
    // 重置密码功能
    onResetPwd(){
      this.$message.warning('重置密码功能暂未开放')
    }
    
  }
}
</script>