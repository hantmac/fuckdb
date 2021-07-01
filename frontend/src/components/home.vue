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
            <el-dropdown-item command="yaml-to-json">yaml-to-json</el-dropdown-item>
            <!-- <el-dropdown-item command="struct-to-sql">struct-to-sql</el-dropdown-item> -->
          </el-dropdown-menu>
        </el-dropdown>
        <iframe
          style="margin-left: 15px; margin-bottom:-7px;"
          frameborder="0" scrolling="0" width="91px" height="20px"
          src="https://ghbtns.com/github-btn.html?user=hantmac&repo=fuckdb&type=star&count=true">
        </iframe>
      </div>
    </div>
    <div class="m-fuckDb" v-show="tabName === 'FuckDb'">
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
              <el-button @click="addConnection" type="primary">Add</el-button>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="8">
            <el-form-item label="table" required>
              <el-select v-model="form.mysql_table">
                <el-option
                  v-for="item in tableList"
                  :key="item"
                  :label="item"
                  :value="item"
                ></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="8">
            <el-form-item prop="package_name" label="package_name">
              <el-input v-model="form.package_name" placeholder="models" clearable></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="8">
            <el-form-item prop="struct_name" label="struct_name">
              <el-input v-model="form.struct_name" clearable></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="annotation">
              <el-checkbox-group v-model="form.annotation_list">
                <el-checkbox label="gorm"></el-checkbox>
                <el-checkbox label="xorm"></el-checkbox>
                <el-checkbox label="json"></el-checkbox>
                <el-checkbox label="xml"></el-checkbox>
                <el-checkbox label="db"></el-checkbox>
              </el-checkbox-group>
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item>
          <div class="f-fl">
            <el-button type="primary" @click="onSubmit('form')">Create</el-button>
            <el-button @click="resetForm('form')">Reset</el-button>
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
        <highlight-code lang="golang">{{ code }}</highlight-code>
      </div>
    </div>
    <div class="m-tool" v-show="tabName == 'json-to-go'">
      <iframe src="https://mholt.github.io/json-to-go/" frameborder="0"></iframe>
    </div>
    <div class="m-tool" v-show="tabName == 'yaml-to-json'">
      <iframe src="https://onlineyamltools.com/convert-yaml-to-json" frameborder="0"></iframe>
    </div>
    <div class="" v-show="tabName == 'struct-to-sql'">
      <div class="m-tabs">
        <el-button-group>
          <el-button size="medium" icon="el-icon-plus">Create</el-button>
          <el-button size="medium" icon="el-icon-delete">Clean</el-button>
          <el-button size="medium" icon="el-icon-copy-document">Copy</el-button>
        </el-button-group>
      </div>
      <editor
        class="m-editor"
        height="400px"
        width="100%"
        ref="editor"
        :content="content"
        :options="{enableBasicAutocompletion: true,enableSnippets: true,enableLiveAutocompletion: true,tabSize:2}"
        :fontSize='12'
        :lang="'golang'"
        :theme="'textmate'"
        @onChange="editorChange"
        @init="editorInit">
      </editor>
      <highlight-code lang="sql">{{ codeSql }}</highlight-code>
    </div>

    <el-dialog title="add new database" :visible.sync="addDBDialogVisible">
      <el-form :model="newConnectionForm" :rules="rules" ref="connForm">
        <el-form-item label="host" label-width="100px" required prop="mysql_host">
          <el-input v-model="newConnectionForm.mysql_host"></el-input>
        </el-form-item>
        <el-form-item label="port" label-width="100px" required prop="mysql_port">
          <el-input v-model.number="newConnectionForm.mysql_port" type="number"></el-input>
        </el-form-item>
        <el-form-item label="user" label-width="100px" required prop="mysql_user">
          <el-input v-model="newConnectionForm.mysql_user"></el-input>
        </el-form-item>
        <el-form-item label="password" label-width="100px" required prop="mysql_passwd">
          <el-input type="password" v-model="newConnectionForm.mysql_passwd"></el-input>
        </el-form-item>
        <el-form-item label="dbname" label-width="100px" required prop="mysql_db">
          <el-input v-model="newConnectionForm.mysql_db"></el-input>
        </el-form-item>
        <el-form-item label-width="100px">
          <el-button type="primary" @click="saveConnection">Add</el-button>
          <el-button @click="detectConnection">test connection</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script>
