<template>
  <div class="page-container">
    <div class="page-header">
      <span class="page-title">留言板管理</span>
    </div>
    <el-table :data="tableData" style="width: 100%" stripe border>
      <el-table-column prop="id" label="ID" width="80" />
      <el-table-column prop="user_name" label="用户" width="100" />
      <el-table-column prop="content" label="留言内容" min-width="250">
        <template #default="scope">
          <div class="text-ellipsis-2">{{ scope.row.content }}</div>
        </template>
      </el-table-column>
      <el-table-column prop="reply" label="回复" min-width="200">
        <template #default="scope">
          <span v-if="scope.row.reply" class="color-primary">{{ scope.row.reply }}</span>
          <el-tag v-else type="info" effect="dark">待回复</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="created_at" label="留言时间" width="180" />
      <el-table-column label="操作" width="150" fixed="right">
        <template #default="scope">
          <el-button type="primary" link size="small" v-if="!scope.row.reply">回复</el-button>
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
  { id: 1, user_name: '张三', content: '请问如何注册成为商家？我想在平台上销售自己种植的蔬菜。', reply: '您好，商家注册需要联系管理员审核，感谢您的关注！', created_at: '2026-05-01 10:30:00' },
  { id: 2, user_name: '李四', content: '平台的商品质量有保证吗？如何保证购买的商品是新鲜的？', reply: '', created_at: '2026-05-02 14:20:00' },
  { id: 3, user_name: '王五', content: '建议增加更多的支付方式，比如微信支付。', reply: '感谢您的建议，我们会考虑增加更多支付方式。', created_at: '2026-05-03 09:15:00' }
])

const total = ref(15)
</script>
