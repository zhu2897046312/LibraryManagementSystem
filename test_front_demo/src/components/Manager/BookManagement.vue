<template>
  <div class="book-management">
    <h2>图书管理</h2>
    
    <button @click="showAddModal = true">新增图书</button>
    <table>
      <thead>
        <tr>
          <th>ID</th>
          <th>书名</th>
          <th>书籍ID</th>
          <th>作者</th>
          <th>出版社</th>
          <th>类型</th>
          <th>出版日期</th>
          <th>价格</th>
          <th>是否借出</th>
          <th>入库时间</th>
          <th>操作</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="book in books" :key="book.id">
          <td>{{ book.id }}</td>
          <td>{{ book.book_name }}</td>
          <td>{{ book.book_id }}</td>
          <td>{{ book.book_author }}</td>
          <td>{{ book.book_publisher }}</td>
          <td>{{ book.book_type }}</td>
          <td>{{ new Date(book.publish_time).toLocaleDateString() }}</td>
          <td>{{ book.book_price.toFixed(2) }}</td>
          <td>{{ book.is_borrowed ? '是' : '否' }}</td>
          <td>{{ new Date(book.inbound_time).toLocaleDateString() }}</td>
          <td>
            <button @click="editBook(book)">修改</button>
            <button @click="deleteBook(book.book_idS)">删除</button>
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
        <h2>编辑图书信息</h2>
        <form @submit.prevent="updateBook">
          <div>
            <label for="book_name">书名:</label>
            <input type="text" v-model="currentBook.book_name" id="book_name" required />
          </div>
          <div>
            <label for="book_id">书籍ID:</label>
            <input type="text" v-model="currentBook.book_id" id="book_id" required />
          </div>
          <div>
            <label for="book_author">作者:</label>
            <input type="text" v-model="currentBook.book_author" id="book_author" required />
          </div>
          <div>
            <label for="book_publisher">出版社:</label>
            <input type="text" v-model="currentBook.book_publisher" id="book_publisher" required />
          </div>
          <div>
            <label for="book_type">类型:</label>
            <input type="text" v-model="currentBook.book_type" id="book_type" required />
          </div>
          <div>
            <label for="publish_time">出版日期:</label>
            <input type="date" v-model="currentBook.publish_time" id="publish_time" required />
          </div>
          <div>
            <label for="book_price">价格:</label>
            <input type="number" v-model="currentBook.book_price" id="book_price" required />
          </div>
          <div>
            <label for="is_borrowed">是否借出:</label>
            <input type="checkbox" v-model="currentBook.is_borrowed" id="is_borrowed" />
          </div>
          <div>
            <label for="inbound_time">入库时间:</label>
            <input type="date" v-model="currentBook.inbound_time" id="inbound_time" required />
          </div>
          <button type="submit">保存</button>
        </form>
      </div>
    </div>

    <div v-if="showAddModal" class="modal">
      <div class="modal-content">
        <span class="close" @click="closeAddModal">&times;</span>
        <h2>新增图书信息</h2>
        <form @submit.prevent="addBook">
          <div>
            <label for="add_book_name">书名:</label>
            <input type="text" v-model="newBook.book_name" id="add_book_name" required />
          </div>
          <div>
            <label for="add_book_id">书籍ID:</label>
            <input type="text" v-model="newBook.book_id" id="add_book_id" required />
          </div>
          <div>
            <label for="add_book_author">作者:</label>
            <input type="text" v-model="newBook.book_author" id="add_book_author" required />
          </div>
          <div>
            <label for="add_book_publisher">出版社:</label>
            <input type="text" v-model="newBook.book_publisher" id="add_book_publisher" required />
          </div>
          <div>
            <label for="add_book_type">类型:</label>
            <input type="text" v-model="newBook.book_type" id="add_book_type" required />
          </div>
          <div>
            <label for="add_publish_time">出版日期:</label>
            <input type="date" v-model="newBook.publish_time" id="add_publish_time" required />
          </div>
          <div>
            <label for="add_book_price">价格:</label>
            <input type="number" v-model="newBook.book_price" id="add_book_price" required />
          </div>
          <div>
            <label for="add_is_borrowed">是否借出:</label>
            <input type="checkbox" v-model="newBook.is_borrowed" id="add_is_borrowed" />
          </div>
          <div>
            <label for="add_inbound_time">入库时间:</label>
            <input type="date" v-model="newBook.inbound_time" id="add_inbound_time" required />
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
  name: 'BookManagement',
  data() {
    return {
      books: [],
      page: 1,
      pageSize: 10,
      totalPages: 1,
      showEditModal: false,
      showAddModal:false,
      currentBook: {
        book_name: '',
        book_id: '',
        book_author: '',
        book_publisher: '',
        book_type: '',
        publish_time: '',
        book_price: 0,
        is_borrowed: false
      },
      newBook: {
        book_name: '',
        book_id: '',
        book_author: '',
        book_publisher: '',
        book_type: '',
        publish_time: '',
        book_price: 0,
        is_borrowed: false,
        inbound_time: ''
      }
    };
  },
  methods: {
    async fetchBooks(page, pageSize) {
      try {
        const response = await axios.get('http://localhost:8081/admin/BookPageQuery', {
          params: {
            page: page,
            page_size: pageSize
          }
        });
        this.books = response.data.data;
        this.totalPages = Math.ceil(response.data.total / pageSize);
      } catch (error) {
        console.error('Error fetching books:', error);
      }
    },
    changePage(newPage) {
      if (newPage > 0 && newPage <= this.totalPages) {
        this.page = newPage;
        this.fetchBooks(this.page, this.pageSize);
      }
    },
    editBook(book) {
      this.currentBook = { ...book };
      this.showEditModal = true;
    },
    closeEditModal() {
      this.showEditModal = false;
    },
    async updateBook() {
      try {
        // Convert dates to ISO 8601 format
    this.currentBook.publish_time = new Date(this.currentBook.publish_time).toISOString();
    this.currentBook.inbound_time = new Date(this.currentBook.inbound_time).toISOString();
        await axios.post(`http://localhost:8081/admin/BookUpdate`, this.currentBook);
        this.fetchBooks(this.page, this.pageSize);
        this.showEditModal = false;
      } catch (error) {
        console.error('Error updating book:', error);
      }
    },
    async deleteBook(bookId) {
      try {
        const confirmed = confirm(`确定要删除书籍 ${bookId} 吗？`);
        if (confirmed) {
          await axios.post(`http://localhost:8081/admin/BookDelete`, { book_id: bookId });
          this.fetchBooks(this.page, this.pageSize);
        }
      } catch (error) {
        console.error('Error deleting book:', error);
      }
    },
    closeAddModal() {
      this.showAddModal = false;
      this.clearNewBookForm();
    },
    async addBook() {
      try {
        // Convert dates to ISO 8601 format
        this.newBook.publish_time = new Date(this.newBook.publish_time).toISOString();
        this.newBook.inbound_time = new Date(this.newBook.inbound_time).toISOString();
        
        await axios.post(`http://localhost:8081/admin/BookInsert`, this.newBook);
        
        this.fetchBooks(this.page, this.pageSize);
        this.showAddModal = false;
        this.clearNewBookForm();
      } catch (error) {
        console.error('Error adding book:', error);
      }
    },
    clearNewBookForm() {
      this.newBook = {
        book_name: '',
        book_id: '',
        book_author: '',
        book_publisher: '',
        book_type: '',
        publish_time: '',
        book_price: 0,
        is_borrowed: false,
        inbound_time: ''
      };
    }
  },
  created() {
    this.fetchBooks(this.page, this.pageSize);
  }
};
</script>



