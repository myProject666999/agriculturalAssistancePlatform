<template>
  <div class="notice-detail-container">
    <van-nav-bar
      title="公告详情"
      left-arrow
      @click-left="onClickLeft"
    />

    <div class="notice-content" v-if="notice">
      <div class="notice-header">
      <div class="notice-tags">
        <van-tag v-if="notice.type === 'important'" type="danger" size="small">重要</van-tag>
        <van-tag v-else type="primary" size="small">公告</van-tag>
      </div>
      <h1 class="notice-title">{{ notice.title }}</h1>
      <div class="notice-meta text-gray">
        <span>{{ notice.created_at }}</span>
      </div>
    </div>
    <div class="notice-body">
      <div class="notice-text">{{ notice.content }}</div>
      <div class="notice-text">
        感谢您对助农平台的关注与支持！我们将继续努力，为您提供更好的服务体验。
      </div>
      <div class="notice-text">
        如果您有任何问题或建议，欢迎通过留言板与我们联系。
      </div>
    </div>
    <div class="notice-footer">
      <div class="divider-line"></div>
      <div class="footer-text text-gray">助农平台运营团队</div>
    </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()
const route = useRoute()

const notice = ref({
  id: 1,
  title: '关于平台系统升级维护的通知',
  content: '为了给您提供更好的服务体验，我们计划于2026年5月10日凌晨2:00-6:00进行系统升级维护。\n\n维护期间，平台将暂停服务，给您带来的不便敬请谅解。\n\n升级内容包括：\n1. 系统性能优化\n2. 新增部分功能\n3. 修复已知问题\n\n请您合理安排使用时间，升级完成后您将获得更好的使用体验。',
  type: 'important',
  created_at: '2026-05-01 00:00:00'
})

const onClickLeft = () => {
  router.back()
}
</script>

<style lang="scss" scoped>
.notice-detail-container {
  background-color: #fff;
  min-height: 100vh;
}

.notice-content {
  padding: 20px;
}

.notice-header {
  margin-bottom: 20px;
  padding-bottom: 20px;
  border-bottom: 1px solid #f5f5f5;
}

.notice-tags {
  margin-bottom: 15px;
}

.notice-title {
  font-size: 20px;
  font-weight: bold;
  color: #333;
  line-height: 1.5;
  margin-bottom: 12px;
}

.notice-meta {
  font-size: 13px;
}

.notice-text {
  font-size: 15px;
  color: #333;
  line-height: 1.8;
  margin-bottom: 15px;
  white-space: pre-wrap;
}

.notice-footer {
  margin-top: 30px;
}

.divider-line {
  height: 1px;
  background: linear-gradient(90deg, transparent, #e5e5e5, transparent);
  margin-bottom: 15px;
}

.footer-text {
  font-size: 13px;
  text-align: right;
}

.text-gray {
  color: #999;
}
</style>
