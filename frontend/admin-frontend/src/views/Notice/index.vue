<template>
  <div class="page-container">
    <div class="page-header flex-between">
      <span class="page-title">公告管理</span>
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>
        添加公告
      </el-button>
    </div>
    <el-table :data="tableData" style="width: 100%" stripe border>
      <el-table-column prop="id" label="ID" width="80" />
      <el-table-column prop="title" label="标题" min-width="250" />
      <el-table-column prop="type" label="类型" width="100">
        <template #default="scope">
          <el-tag :type="scope.row.type === 'important' ? 'danger' : 'primary'" effect="dark">
            {{ scope.row.type === 'important' ? '重要' : '普通' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="status" label="状态" width="100">
        <template #default="scope">
          <el-tag :type="scope.row.status === 'active' ? 'success' : 'info'" effect="dark">
            {{ scope.row.status === 'active' ? '显示' : '隐藏' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="created_at" label="发布时间" width="180" />
      <el-table-column label="操作" width="250" fixed="right">
        <template #default="scope">
          <el-button type="primary" link size="small">编辑</el-button>
          <el-button :type="scope.row.status === 'active' ? 'warning' : 'success'" link size="small">
            {{ scope.row.status === 'active' ? '隐藏' : '显示' }}
          </el-button>
          <el-button type="danger" link size="small">删除</el-button>
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
  { id: 1, title: '关于平台系统升级维护的通知', type: 'important', status: 'active', created_at: '2026-05-01 00:00:00' },
  { id: 2, title: '助农活动火热进行中，快来参与吧！', type: 'normal', status: 'active', created_at: '2026-04-20 10:00:00' },
  { id: 3, title: '新用户注册福利公告', type: 'normal', status: 'active', created_at: '2026-04-15 09:00:00' }
])

const total = ref(8)

const handleAdd = () => {
  ElMessage.info('添加公告功能开发中...')
}
</script>
