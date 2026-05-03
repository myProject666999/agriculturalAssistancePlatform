<template>
  <div class="page-container">
    <div class="page-header">
      <span class="page-title">论坛管理</span>
    </div>
    <el-table :data="tableData" style="width: 100%" stripe border>
      <el-table-column prop="id" label="ID" width="80" />
      <el-table-column prop="title" label="标题" min-width="200" />
      <el-table-column prop="user_name" label="作者" width="100" />
      <el-table-column prop="views" label="浏览量" width="80" />
      <el-table-column prop="replies" label="回复数" width="80" />
      <el-table-column prop="status" label="状态" width="100">
        <template #default="scope">
          <el-tag :type="scope.row.status === 'active' ? 'success' : 'danger'" effect="dark">
            {{ scope.row.status === 'active' ? '正常' : '隐藏' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="created_at" label="发布时间" width="180" />
      <el-table-column label="操作" width="200" fixed="right">
        <template #default="scope">
          <el-button type="primary" link size="small">详情</el-button>
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

const tableData = ref([
  { id: 1, title: '分享一下我的种菜经验', user_name: '张三', views: 1256, replies: 23, status: 'active', created_at: '2026-05-01 10:30:00' },
  { id: 2, title: '求助！这个蔬菜是什么品种？', user_name: '李四', views: 856, replies: 15, status: 'active', created_at: '2026-05-02 14:20:00' },
  { id: 3, title: '推荐几个适合家庭种植的蔬菜', user_name: '王五', views: 2341, replies: 45, status: 'active', created_at: '2026-05-03 09:15:00' }
])

const total = ref(25)
</script>
