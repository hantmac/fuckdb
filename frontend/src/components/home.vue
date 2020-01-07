<template>
   <div class="m-doc">
      <div class="m-header">
        <div class="f-tit">
          <h3 class="f-fl"> FuckDb</h3>
          <el-dropdown class="f-fl m-dropdown" @command="handleCommand">
            <span class="el-dropdown-link">Navigation<i class="el-icon-arrow-down el-icon--right"></i>
            </span>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item command="FuckDb">FuckDb</el-dropdown-item>
              <el-dropdown-item command="json-to-go">json-to-go</el-dropdown-item>
              <!-- <el-dropdown-item command="struct-to-sql">struct-to-sql</el-dropdown-item> -->
            </el-dropdown-menu>
        </el-dropdown>
         <iframe
    style="margin-left: 15px; margin-bottom:-7px;"
    frameborder="0" scrolling="0" width="91px" height="20px"
    src="https://ghbtns.com/github-btn.html?user=hantmac&repo=fuckdb&type=star&count=true" >
</iframe>
        </div>
      </div>
      <div class="m-fuckDb" v-show="tabName == 'FuckDb'">
<el-form ref="form" :model="form" :rules="rules" label-width="130px" size="medium">
  <el-row>
    <el-col :span="24">
      <el-form-item label="Change DB">
         <el-select v-model="value" placeholder="select DB" @change="changeDBList" clearable @clear="clearDBList">
          <el-option
            v-for="item in optionsDBLog"
            :key="item.value"
            :label="item.label"
            :value="item.value"
            >
          </el-option>
        </el-select>
            <el-button @click="cleanDb">Clear</el-button>
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
   <div class="f-fl">
      <el-button type="primary" @click="onSubmit('form')">Create</el-button>
    <el-button  @click="resetForm('form')">Reset</el-button>
   </div>
    <div class="u-tips">
      <p>
        <b>声明：</b> 本工具不会记录任何用户信息，请放心使用
      </p>
      <p>
        <b>Disclaimer:</b> This tool does not record any user information, please use it with confidence
      </p>
    </div>
  </el-form-item>
</el-form>
    <div class="m-body">
      <highlight-code lang="golang">{{code}}</highlight-code>
    </div>
    </div>
   <div class="m-tool" v-show="tabName == 'json-to-go'">
      <iframe src="https://mholt.github.io/json-to-go/" frameborder="0"></iframe>
    </div>
    <div class="" v-show="tabName == 'struct-to-sql'">
      <el-tabs type="border-card">
  <el-tab-pane label="Create"></el-tab-pane>
  <el-tab-pane label="Clean"></el-tab-pane>
  <el-tab-pane label="Copy"></el-tab-pane>
</el-tabs>
     <editor
    height="400px"
    ref="editor"
    :content="content"
    :options="{
        enableBasicAutocompletion: true,
        enableSnippets: true,
        enableLiveAutocompletion: true,
        tabSize:2
    }"
    :fontSize='14'
    :lang="'golang'"
    :theme="'eclipse'"
    @onChange="editorChange"
    @init="editorInit">
    </editor>
     <highlight-code lang="sql">{{codeSql}}</highlight-code>
    </div>
  </div>
</template>

<script>
import service from "../config/api"
import API from "../config/index"
import Editor from 'vue2x-ace-editor'

export default {
  name: 'home',
  components:{
        Editor
  },
  data () {
    return {
      code: "package test",
      content: `
//code...
type testDB struct {
    Status        string \`gorm:"column:Status" json:"Status" xml:"Status"\`
}
`,
      tabName:  "FuckDb", //"struct-to-sql"
      codeSql: `
      CREATE TABLE \`auto_generated\`
      (
        \`registration_date\` varchar(128) NOT NULL ,

      ) engine=innodb DEFAULT charset=utf8mb4;
      `,
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
             { required: true, message: 'mysql_host is empty' , trigger: 'blur'}
          ],
          mysql_db: [
             { required: true, message: 'mysql_db is empty' , trigger: 'blur'}
          ],
          mysql_table: [
            { required: true, message: 'mysql_table is empty' , trigger: 'blur'}
          ],
          mysql_passwd: [
            { required: true, message: 'mysql_passwd is empty' , trigger: 'blur'}
          ],
          mysql_user: [
            { required: true, message: 'mysql_user is empty' , trigger: 'blur'}
          ],
          package_name: [
            { required: true, message: 'package_name is empty' , trigger: 'blur'}
          ],
          struct_name: [
            { required: true, message: 'struct_name is empty' , trigger: 'blur'}
          ],
          mysql_port: [
            {required: true, type: 'number', message: 'port must be integer' , trigger: 'blur'}
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
    editorInit(){
      require("brace/ext/language_tools");
      require(`brace/mode/golang`);
      require(`brace/snippets/golang`);
      require(`brace/theme/eclipse`);
    },
    editorChange(){
      this.$refs.editor.setCompleteData([
        {
          name: "test",//名称
          value: "test",//值
          caption: "test",//字幕，展示在列表的内容
          meta: "test",//展示类型
          type: "local",//类型
          score: 1000 // 分数越大排在越上面
        }
      ]);
    },
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
              obj.mysql_db = data.mysql_db
              obj.mysql_table = data.mysql_table
              obj.package_name = obj.package_name
              obj.struct_name = obj.struct_name

            }
          })

          if(!FuckDbListDistinguish){
            FuckDbList.push(data)
          }
          window.localStorage.setItem("FuckDb_List", JSON.stringify(FuckDbList))
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
      this.code = "package test"
      this.$refs[formName].resetFields();
      this.setFuckDbChangeDBList()
    },
    handleCommand(command){
      if(command === "FuckDb"){
        document.getElementById("app").style.overflowY="auto"
      }else{
        document.getElementById("app").style.overflowY="hidden"
      }
      this.tabName = command;
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

.m-header{
  overflow: hidden;
  padding: 10px 0;
}
.m-tool{
  position: relative;
}
.f-tit{
  margin-left: 10px;
}
.f-fl{
  float: left;
}
.m-dropdown{
  cursor: pointer;
  margin: 5px 0 0 15px;
  cursor: pointer;
  color: #409EFF;
}
.m-tool, .m-tool iframe{
  width: 100%;
  height: 100%;
}
.m-etitor-tool{
  height: 39px;
}
.u-tips{
  padding-left: 20px;
  color: #aaa;
  font-size: 12px;
  float: left;
  line-height: 18px;
}

</style>
