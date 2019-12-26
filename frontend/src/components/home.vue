<template>
   <div class="m-doc">
      <div class="m-header">
        <div class="f-tit">
          <h3>FuckDb</h3>
        </div>
<el-form ref="form" :model="form" :rules="rules" label-width="130px" size="medium">
  <el-row>
    <el-col :span="24">
      <el-form-item label="Change DB">
         <el-select v-model="value" placeholder="请选择" @change="changeDBList" clearable @clear="clearDBList">
          <el-option
            v-for="item in optionsDBLog"
            :key="item.value"
            :label="item.label"
            :value="item.value"
            >
          </el-option>
        </el-select>
            <el-button @click="cleanDb">清除</el-button>
      </el-form-item>
    </el-col>
  </el-row>
  <el-row>
    <el-col :span="8">
      <el-form-item prop="mysql_host"  label="mysql_host"> <el-input v-model="form.mysql_host" clearable></el-input></el-form-item>
    </el-col>

    <el-col :span="8">
     <el-form-item prop="mysql_port" label="mysql_port"><el-input v-model.number="form.mysql_port" clearable></el-input></el-form-item>
    </el-col>

     <el-col :span="8">
     <el-form-item prop="mysql_db" label="mysql_db"><el-input v-model="form.mysql_db" clearable></el-input></el-form-item>
    </el-col>
   </el-row>

   <el-row>
    <el-col :span="8">
      <el-form-item prop="mysql_table" label="mysql_table"> <el-input v-model="form.mysql_table" clearable></el-input></el-form-item>
    </el-col>

    <el-col :span="8">
     <el-form-item prop="mysql_passwd"  label="mysql_passwd"><el-input type="password" v-model="form.mysql_passwd" clearable></el-input></el-form-item>
    </el-col>

     <el-col :span="8">
     <el-form-item prop="mysql_user" label="mysql_user"><el-input v-model="form.mysql_user" clearable></el-input></el-form-item>
    </el-col>
   </el-row>

    <el-row>
    <el-col :span="8">
      <el-form-item prop="package_name"  label="package_name"> <el-input v-model="form.package_name" clearable></el-input></el-form-item>
    </el-col>

    <el-col :span="8">
     <el-form-item prop="struct_name" label="struct_name"><el-input v-model="form.struct_name" clearable></el-input></el-form-item>
    </el-col>

     <el-col :span="8">
     <el-form-item label="json_annotation">
      <el-radio-group v-model="form.json_annotation">
        <el-radio label="true"></el-radio>
        <el-radio label="false"></el-radio>
      </el-radio-group>
     </el-form-item>

    </el-col>
   </el-row>

<el-row>
    <el-col :span="8">
      <el-form-item label="xml_annotation">
           <el-radio-group v-model="form.xml_annotation">
            <el-radio label="true"></el-radio>
            <el-radio label="false"></el-radio>
          </el-radio-group>
      </el-form-item>
    </el-col>

    <el-col :span="8">
     <el-form-item label="gorm_annotation">
        <el-radio-group v-model="form.gorm_annotation">
            <el-radio label="true"></el-radio>
            <el-radio label="false"></el-radio>
          </el-radio-group>
     </el-form-item>
    </el-col>
   </el-row>

  <el-form-item>
    <el-button type="primary" @click="onSubmit('form')">生成</el-button>
    <el-button  @click="resetForm('form')">重置</el-button>
  </el-form-item>
</el-form>
      </div>
    <div class="m-body">
      <highlight-code lang="golang">{{code}}</highlight-code>
    </div>
    <div class="m-tool">
      <div class="u-shelter"></div>
    </div>
  </div>
</template>

