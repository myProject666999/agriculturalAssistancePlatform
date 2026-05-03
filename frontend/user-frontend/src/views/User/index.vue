<template>
  <div class="user-container page-container-with-tabbar">
    <div class="user-header" v-if="userStore.userInfo">
      <div class="user-info">
        <van-icon name="user-o" size="50" color="#fff" />
        <div class="info-text">
          <h3>{{ userStore.userInfo.nickname || userStore.userInfo.username }}</h3>
          <p>余额: ¥{{ userStore.userInfo.balance || 0 }}</p>
        </div>
      </div>
      <van-icon name="setting-o" size="22" color="#fff" @click="goSetting" />
    </div>
    
    <div class="user-header guest-header" v-else @click="goLogin">
      <div class="user-info">
        <van-icon name="user-o" size="50" color="#fff" />
        <div class="info-text">
          <h3>点击登录</h3>
          <p>登录后享受更多功能</p>
        </div>
      </div>
    </div>

    <van-cell-group inset class="order-section">
      <van-cell title="我的订单" is-link @click="$router.push('/order')">
        <template #right-icon>
          <span class="view-all">查看全部</span>
        </template>
      </van-cell>
      <div class="order-menu">
        <van-grid :column-num="4" :border="false">
          <van-grid-item text="待付款" @click="goOrder('pending')">
            <van-icon name="todo-list-o" size="24" color="#666" />
            <van-badge :content="0" offset="5, -5" />
          </van-grid-item>
          <van-grid-item text="待发货" @click="goOrder('paid')">
            <van-icon name="logistics" size="24" color="#666" />
            <van-badge :content="0" offset="5, -5" />
          </van-grid-item>
          <van-grid-item text="待收货" @click="goOrder('shipped')">
            <van-icon name="gift-o" size="24" color="#666" />
            <van-badge :content="0" offset="5, -5" />
          </van-grid-item>
          <van-grid-item text="已完成" @click="goOrder('completed')">
            <van-icon name="description" size="24" color="#666" />
          </van-grid-item>
        </van-grid>
      </div>
    </van-cell-group>

    <van-cell-group inset class="menu-section">
      <van-cell title="我的收藏" is-link icon="star-o" @click="$router.push('/favorite')" />
      <van-cell title="收货地址" is-link icon="location-o" @click="$router.push('/address')" />
      <van-cell title="账户充值" is-link icon="wallet-o" @click="$router.push('/recharge')" />
      <van-cell title="我的留言" is-link icon="envelop-o" @click="$router.push('/message')" />
    </van-cell-group>

    <van-cell-group inset class="menu-section">
      <van-cell title="新闻资讯" is-link icon="newspaper-o" @click="$router.push('/news')" />
      <van-cell title="论坛交流" is-link icon="friends-o" @click="$router.push('/forum')" />
      <van-cell title="公告通知" is-link icon="bell-o" @click="$router.push('/notice')" />
    </van-cell-group>

    <van-cell-group inset class="menu-section" v-if="userStore.userInfo">
      <van-cell title="个人信息" is-link icon="user-circle-o" @click="$router.push('/profile')" />
      <van-cell title="修改密码" is-link icon="key-o" @click="$router.push('/password')" />
    </van-cell-group>

    <div class="logout-section" v-if="userStore.userInfo">
      <van-button round block type="default" @click="onLogout">退出登录</van-button>
    </div>

    <van-tabbar v-model="active" active-color="#4caf50">
      <van-tabbar-item name="home" icon="home-o">首页</van-tabbar-item>
      <van-tabbar-item name="category" icon="apps-o">分类</van-tabbar-item>
      <van-tabbar-item name="cart" icon="shopping-cart-o">购物车</van-tabbar-item>
      <van-tabbar-item name="user" icon="user-o">我的</van-tabbar-item>
    </van-tabbar>

    <van-action-sheet
      v-model:show="showActionSheet"
      :actions="actions"
      cancel-text="取消"
      @select="onSelect"
    />
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'
import { useRouter } from 'vue-router'
import { showConfirmDialog, showToast } from 'vant'
import { useUserStore } from '@/store/user'

const router = useRouter()
const userStore = useUserStore()

const active = ref('user')
const showActionSheet = ref(false)
const actions = [
  { name: '个人信息', value: 'profile' },
  { name: '修改密码', value: 'password' }
]

watch(active, (val) => {
  switch (val) {
    case 'home':
      router.push('/home')
      break
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
  }
})

const goLogin = () => {
  router.push('/login')
}

const goSetting = () => {
  showActionSheet.value = true
}

const onSelect = (action) => {
  showActionSheet.value = false
  router.push(`/${action.value}`)
}

const goOrder = (status) => {
  const token = localStorage.getItem('token')
  if (token) {
    router.push(`/order?status=${status}`)
  } else {
    router.push('/login')
  }
}

const onLogout = () => {
  showConfirmDialog({
    title: '提示',
    message: '确定要退出登录吗？'
  })
    .then(() => {
      userStore.logout()
      showToast('已退出登录')
    })
    .catch(() => {
      // on cancel
    })
}
</script>

<style lang="scss" scoped>
.user-container {
  background-color: #f5f5f5;
  min-height: 100vh;
}

.user-header {
  background: linear-gradient(135deg, #4caf50 0%, #8bc34a 100%);
  padding: 40px 20px 30px;
  display: flex;
  justify-content: space-between;
  align-items: flex-start;

  .user-info {
    display: flex;
    align-items: center;

    .info-text {
      margin-left: 15px;
      color: #fff;

      h3 {
        font-size: 18px;
        font-weight: 500;
        margin-bottom: 5px;
      }

      p {
        font-size: 14px;
        opacity: 0.9;
      }
    }
  }
}

.guest-header {
  cursor: pointer;
}

.order-section {
  margin-top: -15px;
  position: relative;
  z-index: 1;
}

.view-all {
  font-size: 14px;
  color: #999;
}

.order-menu {
  padding: 10px 0;
}

.menu-section {
  margin-top: 10px;
}

.logout-section {
  padding: 30px 20px;
}
</style>
