<template>
  <div class="order-detail-container">
    <van-nav-bar
      title="订单详情"
      left-arrow
      @click-left="onClickLeft"
    />
    <div class="order-status-bar">
      <div class="status-info">
        <van-icon name="success" size="24" color="#4caf50" />
        <div class="status-text">
          <h3>{{ getStatusText(order.status) }}</h3>
          <p>{{ getStatusDesc(order.status) }}</p>
        </div>
      </div>
    </div>

    <van-cell-group inset class="address-section">
      <van-cell>
        <template #default>
          <div class="address-info">
            <div class="address-header">
              <span class="name">{{ order.address?.name }}</span>
              <span class="phone">{{ order.address?.phone }}</span>
            </div>
            <div class="address-detail text-gray">{{ order.address?.province }}{{ order.address?.city }}{{ order.address?.district }}{{ order.address?.detail }}</div>
          </div>
        </template>
      </van-cell>
    </van-cell-group>

    <van-cell-group inset class="product-section">
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
    </van-cell-group>

    <van-cell-group inset class="info-section">
      <van-cell title="订单编号" :value="order.order_no" />
      <van-cell title="下单时间" :value="order.created_at" />
      <van-cell title="支付方式" :value="order.pay_type === 'balance' ? '余额支付' : '在线支付'" />
    </van-cell-group>

    <van-cell-group inset class="price-section">
      <van-cell title="商品金额">
        <template #right-icon>
          <span>¥{{ order.goods_amount }}</span>
        </template>
      </van-cell>
      <van-cell title="运费">
        <template #right-icon>
          <span>{{ order.shipping_fee > 0 ? '¥' + order.shipping_fee : '免运费' }}</span>
        </template>
      </van-cell>
      <van-cell title="实付金额">
        <template #right-icon>
          <span class="color-danger font-bold font-16">¥{{ order.total_amount }}</span>
        </template>
      </van-cell>
    </van-cell-group>

    <van-submit-bar
      v-if="showActions"
      :price="order.total_amount * 100"
      button-text="去支付"
      @submit="handlePay"
    />
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { showToast } from 'vant'

const router = useRouter()
const route = useRoute()

const getDefaultImage = () => {
  return 'https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=fresh%20vegetables%20and%20fruits&image_size=square'
}

const order = ref({
  id: 1,
  order_no: 'ORD202605010001',
  status: 'paid',
  total_amount: 158.50,
  goods_amount: 148.50,
  shipping_fee: 10,
  pay_type: 'balance',
  created_at: '2026-05-01 10:30:00',
  address: {
    name: '张三',
    phone: '13800138000',
    province: '北京市',
    city: '北京市',
    district: '朝阳区',
    detail: '某某街道某某小区1号楼101室'
  },
  items: [
    { id: 1, name: '有机蔬菜套餐', price: 68.00, quantity: 2, image: '' },
    { id: 2, name: '新鲜水果礼盒', price: 22.50, quantity: 1, image: '' }
  ]
})

const showActions = computed(() => {
  return order.value.status === 'pending'
})

const onClickLeft = () => {
  router.back()
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

const getStatusDesc = (status) => {
  const descs = {
    pending: '请尽快完成支付',
    paid: '商家正在准备发货',
    shipped: '商品已发货，请注意查收',
    completed: '交易已完成，感谢您的购买',
    cancelled: '订单已取消'
  }
  return descs[status] || ''
}

const handlePay = () => {
  showToast('支付功能开发中...')
}
</script>

<style lang="scss" scoped>
.order-detail-container {
  background-color: #f5f5f5;
  min-height: 100vh;
  padding-bottom: 60px;
}

.order-status-bar {
  background: linear-gradient(135deg, #4caf50 0%, #8bc34a 100%);
  padding: 20px;
  color: #fff;

  .status-info {
    display: flex;
    align-items: center;
  }

  .status-text {
    margin-left: 15px;

    h3 {
      font-size: 18px;
      font-weight: 500;
      margin-bottom: 5px;
    }

    p {
      font-size: 13px;
      opacity: 0.9;
    }
  }
}

.address-section {
  margin-top: 10px;
}

.address-info {
  .address-header {
    margin-bottom: 8px;

    .name {
      font-size: 15px;
      font-weight: bold;
      color: #333;
    }

    .phone {
      margin-left: 15px;
      font-size: 14px;
      color: #666;
    }
  }

  .address-detail {
    font-size: 14px;
    color: #666;
    line-height: 1.5;
  }
}

.product-section {
  margin-top: 10px;

  .product-item {
    padding: 12px 15px;
    border-bottom: 1px solid #f5f5f5;

    &:last-child {
      border-bottom: none;
    }
  }
}

.product-image {
  width: 80px;
  height: 80px;
  border-radius: 4px;
  object-fit: cover;
}

.product-info {
  margin-left: 12px;
}

.product-name {
  font-size: 14px;
  color: #333;
  margin-bottom: 10px;
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
  font-size: 13px;
  color: #999;
}

.info-section,
.price-section {
  margin-top: 10px;

  :deep(.van-cell__value) {
    color: #333;
    font-size: 14px;
  }
}

.price-section {
  :deep(.van-cell__value) {
    .color-danger {
      font-size: 16px;
    }
  }
}

.text-gray {
  color: #999;
}

.color-danger {
  color: #f44336;
}

.font-bold {
  font-weight: bold;
}

.font-16 {
  font-size: 16px;
}
</style>