<script>
import service from "../config/api"
import API from "../config/index"
export default {
  name: 'home',
  data () {
    return {
      code: "package test",
      form: {
          mysql_host: '',
          mysql_port: 3306,
          mysql_db: '',
          mysql_table: '',
          mysql_passwd: '',
          mysql_user: '',
          package_name: '',
          struct_name: '',
          json_annotation: "true",
          xml_annotation: "false",
          gorm_annotation: "true"
        },
        rules:{
          mysql_host: [
             { required: true, message: 'mysql_host不能为空' , trigger: 'blur'}
          ],
          mysql_db: [
             { required: true, message: 'mysql_db不能为空' , trigger: 'blur'}
          ],
          mysql_table: [
            { required: true, message: 'mysql_table不能为空' , trigger: 'blur'}
          ],
          mysql_passwd: [
            { required: true, message: 'mysql_passwd不能为空' , trigger: 'blur'}
          ],
          mysql_user: [
            { required: true, message: 'mysql_user不能为空' , trigger: 'blur'}
          ],
          package_name: [
            { required: true, message: 'package_name不能为空' , trigger: 'blur'}
          ],
          struct_name: [
            { required: true, message: 'struct_name不能为空' , trigger: 'blur'}
          ],
          mysql_port: [
            {type: 'number', message: 'port必须为整数字' , trigger: 'blur'}
          ]
        },
        optionsDBLog: [],
        FuckDbList: [],
        value: ''
    }
  },
  mounted(){
    this.setFuckDbChangeDBList()
  },
  methods: {
    GetDb2struct(){
      let data = {
          "mysql_host": this.form.mysql_host,
          "mysql_port": this.form.mysql_port,
          "mysql_db": this.form.mysql_db,
          "mysql_table": this.form.mysql_table,
          "mysql_passwd": this.form.mysql_passwd,
          "mysql_user": this.form.mysql_user,
          "package_name": this.form.package_name,
          "struct_name": this.form.struct_name,
          "json_annotation": this.form.json_annotation === "true" ? true : false,
          "xml_annotation": this.form.xml_annotation === "true" ? true : false,
          "gorm_annotation":this.form.gorm_annotation === "true" ? true : false
      }

      service({
        url:  `${API.APIdb2struct}/api/db2struct`,
        method: "post",
        data,
      }).then(res => {
        if(res.data.status === "0"){
          this.code = res.data.data;
          this.setFuckDbList(this.form)
        }
      }).catch(error =>{
        if(error.response.status == 500){
          try {
            this.$message.error(error.response.data.error);
          } catch (error) {
            this.$message.error("异常错误！");
          }
        }else{
          this.$message.error("异常错误！");
        }

      })
    },
    setFuckDbList(data){
      let FuckDbList = window.localStorage.getItem("FuckDb_List")
      if(FuckDbList){
          FuckDbList = JSON.parse(FuckDbList)
          let FuckDbListDistinguish = false

          FuckDbList.forEach(obj => {
            if(obj.mysql_user === data.mysql_user && obj.mysql_host === data.mysql_host){
              FuckDbListDistinguish = true
            }
          })

          if(!FuckDbListDistinguish){
            FuckDbList.push(data)
            window.localStorage.setItem("FuckDb_List", JSON.stringify(FuckDbList))
          }
      }else{
        FuckDbList = [data]
        window.localStorage.setItem("FuckDb_List", JSON.stringify([data]))
      }
      this.optionsDBLog = []
      FuckDbList.forEach((obj, index) =>{
        this.optionsDBLog.push({
          value: index,
          label: `${obj.mysql_user}@${obj.mysql_host}`
        })
         if(obj.mysql_user === data.mysql_user && obj.mysql_host === data.mysql_host){
           this.value = index
          }
      })
      this.FuckDbList = FuckDbList
    },
    setFuckDbChangeDBList(){
      let FuckDbList = window.localStorage.getItem("FuckDb_List")
      if(FuckDbList){
        FuckDbList = JSON.parse(FuckDbList)
        this.FuckDbList = FuckDbList
        this.optionsDBLog = []
        FuckDbList.forEach((obj, index) =>{
          this.optionsDBLog.push({
            value: index,
            label: `${obj.mysql_user}@${obj.mysql_host}`
          })
        })
      }
    },
    changeDBList(index){
      if(index !== ""){
        this.form = this.FuckDbList[index]
      }
    },
    clearDBList(){
      this.code = "package test"
      this.form = {
          "mysql_host": "",
          "mysql_port": 3306,
          "mysql_db": "",
          "mysql_table": "",
          "mysql_passwd": "",
          "mysql_user": "",
          "package_name": "",
          "struct_name": "",
          "json_annotation": "true",
          "xml_annotation":  "true",
          "gorm_annotation": "true"
      }
    },
    cleanDb(){
      this.value = ""
      this.FuckDbList = []
      this.optionsDBLog = []
      this.code = "package test"
      this.form = {
          "mysql_host": "",
          "mysql_port": 3306,
          "mysql_db": "",
          "mysql_table": "",
          "mysql_passwd": "",
          "mysql_user": "",
          "package_name": "",
          "struct_name": "",
          "json_annotation": "true",
          "xml_annotation":  "true",
          "gorm_annotation": "true"
      }

      window.localStorage.setItem("FuckDb_List", "")
    },
    onSubmit(formName){
      this.$refs[formName].validate((valid) => {
          if (valid) {
            this.GetDb2struct()
          } else {
            console.log('error submit!!');
            return false;
          }
        });
    },
    resetForm(formName){
      this.value = ""
      this.$refs[formName].resetFields();
      this.setFuckDbChangeDBList()
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped scss>

h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}

.m-doc{
  position: relative;
  height: 100%;
}

.m-body{
}

.m-header{
  margin-top: 10px;
}
.m-tool{
  position: relative;
}
.u-shelter{
    position: absolute;
    width: 100%;
    height: 36px;
    background: #ffff;
}
.u-shelter-1{
  position: absolute;
  right: 0;
  bottom: 0;
  width: 20px;
  height: 30px;
  background: #000;
}
#ToolIframe{
  display: none;
}
.f-tit{
  margin-left: 10px;
}
</style>
