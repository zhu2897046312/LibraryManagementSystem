<template>
  <div class="borrowing-management">
    <h2>借阅管理</h2>

    <button @click="showAddModal = true">新增借书历史</button>

    <table>
      <thead>
        <tr>
          <th>ID</th>
          <th>借阅者ID</th>
          <th>书籍ID</th>
          <th>借阅者姓名</th>
          <th>书籍名称</th>
          <th>借阅日期</th>
          <th>罚款金额</th>
          <th>操作</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="borrowing in borrowings" :key="borrowing.id">
          <td>{{ borrowing.id }}</td>
          <td>{{ borrowing.borrower_id }}</td>
          <td>{{ borrowing.book_id }}</td>
          <td>{{ borrowing.borrower_name }}</td>
          <td>{{ borrowing.book_name }}</td>
          <td>{{ new Date(borrowing.borrowed_date).toLocaleDateString() }}</td>
          <td>{{ borrowing.fine_amount.toFixed(2) }}</td>
          <td>
            <button @click="editBorrowing(borrowing)">修改</button>
            <button @click="deleteBorrowing(borrowing.borrower_id)">删除</button>
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
        <h2>编辑借阅信息</h2>
        <form @submit.prevent="updateBorrowing">
          <div>
            <label for="borrower_id">借阅者ID:</label>
            <input type="text" v-model="currentBorrowing.borrower_id" id="borrower_id" readonly />
          </div>
          <div>
            <label for="book_id">书籍ID:</label>
            <input type="text" v-model="currentBorrowing.book_id" id="book_id" readonly />
          </div>
          <div>
            <label for="borrowed_date">借阅日期:</label>
            <input type="date" v-model="currentBorrowing.borrowed_date" id="borrowed_date" required />
          </div>
          <div>
            <label for="fine_amount">罚款金额:</label>
            <input type="number" v-model="currentBorrowing.fine_amount" id="fine_amount" required />
          </div>
          <button type="submit">保存</button>
        </form>
      </div>
    </div>

    <div v-if="showAddModal" class="modal">
      <div class="modal-content">
        <span class="close" @click="closeAddModal">&times;</span>
        <h2>编辑借阅信息</h2>
        <form @submit.prevent="addBorrowing">
          <div>
            <label for="add_borrower_id">借阅者ID:</label>
            <input type="text" v-model="newBorrowing.borrower_id" id="add_borrower_id" required />
          </div>
          <div>
            <label for="add_book_id">书籍ID:</label>
            <input type="text" v-model="newBorrowing.book_id" id="add_book_id" required />
          </div>
          <div>
            <label for="add_borrowed_date">借阅日期:</label>
            <input type="date" v-model="newBorrowing.borrowed_date" id="add_borrowed_date" required />
          </div>
          <div>
            <label for="add_fine_amount">罚款金额:</label>
            <input type="number" v-model="newBorrowing.fine_amount" id="add_fine_amount" required />
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
  name: 'BorrowingManagement',
  data() {
    return {
      borrowings: [],
      page: 1,
      pageSize: 10,
      totalPages: 1,
      showEditModal: false,
      showAddModal: false,
      currentBorrowing: {
        borrower_id: '',
        book_id: '',
        borrowed_date: '',
        fine_amount: 0
      },
      newBorrowing: {
        borrower_id: '',
        book_id: '',
        borrowed_date: '',
        fine_amount: 0
      }
    };
  },
  methods: {
    async fetchBorrowings(page, pageSize) {
      try {
        const response = await axios.get('http://localhost:8081/admin/BorrowingPageQuery', {
          params: {
            page: page,
            page_size: pageSize
          }
        });
        this.borrowings = response.data.data;
        this.totalPages = Math.ceil(response.data.total / pageSize);
      } catch (error) {
        console.error('Error fetching borrowings:', error);
      }
    },
    changePage(newPage) {
      if (newPage > 0 && newPage <= this.totalPages) {
        this.page = newPage;
        this.fetchBorrowings(this.page, this.pageSize);
      }
    },
    editBorrowing(borrowing) {
      this.currentBorrowing = { ...borrowing };
      this.showEditModal = true;
    },
    async updateBorrowing() {
      try {
        // Convert issuance_date to ISO string
        this.currentBorrowing.borrowed_date = new Date(this.currentBorrowing.borrowed_date).toISOString();
        await axios.post(`http://localhost:8081/admin/BorrowingUpdate`, this.currentBorrowing);
        this.fetchBorrowings(this.page, this.pageSize);
        this.showEditModal = false;
      } catch (error) {
        console.error('Error updating borrowing:', error);
      }
    },
    closeEditModal() {
      this.showEditModal = false;
    },
    async deleteBorrowing(borrowerId) {
      try {
        const confirmed = confirm(`确定要删除借阅者 ${borrowerId} 的记录吗？`);
        if (confirmed) {
          await axios.post(`http://localhost:8081/admin/BorrowingDelete`, { borrower_id: borrowerId });
          this.fetchBorrowings(this.page, this.pageSize);
        }
      } catch (error) {
        console.error('Error deleting borrowing:', error);
      }
    },
    closeAddModal() {
      this.showAddModal = false;
      this.clearNewBorrowingForm();
    },
    async addBorrowing() {
      try {
        // Convert dates to ISO 8601 format
        this.newBorrowing.borrowed_date = new Date(this.newBorrowing.borrowed_date).toISOString();
        await axios.post(`http://localhost:8081/admin/BorrowingInsert`, this.newBorrowing);
        
        this.fetchBorrowings(this.page, this.pageSize);
        this.showAddModal = false;
        this.clearNewBorrowingForm();
      } catch (error) {
        console.error('Error adding borrower:', error);
      }
    },
    clearNewBorrowingForm() {
      this.newBorrower = {
        borrower_id: '',
        book_id: '',
        borrowed_date: '',
        fine_amount: 0
      };
    }
  },
  created() {
    this.fetchBorrowings(this.page, this.pageSize);
  }
};
</script>

