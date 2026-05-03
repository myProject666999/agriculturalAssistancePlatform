<template>
  <div class="message-container">
    <van-nav-bar
      title="留言板"
      left-arrow
      @click-left="onClickLeft"
    />

    <van-form @submit="onSubmit">
      <van-cell-group inset>
        <van-field
          v-model="messageForm.title"
          label="标题"
          placeholder="请输入留言标题"
          maxlength="50"
          :rules="[{ required: true, message: '请输入标题' }]"
        />
        <van-field
          v-model="messageForm.content"
          label="内容"
          type="textarea"
          placeholder="请输入留言内容"
          maxlength="500"
          show-word-limit
          :autosize="{ maxHeight: 120 }"
          :rules="[{ required: true, message: '请输入内容' }]"
        />
      </van-cell-group>

      <div class="submit-section">
        <van-button type="primary" round block :loading="loading" type="submit">
          提交留言
        </van-button>
      </div>
    </van-form>

    <van-divider>我的留言</van-divider>

    <van-pull-refresh v-model="refreshing" @refresh="onRefresh">
      <van-list
        v-model:loading="listLoading"
        :finished="finished"
        finished-text="没有更多了"
        @load="onLoad"
      >
        <div class="message-list" v-if="messageList.length > 0">
          <div v-for="item in messageList" :key="item.id" class="message-item">
            <div class="message-header flex-between">
              <span class="message-title">{{ item.title }}</span>
              <van-tag :type="item.reply ? 'success' : 'warning'" size="small">
                {{ item.reply ? '已回复' : '待回复' }}
              </van-tag>
            </div>
            <div class="message-content text-gray">{{ item.content }}</div>
            <div class="message-footer flex-between">
              <span class="create-time text-gray">{{ item.created_at }}</span>
            </div>
            <div class="message-reply" v-if="item.reply">
              <div class="reply-label">管理员回复：</div>
              <div class="reply-content">{{ item.reply }}</div>
              <div class="reply-time text-gray">{{ item.reply_time }}</div>
            </div>
          </div>
        </div>
        <van-empty v-else description="暂无留言" />
      </van-list>
    </van-pull-refresh>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { showToast, showSuccessToast } from 'vant'

const router = useRouter()

const loading = ref(false)
const refreshing = ref(false)
const listLoading = ref(false)
const finished = ref(false)

const messageForm = reactive({
  title: '',
  content: ''
})

const messageList = ref([
  {
    id: 1,
    title: '关于支付方式的建议',
    content: '建议增加微信支付和支付宝支付，现在只有余额支付不太方便。',
    created_at: '2026-05-01 10:30',
    reply: '感谢您的建议，我们会尽快增加更多支付方式。',
    reply_time: '2026-05-01 11:00'
  },
  {
    id: 2,
    title: '如何成为商家？',
    content: '我想在平台上销售自己种植的蔬菜，请问如何成为商家？',
    created_at: '2026-05-02 14:20',
    reply: '',
    reply_time: ''
  }
])

const onClickLeft = () => {
  router.back()
}

const handleSubmit = async () => {
  if (!messageForm.title.trim()) {
    showToast('请输入留言标题')
    return
  }
  if (!messageForm.content.trim()) {
    showToast('请输入留言内容')
    return
  }

  loading.value = true
  setTimeout(() => {
    loading.value = false
    showSuccessToast('留言提交成功')
    messageForm.title = ''
    messageForm.content = ''
  }, 1000)
}

const onSubmit = () => {
  handleSubmit()
}

const onLoad = () => {
  setTimeout(() => {
    if (refreshing.value) {
      refreshing.value = false
    }
    listLoading.value = false
    finished.value = true
  }, 1000)
}

const onRefresh = () => {
  onLoad()
}
</script>

<style lang="scss" scoped>
.message-container {
  background-color: #f5f5f5;
  min-height: 100vh;
  padding-bottom: 20px;
}

.submit-section {
  padding: 20px;
}

.message-list {
  padding: 0 10px;
}

.message-item {
  background-color: #fff;
  border-radius: 8px;
  padding: 15px;
  margin-bottom: 10px;
}

.message-header {
  margin-bottom: 10px;
}

.message-title {
  font-size: 15px;
  font-weight: 500;
  color: #333;
}

.message-content {
  font-size: 14px;
  line-height: 1.6;
  margin-bottom: 10px;
}

.message-footer {
  padding-top: 10px;
  border-top: 1px solid #f5f5f5;
}

.create-time {
  font-size: 12px;
}

.message-reply {
  margin-top: 12px;
  padding: 12px;
  background-color: #f5f9f5;
  border-radius: 4px;
}

.reply-label {
  font-size: 13px;
  color: #4caf50;
  margin-bottom: 6px;
}

.reply-content {
  font-size: 14px;
  color: #333;
  line-height: 1.6;
}

.reply-time {
  font-size: 12px;
  margin-top: 6px;
}

.text-gray {
  color: #999;
}

.flex-between {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
