<template>
  <div class="cart-container page-container-with-tabbar">
    <van-nav-bar title="购物车" />
    <van-empty v-if="cartList.length === 0" description="购物车空空如也" />
    <div v-else class="cart-content">
      <van-checkbox-group v-model="checkedIds">
        <van-cell-group inset>
          <div
            v-for="item in cartList"
            :key="item.id"
            class="cart-item"
          >
            <div class="item-left">
              <van-checkbox :name="item.id" shape="square" />
            </div>
            <img :src="item.image || getDefaultImage()" class="item-image" @click="goProduct(item.product_id)" />
            <div class="item-info">
              <div class="item-name text-ellipsis-2" @click="goProduct(item.product_id)">{{ item.name }}</div>
              <div class="item-price-row flex-between">
                <span class="item-price color-danger">¥{{ item.price }}</span>
                <van-stepper
                  v-model="item.quantity"
                  :min="1"
                  :max="100"
                  size="small"
                />
              </div>
            </div>
            <van-icon name="delete" size="18" color="#999" class="item-delete" @click="handleDelete(item.id)" />
          </div>
        </van-cell-group>
      </van-checkbox-group>
    </div>

    <van-submit-bar
      :price="totalPrice * 100"
      button-text="去结算"
      :disabled="checkedIds.length === 0"
      @submit="goCheckout"
    >
      <van-checkbox v-model="checkedAll" shape="square">全选</van-checkbox>
    </van-submit-bar>

    <van-tabbar v-model="active" active-color="#4caf50">
      <van-tabbar-item name="home" icon="home-o">首页</van-tabbar-item>
      <van-tabbar-item name="category" icon="apps-o">分类</van-tabbar-item>
      <van-tabbar-item name="cart" icon="shopping-cart-o">购物车</van-tabbar-item>
      <van-tabbar-item name="user" icon="user-o">我的</van-tabbar-item>
    </van-tabbar>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { useRouter } from 'vue-router'
import { showConfirmDialog, showToast } from 'vant'

const router = useRouter()

const active = ref('cart')
const checkedIds = ref([])

const getDefaultImage = () => {
  return 'https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=fresh%20vegetables%20and%20fruits&image_size=square'
}

const cartList = ref([
  { id: 1, product_id: 1, name: '有机蔬菜套餐', price: 68.00, quantity: 2, image: '' },
  { id: 2, product_id: 2, name: '新鲜水果礼盒', price: 98.00, quantity: 1, image: '' }
])

const checkedAll = computed({
  get: () => cartList.value.length > 0 && checkedIds.value.length === cartList.value.length,
  set: (val) => {
    checkedIds.value = val ? cartList.value.map(item => item.id) : []
  }
})

const totalPrice = computed(() => {
  return cartList.value
    .filter(item => checkedIds.value.includes(item.id))
    .reduce((sum, item) => sum + item.price * item.quantity, 0)
})

const goProduct = (id) => {
  router.push(`/product/${id}`)
}

const handleDelete = (id) => {
  showConfirmDialog({
    title: '提示',
    message: '确定要删除该商品吗？'
  }).then(() => {
    cartList.value = cartList.value.filter(item => item.id !== id)
    checkedIds.value = checkedIds.value.filter(cid => cid !== id)
    showToast('已删除')
  }).catch(() => {})
}

const goCheckout = () => {
  if (checkedIds.value.length === 0) {
    showToast('请选择要结算的商品')
    return
  }
  router.push('/order-confirm')
}

watch(active, (val) => {
  switch (val) {
    case 'home':
      router.push('/home')
      break
    case 'category':
      router.push('/category')
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
.cart-container {
  background-color: #f5f5f5;
  min-height: 100vh;
  padding-bottom: 100px;
}

.cart-content {
  padding: 10px 0;
}

.cart-item {
  display: flex;
  align-items: center;
  padding: 15px;
  border-bottom: 1px solid #eee;

  .item-left {
    margin-right: 10px;
  }

  .item-image {
    width: 80px;
    height: 80px;
    border-radius: 4px;
    object-fit: cover;
  }

  .item-info {
    flex: 1;
    margin-left: 10px;
  }

  .item-name {
    font-size: 14px;
    color: #333;
    margin-bottom: 10px;
  }

  .item-price-row {
    display: flex;
    align-items: center;
  }

  .item-price {
    font-size: 16px;
    font-weight: bold;
  }

  .item-delete {
    margin-left: 10px;
  }
}

:deep(.van-stepper__minus),
:deep(.van-stepper__plus) {
  background-color: #f5f5f5;
  border-color: #ddd;
}

:deep(.van-stepper__input) {
  border-color: #ddd;
}
</style>
