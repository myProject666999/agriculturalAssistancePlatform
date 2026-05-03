<template>
  <div class="order-container">
    <van-nav-bar
      title="我的订单"
      left-arrow
      @click-left="onClickLeft"
    />
    <van-tabs v-model:active="activeStatus" type="card">
      <van-tab title="全部" name="" />
      <van-tab title="待付款" name="pending" />
      <van-tab title="待发货" name="paid" />
      <van-tab title="待收货" name="shipped" />
      <van-tab title="已完成" name="completed" />
    </van-tabs>

    <van-pull-refresh v-model="refreshing" @refresh="onRefresh">
      <van-list
        v-model:loading="loading"
        :finished="finished"
        finished-text="没有更多订单了"
        @load="onLoad"
      >
        <div class="order-list" v-if="orders.length > 0">
          <div
            v-for="order in orders"
            :key="order.id"
            class="order-item"
            @click="goDetail(order.id)"
          >
            <div class="order-header flex-between">
              <span class="order-no">订单号: {{ order.order_no }}</span>
              <van-tag :type="getStatusType(order.status)" effect="dark">
                {{ getStatusText(order.status) }}
              </van-tag>
            </div>
            <div class="order-content">
              <div class="product-item flex" v-for="item in order.items" :key="item.id">
                <img :src="item.image || getDefaultImage()" class="product-image" />
                <div class="product-info flex-1">
                  <div class="product-name text-ellipsis-2">{{ item.name }}</div>
                  <div class="product-price-row flex-between">
                    <span class="price">¥{{ item.price }}</span>
                    <span class="quantity">x{{ item.quantity }}</span>
                  </div>
                </div>
              </div>
            </div>
            <div class="order-footer flex-between">
              <span class="total-text">
                共{{ order.item_count }}件商品 实付:
                <span class="total-price color-danger">¥{{ order.total_amount }}</span>
              </span>
              <div class="order-actions">
                <van-button
                  v-if="order.status === 'pending'"
                  size="small"
                  type="danger"
                  @click.stop="handlePay(order)"
                >去支付</van-button>
                <van-button
                  v-if="order.status === 'pending'"
                  size="small"
                  type="default"
                  @click.stop="handleCancel(order)"
                >取消订单</van-button>
                <van-button
                  v-if="order.status === 'shipped'"
                  size="small"
                  type="primary"
                  @click.stop="handleConfirm(order)"
                >确认收货</van-button>
              </div>
            </div>
          </div>
        </div>
        <van-empty v-else description="暂无订单" />
      </van-list>
    </van-pull-refresh>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { showConfirmDialog, showToast } from 'vant'

const router = useRouter()
const route = useRoute()

const activeStatus = ref(route.query.status || '')
const refreshing = ref(false)
const loading = ref(false)
const finished = ref(false)

const getDefaultImage = () => {
  return 'https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=fresh%20vegetables%20and%20fruits&image_size=square'
}

const orders = ref([
  {
    id: 1,
    order_no: 'ORD202605010001',
    status: 'pending',
    total_amount: 158.50,
    item_count: 2,
    created_at: '2026-05-01 10:30:00',
    items: [
      { id: 1, name: '有机蔬菜套餐', price: 68.00, quantity: 2, image: '' },
      { id: 2, name: '新鲜水果礼盒', price: 22.50, quantity: 1, image: '' }
    ]
  },
  {
    id: 2,
    order_no: 'ORD202605020001',
    status: 'paid',
    total_amount: 98.00,
    item_count: 1,
    created_at: '2026-05-02 14:20:00',
    items: [
      { id: 1, name: '新鲜水果礼盒', price: 98.00, quantity: 1, image: '' }
    ]
  },
  {
    id: 3,
    order_no: 'ORD202605030001',
    status: 'shipped',
    total_amount: 128.00,
    item_count: 3,
    created_at: '2026-05-03 09:15:00',
    items: [
      { id: 1, name: '农家土鸡蛋', price: 45.00, quantity: 2, image: '' },
      { id: 2, name: '绿色大米', price: 38.00, quantity: 1, image: '' }
    ]
  },
  {
    id: 4,
    order_no: 'ORD202604280001',
    status: 'completed',
    total_amount: 58.00,
    item_count: 1,
    created_at: '2026-04-28 16:45:00',
    items: [
      { id: 1, name: '蜂蜜礼盒', price: 58.00, quantity: 1, image: '' }
    ]
  }
])

const onClickLeft = () => {
  router.back()
}

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

const goDetail = (id) => {
  router.push(`/order/${id}`)
}

const handlePay = (order) => {
  showToast('支付功能开发中...')
}

const handleCancel = (order) => {
  showConfirmDialog({
    title: '提示',
    message: '确定要取消该订单吗？'
  }).then(() => {
    showToast('已取消订单')
  }).catch(() => {})
}

const handleConfirm = (order) => {
  showConfirmDialog({
    title: '提示',
    message: '确认已收到商品？'
  }).then(() => {
    showToast('确认收货成功')
  }).catch(() => {})
}

const onLoad = () => {
  setTimeout(() => {
    if (refreshing.value) {
      refreshing.value = false
    }
    loading.value = false
    finished.value = true
  }, 1000)
}

const onRefresh = () => {
  onLoad()
}
</script>

<style lang="scss" scoped>
.order-container {
  background-color: #f5f5f5;
  min-height: 100vh;
}

:deep(.van-tabs__content) {
  padding: 10px;
}

.order-list {
  padding-bottom: 10px;
}

.order-item {
  background-color: #fff;
  border-radius: 8px;
  margin-bottom: 10px;
  padding: 15px;
}

.order-header {
  margin-bottom: 10px;
  padding-bottom: 10px;
  border-bottom: 1px solid #f5f5f5;
}

.order-no {
  font-size: 12px;
  color: #666;
}

.product-item {
  padding: 8px 0;
}

.product-image {
  width: 80px;
  height: 80px;
  border-radius: 4px;
  object-fit: cover;
}

.product-info {
  margin-left: 10px;
}

.product-name {
  font-size: 14px;
  color: #333;
  margin-bottom: 8px;
  line-height: 1.4;
}

.product-price-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.price {
  font-size: 14px;
  color: #f44336;
  font-weight: bold;
}

.quantity {
  font-size: 12px;
  color: #999;
}

.order-footer {
  margin-top: 10px;
  padding-top: 10px;
  border-top: 1px solid #f5f5f5;
}

.total-text {
  font-size: 13px;
  color: #333;
}

.total-price {
  font-size: 16px;
  font-weight: bold;
}

.order-actions {
  :deep(.van-button) {
    margin-left: 8px;
  }
}
</style>
