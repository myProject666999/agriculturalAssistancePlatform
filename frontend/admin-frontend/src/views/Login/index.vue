<template>
  <div class="login-container">
    <div class="login-box">
      <div class="login-header">
        <h2>助农平台管理系统</h2>
        <p>Agricultural Assistance Platform</p>
      </div>
      <el-form
        ref="loginFormRef"
        :model="loginForm"
        :rules="loginRules"
        class="login-form"
        label-width="0px"
      >
        <el-radio-group v-model="loginForm.role" class="role-selector">
          <el-radio value="merchant" border>商家登录</el-radio>
          <el-radio value="admin" border>管理员登录</el-radio>
        </el-radio-group>

        <el-form-item prop="username">
          <el-input
            v-model="loginForm.username"
            placeholder="请输入用户名"
            prefix-icon="User"
            size="large"
          />
        </el-form-item>

        <el-form-item prop="password">
          <el-input
            v-model="loginForm.password"
            type="password"
            placeholder="请输入密码"
            prefix-icon="Lock"
            size="large"
            @keyup.enter="handleLogin"
          />
        </el-form-item>

        <el-form-item>
          <el-button
            type="primary"
            size="large"
            :loading="loading"
            class="login-btn"
            @click="handleLogin"
          >
            登录
          </el-button>
        </el-form-item>
      </el-form>
      <div class="login-tips">
        <p>默认管理员账号: admin / 123456</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/store/user'

const router = useRouter()
const userStore = useUserStore()

const loginFormRef = ref(null)
const loading = ref(false)

const loginForm = reactive({
  username: '',
  password: '',
  role: 'admin'
})

const validateUsername = (rule, value, callback) => {
  if (!value.trim()) {
    callback(new Error('请输入用户名'))
  } else {
    callback()
  }
}

const validatePassword = (rule, value, callback) => {
  if (!value.trim()) {
    callback(new Error('请输入密码'))
  } else {
    callback()
  }
}

const loginRules = {
  username: [{ required: true, validator: validateUsername, trigger: 'blur' }],
  password: [{ required: true, validator: validatePassword, trigger: 'blur' }]
}

const handleLogin = async () => {
  if (!loginFormRef.value) return
  
  await loginFormRef.value.validate(async (valid) => {
    if (valid) {
      loading.value = true
      try {
        const res = await userStore.login(
          loginForm.username,
          loginForm.password,
          loginForm.role
        )
        if (res.code === 200) {
          ElMessage.success('登录成功')
          router.push('/dashboard')
        }
      } catch (error) {
        console.error('Login error:', error)
      } finally {
        loading.value = false
      }
    }
  })
}
</script>

<style lang="scss" scoped>
.login-container {
  width: 100%;
  height: 100vh;
  background: linear-gradient(135deg, #4caf50 0%, #8bc34a 50%, #c8e6c9 100%);
  display: flex;
  justify-content: center;
  align-items: center;
}

.login-box {
  width: 400px;
  background: rgba(255, 255, 255, 0.95);
  border-radius: 10px;
  box-shadow: 0 10px 40px rgba(0, 0, 0, 0.2);
  padding: 40px;
}

.login-header {
  text-align: center;
  margin-bottom: 30px;

  h2 {
    font-size: 24px;
    color: #2e7d32;
    margin-bottom: 10px;
  }

  p {
    font-size: 14px;
    color: #666;
  }
}

.role-selector {
  display: flex;
  justify-content: center;
  margin-bottom: 25px;
}

.login-form {
  .el-form-item {
    margin-bottom: 22px;
  }

  .login-btn {
    width: 100%;
    background: linear-gradient(135deg, #4caf50 0%, #8bc34a 100%);
    border: none;

    &:hover {
      background: linear-gradient(135deg, #43a047 0%, #7cb342 100%);
    }
  }
}

.login-tips {
  text-align: center;
  margin-top: 20px;
  font-size: 12px;
  color: #999;
}
</style>
