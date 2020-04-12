<template>
  <div class="m-fuckDb">
      <el-form ref="form" :model="form" :rules="rules" label-width="130px" size="medium">
        <el-row>
          <el-col :span="24">
            <el-form-item label="Change DB">
              <el-select
                v-model="value"
                placeholder="select DB"
                @change="changeDBList"
                clearable
                @clear="clearDBList"
              >
                <el-option
                  v-for="item in optionsDBLog"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                ></el-option>
              </el-select>
              <el-button @click="cleanDb">Clear</el-button>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="8">
            <el-form-item prop="mysql_host" label="mysql_host">
              <el-input v-model="form.mysql_host" clearable></el-input>
            </el-form-item>
          </el-col>

          <el-col :span="8">
            <el-form-item prop="mysql_port" label="mysql_port">
              <el-input v-model.number="form.mysql_port" clearable></el-input>
            </el-form-item>
          </el-col>

          <el-col :span="8">
            <el-form-item prop="mysql_db" label="mysql_db">
              <el-input v-model="form.mysql_db" clearable></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row>
          <el-col :span="8">
            <el-form-item prop="mysql_table" label="mysql_table">
              <el-input v-model="form.mysql_table" clearable></el-input>
            </el-form-item>
          </el-col>

          <el-col :span="8">
            <el-form-item prop="mysql_passwd" label="mysql_passwd">
              <el-input type="password" v-model="form.mysql_passwd" clearable></el-input>
            </el-form-item>
          </el-col>

          <el-col :span="8">
            <el-form-item prop="mysql_user" label="mysql_user">
              <el-input v-model="form.mysql_user" clearable></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row>
          <el-col :span="8">
            <el-form-item label="请选择输出格式">
              <el-select v-model="form.view_type" placeholder="请选择">
                <el-option
                  v-for="item in viewOptions"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value">
                </el-option>
              </el-select>
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
        <highlight-code lang="golang">{{result}}</highlight-code>
      </div>
    </div>
</template>

<script>

export default {
  // inject:['reload'],
  name: "dump-post",
  props:['result'],
  components: {
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
        view_type:"md",
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
        view_type: [
          { required: true, message: "view_type is empty", trigger: "blur" }
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
      viewOptions: [
        {
          value: 'md',
          label: 'markdown'
        }, {
          value: 'json',
          label: 'json'
        }, {
          value: 'txt',
          label: 'txt'
        },{
          value: 'csv',
          label: 'csv'
        }
      ],
      value: ""
    };
  },
  mounted() {
    this.setFuckDbChangeDBList();
  },
  methods: {
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
        view_type:""
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
        view_type:""
      };
      this.$emit('cleanDb')
    },
    onSubmit(formName) {
      this.$refs[formName].validate(valid => {
        if (valid) {
          this.$emit('formData', this.form)
        } else {
          console.log("error submit!!");
          return false;
        }
      });
    },
    resetForm(formName) {
      this.value = "";
      this.code = "package test";
      this.$refs[formName].resetFields();
      this.setFuckDbChangeDBList();
    }
  }
};
</script>