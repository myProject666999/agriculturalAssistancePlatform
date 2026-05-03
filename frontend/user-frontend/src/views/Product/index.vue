<template>
  <div class="product-container">
    <van-nav-bar
      title="商品详情"
      left-arrow
      @click-left="onClickLeft"
      :fixed="true"
      :placeholder="true"
    />
    <van-swipe :autoplay="3000" indicator-color="#4caf50">
      <van-swipe-item v-for="(image, index) in images" :key="index">
        <img :src="image" class="product-image" />
      </van-swipe-item>
    </van-swipe>

    <van-cell-group inset class="product-info">
      <div class="price-row">
        <span class="price color-danger">¥{{ product.price }}</span>
        <span class="stock text-gray">库存 {{ product.stock }} 件</span>
      </div>
      <div class="product-name">{{ product.name }}</div>
      <div class="product-desc text-gray">{{ product.description }}</div>
    </van-cell-group>

    <van-cell-group inset class="product-section" v-if="product.specs">
      <van-cell title="规格">
        <template #right-icon>
          <span class="text-gray">{{ product.specs }}</span>
        </template>
      </van-cell>
      <van-cell title="产地">
        <template #right-icon>
          <span class="text-gray">{{ product.origin || '中国大陆' }}</span>
        </template>
      </van-cell>
    </van-cell-group>

    <van-cell-group inset class="product-section">
      <van-cell title="商品评价" is-link :badge="comments.length">
        <template #right-icon>
          <span class="text-gray">共 {{ comments.length }} 条</span>
        </template>
      </van-cell>
      <div class="comments" v-if="comments.length > 0">
        <div class="comment-item" v-for="comment in comments" :key="comment.id">
          <div class="comment-header flex-between">
            <span class="user-name">{{ comment.user_name || '匿名用户' }}</span>
            <span class="create-time text-gray">{{ comment.created_at }}</span>
          </div>
          <div class="comment-content">{{ comment.content }}</div>
        </div>
      </div>
    </van-cell-group>

    <van-submit-bar :price="product.price * 100">
      <van-button type="warning" round @click="addToCart">加入购物车</van-button>
      <van-button type="danger" round @click="buyNow">立即购买</van-button>
      <van-button type="primary" square icon="star-o" @click="toggleFavorite">
        {{ isFavorite ? '已收藏' : '收藏' }}
      </van-button>
    </van-submit-bar>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { showToast } from 'vant'

const router = useRouter()
const route = useRoute()

const isFavorite = ref(false)

const getDefaultImage = () => {
  return 'https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=fresh%20vegetables%20and%20fruits&image_size=square'
}

const images = ref([
  getDefaultImage(),
  getDefaultImage(),
  getDefaultImage()
])

const product = ref({
  id: 1,
  name: '有机蔬菜套餐',
  price: 68.00,
  stock: 100,
  sales: 156,
  description: '精选当季新鲜有机蔬菜，营养健康，安全放心。包含白菜、萝卜、青菜、番茄等多种蔬菜，满足您的日常需求。',
  specs: '约3kg/份',
  origin: '山东寿光'
})

const comments = ref([
  { id: 1, user_name: '张**', content: '蔬菜很新鲜，包装也很好，下次还会再来买！', created_at: '2026-05-01' },
  { id: 2, user_name: '李**', content: '收到的蔬菜和图片上的一样，质量不错。', created_at: '2026-04-28' }
])

const onClickLeft = () => {
  router.back()
}

const addToCart = () => {
  showToast('已加入购物车')
}

const buyNow = () => {
  const token = localStorage.getItem('token')
  if (token) {
    router.push('/order-confirm')
  } else {
    router.push('/login')
  }
}

const toggleFavorite = () => {
  const token = localStorage.getItem('token')
  if (!token) {
    router.push('/login')
    return
  }
  isFavorite.value = !isFavorite.value
  showToast(isFavorite.value ? '已收藏' : '已取消收藏')
}
</script>

<style lang="scss" scoped>
.product-container {
  background-color: #f5f5f5;
  min-height: 100vh;
  padding-bottom: 50px;
}

.product-image {
  width: 100%;
  height: 300px;
  object-fit: cover;
}

.product-info {
  margin-top: -20px;
  position: relative;
  z-index: 1;
  border-radius: 20px 20px 0 0;
  padding: 15px;
}

.price-row {
  display: flex;
  align-items: baseline;
  margin-bottom: 10px;
}

.price {
  font-size: 24px;
  font-weight: bold;
}

.stock {
  margin-left: 15px;
  font-size: 14px;
}

.product-name {
  font-size: 18px;
  font-weight: bold;
  color: #333;
  margin-bottom: 8px;
}

.product-desc {
  font-size: 14px;
  line-height: 1.6;
}

.product-section {
  margin-top: 10px;
}

.comments {
  padding: 0 15px 15px;
}

.comment-item {
  padding: 10px 0;
  border-bottom: 1px solid #f5f5f5;

  &:last-child {
    border-bottom: none;
  }
}

.comment-header {
  margin-bottom: 8px;
}

.user-name {
  font-size: 14px;
  color: #333;
}

.create-time {
  font-size: 12px;
}

.comment-content {
  font-size: 14px;
  color: #666;
  line-height: 1.6;
}

.text-gray {
  color: #999;
}

:deep(.van-submit-bar__button) {
  margin-left: 8px;
}

:deep(.van-submit-bar__bar) {
  padding: 0 10px;
}
</style>