<style>
.book-management {
  padding: 20px;
  font-family: Arial, sans-serif;
  background-color: #f8f9fa;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.book-management h2 {
  margin-bottom: 20px;
  color: #343a40;
  text-align: center;
}

.book-management button {
  padding: 10px 15px;
  margin: 10px 5px;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.book-management button:hover {
  background-color: #0056b3;
}

.book-management table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 20px;
  background-color: #fff;
}

.book-management th, .book-management td {
  border: 1px solid #ddd;
  padding: 12px 15px;
  text-align: left;
}

.book-management th {
  background-color: #007bff;
  color: #fff;
}

.book-management td {
  color: #495057;
}

.pagination {
  margin-top: 20px;
  text-align: center;
}

.pagination button {
  padding: 8px 12px;
  margin: 0 5px;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.pagination button:disabled {
  background-color: #c0c0c0;
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
  border: 1px solid #888;
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

.modal-content form div {
  margin-bottom: 15px;
}

.modal-content label {
  display: block;
  margin-bottom: 5px;
  color: #495057;
}

.modal-content input[type="text"],
.modal-content input[type="number"],
.modal-content input[type="date"],
.modal-content input[type="checkbox"] {
  width: 100%;
  padding: 8px;
  margin-top: 5px;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.modal-content button {
  padding: 10px 15px;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.modal-content button:hover {
  background-color: #0056b3;
}
</style>


