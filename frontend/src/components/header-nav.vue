<template>
    <div class="m-header">
      <div class="f-tit">
        <h3 class="f-fl">FuckDb</h3>
        <el-dropdown class="f-fl m-dropdown" @command="handleCommand">
          <span class="el-dropdown-link">
            Navigation
            <i class="el-icon-arrow-down el-icon--right"></i>
          </span>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item command="FuckDb" icon="el-icon-price-tag">
              <router-link :to="{path:'/'}">FuckDb</router-link>
            </el-dropdown-item>
            <el-dropdown-item command="Dump" icon="el-icon-s-tools">Dump</el-dropdown-item>
            <el-dropdown-item command="json-to-go" icon="el-icon-s-tools">json-to-go</el-dropdown-item>
            <el-dropdown-item command="yaml-to-json" icon="el-icon-s-tools">yaml-to-json</el-dropdown-item>
            <!-- <el-dropdown-item command="struct-to-sql">struct-to-sql</el-dropdown-item> -->
          </el-dropdown-menu>
        </el-dropdown>
        <iframe
          style="margin-left: 15px; margin-bottom:-7px;"
          frameborder="0"
          scrolling="0"
          width="91px"
          height="20px"
          src="https://ghbtns.com/github-btn.html?user=hantmac&repo=fuckdb&type=star&count=true"
        ></iframe>
      </div>
    </div>
</template>

<script>
export default {
    name:"header-nav",
    props:[''],
    methods: {
        handleCommand(command) {
            this.$emit('tabName',command)
        }
    },
}
</script>