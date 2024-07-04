<template>
  <div class="user-management">
    <h2>用户管理</h2>
    <table>
      <thead>
        <tr>
          <th>用户名</th>
          <th>管理员</th>
          <th>操作</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="user in users" :key="user.user_name">
          <td>{{ user.user_name }}</td>
          <td>{{ user.is_administrator ? '是' : '否' }}</td>
          <td>
            <button @click="editUser(user)">修改</button>
            <button @click="deleteUser(user.user_name)">删除</button>
          </td>
        </tr>
      </tbody>
    </table>
    <div class="pagination">
      <button @click="changePage(page - 1)" :disabled="page === 1">上一页</button>
      <span>第 {{ page }} 页</span>
      <button @click="changePage(page + 1)" :disabled="page >= totalPages">下一页</button>
    </div>

    <div v-if="showEditModal" class="modal">
      <div class="modal-content">
        <span class="close" @click="closeEditModal">&times;</span>
        <h2>编辑用户信息</h2>
        <form @submit.prevent="updateUser">
          <div>
            <label for="user_name" >用户名:</label>
            <input type="text" v-model="currentUser.user_name" id="user_name" required />
          </div>
          <div>
            <label for="is_administrator">管理员:</label>
            <input type="checkbox" v-model="currentUser.is_administrator" id="is_administrator" />
          </div>
          <button type="submit">保存</button>
        </form>
      </div>
    </div>
  </div>
</template>

  
<script>
import axios from 'axios';

export default {
  name: 'UserManagement',
  data() {
    return {
      users: [],
      page: 1,
      pageSize: 6,
      totalPages: 1,
      showEditModal: false,
      currentUser: {
        user_name: '',
        is_administrator: false
      }
    };
  },
  methods: {
    async fetchUsers(page, pageSize) {
      try {
        const response = await axios.get('http://localhost:8081/admin/UserPageQuery', {
          params: {
            page: page,
            page_size: pageSize
          }
        });
        this.users = response.data.data;
        this.totalPages = Math.ceil(response.data.total / pageSize);
      } catch (error) {
        console.error('Error fetching users:', error);
      }
    },
    changePage(newPage) {
      if (newPage > 0 && newPage <= this.totalPages) {
        this.page = newPage;
        this.fetchUsers(this.page, this.pageSize);
      }
    },
    editUser(user) {
      this.currentUser = { ...user };
      this.showEditModal = true;
    },
    closeEditModal() {
      this.showEditModal = false;
    },
    async updateUser() {
      try {
        await axios.post(`http://localhost:8081/admin/UserUpdate`, this.currentUser);
        this.fetchUsers(this.page, this.pageSize);
        this.showEditModal = false;
      } catch (error) {
        console.error('Error updating user:', error);
      }
    },
    async deleteUser(userName) {
      try {
        const confirmed = confirm(`确定要删除用户 ${userName} 吗？`);
        if (confirmed) {
          await axios.post(`http://localhost:8081/admin/UserDelete`, { user_name: userName } );
          this.fetchUsers(this.page, this.pageSize);
        }
      } catch (error) {
        console.error('Error deleting user:', error);
      }
    }
  },
  created() {
    this.fetchUsers(this.page, this.pageSize);
  }
};
</script>

  
<style>
.user-management {
  padding: 20px;
}

.user-management table {
  width: 100%;
  border-collapse: collapse;
}

.user-management th, .user-management td {
  border: 1px solid #ddd;
  padding: 8px;
}

.user-management th {
  background-color: #f4f4f4;
}

.pagination {
  margin-top: 20px;
  text-align: center;
}

.pagination button {
  padding: 5px 10px;
  margin: 0 5px;
}

button {
  padding: 5px 10px;
  margin: 0 5px;
  cursor: pointer;
}

.modal {
  display: flex;
  justify-content: center;
  align-items: center;
  position: fixed;
  z-index: 1;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  overflow: auto;
  background-color: rgba(0,0,0,0.4);
}

.modal-content {
  background-color: #fff;
  padding: 20px;
  border: 1px solid #888;
  width: 400px;
  position: relative;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  animation: modalopen 0.4s;
}

@keyframes modalopen {
  from { opacity: 0; }
  to { opacity: 1; }
}

.close {
  color: #aaa;
  float: right;
  font-size: 28px;
  font-weight: bold;
}

.close:hover,
.close:focus {
  color: #000;
  text-decoration: none;
  cursor: pointer;
}
</style>
