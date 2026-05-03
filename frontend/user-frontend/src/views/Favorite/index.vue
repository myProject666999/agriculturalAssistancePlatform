<template>
  <div class="favorite-container">
    <van-nav-bar
      title="我的收藏"
      left-arrow
      @click-left="onClickLeft"
    />

    <van-pull-refresh v-model="refreshing" @refresh="onRefresh">
      <van-list
        v-model:loading="loading"
        :finished="finished"
        finished-text="没有更多了"
        @load="onLoad"
      >
        <div class="favorite-list" v-if="favoriteList.length > 0">
          <van-grid :column-num="2" gutter="10">
            <van-grid-item
              v-for="item in favoriteList"
              :key="item.id"
              @click="goProduct(item.product_id)"
            >
              <van-card
                :tag="item.is_hot ? '热销' : item.is_new ? '新品' : ''"
                :num="item.stock"
                :price="item.price"
                :desc="item.description"
                :title="item.name"
                :thumb="item.image || getDefaultImage()"
              >
                <template #footer>
                  <van-button size="mini" type="danger" @click.stop="handleRemove(item)">取消收藏</van-button>
                </template>
              </van-card>
            </van-grid-item>
          </van-grid>
        </div>
        <van-empty v-else description="暂无收藏商品" />
      </van-list>
    </van-pull-refresh>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { showConfirmDialog, showToast } from 'vant'

const router = useRouter()

const refreshing = ref(false)
const loading = ref(false)
const finished = ref(false)

const getDefaultImage = () => {
  return 'https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=fresh%20vegetables%20and%20fruits&image_size=square'
}

const favoriteList = ref([
  { id: 1, product_id: 1, name: '有机蔬菜套餐', price: 68.00, stock: 100, sales: 156, description: '新鲜有机蔬菜，营养健康', is_hot: true, is_new: false, image: '' },
  { id: 2, product_id: 2, name: '新鲜水果礼盒', price: 98.00, stock: 50, sales: 128, description: '精选当季水果礼盒装', is_hot: false, is_new: true, image: '' },
  { id: 3, product_id: 3, name: '农家土鸡蛋', price: 45.00, stock: 200, sales: 95, description: '农家散养土鸡蛋', is_hot: false, is_new: false, image: '' }
])

const onClickLeft = () => {
  router.back()
}

const goProduct = (id) => {
  router.push(`/product/${id}`)
}

const handleRemove = (item) => {
  showConfirmDialog({
    title: '提示',
    message: '确定要取消收藏吗？'
  }).then(() => {
    favoriteList.value = favoriteList.value.filter(f => f.id !== item.id)
    showToast('已取消收藏')
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
.favorite-container {
  background-color: #f5f5f5;
  min-height: 100vh;
}

.favorite-list {
  padding: 10px;
}

:deep(.van-card__footer) {
  padding: 8px 0 0;
  border-top: 1px solid #f5f5f5;
  margin-top: 8px;
}
</style>
