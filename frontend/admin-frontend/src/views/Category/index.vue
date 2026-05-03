<template>
  <div class="page-container">
    <div class="page-header flex-between">
      <span class="page-title">分类管理</span>
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>
        添加分类
      </el-button>
    </div>
    <el-table :data="tableData" style="width: 100%" stripe border row-key="id">
      <el-table-column prop="id" label="ID" width="80" />
      <el-table-column label="图标" width="100">
        <template #default="scope">
          <el-tag>{{ scope.row.icon || '默认' }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="name" label="分类名称" />
      <el-table-column prop="sort" label="排序" width="100" />
      <el-table-column prop="created_at" label="创建时间" width="180" />
      <el-table-column label="操作" width="200" fixed="right">
        <template #default="scope">
          <el-button type="primary" link size="small">编辑</el-button>
          <el-button type="danger" link size="small">删除</el-button>
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
  { id: 1, name: '蔬菜', icon: '🥬', sort: 1, created_at: '2026-01-01 10:00:00' },
  { id: 2, name: '水果', icon: '🍎', sort: 2, created_at: '2026-01-01 10:00:00' },
  { id: 3, name: '粮油', icon: '🌾', sort: 3, created_at: '2026-01-01 10:00:00' },
  { id: 4, name: '肉禽', icon: '🥩', sort: 4, created_at: '2026-01-01 10:00:00' },
  { id: 5, name: '其他', icon: '📦', sort: 5, created_at: '2026-01-01 10:00:00' }
])

const handleAdd = () => {
  ElMessage.info('添加分类功能开发中...')
}
</script>
