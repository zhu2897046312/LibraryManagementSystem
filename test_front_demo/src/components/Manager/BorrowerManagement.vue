<template>
  <div class="borrower-management">
    <h2>借阅者管理</h2>

    <button @click="showAddModal = true">新增借阅者</button>

    <table>
      <thead>
        <tr>
          <th>ID</th>
          <th>借阅者ID</th>
          <th>姓名</th>
          <th>地址</th>
          <th>电话</th>
          <th>类型</th>
          <th>借阅数量</th>
          <th>身份证号</th>
          <th>发证日期</th>
          <th>挂失</th>
          <th>性别</th>
          <th>单位</th>
          <th>操作</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="borrower in borrowers" :key="borrower.id">
          <td>{{ borrower.id }}</td>
          <td>{{ borrower.borrower_id }}</td>
          <td>{{ borrower.borrower_name }}</td>
          <td>{{ borrower.borrower_address }}</td>
          <td>{{ borrower.borrower_phone }}</td>
          <td>{{ borrower.borrower_type }}</td>
          <td>{{ borrower.borrower_books_numbers }}</td>
          <td>{{ borrower.identity_card_id }}</td>
          <td>{{ new Date(borrower.issuance_date).toLocaleDateString() }}</td>
          <td>{{ borrower.is_loss ? '是' : '否' }}</td>
          <td>{{ borrower.gender === 1 ? '男' : '女' }}</td>
          <td>{{ borrower.unit }}</td>
          <td>
            <button @click="editBorrower(borrower)">修改</button>
            <button @click="deleteBorrower(borrower.borrower_id)">删除</button>
          </td>
        </tr>
      </tbody>
    </table>
    <div class="pagination">
      <button @click="changePage(page - 1)" :disabled="page === 1">上一页</button>
      <span>第 {{ page }} 页</span>
      <button @click="changePage(page + 1)" :disabled="page >= totalPages">下一页</button>
    </div>

    <!-- 编辑模态框 -->
    <div v-if="showEditModal" class="modal">
      <div class="modal-content">
        <span class="close" @click="closeEditModal">&times;</span>
        <h2>编辑借阅者信息</h2>
        <form @submit.prevent="updateBorrower">
          <div>
            <label for="borrower_id">借阅者ID:</label>
            <input type="text" v-model="currentBorrower.borrower_id" id="borrower_id" readonly />
          </div>
          <div>
            <label for="borrower_name">姓名:</label>
            <input type="text" v-model="currentBorrower.borrower_name" id="borrower_name" required />
          </div>
          <div>
            <label for="borrower_address">地址:</label>
            <input type="text" v-model="currentBorrower.borrower_address" id="borrower_address" required />
          </div>
          <div>
            <label for="borrower_phone">电话:</label>
            <input type="text" v-model="currentBorrower.borrower_phone" id="borrower_phone" required />
          </div>
          <div>
            <label for="borrower_type">类型:</label>
            <input type="text" v-model="currentBorrower.borrower_type" id="borrower_type" required />
          </div>
          <div>
            <label for="borrower_books_numbers">借阅数量:</label>
            <input type="number" v-model="currentBorrower.borrower_books_numbers" id="borrower_books_numbers" required />
          </div>
          <div>
            <label for="identity_card_id">身份证号:</label>
            <input type="text" v-model="currentBorrower.identity_card_id" id="identity_card_id" required />
          </div>
          <div>
            <label for="issuance_date">发证日期:</label>
            <input type="date" v-model="currentBorrower.issuance_date" id="issuance_date" required />
          </div>
          <div>
            <label for="is_loss">挂失:</label>
            <input type="checkbox" v-model="currentBorrower.is_loss" id="is_loss" />
          </div>
          <div>
            <label for="gender">性别:</label>
            <select v-model="currentBorrower.gender" id="gender" required>
              <option value="1">男</option>
              <option value="2">女</option>
            </select>
          </div>
          <div>
            <label for="unit">单位:</label>
            <input type="text" v-model="currentBorrower.unit" id="unit" required />
          </div>
          <button type="submit">保存</button>
        </form>
      </div>
    </div>

    <div v-if="showAddModal" class="modal">
      <div class="modal-content">
        <span class="close" @click="closeAddModal">&times;</span>
        <h2>编辑借阅者信息</h2>
        <form @submit.prevent="addBorrower">
          <div>
            <label for="add_borrower_id">借阅者ID:</label>
            <input type="text" v-model="newBorrower.borrower_id" id="add_borrower_id" required />
          </div>
          <div>
            <label for="add_borrower_name">姓名:</label>
            <input type="text" v-model="newBorrower.borrower_name" id="add_borrower_name" required />
          </div>
          <div>
            <label for="add_borrower_address">地址:</label>
            <input type="text" v-model="newBorrower.borrower_address" id="add_borrower_address" required />
          </div>
          <div>
            <label for="add_borrower_phone">电话:</label>
            <input type="text" v-model="newBorrower.borrower_phone" id="add_borrower_phone" required />
          </div>
          <div>
            <label for="add_borrower_type">类型:</label>
            <input type="text" v-model="newBorrower.borrower_type" id="add_borrower_type" required />
          </div>
          <div>
            <label for="add_borrower_books_numbers">借阅数量:</label>
            <input type="number" v-model="newBorrower.borrower_books_numbers" id="add_borrower_books_numbers" required />
          </div>
          <div>
            <label for="add_identity_card_id">身份证号:</label>
            <input type="text" v-model="newBorrower.identity_card_id" id="add_identity_card_id" required />
          </div>
          <div>
            <label for="add_issuance_date">发证日期:</label>
            <input type="date" v-model="newBorrower.issuance_date" id="add_issuance_date" required />
          </div>
          <div>
            <label for="add_is_loss">挂失:</label>
            <input type="checkbox" v-model="newBorrower.is_loss" id="add_is_loss" />
          </div>
          <div>
            <label for="add_gender">性别:</label>
            <select v-model="newBorrower.gender" id="add_gender" required>
              <option value="1">男</option>
              <option value="2">女</option>
            </select>
          </div>
          <div>
            <label for="add_unit">单位:</label>
            <input type="text" v-model="newBorrower.unit" id="add_unit" required />
          </div>
          <button type="submit">添加</button>
        </form>
      </div>
    </div>
    
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'BorrowerManagement',
  data() {
    return {
      borrowers: [],
      page: 1,
      pageSize: 10,
      totalPages: 1,
      showEditModal: false,
      showAddModal: false,
      currentBorrower: {
        borrower_id: '',
        borrower_name: '',
        borrower_address: '',
        borrower_phone: '',
        borrower_type: '',
        borrower_books_numbers: 0,
        identity_card_id: '',
        issuance_date: '',
        is_loss: false,
        gender: 1,
        unit: ''
      },
      newBorrower: {
        borrower_id: '',
        borrower_name: '',
        borrower_address: '',
        borrower_phone: '',
        borrower_type: '',
        borrower_books_numbers: 0,
        identity_card_id: '',
        issuance_date: '',
        is_loss: false,
        gender: 1,
        unit: ''
      }
    };
  },
  methods: {
    async fetchBorrowers(page, pageSize) {
      try {
        const response = await axios.get('http://localhost:8081/admin/BorrowerPageQuery', {
          params: {
            page: page,
            page_size: pageSize
          }
        });
        this.borrowers = response.data.data;
        this.totalPages = Math.ceil(response.data.total / pageSize);
      } catch (error) {
        console.error('Error fetching borrowers:', error);
      }
    },
    changePage(newPage) {
      if (newPage > 0 && newPage <= this.totalPages) {
        this.page = newPage;
        this.fetchBorrowers(this.page, this.pageSize);
      }
    },
    editBorrower(borrower) {
      this.currentBorrower = { ...borrower };
      this.showEditModal = true;
    },
    async updateBorrower() {
      try {
        this.currentBorrower.issuance_date = new Date(this.currentBorrower.issuance_date).toISOString();
        await axios.post(`http://localhost:8081/admin/BorrowerUpdate`, this.currentBorrower);
        this.fetchBorrowers(this.page, this.pageSize);
        this.showEditModal = false;
      } catch (error) {
        console.error('Error updating borrower:', error);
      }
    },
    closeEditModal() {
      this.showEditModal = false;
    },
    async deleteBorrower(borrowerId) {
      try {
        const confirmed = confirm(`确定要删除借阅者 ${borrowerId} 吗？`);
        if (confirmed) {
          await axios.post(`http://localhost:8081/admin/BorrowerDelete`, { borrower_id: borrowerId });
          this.fetchBorrowers(this.page, this.pageSize);
        }
      } catch (error) {
        console.error('Error deleting borrower:', error);
      }
    },
    closeAddModal() {
      this.showAddModal = false;
      this.clearNewBorrowerForm();
    },
    async addBorrower() {
      try {
        // Convert dates to ISO 8601 format
        this.newBorrower.issuance_date = new Date(this.newBorrower.issuance_date).toISOString();
        await axios.post(`http://localhost:8081/admin/BorrowerInsert`, this.newBorrower);
        
        this.fetchBorrowers(this.page, this.pageSize);
        this.showAddModal = false;
        this.clearNewBorrowerForm();
      } catch (error) {
        console.error('Error adding borrower:', error);
      }
    },
    clearNewBorrowerForm() {
      this.newBorrower = {
        borrower_id: '',
        borrower_name: '',
        borrower_address: '',
        borrower_phone: '',
        borrower_type: '',
        borrower_books_numbers: 0,
        identity_card_id: '',
        issuance_date: '',
        is_loss: false,
        gender: 1,
        unit: ''
      };
      
    }
  },
  created() {
    this.fetchBorrowers(this.page, this.pageSize);
  }
};
</script>

