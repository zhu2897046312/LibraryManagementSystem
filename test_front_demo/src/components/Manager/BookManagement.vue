<template>
    <div>
      <input v-model="searchQuery" placeholder="搜索图书" class="search-box">
      <ul class="book-list">
        <li v-for="book in filteredBooks" :key="book.id">
          {{ book.title }}
        </li>
      </ul>
    </div>
  </template>
  
  <script>
  export default {
    name: 'BookManagement',
    data() {
      return {
        searchQuery: '',
        books: [
          { id: 1, title: 'Book 1' },
          { id: 2, title: 'Book 2' },
          { id: 3, title: 'Book 3' },
          { id: 4, title: 'Book 4' },
          { id: 5, title: 'Book 5' }
        ]
      };
    },
    computed: {
      filteredBooks() {
        return this.books.filter(book =>
          book.title.toLowerCase().includes(this.searchQuery.toLowerCase())
        );
      }
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
  
  .book-list {
    list-style: none;
    padding: 0;
  }
  
  .book-list li {
    padding: 10px;
    background-color: #f4f4f4;
    border: 1px solid #e7e7e7;
    margin-bottom: 10px;
    border-radius: 4px;
  }
  </style>
  