<template>
  <router-view />
</template>

<script setup>
import { onMounted } from 'vue'
import { useUserStore } from './store/user'
import { useRoute, useRouter } from 'vue-router'

const userStore = useUserStore()
const route = useRoute()
const router = useRouter()

onMounted(() => {
  const token = localStorage.getItem('token')
  if (token) {
    userStore.setToken(token)
    userStore.getUserInfo().catch(() => {
      localStorage.removeItem('token')
      userStore.setToken('')
      userStore.setUserInfo(null)
      if (route.meta.requiresAuth) {
        router.push('/login')
      }
    })
  }
})
</script>

<style lang="scss">
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  background-color: #f5f5f5;
  color: #333;
}

a {
  text-decoration: none;
  color: inherit;
}

img {
  max-width: 100%;
  display: block;
}

.page-container {
  min-height: 100vh;
  padding-bottom: 50px;
  background-color: #f5f5f5;
}

.page-container-with-tabbar {
  padding-bottom: 50px;
}

.van-nav-bar {
  background: linear-gradient(135deg, #4caf50 0%, #8bc34a 100%);
}

.van-nav-bar__title {
  color: #fff;
  font-weight: 500;
}

.van-nav-bar__text {
  color: #fff;
}

.van-tabbar {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  z-index: 100;
}

.van-tabbar-item--active {
  color: #4caf50;
}

.van-button--primary {
  background: linear-gradient(135deg, #4caf50 0%, #8bc34a 100%);
  border: none;
}

.van-card__price {
  color: #f44336;
}

.van-tag--primary {
  background-color: #4caf50;
  border-color: #4caf50;
}

.loading-container {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 20px;
}

.empty-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 60px 20px;
}

.swiper-image {
  width: 100%;
  height: 180px;
  object-fit: cover;
}

.section-title {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 15px;
  background-color: #fff;
  margin-bottom: 10px;

  .title-text {
    font-size: 16px;
    font-weight: 600;
    color: #333;
  }

  .more-text {
    font-size: 14px;
    color: #999;
  }
}
</style>