<style>
.borrower-management {
  padding: 20px;
  font-family: 'Arial', sans-serif;
}

.borrower-management h2 {
  font-size: 24px;
  margin-bottom: 20px;
  color: #333;
}

.borrower-management button {
  background-color: #4CAF50;
  color: white;
  border: none;
  padding: 10px 20px;
  font-size: 16px;
  cursor: pointer;
  border-radius: 5px;
  transition: background-color 0.3s;
}

.borrower-management button:hover {
  background-color: #45a049;
}

.borrower-management table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 20px;
}

.borrower-management th, .borrower-management td {
  border: 1px solid #ddd;
  padding: 12px;
  text-align: left;
}

.borrower-management th {
  background-color: #f4f4f4;
  color: #333;
  font-weight: bold;
}

.borrower-management td {
  background-color: #fff;
}

.pagination {
  margin-top: 20px;
  text-align: center;
}

.pagination button {
  background-color: #007BFF;
  color: white;
  border: none;
  padding: 10px 15px;
  font-size: 16px;
  cursor: pointer;
  border-radius: 5px;
  transition: background-color 0.3s;
  margin: 0 5px;
}

.pagination button:disabled {
  background-color: #ccc;
  cursor: not-allowed;
}

.pagination button:hover:not(:disabled) {
  background-color: #0056b3;
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
  border-radius: 10px;
  width: 500px;
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

form {
  display: flex;
  flex-direction: column;
}

form div {
  margin-bottom: 10px;
}

form label {
  font-size: 14px;
  color: #333;
  margin-bottom: 5px;
}

form input[type="text"],
form input[type="number"],
form input[type="date"],
form select {
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 4px;
  width: 100%;
  box-sizing: border-box;
}

form input[type="checkbox"] {
  width: auto;
}

form button {
  background-color: #4CAF50;
  color: white;
  border: none;
  padding: 10px;
  font-size: 16px;
  cursor: pointer;
  border-radius: 5px;
  transition: background-color 0.3s;
  margin-top: 10px;
}

form button:hover {
  background-color: #45a049;
}

</style>
