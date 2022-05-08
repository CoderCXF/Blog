<template>
  <div class="container">
    <el-col :span="12">
      <el-card>
        <div class="avatr">
          <el-image :src="myInfo.avatar" alt="未设置头像" style="width:100px;height:100px;border-radius:50px;margin: 20px;"></el-image>
          <div style="color:rgba(0, 0, 0, 0.85);font-weight: 500;line: height 1.25rem;">{{myInfo.username}}</div>
          <div>{{myInfo.introduce}}</div>
        </div>
        <el-row>
          <i class="el-icon-link icon1"></i>
          <span>http://dancingcode.cn</span> 
        </el-row>
        <el-row>
          <i class="el-icon-message icon1"></i>
          <span>hi@gin.run</span> 
        </el-row>
        <el-row>
          <i class="el-icon-date icon1"></i>
          <span>{{myInfo.regTime}}天</span> 
        </el-row>
        <el-divider></el-divider>
        <el-row>
          <i class="el-icon-info" style="padding-right: 10px"></i>
          <span>累计发表了1篇文章</span>
        </el-row>
        <el-divider></el-divider>
        <el-row>
          <i class="el-icon-info" style="padding-right: 10px"></i>
          <span>累计发表了1篇文章</span>
        </el-row>
        <el-divider></el-divider>
        <el-row>
          <i class="el-icon-info" style="padding-right: 10px"></i>
          <span>累计发表了1篇文章</span>
        </el-row>
        <el-divider></el-divider>
        <el-row>
          <i class="el-icon-info" style="padding-right: 10px"></i>
          <span>累计发表了1篇文章</span>
        </el-row>
        <el-divider></el-divider>
        <el-row>
          <i class="el-icon-info" style="padding-right: 10px"></i>
          <span>累计发表了1篇文章</span>
        </el-row>
      </el-card>
    </el-col>
    <!-- 右半部分 -->
    <el-col :span="12">
      <el-card>
        <div style="padding-bottom:20px;">编辑资料</div>
        <el-tabs type="border-card">
        <el-tab-pane>
          <span slot="label"><i class="el-icon-s-custom"></i> 基本资料</span>
            <el-form :model="myInfo">
              <el-form-item label="用户名">
                <el-input v-model="myInfo.username"></el-input>
              </el-form-item>
              <el-form-item label="昵称">
                <el-input v-model="myInfo.nickname"></el-input>
              </el-form-item>
              <el-form-item label="邮箱">
                <el-input v-model="myInfo.email"></el-input>
              </el-form-item>
              <el-form-item label="博客说明">
                <el-input type="textarea" v-model="myInfo.introduce"></el-input>
              </el-form-item>
            </el-form>
            <el-button type="danger">保存</el-button>
        </el-tab-pane>
        <el-tab-pane>
          <span slot="label"><i class="el-icon-lock"></i> 密码</span>
          <el-form :model="myInfo">
            <el-form-item label="原始密码">
              <el-input v-model="myInfo.username"></el-input>
            </el-form-item>
            <el-form-item label="新密码">
              <el-input v-model="myInfo.nickname"></el-input>
            </el-form-item>
            <el-form-item label="确认密码">
              <el-input v-model="myInfo.email"></el-input>
            </el-form-item>
          </el-form>
          <el-button type="danger">确认更改</el-button>
        </el-tab-pane>
        <el-tab-pane>
          <span slot="label"><i class="el-icon-notebook-1"></i>任务活动</span>
          暂未开放任务活动
        </el-tab-pane>
      </el-tabs>
      </el-card>
    </el-col>
  </div>
</template>

<script>
export default {
  data(){
    return{
      myInfo: {
        avatar: require('../../assets/avatar.jpg'),
        username: 'Admin',
        nickname: '',
        password: '',
        email: '',
        regTime: 1,
        introduce: '自我介绍',
      }
    }
  }
}
</script>

<style scoped>
.container{
  margin-top: 20px;
  font-size: 14px;
}
.avatr{
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.icon1{
  padding-right: 10px;
  margin: 10px 0;
}
</style>