<template>
  <div class="container">
    <div class="btn-wrap">
        <el-button  class="btn1" type="primary" size="medium" @click="saveToDrafts">保存草稿</el-button>
        <el-button type="danger" size="medium" class="btn2" @click="postArticle">发布</el-button>
    </div>
    <el-form class="form" :model="artInfo" :rules="formRules">
      <el-form-item prop="title">
        <el-input placeholder="请输入文章标题" v-model="artInfo.title" class="input_title"></el-input>
      </el-form-item>
      <el-form-item class="form-item" prop="content">
        <mavon-editor v-model="artInfo.content" :ishljs="true" ref=md @imgAdd="$imgAdd" @imgDel="$imgDel" class="editor"/>
      </el-form-item>
    </el-form>

    <!-- 点击发布显示的对话框按钮 -->
    <el-dialog title="文章发布" :visible="dialogPostVisible" show-close @close="onCloseEdit">
      <!-- <el-divider></el-divider> -->
      <el-tabs type="border-card">
        <el-tab-pane label="常规" class="tabsPaneNorm">
          <el-form :v-model="artInfo" label-width="120px" :rules="dialogFormRule" ref="dialogFormRef">
            <el-form-item label="文章标题：" class="form-item-dialog" prop="title">
              <el-input v-model="artInfo.title"></el-input>
            </el-form-item>
            <el-form-item label="分类目录：" class="form-item-dialog" prop="cid">
              <el-radio-group v-model="artInfo.cid" class="btn-group">
                <el-radio class="radio" v-for="(item, index) in cateList" :key="index" :label="item.id">{{item.name}}</el-radio>
              </el-radio-group>
              <el-button size="small" class="add-cate-btn" @click="addNewCate">新增分类</el-button>
            </el-form-item>
            <el-form-item label="标签：" class="form-item-dialog" prop="tag">
              <!-- tags中保存的是选中的标签的ID -->
              <el-select 
                v-model="artInfo.tag" 
                clearable 
                multiple 
                filterable 
                allow-create 
                default-first-option 
                :multiple-limit="5" 
                placeholder="选择标签，最多显示5个，可以用回车添加"
                @blur="showSelected"
                class="sel-tag"
                @change="selChange"
                >
                <el-option v-for="item in tagList" :key="item.ID" :label="item.name" :value="item.ID"></el-option>
              </el-select>
            </el-form-item>
            <!-- 描述 -->
            <el-form-item label="文章摘要：" class="form-item-dialog" prop="desc">
              <el-input v-model="artInfo.desc" type="textarea" placeholder="请输入文章摘要，如果不写的话则为文章标题">
              </el-input>
            </el-form-item>
            <el-form-item label="封面图：" prop="img">
              <el-upload
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
            </el-form-item>
          </el-form>
          <div class="btn-group2">
            <el-button class="post-or-save-btn">保存至草稿</el-button>
            <el-button class="post-or-save-btn" @click="postArticleFinal">发布</el-button>
            <el-button size="small" class="closeBtn" @click="onCloseBtn">关闭</el-button>
          </div>
        </el-tab-pane>
        <el-tab-pane label="消息中心">高级</el-tab-pane>
      </el-tabs>
    </el-dialog>

    <!-- 预览封面图 -->
    <el-dialog :visible.sync="showPreview" center width="40%">
        <el-image :src="artInfo.img" fit="fill" alt="加载失败"></el-image>
    </el-dialog>

    <!-- 新增用户区域 dialog-->
    <!-- dialog的关闭按钮需要绑定一个close事件，不绑定的话点击关闭按钮没有反应 -->
    <el-dialog title="添加分类" :visible="dialogFormVisible" show-close @close="onCloseAddNewCate">
      <el-form :model="newCateInfo" ref="addCateFormRef" :rules="addCateFormRules">
        <el-form-item label="分类名" prop="name">
          <el-input v-model="newCateInfo.name" placeholder="请输入分类名"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancelAddCate">取 消</el-button>
        <el-button type="primary" @click="confirmAddCate">确 定</el-button>
      </div>
    </el-dialog>

  </div>
</template>