<style>
.borrowing-management {
  padding: 20px;
  font-family: Arial, sans-serif;
  background-color: #f9f9f9;
}

.borrowing-management h2 {
  color: #333;
  margin-bottom: 20px;
}

.borrowing-management table {
  width: 100%;
  border-collapse: collapse;
  margin-bottom: 20px;
  background-color: #fff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.borrowing-management th, .borrowing-management td {
  border: 1px solid #ddd;
  padding: 12px;
  text-align: left;
}

.borrowing-management th {
  background-color: #f4f4f4;
  color: #333;
  font-weight: bold;
}

.borrowing-management td {
  background-color: #fff;
}

.borrowing-management button {
  padding: 10px 15px;
  margin: 5px;
  cursor: pointer;
  border: none;
  border-radius: 4px;
  background-color: #007bff;
  color: white;
  transition: background-color 0.3s ease;
}

.borrowing-management button:hover {
  background-color: #0056b3;
}

.pagination {
  margin-top: 20px;
  text-align: center;
}

.pagination button {
  padding: 8px 12px;
  margin: 0 5px;
  cursor: pointer;
  border: none;
  border-radius: 4px;
  background-color: #007bff;
  color: white;
  transition: background-color 0.3s ease;
}

.pagination button:disabled {
  background-color: #ddd;
  cursor: not-allowed;
}

.pagination button:hover:enabled {
  background-color: #0056b3;
}

.pagination span {
  font-size: 16px;
  color: #333;
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
  background-color: rgba(0, 0, 0, 0.4);
}

.modal-content {
  background-color: #fff;
  padding: 20px;
  border-radius: 8px;
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

.modal form {
  display: flex;
  flex-direction: column;
}

.modal form div {
  margin-bottom: 15px;
}

.modal form label {
  font-weight: bold;
  margin-bottom: 5px;
  display: block;
}

.modal form input[type="text"],
.modal form input[type="date"],
.modal form input[type="number"] {
  width: calc(100% - 22px);
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.modal form button {
  align-self: flex-end;
  padding: 10px 15px;
  cursor: pointer;
  border: none;
  border-radius: 4px;
  background-color: #007bff;
  color: white;
  transition: background-color 0.3s ease;
}

.modal form button:hover {
  background-color: #0056b3;
}

</style>
