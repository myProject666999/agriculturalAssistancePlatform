<template>
  <div class="product-list-container">
    <van-nav-bar
      title="商品列表"
      left-arrow
      @click-left="onClickLeft"
    >
      <template #right>
        <van-icon name="search" size="20" @click="showSearch = true" />
      </template>
    </van-nav-bar>

    <van-tabs v-model:active="activeTab" sticky>
      <van-tab title="综合" name="default" />
      <van-tab title="销量" name="sales" />
      <van-tab title="价格" name="price" />
      <van-tab title="最新" name="new" />
    </van-tabs>

    <van-pull-refresh v-model="refreshing" @refresh="onRefresh">
      <van-list
        v-model:loading="loading"
        :finished="finished"
        finished-text="没有更多了"
        @load="onLoad"
      >
        <div class="product-grid">
          <van-grid :column-num="2" gutter="10">
            <van-grid-item
              v-for="product in products"
              :key="product.id"
              @click="goProduct(product.id)"
            >
              <van-card
                :tag="product.is_hot ? '热销' : product.is_new ? '新品' : ''"
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

    <van-popup v-model:show="showSearch" position="top" round>
      <div class="search-popup">
        <div class="search-header">
          <van-search
            v-model="keyword"
            placeholder="搜索商品"
            show-action
            @search="onSearch"
            @cancel="showSearch = false"
          />
        </div>
        <div class="search-history" v-if="historyList.length > 0">
          <div class="history-header flex-between">
            <span>搜索历史</span>
            <van-icon name="delete" size="16" color="#999" @click="clearHistory" />
          </div>
          <div class="history-tags">
            <van-tag
              v-for="(item, index) in historyList"
              :key="index"
              size="medium"
              @click="onHistoryClick(item)"
            >
              {{ item }}
            </van-tag>
          </div>
        </div>
      </div>
    </van-popup>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()
const route = useRoute()

const activeTab = ref('default')
const refreshing = ref(false)
const loading = ref(false)
const finished = ref(false)
const showSearch = ref(false)
const keyword = ref('')
const page = ref(1)
const pageSize = 10

const getDefaultImage = () => {
  return 'https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=fresh%20vegetables%20and%20fruits&image_size=square'
}

const products = ref([
  { id: 1, name: '有机蔬菜套餐', price: 68.00, stock: 100, sales: 156, description: '新鲜有机蔬菜，营养健康', is_hot: true, is_new: false, image: '' },
  { id: 2, name: '新鲜水果礼盒', price: 98.00, stock: 50, sales: 128, description: '精选当季水果礼盒装', is_hot: false, is_new: true, image: '' },
  { id: 3, name: '农家土鸡蛋', price: 45.00, stock: 200, sales: 95, description: '农家散养土鸡蛋', is_hot: false, is_new: false, image: '' },
  { id: 4, name: '绿色大米', price: 38.00, stock: 300, sales: 82, description: '原生态绿色大米', is_hot: true, is_new: false, image: '' },
  { id: 5, name: '蜂蜜礼盒', price: 58.00, stock: 80, sales: 67, description: '纯天然农家蜂蜜', is_hot: false, is_new: true, image: '' },
  { id: 6, name: '有机白菜', price: 8.50, stock: 500, sales: 345, description: '新鲜有机白菜', is_hot: true, is_new: false, image: '' }
])

const historyList = ref(['蔬菜', '水果', '鸡蛋'])

const onClickLeft = () => {
  router.back()
}

const goProduct = (id) => {
  router.push(`/product/${id}`)
}

const onLoad = () => {
  setTimeout(() => {
    if (refreshing.value) {
      products.value = []
      refreshing.value = false
    }

    page.value++
    loading.value = false

    if (page.value > 3) {
      finished.value = true
    }
  }, 1000)
}

const onRefresh = () => {
  page.value = 1
  finished.value = false
  onLoad()
}

const onSearch = () => {
  if (keyword.value.trim()) {
    historyList.value.unshift(keyword.value.trim())
    showSearch.value = false
  }
}

const onHistoryClick = (item) => {
  keyword.value = item
  showSearch.value = false
}

const clearHistory = () => {
  historyList.value = []
}
</script>

<style lang="scss" scoped>
.product-list-container {
  background-color: #f5f5f5;
  min-height: 100vh;
}

.product-grid {
  padding: 10px;
  background-color: #fff;
}

.search-popup {
  background-color: #fff;
  min-height: 300px;
}

.search-header {
  padding: 10px;
  background-color: #f7f8fa;
}

.search-history {
  padding: 15px;
}

.history-header {
  margin-bottom: 10px;
  font-size: 14px;
  color: #333;
}

.history-tags {
  :deep(.van-tag) {
    margin: 5px 10px 5px 0;
  }
}
</style>
