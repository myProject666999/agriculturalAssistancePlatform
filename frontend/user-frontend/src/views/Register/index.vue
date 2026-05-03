<template>
  <div class="register-container">
    <van-nav-bar title="注册" left-arrow @click-left="onClickLeft" />
    <div class="logo-section">
      <div class="logo">
        <van-icon name="add" size="50" color="#4caf50" />
      </div>
      <h2>注册账号</h2>
    </div>
    <div class="form-section">
      <van-cell-group inset>
        <van-field
          v-model="formData.username"
          name="username"
          label="用户名"
          placeholder="请输入用户名"
          :rules="[{ required: true, message: '请填写用户名' }]"
        />
        <van-field
          v-model="formData.nickname"
          name="nickname"
          label="昵称"
          placeholder="请输入昵称"
          :rules="[{ required: true, message: '请填写昵称' }]"
        />
        <van-field
          v-model="formData.phone"
          name="phone"
          label="手机号"
          placeholder="请输入手机号"
          :rules="[{ required: true, message: '请填写手机号' }, { pattern: /^1[3-9]\d{9}$/, message: '手机号格式错误' }]"
        />
        <van-field
          v-model="formData.password"
          type="password"
          name="password"
          label="密码"
          placeholder="请输入密码"
          :rules="[{ required: true, message: '请填写密码' }, { min: 6, message: '密码长度至少6位' }]"
        />
        <van-field
          v-model="formData.confirmPassword"
          type="password"
          name="confirmPassword"
          label="确认密码"
          placeholder="请再次输入密码"
          :rules="[{ required: true, message: '请确认密码' }]"
        />
      </van-cell-group>
      <div style="margin: 16px; padding-top: 20px;">
        <van-button type="primary" round block :loading="loading" @click="onSubmit">
          注册
        </van-button>
      </div>
      <div class="link-section">
        <router-link to="/login" class="link">已有账号？立即登录</router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { showToast, showLoadingToast, closeToast } from 'vant'
import { useUserStore } from '@/store/user'

const router = useRouter()
const userStore = useUserStore()

const formData = reactive({
  username: '',
  nickname: '',
  phone: '',
  password: '',
  confirmPassword: ''
})

const loading = ref(false)

const onClickLeft = () => {
  router.back()
}

const onSubmit = async () => {
  if (!formData.username.trim()) {
    showToast('请输入用户名')
    return
  }
  if (!formData.nickname.trim()) {
    showToast('请输入昵称')
    return
  }
  if (!formData.phone.trim()) {
    showToast('请输入手机号')
    return
  }
  if (!/^1[3-9]\d{9}$/.test(formData.phone.trim())) {
    showToast('手机号格式错误')
    return
  }
  if (!formData.password.trim()) {
    showToast('请输入密码')
    return
  }
  if (formData.password.length < 6) {
    showToast('密码长度至少6位')
    return
  }
  if (formData.password !== formData.confirmPassword) {
    showToast('两次输入的密码不一致')
    return
  }

  loading.value = true
  showLoadingToast({
    message: '注册中...',
    forbidClick: true
  })

  try {
    const res = await userStore.register(
      formData.username,
      formData.password,
      formData.nickname,
      formData.phone
    )
    if (res.code === 200) {
      closeToast()
      showToast('注册成功，请登录')
      setTimeout(() => {
        router.replace('/login')
      }, 1000)
    }
  } catch (error) {
    console.error('Register error:', error)
  } finally {
    loading.value = false
    closeToast()
  }
}
</script>

<style lang="scss" scoped>
.register-container {
  min-height: 100vh;
  background: linear-gradient(135deg, #e8f5e9 0%, #c8e6c9 100%);
}

.logo-section {
  text-align: center;
  padding: 40px 0 30px;

  .logo {
    width: 100px;
    height: 100px;
    background-color: #fff;
    border-radius: 50%;
    display: flex;
    justify-content: center;
    align-items: center;
    margin: 0 auto 15px;
    box-shadow: 0 4px 20px rgba(76, 175, 80, 0.3);
  }

  h2 {
    font-size: 20px;
    font-weight: bold;
    color: #2e7d32;
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
