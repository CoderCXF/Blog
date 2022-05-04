<template>
  <div>
    <el-card>
      <h3>{{id?'编辑文章':'添加文章'}}</h3>
      <el-form :model="artInfo" :rules="formRules" ref="artInfoRef">
        <el-form-item label="文章标题" prop="title">
          <br>
          <el-input v-model="artInfo.title" placeholder="请输入文章标题"></el-input>
        </el-form-item>
        <el-form-item label="文章分类" prop="cid">
          <br>
          <!-- 目前分类只支持选择一类，可以扩展为多类,默认的分类就是原来文章的分类 -->
          <el-select v-model="cid" :placeholder="cidName? cidName:'请选择文章分类'" clearable style="width: 500px;" @change="optionChange">
            <el-option
              v-for="item in cateList"
              :key="item.id"
              :label="item.name"
              :value="item.id">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="文章描述" prop="desc">
          <el-input type="textarea" v-model="artInfo.desc" placeholder="最多用150字描述本篇文章"></el-input>
        </el-form-item>
        <el-form-item label="文章缩略图" prop="img">
          <!-- TODO: 文件上传 -->
          <br>
          <!-- 必须设置上传字段的文件名，与后端保持一致，后端为'File' -->
          <el-upload
            ref='upLoadRef'
            class="upload-demo"
            drag
            :limit="1"
            accept="image/jpeg,image/gif,image/png"
            list-type="picture"
            name='File'
            :action=upRUL
            :headers=headers
            :on-preview="handleOnPreview"
            :on-success="handleSuccess"
            :on-error="handleError"
            >
            <i class="el-icon-upload"></i>
            <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
            <div class="el-upload__tip" slot="tip">
              <i class="el-icon-warning" style="color:red;"></i>
              只能上传jpg/png文件，且不超过500kb，只能上传一张哦
            </div>
            
          </el-upload>
          <!-- 理论上是不显示的 -->
          <template v-if="id">
            <img :src="artInfo.img" style="width:120px;height:100px;" alt="暂未添加缩略图" />
          </template>
        </el-form-item>
        <el-form-item label="文章内容" prop="content">
          <br>
          <Editor v-model="artInfo.content" placeholder="请编辑文章内容"></Editor>
        </el-form-item>
        <el-form-item class="loginBtn">
          <el-button type="primary" @click="submitArt(artInfo.id)">{{artInfo.id?'更新':'提交'}}</el-button>
          <el-button @click="cancelSubmit">取消</el-button>
        </el-form-item>
      </el-form>
    </el-card>
    <!-- 图片预览 -->
    <el-dialog :visible.sync="showPreview" center width="40%">
        <el-image :src="artInfo.img" fit="fill" alt="加载失败"></el-image>
    </el-dialog>
  </div>
</template>

<script>
import {URL} from '../../plugin/http'
import Editor from '../editor/index.vue'
export default {
  components:{ Editor},
  // 是从文章列表页面传递的参数
  props:['id'],
  data(){
    return{
      artInfo:{
        id:0,
        title:'',
        cid:0,
        desc:'',
        content:'',
        img:''
      },
      cid:'',
      cidName:'',
      cateList:[],
      selectPlh:'请选择文章分类',
      upRUL: URL + '/upload', // 上传地址，这里就是后端api
      headers:{},
      showPreview: false,
      formRules:{
        title:[{ required: true, message: '请输入文章标题', trigger: 'blur' },
               { min: 1, max: 35, message: '长度在 1 到 35 个字符', trigger: 'blur' }],
        cid:  [{ required: true, message: '请选择文章分类', trigger: 'blur' }],
        desc: [{min:0, max: 150, message: '最多用150个字描述', trigger: 'blur' }],
        content:[{ required: true, message: '请输入文章内容', trigger: 'blur' }],
        img:[{ required: true, message: '为了优美，请上传一张缩略图', trigger: 'blur'}]
      }
    }
  },
  created(){
    // 是从编辑页面过来的
    this.getCateList()
    // 请求头就是设置的JWT的token
    this.headers = {Authorization: `Bearer ${window.sessionStorage.getItem('token')}`}
    if(this.id){
      this.getArtInfo(this.id)
    }
  },
  methods:{
    async getArtInfo(id){
      const {data:res} = await this.$http.get(`article/info/${id}`)
      if(res.status != 200) return this.$message.error('获取该文章信息失败')
      this.artInfo = res.data
      this.artInfo.id = res.data.ID
      this.cidName = res.data.Category.name
      // console.log(res.data)
      this.$message.success('文章同步成功，开始编辑')
    },
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
    // 获取所有的分类以供选择
    async getCateList() {
      const {data:res} = await this.$http.get('category')
      if(res.status != 200) return this.$message.error('获取用户列表失败')
      this.cateList = res.data
    },
    // 文章分类：选择事件
    optionChange(value){
      this.artInfo.cid = value
    },

    handleOnPreview(){
      this.showPreview = true
    },
    // 上传缩略图
    // on-chage钩子函数官网描述：
    // 文件状态改变时的钩子，添加文件、上传成功和上传失败时都会被调用
    // 所以该函数会被调用两次
    handleSuccess(response, file, fileList){
      let status = response.status
      if(status != 200) {
        return this.$message.error('图片上传失败')
      }
      this.$message.success('图片上传成功')
      // 记得一定要处理缩略图，因为上面的表单的upload没有双向绑定
      this.artInfo.img = response.url
    },
    handleError(){
      this.$message.error('图片上传失败')
    },
    // 提交或更新表单
    submitArt(id){
      this.$refs.artInfoRef.validate(async(valid)=>{
        if(id === 0){
          const {data:res} = await this.$http.post('article/add',this.artInfo)
          if(res.status !== 200){
            return this.$message.error(res.message)
          }
          this.$message.success('添加文章成功')
          this.$router.push('/artlist')
        } else {
          const {data:res} = await this.$http.put(`article/${id}`,this.artInfo)
          if(res.status !== 200){
            return this.$message.error(res.message)
          }
          this.$message.success('更新文章成功')
          this.$router.push('/artlist')
        }
      })
    },
    cancelSubmit(){
      this.$refs.artInfoRef.resetFields()
    },
  }
}
</script>

