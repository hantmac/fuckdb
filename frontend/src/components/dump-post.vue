<template>
    <div class="m-fuckDb" v-show="tabName == 'dump'">
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
              <select v-model="form.view_type">
                <option disabled value="">请选择</option>
                <option value="md">markdown</option>
                <option value="json">json</option>
                <option value="txt">txt</option>
                <option value="csv">csv</option>
              </select>
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
        <highlight-code lang="golang">{{code}}</highlight-code>
      </div>
    </div>
</template>