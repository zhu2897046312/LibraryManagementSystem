<template>
  <div class="dashboard" @click="hideProfileMenu">
    <aside class="sidebar">
      <h1>工作台</h1>
      <nav>
        <ul>
          <li @click="currentView = 'BorrowerManagement'">借阅者管理</li>
          <li @click="currentView = 'BookManagement'">图书管理</li>
          <li @click="currentView = 'LoanManagement'">借阅管理</li>
          <li @click.stop="currentView = 'UserManagement'">用户管理</li>
          <!-- <li @click="showOtherInfo = toggleOtherInfo" @mouseleave="hideOtherInfo">
            其他信息管理
            <ul v-if="showOtherInfo">
              <li @click.stop="currentView = 'BorrowerCategoryManagement'">借阅者类别信息管理</li>
              <li @click.stop="currentView = 'BookCategoryManagement'">图书类别信息管理</li>
            </ul>
          </li> -->
        </ul>
      </nav>
    </aside>
    <main class="content">
      <header class="header">
        <div class="profile" @click.stop="toggleProfileMenu">
          <img src="../assets/libary_logo.jpg" alt="User Avatar" class="avatar">
          <ul v-if="showProfileMenu" class="profile-menu">
            <li @click="currentView = 'UserProfile'">个人信息</li>
            <li @click="logout">退出登录</li>
          </ul>
        </div>
      </header>
      <component :is="currentView"/>
    </main>
  </div>
</template>

<script>
import Login from './Login.vue';
import BorrowerManagement from '../components/Manager/BorrowerManagement.vue';
import BookManagement from '../components/Manager/BookManagement.vue';
import LoanManagement from '../components/Manager/LoanManagement.vue';
import UserManagement from '../components/Manager/UserManagement.vue';

export default {
  name: 'Dashboard',
  components: {
    Login,
    BorrowerManagement,
    BookManagement,
    LoanManagement,
    UserManagement
  },
  data() {
    return {
      currentView: 'BorrowerManagement',
      showOtherInfo: false,
      showProfileMenu: false,
    };
  },
  methods: {
    toggleOtherInfo() {
      this.showOtherInfo = !this.showOtherInfo;
    },
    hideOtherInfo() {
      this.showOtherInfo = false;
    },
    toggleProfileMenu() {
      this.showProfileMenu = !this.showProfileMenu;
    },
    hideProfileMenu() {
      this.showProfileMenu = false;
    },
    logout() {
      // Clear user data (if any)
      // Perform any necessary cleanup actions

      // Redirect to login page
      this.$router.push({ path: '/login' })
    }
  }
};
</script>

<style>
.dashboard {
  display: flex;
}

.sidebar {
  width: 200px;
  background-color: #f4f4f4;
  padding: 20px;
  box-shadow: 2px 0 5px rgba(0,0,0,0.1);
}

.sidebar h1 {
  font-size: 18px;
  margin-bottom: 20px;
}

.sidebar nav ul {
  list-style: none;
  padding: 0;
}

.sidebar nav ul li {
  margin-bottom: 10px;
  cursor: pointer;
  padding: 10px;
  background-color: #e7e7e7;
  border-radius: 4px;
  text-align: center;
  position: relative;
}

.sidebar nav ul li:hover {
  background-color: #d4d4d4;
}

.sidebar nav ul li ul {
  list-style: none;
  padding: 0;
  margin: 10px 0 0 0;
  background-color: #f4f4f4;
  border-radius: 4px;
}

.sidebar nav ul li ul li {
  padding: 10px;
  background-color: #e7e7e7;
  border-radius: 4px;
  margin: 5px 0;
  text-align: left;
}

.sidebar nav ul li ul li:hover {
  background-color: #d4d4d4;
}

.content {
  flex: 1;
  padding: 20px;
}

.header {
  display: flex;
  justify-content: flex-end;
  padding: 10px 20px;
  background-color: #f4f4f4;
  box-shadow: 0 2px 5px rgba(0,0,0,0.1);
}

.profile {
  position: relative;
  cursor: pointer;
}

.avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
}

.profile-menu {
  list-style: none;
  padding: 0;
  margin: 10px 0 0 0;
  background-color: #f4f4f4;
  border: 1px solid #e7e7e7;
  border-radius: 4px;
  position: absolute;
  right: 0;
  top: 50px;
  width: 120px;
  box-shadow: 0 2px 5px rgba(0,0,0,0.1);
}

.profile-menu li {
  padding: 10px;
  text-align: left;
}

.profile-menu li:hover {
  background-color: #d4d4d4;
}
</style>
