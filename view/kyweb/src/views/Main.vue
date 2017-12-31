<template>
  <el-container>
    <el-header class="fixed-header">
      <el-row type="flex" class="row-bg" style="height: inherit">
      <el-col :span="4">
        <span  class="minename">{{MineName}}</span>
      </el-col>
        <el-col :span="17">
          <el-menu class="el-menu-demo" mode="horizontal" :router="true" :default-active="CurrentPath" @select="selectMenu">
              <el-menu-item index="/WorkPlace" >系统</el-menu-item>
            <el-submenu index="2">
              <template slot="title">配置</template>
              <el-menu-item index="/WorkPlace" >工作面配置</el-menu-item>
              <el-menu-item index="/Config1" >配置2</el-menu-item>
            </el-submenu>
            <el-menu-item index="/Demo" >Demo</el-menu-item>
          </el-menu>
        </el-col>
        <el-col :span="3">
          <div class="user">
            <a href="javascript:void(0)" @click="logout">退出</a>
            <span>{{UserName}}</span>
          </div>
        </el-col>
      </el-row>
    </el-header>
    <el-container>
      <el-aside width="30px">

      </el-aside>
      <el-main>
        <router-view/>
      </el-main>
    </el-container>

    <el-footer  class="fixed-footer">

    </el-footer>
  </el-container>
</template>
<script>
  import {GetStore, SetStore} from '../vuex/store'
  import ElFooter from '../../node_modules/element-ui/packages/footer/src/main.vue'
  import ElContainer from '../../node_modules/element-ui/packages/container/src/main.vue'
  import ElAside from '../../node_modules/element-ui/packages/aside/src/main.vue'
  export default {
    components: {
      ElAside,
      ElContainer,
      ElFooter},
    data () {
      return {
        MineName: '量化投资',
        UserName: 'admin',
        CurrentPath: ''
      }
    },
    methods: {
      logout () {
        SetStore('Token', 0)
        SetStore('UserName', '')
        SetStore('MineName', '')
        this.$router.push('/Login')
      },
      selectMenu (index) {
        SetStore('CurrentPath', index)
      }
    },
    created () {
      this.UserName = GetStore('UserName')
      this.MineName = GetStore('MineName')
      this.CurrentPath = GetStore('CurrentPath')
      if (this.CurrentPath === '') {
        this.CurrentPath = '/Demo'
      }
      this.defaultActive = this.CurrentPath
    }
  }
</script>

<style>
  .minename{
    text-align: center;
    vertical-align: middle;
    line-height:60px;
    font-size: 28px;
    color: #FFFFFF;
  }

  .user span, a{
    vertical-align: middle;
    line-height: 60px;
    text-align: right;
    font-size: 16px;
    color: #FFFFFF;
  }

  .user a {
    margin-left: 25px;
    float: right;
  }

  .user span {
    float: right;
  }







  .el-menu {
    background-color: rgba(19, 116, 157, 1);
  }
.el-menu--horizontal {
  float: left;
}

  .el-menu--horizontal .el-submenu .el-menu-item{
    background-color: rgba(19, 116, 157, 1);
    font-size: 16px;
    padding:0;
  }

  .el-menu--horizontal .el-menu-item {
  font-size: 16px;
  color: #ffffff;
}



.el-menu--horizontal>.el-menu-item.is-active, .el-menu--horizontal>.el-submenu .el-submenu__title.is-active {
  background-color: #0E5472;
  color: #ffffff;
}
  .el-menu--horizontal .el-submenu .el-submenu__title:focus, .el-menu--horizontal .el-submenu .el-submenu__title:hover {
    background-color: #0E5472;
    color: #ffffff;
  }

.el-menu--horizontal .el-menu-item:focus, .el-menu--horizontal .el-menu-item:hover {
  background-color: #0E5472;
  color: #ffffff;
}
  .el-menu--horizontal .el-submenu .el-submenu__title{
    color: #ffffff;
  }

</style>
