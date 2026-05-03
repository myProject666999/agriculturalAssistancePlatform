<template>
  <div class="home-container page-container-with-tabbar">
    <van-pull-refresh v-model="refreshing" @refresh="onRefresh">
      <van-list
        v-model:loading="loading"
        :finished="finished"
        finished-text="没有更多了"
        @load="onLoad"
      >
        <div class="header-section">
          <div class="search-bar" @click="goSearch">
            <van-icon name="search" size="16" color="#999" />
            <span>搜索助农商品</span>
          </div>
        </div>

        <van-swipe class="swipe-section" :autoplay="3000" indicator-color="#4caf50">
          <van-swipe-item v-for="banner in banners" :key="banner.id" @click="goBanner(banner)">
            <img :src="banner.image || getDefaultImage()" class="swiper-image" :alt="banner.title" />
          </van-swipe-item>
        </van-swipe>

        <div class="category-section">
          <van-grid :column-num="5" :border="false">
            <van-grid-item icon="newspaper-o" text="新闻资讯" @click="$router.push('/news')" />
            <van-grid-item icon="friends-o" text="论坛" @click="$router.push('/forum')" />
            <van-grid-item icon="bell-o" text="公告" @click="$router.push('/notice')" />
            <van-grid-item icon="envelop-o" text="留言板" @click="goMessage" />
            <van-grid-item icon="hot-o" text="热销商品" @click="goHotProducts" />
          </van-grid>
        </div>

        <div class="section-title" v-if="hotProducts.length > 0">
          <span class="title-text">热销商品</span>
          <span class="more-text" @click="goHotProducts">查看更多</span>
        </div>
        <div class="product-list" v-if="hotProducts.length > 0">
          <van-grid :column-num="2" gutter="10">
            <van-grid-item v-for="product in hotProducts" :key="product.id" @click="goProduct(product.id)">
              <van-card
                :tag="product.is_hot ? '热销' : ''"
                :num="product.stock"
                :price="product.price"
                :desc="product.description"
                :title="product.name"
                :thumb="product.image || getDefaultImage()"
              />
            </van-grid-item>
          </van-grid>
        </div>

        <div class="section-title" v-if="newProducts.length > 0">
          <span class="title-text">新品上市</span>
          <span class="more-text" @click="goNewProducts">查看更多</span>
        </div>
        <div class="product-list" v-if="newProducts.length > 0">
          <van-grid :column-num="2" gutter="10">
            <van-grid-item v-for="product in newProducts" :key="product.id" @click="goProduct(product.id)">
              <van-card
                :tag="product.is_new ? '新品' : ''"
                :num="product.stock"
                :price="product.price"
                :desc="product.description"
                :title="product.name"
                :thumb="product.image || getDefaultImage()"
              />
            </van-grid-item>
          </van-grid>
        </div>

        <div class="section-title" v-if="products.length > 0">
          <span class="title-text">全部商品</span>
        </div>
        <div class="product-list" v-if="products.length > 0">
          <van-grid :column-num="2" gutter="10">
            <van-grid-item v-for="product in products" :key="product.id" @click="goProduct(product.id)">
              <van-card
                :num="product.stock"
                :price="product.price"
                :desc="product.description"
                :title="product.name"
                :thumb="product.image || getDefaultImage()"
              />
            </van-grid-item>
          </van-grid>
        </div>
      </van-list>
    </van-pull-refresh>

    <van-tabbar v-model="active" active-color="#4caf50">
      <van-tabbar-item name="home" icon="home-o">首页</van-tabbar-item>
      <van-tabbar-item name="category" icon="apps-o">分类</van-tabbar-item>
      <van-tabbar-item name="cart" icon="shopping-cart-o">购物车</van-tabbar-item>
      <van-tabbar-item name="user" icon="user-o">我的</van-tabbar-item>
    </van-tabbar>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import { showToast } from 'vant'
import { get } from '@/api/request'

