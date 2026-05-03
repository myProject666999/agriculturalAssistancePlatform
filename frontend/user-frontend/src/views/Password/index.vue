<template>
  <div class="password-container">
    <van-nav-bar
      title="修改密码"
      left-arrow
      @click-left="onClickLeft"
    />

    <van-form @submit="onSubmit">
      <van-cell-group inset>
        <van-field
          v-model="form.old_password"
          type="password"
          label="原密码"
          placeholder="请输入原密码"
          :rules="[{ required: true, message: '请输入原密码' }]"
        />
        <van-field
          v-model="form.new_password"
          type="password"
          label="新密码"
          placeholder="请输入新密码（6-16位）"
          :rules="[{ required: true, message: '请输入新密码' }, { min: 6, message: '密码长度至少6位' }]"
        />
        <van-field
          v-model="form.confirm_password"
          type="password"
          label="确认密码"
          placeholder="请再次输入新密码"
          :rules="[{ required: true, message: '请确认新密码' }]"
        />
      </van-cell-group>

      <van-cell-group inset class="tips-section">
        <van-cell title="密码规则">
          <template #default>
            <div class="tips-content text-gray">
              <p>• 密码长度为 6-16 位</p>
              <p>• 建议包含字母、数字和特殊字符</p>
              <p>• 请勿使用简单密码</p>
            </div>
          </template>
        </van-cell>
      </van-cell-group>

      <div class="submit-section">
        <van-button type="primary" round block :loading="loading" type="submit">
          确认修改
        </van-button>
      </div>
    </van-form>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { showToast } from 'vant'

const router = useRouter()

const loading = ref(false)

const form = reactive({
  old_password: '',
  new_password: '',
  confirm_password: ''
})

const onClickLeft = () => {
  router.back()
}

const handleSubmit = async () => {
  if (!form.old_password.trim()) {
    showToast('请输入原密码')
    return
  }
  if (!form.new_password.trim()) {
    showToast('请输入新密码')
    return
  }
  if (form.new_password.length < 6) {
    showToast('密码长度至少6位')
    return
  }
  if (form.new_password !== form.confirm_password) {
    showToast('两次输入的密码不一致')
    return
  }

  loading.value = true
  setTimeout(() => {
    loading.value = false
    showToast('密码修改成功')
    setTimeout(() => {
      router.back()
    }, 1000)
  }, 1000)
}

const onSubmit = () => {
  handleSubmit()
}
</script>

<style lang="scss" scoped>
.password-container {
  background-color: #f5f5f5;
  min-height: 100vh;
}

.tips-section {
  margin-top: 10px;
}

.tips-content {
  font-size: 13px;
  line-height: 1.6;

  p {
    margin: 4px 0;
  }
}

.submit-section {
  padding: 20px;
}

.text-gray {
  color: #999;
}
</style>
