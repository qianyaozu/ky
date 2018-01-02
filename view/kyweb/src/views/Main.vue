<template>
  <el-container>
    <el-header class="fixed-header">
      <el-row type="flex" class="row-bg" style="height: inherit">
      <el-col :span="4">
        <span  class="minename">{{MineName}}</span>
      </el-col>
        <el-col :span="17">
          <el-menu class="el-menu-demo" mode="horizontal" :router="true" :default-active="CurrentPath" @select="selectMenu">
            <el-submenu index="2">
              <template slot="title">配置</template>
              <el-menu-item index="/WorkPlace" >工作面</el-menu-item>
              <el-menu-item index="/FrameSet" >支架</el-menu-item>
              <el-menu-item index="/PillarSet">支柱</el-menu-item>
              <el-menu-item index="/DipSet">倾角</el-menu-item>
            </el-submenu>
            <el-submenu index="3">
              <template slot="title">实时</template>
              <el-menu-item index="/FrameReal" >支架</el-menu-item>
              <el-menu-item index="/PillarReal" >支柱</el-menu-item>
            </el-submenu>
            <el-submenu index="4">
              <template slot="title">历史</template>
              <el-menu-item index="/FrameHis">支架</el-menu-item>
              <el-menu-item index="/PillarHis" >支柱</el-menu-item>
            </el-submenu>
            <el-submenu index="5">
              <template slot="title">报表</template>
              <el-menu-item index="/FrameHis">支架</el-menu-item>
            </el-submenu>
            <el-submenu index="6">
              <template slot="title">工具</template>
              <el-menu-item index="/Com">串口通讯</el-menu-item>
              <el-menu-item index="/Adb">手机通讯</el-menu-item>
              <el-menu-item index="/Backup">备份</el-menu-item>
              <el-menu-item index="/Restore">还原</el-menu-item>
            </el-submenu>
            <el-submenu index="7">
              <template slot="title">帮助</template>
              <el-menu-item index="/Help">帮助</el-menu-item>
              <el-menu-item index="/About">关于</el-menu-item>
            </el-submenu>
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

    <el-footer  class="fixed-footer" >

    </el-footer>
  </el-container>
</template>
<script>
  import {GetStore, SetStore} from '../vuex/store'
  export default {
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
      },
      GetAlarm(){

      }
    },
    created () {
      this.UserName = GetStore('UserName')
      this.MineName = GetStore('MineName')
      this.CurrentPath = GetStore('CurrentPath')
      if (this.CurrentPath === '') {
        this.CurrentPath = '/Demo'
      }
      // window.setInterval(function(){
      //   this.GetAlarm()
      // }, 5000)
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
