<template>
  <div class="page-container">
    <div class="page-header flex-between">
      <span class="page-title">商家管理</span>
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>
        添加商家
      </el-button>
    </div>
    <el-table :data="tableData" style="width: 100%" stripe border>
      <el-table-column prop="id" label="ID" width="80" />
      <el-table-column prop="username" label="用户名" />
      <el-table-column prop="name" label="商家名称" />
      <el-table-column prop="contact" label="联系人" />
      <el-table-column prop="phone" label="手机号" width="130" />
      <el-table-column prop="status" label="状态" width="100">
        <template #default="scope">
          <el-tag :type="scope.row.status === 'active' ? 'success' : 'danger'" effect="dark">
            {{ scope.row.status === 'active' ? '正常' : '禁用' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="created_at" label="注册时间" width="180" />
      <el-table-column label="操作" width="200" fixed="right">
        <template #default="scope">
          <el-button type="primary" link size="small">编辑</el-button>
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
import { ElMessage } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'

const tableData = ref([
  { id: 1, username: 'merchant1', name: '绿源农场', contact: '王经理', phone: '13800138011', status: 'active', created_at: '2026-01-10 10:00:00' },
  { id: 2, username: 'merchant2', name: '鲜果园', contact: '李经理', phone: '13800138012', status: 'active', created_at: '2026-01-15 11:00:00' },
  { id: 3, username: 'merchant3', name: '乡村优品', contact: '张经理', phone: '13800138013', status: 'inactive', created_at: '2026-02-01 09:00:00' }
])

const total = ref(15)

const handleAdd = () => {
  ElMessage.info('添加商家功能开发中...')
}
</script>
