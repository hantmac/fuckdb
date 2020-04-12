<template>
  <div class="m-doc">
    <HeaderNav @tabName="handleCommand"></HeaderNav>
    <div class="m-fuckDb" v-show="tabName == 'FuckDb'">
      <FuckDbPost @formData=GetDb2struct @cleanDb="cleanDb" :result="code"></FuckDbPost>
    </div>

    <div class="m-fuckDb" v-show="tabName == 'Dump'">
      <DumpPost @formData=GetDb2View  @cleanDb="cleanDb" :result="code" ></DumpPost>
    </div>

    <div class="m-tool" v-show="tabName == 'json-to-go'">
      <iframe src="https://mholt.github.io/json-to-go/" frameborder="0"></iframe>
    </div>
    <div class="m-tool" v-show="tabName == 'yaml-to-json'">
      <iframe src="https://onlineyamltools.com/convert-yaml-to-json" frameborder="0"></iframe>
    </div>
    <div class v-show="tabName == 'struct-to-sql'">
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
        :options="{
        enableBasicAutocompletion: true,
        enableSnippets: true,
        enableLiveAutocompletion: true,
        tabSize:2
    }"
        :fontSize="12"
        :lang="'golang'"
        :theme="'textmate'"
        @onChange="editorChange"
        @init="editorInit"
      ></editor>
      <highlight-code lang="sql">{{codeSql}}</highlight-code>
    </div>
  </div>
</template>

<script>
import service from "../config/api";
import API from "../config/index";
import Editor from "vue2x-ace-editor";
import HeaderNav from "./header-nav"
import FuckDbPost from "./fuckdb-post"
import DumpPost from "./dump-post"