import service from "../config/api"
import Editor from 'vue2x-ace-editor'

export default {
  name: 'home',
  components: {
    Editor
  },
  data() {
    return {
      code: "package models",
      addDBDialogVisible: false,
      annotationList: [],
      tableList: [],
      content: `
//code...
type testDB struct {
    Status        string \`gorm:"column:Status" json:"Status" xml:"Status"\`
}
`,
      tabName: "FuckDb", //"struct-to-sql"
      codeSql: `
        CREATE TABLE \`auto_generated\`
        (
          \`registration_date\` varchar(128) NOT NULL,

        ) engine=innodb DEFAULT charset=utf8mb4;
      `,
      form: {
        mysql_host: '',
        mysql_port: 3306,
        mysql_db: '',
        mysql_table: '',
        mysql_passwd: '',
        mysql_user: 'root',
        package_name: '',
        struct_name: '',
        annotation_list: ['gorm', 'json'],
      },
      newConnectionForm: {
        mysql_host: '',
        mysql_port: 3306,
        mysql_user: 'root',
        mysql_passwd: '',
        mysql_db: '',
      },
      rules: {
        mysql_host: [
          {required: true, message: 'mysql_host is empty', trigger: 'blur'}
        ],
        mysql_db: [
          {required: true, message: 'mysql_db is empty', trigger: 'blur'}
        ],
        mysql_passwd: [
          {required: true, message: 'mysql_passwd is empty', trigger: 'blur'}
        ],
        mysql_user: [
          {required: true, message: 'mysql_user is empty', trigger: 'blur'}
        ],
        mysql_port: [
          {required: true, type: 'number', message: 'port must be integer', trigger: 'blur'}
        ]
      },
      optionsDBLog: [],
      FuckDbList: [],
      value: ''
    }
  },
  mounted() {
    this.setFuckDbChangeDBList()
  },
  methods: {
    editorInit() {
      require("brace/ext/language_tools");
      require(`brace/mode/golang`);
      require(`brace/snippets/golang`);
      require(`brace/theme/textmate`);
    },
    editorChange() {
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
    saveConnection() {
      this.$refs['connForm'].validate(valid => {
        if (valid) {
          this.setFuckDbList(this.newConnectionForm)
          this.addDBDialogVisible = false
          return
        }
        return false
      })
    },
    detectConnection() {
      this.$refs['connForm'].validate(valid => {
        if (valid) {
          let data = this.newConnectionForm
          const vueInstance = this
          service({
            url: "/api/testconn",
            method: "post",
            data,
          }).then(res => {
            if (res.data.status === "0") {
              vueInstance.$message.success("connect to database success!")
            } else {
              vueInstance.$message.error("connect to db error: " + res.data.error)
            }
          }).catch(err => {
            vueInstance.$message.error(err)
          })
        }
      })
    },
    GetDb2struct() {
      console.log(this.form)
      let data = {
        "mysql_host": this.form.mysql_host,
        "mysql_port": this.form.mysql_port,
        "mysql_db": this.form.mysql_db,
        "mysql_table": this.form.mysql_table,
        "mysql_passwd": this.form.mysql_passwd,
        "mysql_user": this.form.mysql_user,
        "package_name": this.form.package_name,
        "struct_name": this.form.struct_name,
        "json_annotation": this.form.annotation_list.includes('json'),
        "db_annotation": this.form.annotation_list.includes('db'),
        "xml_annotation": this.form.annotation_list.includes('xml'),
        "gorm_annotation": this.form.annotation_list.includes('gorm'),
        "xorm_annotation": this.form.annotation_list.includes('xorm'),
      }

      service({
        url: `/api/db2struct`,
        method: "post",
        data,
      }).then(res => {
        if (res.data.status === "0") {
          this.code = res.data.data;
        }
      }).catch(error => {
        if (error.response.status === 500) {
          try {
            this.$message.error(error.response.data.error);
          } catch (error) {
            this.$message.error("异常错误！");
          }
        } else {
          this.$message.error("异常错误！");
        }
      })
    },
    setFuckDbList(data) {
      console.log("new connnection:", data)
      let FuckDbList = window.localStorage.getItem("FuckDb_List")
      if (FuckDbList) {
        FuckDbList = JSON.parse(FuckDbList)

        const exist = FuckDbList.some(item => item.mysql_user === data.mysql_user
          && item.mysql_host === data.mysql_host)

        if (!exist) {
          FuckDbList.push(data)
        }
        window.localStorage.setItem("FuckDb_List", JSON.stringify(FuckDbList))
      } else {
        FuckDbList = [data]
        window.localStorage.setItem("FuckDb_List", JSON.stringify([data]))
      }
      this.optionsDBLog = []
      FuckDbList.forEach((obj, index) => {
        this.optionsDBLog.push({
          value: index,
          label: `${obj.mysql_user}@${obj.mysql_host}`
        })
        if (obj.mysql_user === data.mysql_user && obj.mysql_host === data.mysql_host) {
          this.value = index
        }
      })
      this.FuckDbList = FuckDbList
    },
    setFuckDbChangeDBList() {
      let FuckDbList = window.localStorage.getItem("FuckDb_List")
      if (FuckDbList) {
        FuckDbList = JSON.parse(FuckDbList)
        this.FuckDbList = FuckDbList
        this.optionsDBLog = []
        FuckDbList.forEach((obj, index) => {
          this.optionsDBLog.push({
            value: index,
            label: `${obj.mysql_user}@${obj.mysql_host}`
          })
        })
      }
    },
    async changeDBList(index) {
      if (index !== "") {
        const db = this.FuckDbList[index]
        this.form.mysql_host = db.mysql_host
        this.form.mysql_port = db.mysql_port
        this.form.mysql_db = db.mysql_db
        this.form.mysql_passwd = db.mysql_passwd
        this.form.mysql_user = db.mysql_user
        this.form.package_name = ""
        this.form.struct_name = ""
      }
      await this.fetchTableList()
    },
    async fetchTableList() {
      let data = {
        "mysql_host": this.form.mysql_host,
        "mysql_port": this.form.mysql_port,
        "mysql_db": this.form.mysql_db,
        "mysql_passwd": this.form.mysql_passwd,
        "mysql_user": this.form.mysql_user,
      }

      service({
        url: `/api/tables`,
        method: "post",
        data,
      }).then(res => {
        if (res.data.status === "0") {
          this.tableList = res.data.data.map(item => item.name)
        }
      }).catch(error => {
        if (error.response.status === 500) {
          try {
            this.$message.error(error.response.data.error);
          } catch (error) {
            this.$message.error("异常错误！");
          }
        } else {
          this.$message.error("异常错误！");
        }
      })

    },
    clearDBList() {
      this.code = "package models"
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
        "db_annotation": "true",
        "xml_annotation": "true",
        "gorm_annotation": "true"
      }
    },
    cleanDb() {
      this.value = ""
      this.FuckDbList = []
      this.optionsDBLog = []
      this.code = "package models"
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
        "db_annotation": "true",
        "xml_annotation": "true",
        "gorm_annotation": "true"
      }

      window.localStorage.setItem("FuckDb_List", "")
    },
    addConnection() {
      this.newConnectionForm = {
        mysql_host: '',
        mysql_port: 3306,
        mysql_user: 'root',
        mysql_passwd: '',
        mysql_db: '',
      }
      this.addDBDialogVisible = true
    },
    onSubmit(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.GetDb2struct()
        } else {
          console.log('error submit!!');
          return false;
        }
      });
    },
    resetForm(formName) {
      this.value = ""
      this.code = "package test"
      this.$refs[formName].resetFields();
      this.setFuckDbChangeDBList()
    },
    handleCommand(command) {
      if (command === "FuckDb") {
        document.getElementById("app").style.overflowY = "auto"
      } else {
        document.getElementById("app").style.overflowY = "hidden"
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

.m-doc {
  position: relative;
  height: 100%;
}

.m-header {
  overflow: hidden;
  padding: 10px 0;
}

.m-tool {
  position: relative;
}

.f-tit {
  margin-left: 10px;
}

.f-fl {
  float: left;
}

.m-dropdown {
  cursor: pointer;
  margin: 5px 0 0 15px;
  cursor: pointer;
  color: #409EFF;
}

.m-tool, .m-tool iframe {
  width: 100%;
  height: 100%;
}

.m-etitor-tool {
  height: 39px;
}

.u-tips {
  padding-left: 20px;
  color: #aaa;
  font-size: 12px;
  float: left;
  line-height: 18px;
}

.m-tabs {
  padding: 10px 0 10px 20px;
  background: #f5f7fa;
}

.m-editor {

}
</style>
