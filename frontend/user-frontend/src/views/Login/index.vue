<template>
  <div class="login-container">
    <van-nav-bar title="登录" left-arrow @click-left="onClickLeft" />
    <div class="logo-section">
      <div class="logo">
        <van-icon name="shopping-cart" size="60" color="#4caf50" />
      </div>
      <h2>助农平台</h2>
      <p>助力农业，服务三农</p>
    </div>
    <div class="form-section">
      <van-cell-group inset>
        <van-field
          v-model="username"
          name="username"
          label="用户名"
          placeholder="请输入用户名"
          :rules="[{ required: true, message: '请填写用户名' }]"
        />
        <van-field
          v-model="password"
          type="password"
          name="password"
          label="密码"
          placeholder="请输入密码"
          :rules="[{ required: true, message: '请填写密码' }]"
        />
      </van-cell-group>
      <div style="margin: 16px; padding-top: 20px;">
        <van-button type="primary" round block :loading="loading" @click="onSubmit">
          登录
        </van-button>
      </div>
      <div class="link-section">
        <router-link to="/register" class="link">还没有账号？立即注册</router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { showToast, showLoadingToast, closeToast } from 'vant'
import { useUserStore } from '@/store/user'

const router = useRouter()
const route = useRoute()
const userStore = useUserStore()

const username = ref('')
const password = ref('')
const loading = ref(false)

const onClickLeft = () => {
  router.back()
}

const onSubmit = async () => {
  if (!username.value.trim()) {
    showToast('请输入用户名')
    return
  }
  if (!password.value.trim()) {
    showToast('请输入密码')
    return
  }

  loading.value = true
  showLoadingToast({
    message: '登录中...',
    forbidClick: true
  })

  try {
    const res = await userStore.login(username.value, password.value)
    if (res.code === 200) {
      closeToast()
      showToast('登录成功')
      const redirect = route.query.redirect || '/home'
      setTimeout(() => {
        router.replace(redirect)
      }, 1000)
    }
  } catch (error) {
    console.error('Login error:', error)
  } finally {
    loading.value = false
    closeToast()
  }
}
</script>

<style lang="scss" scoped>
.login-container {
  min-height: 100vh;
  background: linear-gradient(135deg, #e8f5e9 0%, #c8e6c9 100%);
}

.logo-section {
  text-align: center;
  padding: 50px 0 40px;

  .logo {
    width: 120px;
    height: 120px;
    background-color: #fff;
    border-radius: 50%;
    display: flex;
    justify-content: center;
    align-items: center;
    margin: 0 auto 20px;
    box-shadow: 0 4px 20px rgba(76, 175, 80, 0.3);
  }

  h2 {
    font-size: 24px;
    font-weight: bold;
    color: #2e7d32;
    margin-bottom: 10px;
  }

  p {
    font-size: 14px;
    color: #666;
  }
}

.form-section {
  padding: 0 20px;
}

.link-section {
  text-align: center;
  margin-top: 20px;

  .link {
    font-size: 14px;
    color: #4caf50;
  }
}
</style>