export default {
  // inject:['reload'],
  name: "home",
  components: {
    Editor, HeaderNav,FuckDbPost,DumpPost
  },
  data() {
    return {
      code: "package test",
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
        \`registration_date\` varchar(128) NOT NULL ,

      ) engine=innodb DEFAULT charset=utf8mb4;
      `,
      form: {
        mysql_host: "",
        mysql_port: 3306,
        mysql_db: "",
        mysql_table: "",
        mysql_passwd: "",
        mysql_user: "",
        package_name: "",
        struct_name: "",
        json_annotation: "true",
        xml_annotation: "false",
        gorm_annotation: "true"
      },
      rules: {
        mysql_host: [
          { required: true, message: "mysql_host is empty", trigger: "blur" }
        ],
        mysql_db: [
          { required: true, message: "mysql_db is empty", trigger: "blur" }
        ],
        mysql_table: [
          { required: true, message: "mysql_table is empty", trigger: "blur" }
        ],
        mysql_passwd: [
          { required: true, message: "mysql_passwd is empty", trigger: "blur" }
        ],
        mysql_user: [
          { required: true, message: "mysql_user is empty", trigger: "blur" }
        ],
        package_name: [
          { required: true, message: "package_name is empty", trigger: "blur" }
        ],
        struct_name: [
          { required: true, message: "struct_name is empty", trigger: "blur" }
        ],
        mysql_port: [
          {
            required: true,
            type: "number",
            message: "port must be integer",
            trigger: "blur"
          }
        ]
      },
      optionsDBLog: [],
      FuckDbList: [],
      value: ""
    };
  },
  mounted() {
    this.setFuckDbChangeDBList();
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
          name: "test", //名称
          value: "test", //值
          caption: "test", //字幕，展示在列表的内容
          meta: "test", //展示类型
          type: "local", //类型
          score: 1000 // 分数越大排在越上面
        }
      ]);
    },
    GetDb2View(form) {
      let data = {
        mysql_host: form.mysql_host,
        mysql_port: form.mysql_port,
        mysql_db: form.mysql_db,
        mysql_table: form.mysql_table,
        mysql_passwd: form.mysql_passwd,
        mysql_user: form.mysql_user,
        view_type:form.view_type
      };

      service({
        url: `${API.APIdb2struct}/api/db2view`,
        method: "post",
        data
      })
        .then(res => {
          if (res.data.status === "0") {
            this.code = res.data.data;
            this.setFuckDbList(form);
          }
        })
        .catch(error => {
          if (error.response.status == 500) {
            try {
              this.$message.error(error.response.data.error);
            } catch (error) {
              this.$message.error("异常错误！");
            }
          } else {
            this.$message.error("异常错误！");
          }
        });
    },
    GetDb2struct(form) {
      let data = {
        mysql_host: form.mysql_host,
        mysql_port: form.mysql_port,
        mysql_db: form.mysql_db,
        mysql_table: form.mysql_table,
        mysql_passwd: form.mysql_passwd,
        mysql_user: form.mysql_user,
        package_name: form.package_name,
        struct_name: form.struct_name,
        json_annotation: form.json_annotation === "true" ? true : false,
        xml_annotation: form.xml_annotation === "true" ? true : false,
        gorm_annotation: form.gorm_annotation === "true" ? true : false
      };

      service({
        url: `${API.APIdb2struct}/api/db2struct`,
        method: "post",
        data
      })
        .then(res => {
          if (res.data.status === "0") {
            this.code = res.data.data;
            this.setFuckDbList(form);
          }
        })
        .catch(error => {
          if (error.response.status == 500) {
            try {
              this.$message.error(error.response.data.error);
            } catch (error) {
              this.$message.error("异常错误！");
            }
          } else {
            this.$message.error("异常错误！");
          }
        });
    },
    setFuckDbList(data) {
      let FuckDbList = window.localStorage.getItem("FuckDb_List");
      if (FuckDbList) {
        FuckDbList = JSON.parse(FuckDbList);
        let FuckDbListDistinguish = false;

        FuckDbList.forEach(obj => {
          if (
            obj.mysql_user === data.mysql_user &&
            obj.mysql_host === data.mysql_host
          ) {
            FuckDbListDistinguish = true;
            obj.mysql_db = data.mysql_db;
            obj.mysql_table = data.mysql_table;
            obj.package_name = obj.package_name;
            obj.struct_name = obj.struct_name;
          }
        });

        if (!FuckDbListDistinguish) {
          FuckDbList.push(data);
        }
        window.localStorage.setItem("FuckDb_List", JSON.stringify(FuckDbList));
      } else {
        FuckDbList = [data];
        window.localStorage.setItem("FuckDb_List", JSON.stringify([data]));
      }
      this.optionsDBLog = [];
      FuckDbList.forEach((obj, index) => {
        this.optionsDBLog.push({
          value: index,
          label: `${obj.mysql_user}@${obj.mysql_host}`
        });
        if (
          obj.mysql_user === data.mysql_user &&
          obj.mysql_host === data.mysql_host
        ) {
          this.value = index;
        }
      });
      this.FuckDbList = FuckDbList;
    },
    setFuckDbChangeDBList() {
      let FuckDbList = window.localStorage.getItem("FuckDb_List");
      if (FuckDbList) {
        FuckDbList = JSON.parse(FuckDbList);
        this.FuckDbList = FuckDbList;
        this.optionsDBLog = [];
        FuckDbList.forEach((obj, index) => {
          this.optionsDBLog.push({
            value: index,
            label: `${obj.mysql_user}@${obj.mysql_host}`
          });
        });
      }
    },
    changeDBList(index) {
      if (index !== "") {
        this.form = this.FuckDbList[index];
      }
    },
    clearDBList() {
      this.code = "package test";
      this.form = {
        mysql_host: "",
        mysql_port: 3306,
        mysql_db: "",
        mysql_table: "",
        mysql_passwd: "",
        mysql_user: "",
        package_name: "",
        struct_name: "",
        json_annotation: "true",
        xml_annotation: "true",
        gorm_annotation: "true"
      };
    },
    cleanDb() {
      this.value = "";
      this.FuckDbList = [];
      this.optionsDBLog = [];
      this.code = "package test";
      this.form = {
        mysql_host: "",
        mysql_port: 3306,
        mysql_db: "",
        mysql_table: "",
        mysql_passwd: "",
        mysql_user: "",
        package_name: "",
        struct_name: "",
        json_annotation: "true",
        xml_annotation: "true",
        gorm_annotation: "true"
      };

      window.localStorage.setItem("FuckDb_List", "");
    },

    resetForm(formName) {
      this.value = "";
      this.code = "package test";
      this.$refs[formName].resetFields();
      this.setFuckDbChangeDBList();
    },
    handleCommand(command) {
      if (command === "FuckDb" || command === "Dump") {
        // this.reload()
        document.getElementById("app").style.overflowY = "auto";
      } else {
        document.getElementById("app").style.overflowY = "hidden";
      }
      this.tabName = command;
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scss>
h1,
h2 {
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
  color: #409eff;
}
.m-tool,
.m-tool iframe {
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
