<template>
  <div class="login_container">
    <div class="login_box" style="display: block;">
      <div class="form">
        <span>图书管理系统</span>
          <form>
            <input placeholder="邮箱地址" type="text" v-model="loginData.user_name"><br>
            <input placeholder="密码" type="password" v-model="loginData.password"><br>
            <button type="button" @click="login">登录</button>
            <a href="#">还没有账户？点击注册</a>
          </form>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  data(){
    return {
      // 登录表单的数据绑定对象
      loginData: {
        user_name: '',
        password: ''
      },
      // // 注册表单验证对象
      // registerData: {
      //   stu_id: '',
      //   name: '',
      //   user_name: '',
      //   vali: '',
      //   password: ''
      // }
    }
  },
  methods: {
    // 登录表单提交事件
    login(){
      // 发送登录请求
      let submit = {
          "user": {
            "user_name": this.loginData.user_name,
            "password": this.loginData.password
          }
        }
        axios({
        method: "POST",
        url: "http://localhost:8081/admin/login",
        data: JSON.stringify(submit),
        headers: {
          'Content-Type': 'application/json'
        }
      }).then(res => {
        console.log(res);
        console.log("登录成功");
      }).catch(err => {
        console.log("登录失败");
      });
    },
    // 注册表单提交事件
    registerSubmit(){
      // 发送注册请求
    }
  }
}
</script>