<template>
    <div>
      <h2>借书</h2>
      <!-- 借书表单 -->
  
      <h2>还书</h2>
      <!-- 还书表单 -->
  
      <h2>借还记录</h2>
      <ul class="loan-history">
        <li v-for="record in loanRecords" :key="record.id">
          {{ record.borrower }} 借阅了 {{ record.book }}，借阅时间：{{ record.borrowDate }}，归还时间：{{ record.returnDate }}
        </li>
      </ul>
    </div>
  </template>
  
  <script>
  export default {
    name: 'LoanManagement',
    data() {
      return {
        loanRecords: [
          { id: 1, borrower: 'Alice', book: 'Book 1', borrowDate: '2024-05-01', returnDate: '2024-05-10' },
          { id: 2, borrower: 'Bob', book: 'Book 2', borrowDate: '2024-05-03', returnDate: '2024-05-12' },
          { id: 3, borrower: 'Charlie', book: 'Book 3', borrowDate: '2024-05-05', returnDate: '2024-05-15' },
          { id: 4, borrower: 'David', book: 'Book 4', borrowDate: '2024-05-07', returnDate: '2024-05-18' },
          { id: 5, borrower: 'Eve', book: 'Book 5', borrowDate: '2024-05-09', returnDate: '2024-05-20' }
        ]
      };
    }
  };
  </script>
  
  <style>
  /* 样式未改变 */
  </style>
  