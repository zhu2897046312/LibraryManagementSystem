<template>
  <div class="login_container">
    <div class="login_box" v-if="isOnLogin">
      <div class="form">
        <span>图书管理系统</span>
        <form>
          <input placeholder="邮箱地址" type="text" v-model="loginData.user_name"><br>
          <input placeholder="密码" type="password" v-model="loginData.password"><br>
          <button type="button" @click="login">登录</button>
          <a href="#" @click="switchToRegister">还没有账户？点击注册</a>
        </form>
      </div>
    </div>
    <div class="register_box" v-if="!isOnLogin">
      <span>新用户注册</span>
      <form>
        <input placeholder="用户名" v-model="registerData.user_name"><br>
        <input placeholder="密码" type="password" v-model="registerData.password"><br>
        <div class="radio_group">
          <label>
            <input type="radio" v-model="registerData.is_administrator" value="true"> 是
          </label>
          <label>
            <input type="radio" v-model="registerData.is_administrator" value="false"> 否
          </label>
        </div>
        <button type="button" @click="register">确认注册</button>
      </form>
      <a href="#" @click="switchToRegister">已有账户？ 点击登录</a>
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
      // 注册表单验证对象
      registerData: {
        user_name: '',
        password: '',
        is_administrator: ''
      },
      // 用户当前是否正准备登录
      isOnLogin: true,
    }
  },
  methods: {
    // 登录表单提交事件
    login(){
      // 发送登录请求
      let submit = {
          "user_name": this.loginData.user_name,
          "password": this.loginData.password
      }
      console.log(submit);
      axios({
        method: "POST",
        url: "http://localhost:8081/admin/login",
        data: submit,
        headers: {
          'Content-Type': 'application/json'
        }
      }).then(res => {
        console.log(submit); // 打印登录请求的数据
        console.log(res.data); // 打印后端返回的数据
        console.log("登录成功");
        this.$router.push({ path: '/' })
      }).catch(err => {
        console.log("登录失败");
      });
    },
    switchStatus() {
        // 切换界面
        this.isOnLogin = !this.isOnLogin
    },
    switchToRegister(event) {
      event.preventDefault(); // 阻止默认行为
      this.switchStatus(); // 调用切换方法
    },   
    // 注册表单提交事件
    register(){
      // 发送注册请求
      let submit = {
          "user_name": this.registerData.user_name,
          "password":  this.registerData.password,
          "is_administrator": this.registerData.is_administrator === 'true'
      }
      console.log(submit);
      axios({
        method: "POST",
        url: "http://localhost:8081/admin/register",
        data: submit,
        headers: {
          'Content-Type': 'application/json'
        }
      }).then(res => {
        console.log(submit); // 打印注册请求的数据
        console.log(res.data); // 打印后端返回的数据
        console.log("注册成功");
        this.$router.push({ path: '/login' })
      }).catch(err => {
        console.log("注册失败");
      });
    }
  }
}
</script>

<style scoped>
.login_container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

.login_box, .register_box {
  background-color: #f8f8f8;
  border-radius: 10px;
  padding: 20px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  width: 300px;
  text-align: center;
}

.form {
  margin-bottom: 20px;
}

.form input {
  margin-bottom: 10px;
  padding: 10px;
  width: 100%;
  border-radius: 5px;
  border: 1px solid #ccc;
}

.form button {
  padding: 10px 20px;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.radio_group {
  display: flex;
  justify-content: center;
  margin-bottom: 10px;
}

.radio_group label {
  margin-right: 10px;
}

.radio_group input[type="radio"] {
  margin-right: 5px;
}

a {
  display: block;
  color: #007bff;
  margin-top: 10px;
}
</style>