const router = useRouter()

const active = ref('home')
const refreshing = ref(false)
const loading = ref(false)
const finished = ref(false)
const page = ref(1)
const pageSize = 10

const banners = ref([])
const hotProducts = ref([])
const newProducts = ref([])
const products = ref([])

const getDefaultImage = () => {
  return 'https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=agriculture%20product%20fresh%20vegetables%20and%20fruits&image_size=square'
}

const fetchBanners = async () => {
  try {
    const res = await get('/api/banners')
    if (res.code === 200) {
      banners.value = res.data || []
    }
  } catch (error) {
    console.error('Fetch banners error:', error)
  }
}

const fetchHotProducts = async () => {
  try {
    const res = await get('/api/products/hot', { limit: 4 })
    if (res.code === 200) {
      hotProducts.value = res.data || []
    }
  } catch (error) {
    console.error('Fetch hot products error:', error)
  }
}

const fetchNewProducts = async () => {
  try {
    const res = await get('/api/products/new', { limit: 4 })
    if (res.code === 200) {
      newProducts.value = res.data || []
    }
  } catch (error) {
    console.error('Fetch new products error:', error)
  }
}

const fetchProducts = async () => {
  if (loading.value) return
  
  loading.value = true
  
  try {
    const res = await get('/api/products', { page: page.value, page_size: pageSize })
    if (res.code === 200) {
      const newProducts = res.data?.list || []
      if (page.value === 1) {
        products.value = newProducts
      } else {
        products.value = [...products.value, ...newProducts]
      }
      
      if (newProducts.length < pageSize) {
        finished.value = true
      } else {
        page.value++
      }
    }
  } catch (error) {
    console.error('Fetch products error:', error)
  } finally {
    loading.value = false
    refreshing.value = false
  }
}

const onLoad = () => {
  fetchProducts()
}

const onRefresh = async () => {
  page.value = 1
  finished.value = false
  await Promise.all([
    fetchBanners(),
    fetchHotProducts(),
    fetchNewProducts()
  ])
  fetchProducts()
}

const goSearch = () => {
  router.push('/product-list')
}

const goBanner = (banner) => {
  if (banner.product_id) {
    router.push(`/product/${banner.product_id}`)
  } else if (banner.link) {
    window.location.href = banner.link
  }
}

const goProduct = (id) => {
  router.push(`/product/${id}`)
}

const goHotProducts = () => {
  router.push('/product-list?type=hot')
}

const goNewProducts = () => {
  router.push('/product-list?type=new')
}

const goMessage = () => {
  const token = localStorage.getItem('token')
  if (token) {
    router.push('/message')
  } else {
    router.push('/login')
  }
}

onMounted(() => {
  fetchBanners()
  fetchHotProducts()
  fetchNewProducts()
})

watch(active, (val) => {
  switch (val) {
    case 'category':
      router.push('/category')
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
.home-container {
  background-color: #f5f5f5;
  min-height: 100vh;
}

.header-section {
  background: linear-gradient(135deg, #4caf50 0%, #8bc34a 100%);
  padding: 15px 15px 20px;
}

.search-bar {
  display: flex;
  align-items: center;
  background-color: #fff;
  border-radius: 20px;
  padding: 10px 15px;
  font-size: 14px;
  color: #999;

  span {
    margin-left: 8px;
  }
}

.swipe-section {
  margin: -10px 15px 15px;
  border-radius: 10px;
  overflow: hidden;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.category-section {
  background-color: #fff;
  margin-bottom: 10px;
  padding: 15px 0;
}

.product-list {
  padding: 0 10px 20px;
  background-color: #fff;
}

:deep(.van-card__title) {
  font-size: 14px;
  color: #333;
}

:deep(.van-card__desc) {
  font-size: 12px;
  color: #999;
  margin-top: 4px;
}

:deep(.van-card__price) {
  font-size: 16px;
}
</style>
