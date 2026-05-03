<template>
  <div class="page-container">
    <div class="page-header flex-between">
      <span class="page-title">商品管理</span>
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>
        添加商品
      </el-button>
    </div>
    <el-table :data="tableData" style="width: 100%" stripe border>
      <el-table-column prop="id" label="ID" width="80" />
      <el-table-column label="图片" width="100">
        <template #default="scope">
          <img :src="scope.row.image || getDefaultImage()" style="width: 60px; height: 60px; object-fit: cover; border-radius: 4px;" />
        </template>
      </el-table-column>
      <el-table-column prop="name" label="商品名称" min-width="150" />
      <el-table-column prop="category_name" label="分类" width="100" />
      <el-table-column prop="price" label="价格" width="100">
        <template #default="scope">
          <span class="color-danger">¥{{ scope.row.price }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="stock" label="库存" width="80" />
      <el-table-column prop="sales" label="销量" width="80" />
      <el-table-column prop="status" label="状态" width="100">
        <template #default="scope">
          <el-tag :type="scope.row.status === 'active' ? 'success' : 'info'" effect="dark">
            {{ scope.row.status === 'active' ? '上架' : '下架' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="200" fixed="right">
        <template #default="scope">
          <el-button type="primary" link size="small">编辑</el-button>
          <el-button :type="scope.row.status === 'active' ? 'warning' : 'success'" link size="small">
            {{ scope.row.status === 'active' ? '下架' : '上架' }}
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
import { ElMessageBox, ElMessage } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'

const getDefaultImage = () => {
  return 'https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=fresh%20vegetables%20and%20fruits&image_size=square'
}

const tableData = ref([
  { id: 1, name: '有机蔬菜套餐', category_name: '蔬菜', price: 68.00, stock: 100, sales: 156, status: 'active', image: '' },
  { id: 2, name: '新鲜水果礼盒', category_name: '水果', price: 98.00, stock: 50, sales: 128, status: 'active', image: '' },
  { id: 3, name: '农家土鸡蛋', category_name: '肉禽', price: 45.00, stock: 200, sales: 95, status: 'active', image: '' },
  { id: 4, name: '绿色大米', category_name: '粮油', price: 38.00, stock: 300, sales: 82, status: 'active', image: '' },
  { id: 5, name: '蜂蜜礼盒', category_name: '其他', price: 58.00, stock: 0, sales: 67, status: 'inactive', image: '' }
])

const total = ref(35)

const handleAdd = () => {
  ElMessage.info('添加商品功能开发中...')
}
</script>
