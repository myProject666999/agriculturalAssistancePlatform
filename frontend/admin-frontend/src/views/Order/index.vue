<template>
  <div class="page-container">
    <div class="page-header">
      <span class="page-title">订单管理</span>
    </div>
    <el-table :data="tableData" style="width: 100%" stripe border>
      <el-table-column prop="id" label="订单ID" width="100" />
      <el-table-column prop="order_no" label="订单号" width="200" />
      <el-table-column prop="user_name" label="用户" />
      <el-table-column prop="total_amount" label="金额" width="120">
        <template #default="scope">
          <span class="color-primary">¥{{ scope.row.total_amount }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="status" label="状态" width="100">
        <template #default="scope">
          <el-tag :type="getStatusType(scope.row.status)" effect="dark">
            {{ getStatusText(scope.row.status) }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="created_at" label="创建时间" width="180" />
      <el-table-column label="操作" width="200" fixed="right">
        <template #default="scope">
          <el-button type="primary" link size="small">详情</el-button>
          <el-button type="warning" link size="small" v-if="scope.row.status === 'paid'">发货</el-button>
          <el-button type="success" link size="small" v-if="scope.row.status === 'shipped'">完成</el-button>
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
  { id: 1, order_no: 'ORD202605010001', user_name: '张三', total_amount: 158.50, status: 'pending', created_at: '2026-05-01 10:30:00' },
  { id: 2, order_no: 'ORD202605010002', user_name: '李四', total_amount: 89.00, status: 'paid', created_at: '2026-05-01 11:20:00' },
  { id: 3, order_no: 'ORD202605020001', user_name: '王五', total_amount: 256.80, status: 'shipped', created_at: '2026-05-02 09:15:00' },
  { id: 4, order_no: 'ORD202605020002', user_name: '赵六', total_amount: 120.00, status: 'completed', created_at: '2026-05-02 14:45:00' },
  { id: 5, order_no: 'ORD202605030001', user_name: '钱七', total_amount: 320.00, status: 'cancelled', created_at: '2026-05-03 08:00:00' }
])

const total = ref(25)

const getStatusType = (status) => {
  const types = {
    pending: 'warning',
    paid: 'primary',
    shipped: 'info',
    completed: 'success',
    cancelled: 'danger'
  }
  return types[status] || 'info'
}

const getStatusText = (status) => {
  const texts = {
    pending: '待付款',
    paid: '待发货',
    shipped: '待收货',
    completed: '已完成',
    cancelled: '已取消'
  }
  return texts[status] || status
}
</script>
