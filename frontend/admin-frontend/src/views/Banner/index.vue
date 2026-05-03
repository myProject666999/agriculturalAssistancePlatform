<template>
  <div class="page-container">
    <div class="page-header flex-between">
      <span class="page-title">轮播图管理</span>
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>
        添加轮播图
      </el-button>
    </div>
    <el-table :data="tableData" style="width: 100%" stripe border>
      <el-table-column prop="id" label="ID" width="80" />
      <el-table-column label="图片" width="200">
        <template #default="scope">
          <img :src="scope.row.image || getDefaultImage()" style="width: 180px; height: 80px; object-fit: cover; border-radius: 4px;" />
        </template>
      </el-table-column>
      <el-table-column prop="title" label="标题" min-width="150" />
      <el-table-column prop="sort" label="排序" width="80" />
      <el-table-column prop="status" label="状态" width="100">
        <template #default="scope">
          <el-tag :type="scope.row.status === 'active' ? 'success' : 'info'" effect="dark">
            {{ scope.row.status === 'active' ? '启用' : '禁用' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="250" fixed="right">
        <template #default="scope">
          <el-button type="primary" link size="small">编辑</el-button>
          <el-button :type="scope.row.status === 'active' ? 'warning' : 'success'" link size="small">
            {{ scope.row.status === 'active' ? '禁用' : '启用' }}
          </el-button>
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

const getDefaultImage = () => {
  return 'https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=banner%20agriculture%20farm%20products&image_size=landscape_16_9'
}

const tableData = ref([
  { id: 1, title: '春季蔬菜上市', sort: 1, status: 'active', image: '' },
  { id: 2, title: '新鲜水果限时特惠', sort: 2, status: 'active', image: '' },
  { id: 3, title: '助农活动火热进行中', sort: 3, status: 'inactive', image: '' }
])

const handleAdd = () => {
  ElMessage.info('添加轮播图功能开发中...')
}
</script>
