<template>
  <div class="main" v-loading="loading">
   <el-row>
      <el-col :span="24">
          <div class="title" >量化投资登录</div>
      </el-col>
   </el-row>
    <el-row>
      <el-col :span="24">
        <el-input placeholder="用户名" v-model="username"></el-input>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="24">
        <el-input type="password" placeholder="密码" v-model="password"></el-input>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="24">
        <el-button class="loginbtn" type="primary" @click="login">登录</el-button>
      </el-col>
    </el-row>
  </div>
</template>

<script>
  import axios from 'axios'
  import {SetStore, GetStore} from '../vuex/store'
  export default {
    data () {
      return {
        username: '',
        password: '',
        loading: false
      }
    },
    methods: {
      login () {
        if (this.username === '' || this.password === '') {
          this.$message.error('请输入用户名或密码')
          return
        }
        this.loading = true
        axios.post('/api/login', {UserName: this.username, Password: this.password}).then(res => {
          this.loading = false
          if (!res.data) {
            this.$message.error('登录失败')
            return
          }
          if (res.data.Stats === 0) {
            this.$message.error(res.data.Messgae)
            return
          }
          SetStore('UserName', res.data.Data.UserName)
          SetStore('MineName', res.data.Data.MineName)
          SetStore('Token', res.data.Data.Token)
          this.$router.push('/Main')
        }).catch(err => {
          this.loading = false
          this.$message.error('登录失败')
          console.log(err)
        })
      }
    },
    created () {
      var token = GetStore('Token')
      if (token > 0) {
        this.$router.push('/Main')
      }
    }
  }
</script>

<style>
  body {
    margin: 0;
    padding: 0;
  }

  .el-row {
    margin-bottom: 20px;
  }
  .main {
    position: absolute;
    top: 50%;
    left: 50%;
    width: 440px;
    text-align: center;
    transform: translate(-50%,-50%);
  }

  .main .title {
    font-size: 30px;
    letter-spacing: 8px;
  }

  .main .loginbtn {
    width: 100%;
    font-size: 14px;
    letter-spacing: 15px;
    font-weight: bolder
  }
</style>
