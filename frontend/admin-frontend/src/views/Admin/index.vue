<template>
  <div class="page-container">
    <div class="page-header flex-between">
      <span class="page-title">管理员管理</span>
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>
        添加管理员
      </el-button>
    </div>
    <el-table :data="tableData" style="width: 100%" stripe border>
      <el-table-column prop="id" label="ID" width="80" />
      <el-table-column prop="username" label="用户名" />
      <el-table-column prop="nickname" label="昵称" />
      <el-table-column prop="phone" label="手机号" width="130" />
      <el-table-column prop="status" label="状态" width="100">
        <template #default="scope">
          <el-tag :type="scope.row.status === 'active' ? 'success' : 'danger'" effect="dark">
            {{ scope.row.status === 'active' ? '正常' : '禁用' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="created_at" label="创建时间" width="180" />
      <el-table-column label="操作" width="150" fixed="right">
        <template #default="scope">
          <el-button :type="scope.row.status === 'active' ? 'danger' : 'success'" link size="small">
            {{ scope.row.status === 'active' ? '禁用' : '启用' }}
          </el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'

const tableData = ref([
  { id: 1, username: 'admin', nickname: '超级管理员', phone: '13800000000', status: 'active', created_at: '2026-01-01 00:00:00' },
  { id: 2, username: 'admin2', nickname: '管理员二', phone: '13800000002', status: 'active', created_at: '2026-02-15 10:00:00' }
])

const handleAdd = () => {
  ElMessage.info('添加管理员功能开发中...')
}
</script>