<script>
import {URL} from '../../plugin/http'
export default {
  data(){
    return {
      artInfo:{
        id:0,
        title:'',
        cid:0,
        desc:'',
        content:'',
        img:'',
        tid: '',
        tag:[],
      },
      newCateInfo:{
        name:'',
        role:'',
      },
      role:0,
      title:'',
      formRules:{
        title:[{ required: true, message: '请输入文章标题', trigger: 'blur' },
               { min: 10, max: 35, message: '长度在 10 到 35 个字符', trigger: 'blur' }],
        content:[{ required: true, message: '请输入文章内容', trigger: 'blur' }],
      },
      dialogFormRule:{
            title:[{ required: true, message: '请输入文章标题', trigger: 'blur' },
                   { min: 10, max: 35, message: '长度需在 1 到 35 个字符', trigger: 'blur' }],
            cid:  [{ required: true, message: '请选择文章分类', trigger: 'blur' }],
            desc: [{min:0, max: 150, message: '最多用150个字描述', trigger: 'blur' }],
            // img:[{ required: true, message: '为了优美，请上传一张缩略图', trigger: 'blur'}],
            content:[{ required: true, message: '请输入文章内容', trigger: 'blur' }],

      },
      addCateFormRules: {
        name: [
          { required: true, message: '请输入分类名', trigger: 'blur' }
        ]
      },
      dialogPostVisible:true,
      // 分类列表
      cateList:[],
      pagesize: 0,
      pagenum: 20,
      curCate:0,
      // 标签列表
      tagList: [],
      choiceOfTag: '',
      tags:[],
      upRUL: URL + '/upload', // 上传地址，这里就是后端api
      headers:{},
      showPreview: false,
      dialogFormVisible:false,
    }
  },
  created(){
    this.getCateList()
    this.getTagList()
    this.headers = {Authorization: `Bearer ${window.sessionStorage.getItem('token')}`}
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
        this.cateList = res.data
    },

    async getTagList() {
      const {data:res} = await this.$http.get('tags')
      if(res.status !== 200){
        return this.$$message.error('获取标签列表失败')
      }
      this.tagList = res.data
    },

    $imgAdd(pos, $file){
      var formdata = new FormData();
      formdata.append('File', $file);
      this.img_file[pos] = $file;
      this.$http({
        url: '/upload',
        method: 'post',
        data: formdata,
        headers: { 'Content-Type': 'multipart/form-data' },
      }).then((res) => {
        // console.log(res)
        let _res = res.data;
        // 第二步.将返回的url替换到文本原位置![...](0) -> ![...](url)
        this.$refs.md.$img2Url(pos, _res.url);
      })
    },
    $imgDel(pos, $file){
      // console.log(pos, $file)
    },
    saveToDrafts(){
      console.log('保存至草稿箱')
    },
    
    // 关闭按钮和右上角叉掉按钮
    onCloseEdit(){
      this.dialogPostVisible = false
    },
    onCloseBtn(){
      this.dialogPostVisible = false
    },
    // selected
    showSelected(e){
      // console.log(e)
    },
    // 图片上传
    handleOnPreview(){
      this.showPreview = true
    },
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
    // 新增分类
    addNewCate(){
      this.dialogFormVisible = true
    },
    onCloseAddNewCate(){
      this.dialogFormVisible = false
    },
    cancelAddCate(){
      this.dialogFormVisible = false
    },
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
        this.$message.success('添加分类成功')
        this.dialogFormVisible = false
        this.getCateList()
      })
    },
    // 发布文章:弹出选择分类标签等对话框页面
    postArticle(){
      this.dialogPostVisible = true
    },
    // dialog中发布文章
    async postArticleFinal(){
      this.$refs.dialogFormRef.validate((valid)=>{
        if(!valid){
          return this.$message.error('参数格式不确定')
        }
      })
      console.log('点击了发布文章按钮')
      // 处理摘要部分
      if(!this.artInfo.desc){
        this.artInfo.desc = this.artInfo.title
      }
      // 处理标签部分，获取所有标签，拼接为字符串保存在数据库
      // console.log(this.tagList)
      this.artInfo.tag.forEach((value)=>{
        let str = value.toString()
        this.artInfo.tid = this.artInfo.tid + str + ','
      })
      // 添加文章
      
      console.log(this.artInfo)
    },
    // 下拉选择标签选中值事件
    async selChange(val) {
      // this.$message.success(val)
      // 处理标签部分（重点处理的是是否添加了新的标签，如果有新的标签的话，需要将其保存至数据库中）
      var tag = this.artInfo.tag[this.artInfo.tag.length - 1]
      if(typeof tag === 'string') {
        // 将新的标签保存在数据库中
        // 添加新的标签，弹出dialog
        const confirmRes = await this.$confirm('此标签不存在, 是否添加新标签?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
          }).catch(err => err)  
        if ('cancel' === confirmRes) {
          return
        }
        if ('confirm' === confirmRes) { 
          const {data:res} = await this.$http.post('tag/add',{
          name: tag
          })
          if(res.status !== 200) {
            // return this.$message.error(res.message)
          }
          this.getTagList()
        }
      }
    },
  }
}
</script>

<style scoped>
.container .editor{
  /* margin-top: 50px; */
  width: 100%;
  height: 500px;
}

.container>>>.el-form .el-form-item{
  margin: 20px 0px;
}

.el-input{
  /* border-radius: 0; */
  margin-bottom: 10px;
}

.btn-wrap{
  display: flex;
  flex-direction: row;
  flex-wrap: wrap-reverse;
  justify-content: flex-end;
  align-items: center;
  margin-bottom: 10px;
}

.btn1{
  background-color: white;
  color: #000; 
  border: 1px solid rgb(217, 217, 217);
}

.btn1:hover{
  background-color: #fff;
  color: #ff7a45;
  border: 1px solid #ff7a45;
}

.btn2{
  background-color: red;
}

/* 深度选择器 */
.container>>>.el-input__inner{
  border-radius: 0;
}

/* 鼠标悬停状态 */
.container>>>.el-input__inner:hover{
  border: 1px solid red;
}
/* 选中状态 */
.container>>>.el-input__inner:focus{
  border: 1px solid red;
}

.container>>>.el-textarea__inner{
  border-radius: 0;
}

/* 鼠标悬停状态 */
.container>>>.el-textarea__inner:hover{
  border: 1px solid red;
}
/* 选中状态 */
.container>>>.el-textarea__inner:focus{
  border: 1px solid red;
}

.tabsPaneNorm>>>.el-form{
  white-space: nowrap;
}

.btn-group{
  display: flex;
  flex-wrap: wrap;
}

.btn-group .radio{
  padding: 5px;
}

.add-cate-btn{
  background-color:#fff;
  color:#000;
  border: 1px dashed #e8e8e8;
}

.add-cate-btn:hover{
  background-color:#fff;
  color: #ff7a45;
  border: 1px dashed #ff7a45;
}

.sel-tag{
  width: 100%;
}

.closeBtn:hover{
  background-color: #fff;
  color: #ff7a45;
  border: 1px solid #ff7a45;
}

.btn-group2{
  display: flex;
  justify-content: flex-end;
}

.post-or-save-btn{
  background-color: red;
  color: #fff;

}

.post-or-save-btn:hover{
  background-color: red;
  color: #fff;
}



</style>