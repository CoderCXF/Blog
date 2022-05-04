<template>
<div class="container">
  <div class="loginbox">
    <el-form :model="formData" :rules="rules" ref="formData" class="loginForm">
      <el-form-item prop="username" class="loginInput">
        <el-input v-model="formData.username" prefix-icon="el-icon-user" placeholder="用户名"></el-input>
      </el-form-item>
      <el-form-item prop="password" class="loginInput">
        <el-input type="password" 
        v-model="formData.password" 
        prefix-icon="el-icon-lock" 
        placeholder="密码"
        v-on:keyup.enter="onLogin"
        >
        </el-input>
      </el-form-item>
      <el-form-item class="loginBtn">
        <el-button type="primary" @click="onLogin">登录</el-button>
        <el-button @click="resetForm">取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</div>
</template>


<script>
  export default {
    data() {
      return {
        formData: {
          username: '',
          password: ''
        },
        rules: {
          username: [
            { required: true, message: '请输入用户名', trigger: 'blur' },
            { min: 4, max: 12, message: '长度在 4 到 12 个字符', trigger: 'blur' }
          ],
          password: [
            { required: true, message: '请输入密码', trigger: 'blur' },
            { min: 6, max: 12, message: '长度在 6 到 12 个字符', trigger: 'blur' }
          ],
        }
      };
    },
    methods: {
      resetForm() {
        this.$refs.formData.resetFields();
      },
      onLogin(){
        // 登录的时候进行验证
        this.$refs.formData.validate(async (valid)=>{
          if(valid){
            // 请求localhost:3000/api/v1/login
            // main.js中已经规定axios.defaults.baseURL = 'http://localhost:3000/api/v1'，所以这里就是直接写login
            var {data:res} = await this.$http.post('login',this.formData)
            console.log(res)
            if (res.status != 200){
              return this.$message.error(res.message)
            }
            window.sessionStorage.setItem('token', res.token)
            // 编程式导航：前往admin页面
            this.$router.push('/index')
          } else {
            this.$message({
              showClose: true,
              message: '用户名或密码格式错误',
              type: 'error',
              duration: 1000
            })
          }
        })
      }
    }
  }
</script>

<style scoped>
.container{
    height: 100%;
    background-color: #282c34;
}

.loginbox{
    width: 450px;
    height: 300px;
    background-color: #fff;
    /*没有使用flex，因为兼容性不好，这里使用相对定位*/
    position: absolute;
    left: 70%;
    top: 50%;
    transform: translate(-50%,-50%);
    border-radius: 9px;
}
.loginForm{
    width: 100%;
    position: absolute;
    bottom: 5%;
    padding: 0 20px;
    box-sizing: border-box;
} 

.loginInput{

}

.loginBtn{
    display: flex;
    justify-content: flex-end;
}

</style>
