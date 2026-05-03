<template>
  <div class="profile-container">
    <van-nav-bar
      title="个人信息"
      left-arrow
      @click-left="onClickLeft"
    >
      <template #right>
        <van-button type="primary" size="small" plain @click="handleSave">保存</van-button>
      </template>
    </van-nav-bar>

    <van-form @submit="onSubmit">
      <van-cell-group inset class="avatar-section">
        <van-cell title="头像" is-link>
          <template #default>
            <van-avatar :size="50" icon="user" />
          </template>
        </van-cell>
      </van-cell-group>

      <van-cell-group inset>
        <van-field
          v-model="form.nickname"
          label="昵称"
          placeholder="请输入昵称"
          :rules="[{ required: true, message: '请输入昵称' }]"
        />
        <van-field
          v-model="form.phone"
          label="手机号"
          placeholder="请输入手机号"
          disabled
        />
        <van-field
          v-model="form.email"
          label="邮箱"
          placeholder="请输入邮箱（选填）"
        />
        <van-cell title="性别" is-link @click="showGenderPicker = true">
          <template #right-icon>
            <span class="text-gray">{{ form.gender || '请选择' }}</span>
          </template>
        </van-cell>
        <van-cell title="生日" is-link @click="showBirthdayPicker = true">
          <template #right-icon>
            <span class="text-gray">{{ form.birthday || '请选择' }}</span>
          </template>
        </van-cell>
      </van-cell-group>

      <div class="submit-section">
        <van-button type="primary" round block :loading="loading" type="submit">
          保存修改
        </van-button>
      </div>
    </van-form>

    <van-action-sheet
      v-model:show="showGenderPicker"
      :actions="genderActions"
      cancel-text="取消"
      @select="onSelectGender"
    />

    <van-popup v-model:show="showBirthdayPicker" position="bottom">
      <van-datetime-picker
        v-model="currentDate"
        type="date"
        title="选择生日"
        :max-date="new Date()"
        @confirm="onConfirmBirthday"
        @cancel="showBirthdayPicker = false"
      />
    </van-popup>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { showToast } from 'vant'

const router = useRouter()

const loading = ref(false)
const showGenderPicker = ref(false)
const showBirthdayPicker = ref(false)
const currentDate = ref(new Date())

const form = reactive({
  nickname: '张三',
  phone: '13800138000',
  email: '',
  gender: '',
  birthday: ''
})

const genderActions = ref([
  { name: '男', value: 'male' },
  { name: '女', value: 'female' },
  { name: '保密', value: 'secret' }
])

const onClickLeft = () => {
  router.back()
}

const onSelectGender = (action) => {
  form.gender = action.name
}

const onConfirmBirthday = (value) => {
  form.birthday = value.toISOString().split('T')[0]
  showBirthdayPicker.value = false
}

const handleSave = async () => {
  if (!form.nickname.trim()) {
    showToast('请输入昵称')
    return
  }

  loading.value = true
  setTimeout(() => {
    loading.value = false
    showToast('保存成功')
    setTimeout(() => {
      router.back()
    }, 1000)
  }, 1000)
}

const onSubmit = () => {
  handleSave()
}
</script>

<style lang="scss" scoped>
.profile-container {
  background-color: #f5f5f5;
  min-height: 100vh;
}

.avatar-section {
  margin-top: 10px;
}

.submit-section {
  padding: 20px;
}

.text-gray {
  color: #999;
}
</style>
