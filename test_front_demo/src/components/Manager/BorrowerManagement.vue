<template>
  <div>
    <input v-model="searchQuery" placeholder="搜索借阅者" class="search-box">
    <button @click="showAddForm = true" class="add-button">添加借阅者</button>
    <ul class="borrower-list">
      <li v-for="borrower in filteredBorrowers" :key="borrower.id">
        <span>{{ borrower.borrowerName }}</span>
        <button @click="editBorrower(borrower)" class="edit-button">编辑</button>
        <button @click="deleteBorrower(borrower.id)" class="delete-button">删除</button>
      </li>
    </ul>
    
    <div v-if="showAddForm" class="form-overlay">
      <div class="form-container">
        <h3>{{ isEditing ? '编辑借阅者' : '添加借阅者' }}</h3>
        <form @submit.prevent="isEditing ? updateBorrower() : addBorrower()">
          <label for="borrowerName">姓名:</label>
          <input v-model="currentBorrower.borrowerName" id="borrowerName" required>
          <!-- Add more fields as necessary -->
          <button type="submit">{{ isEditing ? '更新' : '添加' }}</button>
          <button type="button" @click="cancelEdit">取消</button>
        </form>
      </div>
    </div>

    <div class="pagination">
      <button @click="prevPage" :disabled="currentPage === 1">上一页</button>
      <span>第 {{ currentPage }} 页，共 {{ totalPages }} 页</span>
      <button @click="nextPage" :disabled="currentPage === totalPages">下一页</button>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'BorrowerManagement',
  data() {
    return {
      searchQuery: '',
      borrowers: [],
      showAddForm: false,
      isEditing: false,
      currentBorrower: this.getEmptyBorrower(),
      currentPage: 1,
      totalPages: 1,
      pageSize: 10
    };
  },
  computed: {
    filteredBorrowers() {
      return this.borrowers.filter(borrower =>
        borrower.borrowerName.toLowerCase().includes(this.searchQuery.toLowerCase())
      );
    }
  },
  methods: {
    async fetchBorrowers(page = 1) {
      try {
        const response = await axios.get(`/api/borrowers?page=${page}&pageSize=${this.pageSize}`);
        this.borrowers = response.data.borrowers;
        this.totalPages = response.data.totalPages;
      } catch (error) {
        console.error("Error fetching borrowers:", error);
      }
    },
    getEmptyBorrower() {
      return {
        id: null,
        borrowerName: '',
        // initialize other fields as necessary
      };
    },
    async addBorrower() {
      try {
        const response = await axios.post('/api/borrowers', this.currentBorrower);
        this.fetchBorrowers(this.currentPage); // Refresh the current page
        this.showAddForm = false;
        this.currentBorrower = this.getEmptyBorrower();
      } catch (error) {
        console.error("Error adding borrower:", error);
      }
    },
    async updateBorrower() {
      try {
        const response = await axios.put(`/api/borrowers/${this.currentBorrower.id}`, this.currentBorrower);
        this.fetchBorrowers(this.currentPage); // Refresh the current page
        this.showAddForm = false;
        this.currentBorrower = this.getEmptyBorrower();
        this.isEditing = false;
      } catch (error) {
        console.error("Error updating borrower:", error);
      }
    },
    editBorrower(borrower) {
      this.currentBorrower = { ...borrower };
      this.showAddForm = true;
      this.isEditing = true;
    },
    async deleteBorrower(id) {
      try {
        await axios.delete(`/api/borrowers/${id}`);
        this.fetchBorrowers(this.currentPage); // Refresh the current page
      } catch (error) {
        console.error("Error deleting borrower:", error);
      }
    },
    cancelEdit() {
      this.showAddForm = false;
      this.currentBorrower = this.getEmptyBorrower();
      this.isEditing = false;
    },
    prevPage() {
      if (this.currentPage > 1) {
        this.currentPage--;
        this.fetchBorrowers(this.currentPage);
      }
    },
    nextPage() {
      if (this.currentPage < this.totalPages) {
        this.currentPage++;
        this.fetchBorrowers(this.currentPage);
      }
    }
  },
  created() {
    this.fetchBorrowers();
  }
};
</script>

<style>
.search-box {
  width: 100%;
  padding: 10px;
  margin-bottom: 20px;
  box-sizing: border-box;
}

.add-button {
  padding: 10px;
  margin-bottom: 20px;
}

.borrower-list {
  list-style: none;
  padding: 0;
}

.borrower-list li {
  padding: 10px;
  background-color: #f4f4f4;
  border: 1px solid #e7e7e7;
  margin-bottom: 10px;
  border-radius: 4px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.edit-button, .delete-button {
  margin-left: 10px;
}

.form-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
}

.form-container {
  background: #fff;
  padding: 20px;
  border-radius: 5px;
}

.pagination {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 20px;
}

.pagination button {
  margin: 0 10px;
  padding: 10px;
}
</style>
