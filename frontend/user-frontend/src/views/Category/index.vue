<template>
  <div class="category-container page-container-with-tabbar">
    <van-nav-bar title="商品分类" />
    <div class="category-content">
      <van-sidebar v-model="activeCategory">
        <van-sidebar-item
          v-for="(category, index) in categories"
          :key="category.id"
          :title="category.name"
        />
      </van-sidebar>
      <div class="category-right">
        <div class="section-title">{{ categories[activeCategory]?.name }}</div>
        <van-grid :column-num="3" :border="false">
          <van-grid-item
            v-for="product in products"
            :key="product.id"
            @click="goProduct(product.id)"
          >
            <img :src="product.image || getDefaultImage()" class="product-image" />
            <div class="product-name text-ellipsis">{{ product.name }}</div>
            <div class="product-price color-danger">¥{{ product.price }}</div>
          </van-grid-item>
        </van-grid>
      </div>
    </div>

    <van-tabbar v-model="active" active-color="#4caf50">
      <van-tabbar-item name="home" icon="home-o">首页</van-tabbar-item>
      <van-tabbar-item name="category" icon="apps-o">分类</van-tabbar-item>
      <van-tabbar-item name="cart" icon="shopping-cart-o">购物车</van-tabbar-item>
      <van-tabbar-item name="user" icon="user-o">我的</van-tabbar-item>
    </van-tabbar>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

const active = ref('category')
const activeCategory = ref(0)

const getDefaultImage = () => {
  return 'https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=fresh%20vegetables%20and%20fruits&image_size=square'
}

const categories = ref([
  { id: 1, name: '蔬菜' },
  { id: 2, name: '水果' },
  { id: 3, name: '粮油' },
  { id: 4, name: '肉禽' },
  { id: 5, name: '其他' }
])

const products = ref([
  { id: 1, name: '有机小白菜', price: 8.50, image: '' },
  { id: 2, name: '新鲜番茄', price: 12.00, image: '' },
  { id: 3, name: '有机黄瓜', price: 9.80, image: '' },
  { id: 4, name: '胡萝卜', price: 6.50, image: '' },
  { id: 5, name: '青椒', price: 7.80, image: '' },
  { id: 6, name: '茄子', price: 8.00, image: '' }
])

const goProduct = (id) => {
  router.push(`/product/${id}`)
}

watch(active, (val) => {
  switch (val) {
    case 'home':
      router.push('/home')
      break
    case 'cart':
      const token = localStorage.getItem('token')
      if (token) {
        router.push('/cart')
      } else {
        router.push('/login')
      }
      break
    case 'user':
      const userToken = localStorage.getItem('token')
      if (userToken) {
        router.push('/user')
      } else {
        router.push('/login')
      }
      break
  }
})
</script>

<style lang="scss" scoped>
.category-container {
  background-color: #f5f5f5;
  min-height: 100vh;
}

.category-content {
  display: flex;
  min-height: calc(100vh - 50px);
}

:deep(.van-sidebar) {
  width: 80px;
  background-color: #f7f8fa;
}

.category-right {
  flex: 1;
  padding: 10px;
  background-color: #fff;
}

.section-title {
  font-size: 16px;
  font-weight: bold;
  color: #333;
  margin-bottom: 15px;
}

.product-image {
  width: 70px;
  height: 70px;
  border-radius: 4px;
  object-fit: cover;
}

.product-name {
  font-size: 12px;
  color: #333;
  margin-top: 8px;
  text-align: center;
}

.product-price {
  font-size: 14px;
  font-weight: bold;
  margin-top: 4px;
  text-align: center;
}
</style>
