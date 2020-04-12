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
            <el-form-item prop="package_name" label="package_name">
              <el-input v-model="form.package_name" clearable></el-input>
            </el-form-item>
          </el-col>

          <el-col :span="8">
            <el-form-item prop="struct_name" label="struct_name">
              <el-input v-model="form.struct_name" clearable></el-input>
            </el-form-item>
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
  name: "fuckdb-post",
  props:['result'],
  components: {

  },
  data() {
    return {
      // code: "package test",
      form: {
        mysql_host: "127.0.0.1",
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
      // this.code = "package test";
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
      this.$refs[formName].resetFields();
      this.setFuckDbChangeDBList();
    },
  }
}
</script>