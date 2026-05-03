<template>
  <div class="forum-detail-container">
    <van-nav-bar
      title="帖子详情"
      left-arrow
      @click-left="onClickLeft"
    />

    <div class="post-content" v-if="post">
      <div class="post-header">
        <van-avatar :size="40" icon="user" />
        <div class="user-info">
          <div class="user-name">{{ post.user_name || '匿名用户' }}</div>
          <div class="create-time text-gray">{{ post.created_at }}</div>
        </div>
      </div>
      <div class="post-body">
        <h3 class="post-title">{{ post.title }}</h3>
        <div class="post-text">{{ post.content }}</div>
        <div class="post-footer">
          <span class="count-item text-gray">
            <van-icon name="eye-o" size="14" />
            {{ post.views }} 阅读
          </span>
        </div>
      </div>
    </div>

    <van-divider>评论 ({{ replies.length }})</van-divider>

    <div class="reply-list" v-if="replies.length > 0">
      <div v-for="reply in replies" :key="reply.id" class="reply-item">
        <van-avatar :size="32" icon="user" />
        <div class="reply-content">
          <div class="reply-header">
            <span class="user-name">{{ reply.user_name || '匿名用户' }}</span>
            <span class="create-time text-gray">{{ reply.created_at }}</span>
          </div>
          <div class="reply-text">{{ reply.content }}</div>
        </div>
      </div>
    </div>
    <van-empty v-else description="暂无评论" />

    <div class="reply-bar">
      <van-field
        v-model="replyContent"
        placeholder="发表评论..."
        @focus="checkLogin"
      />
      <van-button type="primary" :disabled="!replyContent.trim()" @click="submitReply">发送</van-button>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { showToast } from 'vant'

const router = useRouter()

const replyContent = ref('')

const post = ref({
  id: 1,
  title: '分享一下我的种菜经验',
  content: '今年种了西红柿、黄瓜和茄子，收成还不错。跟大家分享一下我的种植经验：\n\n1. 土壤准备：选择肥沃、排水良好的土壤，种植前施足基肥。\n\n2. 播种时间：根据不同蔬菜的生长习性选择合适的播种时间。\n\n3. 日常管理：定期浇水、施肥、除草、防治病虫害。\n\n4. 收获时机：及时收获，保持蔬菜的新鲜度和口感。\n\n希望对大家有帮助！',
  user_name: '张三',
  views: 1256,
  created_at: '2026-05-01 10:30'
})

const replies = ref([
  { id: 1, user_name: '李四', content: '非常实用的经验分享，感谢楼主！', created_at: '2026-05-01 11:00' },
  { id: 2, user_name: '王五', content: '请问西红柿什么时候种比较好？', created_at: '2026-05-01 12:30' },
  { id: 3, user_name: '赵六', content: '学到了很多，今年也打算试试种一些蔬菜。', created_at: '2026-05-02 09:15' }
])

const onClickLeft = () => {
  router.back()
}

const checkLogin = () => {
  const token = localStorage.getItem('token')
  if (!token) {
    router.push('/login')
  }
}

const submitReply = () => {
  if (!replyContent.value.trim()) {
    showToast('请输入评论内容')
    return
  }
  showToast('评论发布成功')
  replyContent.value = ''
}
</script>

<style lang="scss" scoped>
.forum-detail-container {
  background-color: #f5f5f5;
  min-height: 100vh;
  padding-bottom: 60px;
}

.post-content {
  background-color: #fff;
  padding: 15px;
}

.post-header {
  display: flex;
  align-items: center;
  margin-bottom: 15px;
}

.user-info {
  margin-left: 12px;
}

.user-name {
  font-size: 14px;
  color: #333;
  font-weight: 500;
}

.create-time {
  font-size: 12px;
  margin-top: 3px;
}

.post-title {
  font-size: 18px;
  font-weight: bold;
  color: #333;
  margin-bottom: 15px;
  line-height: 1.5;
}

.post-text {
  font-size: 15px;
  color: #333;
  line-height: 1.8;
  white-space: pre-wrap;
}

.post-footer {
  margin-top: 15px;
  padding-top: 15px;
  border-top: 1px solid #f5f5f5;
}

.count-item {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 13px;
}

.reply-list {
  background-color: #fff;
  padding: 0 15px;
}

.reply-item {
  display: flex;
  padding: 15px 0;
  border-bottom: 1px solid #f5f5f5;

  &:last-child {
    border-bottom: none;
  }
}

.reply-content {
  flex: 1;
  margin-left: 10px;
}

.reply-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 6px;
}

.reply-text {
  font-size: 14px;
  color: #333;
  line-height: 1.5;
}

.reply-bar {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  display: flex;
  align-items: center;
  padding: 10px 15px;
  background-color: #fff;
  box-shadow: 0 -2px 10px rgba(0, 0, 0, 0.05);

  .van-field {
    flex: 1;
    margin-right: 10px;

    :deep(.van-field__control) {
      background-color: #f5f5f5;
      border-radius: 15px;
      padding: 8px 15px;
    }
  }
}

.text-gray {
  color: #999;
}
</style>
