<template>
  <div class="page-container">
    <div class="page-header">
      <span class="page-title">用户管理</span>
    </div>
    <el-table :data="tableData" style="width: 100%" stripe border>
      <el-table-column prop="id" label="ID" width="80" />
      <el-table-column prop="username" label="用户名" />
      <el-table-column prop="nickname" label="昵称" />
      <el-table-column prop="phone" label="手机号" width="130" />
      <el-table-column prop="balance" label="余额" width="100">
        <template #default="scope">
          <span class="color-primary">¥{{ scope.row.balance }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="status" label="状态" width="100">
        <template #default="scope">
          <el-tag :type="scope.row.status === 'active' ? 'success' : 'danger'" effect="dark">
            {{ scope.row.status === 'active' ? '正常' : '禁用' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="created_at" label="注册时间" width="180" />
      <el-table-column label="操作" width="150" fixed="right">
        <template #default="scope">
          <el-button :type="scope.row.status === 'active' ? 'danger' : 'success'" link size="small">
            {{ scope.row.status === 'active' ? '禁用' : '启用' }}
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <div class="pagination-container">
      <el-pagination
        background
        layout="prev, pager, next, total"
        :total="total"
        :page-size="10"
      />
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const tableData = ref([
  { id: 1, username: 'zhangsan', nickname: '张三', phone: '13800138001', balance: 128.50, status: 'active', created_at: '2026-01-15 10:30:00' },
  { id: 2, username: 'lisi', nickname: '李四', phone: '13800138002', balance: 0.00, status: 'active', created_at: '2026-02-20 14:20:00' },
  { id: 3, username: 'wangwu', nickname: '王五', phone: '13800138003', balance: 580.00, status: 'inactive', created_at: '2026-03-10 09:15:00' },
  { id: 4, username: 'zhaoliu', nickname: '赵六', phone: '13800138004', balance: 215.80, status: 'active', created_at: '2026-04-05 16:45:00' }
])

const total = ref(50)
</script>
