<template>
  <div class="container">
    <el-col :span="12">
      <el-card style="border-radius: 0">
        <span class="card-head-title">{{ID === 0? '添加标签':'修改标签'}}</span>
        <el-divider></el-divider>
        <el-form :model="tagInfo" ref="tagInfoRef" :rules="formRules">
          <el-form-item label="标签名：" prop="name">
            <el-input v-model="tagInfo.name" placeholder="请输入新的标签名，注意不要添加已经存在的标签"></el-input>
          </el-form-item>
          <el-form-item label="颜色：" prop="color">
            <br>
            <el-row>
              <el-col :span="22">
                <el-input v-model="tagInfo.color"></el-input>
              </el-col>
              <el-col :span="1">
                <el-color-picker v-model="tagInfo.color" :predefine="predefineColors" style="padding-left:5px;"></el-color-picker>
              </el-col>
            </el-row>
          </el-form-item>
          <el-form-item class="loginBtn">
            <el-button type="danger" v-if="ID===0" style="background-color:red;" @click="saveTag">保存</el-button>
            <div v-if="ID > 0">
              <el-button type="danger" style="background-color:red; margin-right: 0;" @click="updateTag">更新</el-button>
              <el-button type="primary" @click="backToAdd" class="backToAddBtn">返回添加</el-button>
            </div>
            
        </el-form-item>
        </el-form>

      </el-card>
    </el-col>
    <el-col :span="12">
      <el-card>
        <span class="card-head-title">所有标签</span>
        <el-divider></el-divider>
        <div class="tagListContainer">
          <el-tag 
            v-for="item in tagList"
            :key="item.ID"
            closable
            size="small"
            :color="item.color"
            @click="clickTag(item.name, item.color, item.ID)"
            @close="deleteTag(item.ID)"
          >
          {{item.name}}
          </el-tag>
        </div>
        
      </el-card>
    </el-col>
  </div>
</template>

<script>
export default {
  data() {
    return {
      tagInfo:{
        ID: 1,
        name: '',
        color: '#CFD3D7'
      },
      // 通过ID判断是新的标签（添加标签），还是通过点击右侧的标签取到的（修改标签）
      ID: 0,  
      tagList:[],
      // 预定义颜色
      predefineColors:[
        '#ff4500',
        '#ff8c00',
        '#ffd700',
        '#90ee90',
        '#00ced1',
        '#1e90ff',
        '#c71585',
        'rgba(255, 69, 0, 0.68)',
        'rgb(255, 120, 0)',
        'hsv(51, 100, 98)',
        'hsva(120, 40, 94, 0.5)',
        'hsl(181, 100%, 37%)',
        'hsla(209, 100%, 56%, 0.73)',
        '#c7158577'
      ],
      formRules:{
        name:[
          {required: true, message: '请输入标签名', trigger: 'blur'},
          {min: 1, max: 20, message: '长度为 1 到 20 个字符', trigger: 'blur'}
        ],
        color:[{required: true, message: '请选择或输入标签颜色', trigger: 'blur'}]
      }
    }
  },
  created(){
    this.getTagList()
  },
  methods:{
    async getTagList(){
      const {data:res} = await this.$http.get('tags')
      if(res.status !== 200){
        return this.$$message.error('获取标签列表失败')
      }
      this.tagList = res.data
    },
    // 添加新的标签
    saveTag() {
    // post请求携带参数，传入的是一个对象
    // 需要验证
      this.$refs.tagInfoRef.validate(async(valid)=>{
        if(!valid) return
        const {data:res} = await this.$http.post('tag/add',{
          name: this.tagInfo.name,
          color: this.tagInfo.color
        })
        if(res.status !== 200) {
          return this.$message.error(res.message)
        }
        this.$message.success('标签添加成功')
        this.getTagList()
      })
    },
    // 修改标签
    async updateTag(){
        this.$refs.tagInfoRef.validate(async(valid)=>{
        if(!valid) return
        const {data:res} = await this.$http.put(`tag/${this.ID}`,{
          name : this.tagInfo.name,
          color: this.tagInfo.color        
        })
        if(res.status !== 200){
          return this.$message.error(res.message)
        }
        this.$message.success('标签修改成功')
        this.getTagList()
      })
    },
    // 点击标签
    clickTag(name, color, ID) {
      this.tagInfo.name = name
      this.tagInfo.color = color
      this.ID = ID
    },
    // 删除标签
    async deleteTag(id){
      const confirmRes = await this.$confirm('此操作将永久删除该标签, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
        }).catch(err => err)
      if ('cancel' === confirmRes) {// 用户点击了取消
        return this.$message.success('取消删除')
      }
      if ('confirm' === confirmRes) {
        const {data:res} = await this.$http.delete(`tag/${id}`)
        if(res.status !== 200){
          return this.$message.error(res.message)
        }
        this.$message.success('标签删除成功')
        this.getTagList()
      }
    },
    backToAdd(){
      // 要使resetField生效，还要定义prop
      this.$refs.tagInfoRef.resetFields() 
      this.ID = 0
    }
  }
}
</script>

<style scoped>
.tagListContainer{
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
}
.card-head-title{
  color: rgba(0,0,0,.85);
  font-weight: 500;
  font-size: 16px;
}

.el-divider{
  margin-top: 15px;
}

.el-tag{
  color: #000;
  font-size: 12px;
  margin: 5px;
  /* border: 1px dashed gray; */
  border: 0;
}

/* TODO:点击样式有点问题 */

.backToAddBtn{
  background-color: #fff;
  color: black;
  border: 1px dashed gray;
}

.backToAddBtn:hover{
  background-color: #fff;
  border: 1px dashed red;
  color: orangered;
}

</style>