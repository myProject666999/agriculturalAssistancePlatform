<template>
  <div class="forum-add-container">
    <van-nav-bar
      title="发布帖子"
      left-arrow
      @click-left="onClickLeft"
    >
      <template #right>
        <van-button type="primary" size="small" round :loading="loading" @click="handleSubmit">发布</van-button>
      </template>
    </van-nav-bar>

    <van-form @submit="onSubmit">
      <van-cell-group inset>
        <van-field
          v-model="form.title"
          label="标题"
          placeholder="请输入帖子标题"
          maxlength="50"
          show-word-limit
          :rules="[{ required: true, message: '请输入帖子标题' }]"
        />
        <van-field
          v-model="form.content"
          label="内容"
          type="textarea"
          placeholder="请输入帖子内容"
          maxlength="500"
          show-word-limit
          :autosize="{ maxHeight: 200 }"
          :rules="[{ required: true, message: '请输入帖子内容' }]"
        />
      </van-cell-group>

      <van-cell-group inset class="tips-section">
        <van-cell title="温馨提示">
          <template #default>
            <div class="tips-content text-gray">
              <p>• 请遵守社区规范，文明发帖</p>
              <p>• 禁止发布违法违规内容</p>
              <p>• 禁止发布广告和垃圾信息</p>
            </div>
          </template>
        </van-cell>
      </van-cell-group>
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
  title: '',
  content: ''
})

const onClickLeft = () => {
  router.back()
}

const handleSubmit = async () => {
  if (!form.title.trim()) {
    showToast('请输入帖子标题')
    return
  }
  if (!form.content.trim()) {
    showToast('请输入帖子内容')
    return
  }

  loading.value = true
  setTimeout(() => {
    loading.value = false
    showToast('发布成功')
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
.forum-add-container {
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

.text-gray {
  color: #999;
}
</style>
