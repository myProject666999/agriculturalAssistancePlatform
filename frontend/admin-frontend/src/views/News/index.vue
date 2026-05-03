<template>
  <div class="page-container">
    <div class="page-header flex-between">
      <span class="page-title">资讯管理</span>
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>
        添加资讯
      </el-button>
    </div>
    <el-table :data="tableData" style="width: 100%" stripe border>
      <el-table-column prop="id" label="ID" width="80" />
      <el-table-column label="封面" width="120">
        <template #default="scope">
          <img :src="scope.row.cover || getDefaultImage()" style="width: 100px; height: 60px; object-fit: cover; border-radius: 4px;" />
        </template>
      </el-table-column>
      <el-table-column prop="title" label="标题" min-width="200" />
      <el-table-column prop="category" label="分类" width="100" />
      <el-table-column prop="views" label="浏览量" width="80" />
      <el-table-column prop="status" label="状态" width="100">
        <template #default="scope">
          <el-tag :type="scope.row.status === 'active' ? 'success' : 'info'" effect="dark">
            {{ scope.row.status === 'active' ? '发布' : '草稿' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="created_at" label="发布时间" width="180" />
      <el-table-column label="操作" width="250" fixed="right">
        <template #default="scope">
          <el-button type="primary" link size="small">编辑</el-button>
          <el-button type="info" link size="small">详情</el-button>
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

const getDefaultImage = () => {
  return 'https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=news%20agriculture%20farm&image_size=square'
}

const tableData = ref([
  { id: 1, title: '春季蔬菜种植技术指南', category: '种植技术', views: 1256, status: 'active', created_at: '2026-05-01 10:30:00', cover: '' },
  { id: 2, title: '助农政策解读：如何申请农业补贴', category: '政策解读', views: 856, status: 'active', created_at: '2026-05-02 14:20:00', cover: '' },
  { id: 3, title: '市场行情：本周蔬菜价格走势分析', category: '市场行情', views: 2341, status: 'active', created_at: '2026-05-03 09:15:00', cover: '' }
])

const total = ref(35)

const handleAdd = () => {
  ElMessage.info('添加资讯功能开发中...')
}
</script>
