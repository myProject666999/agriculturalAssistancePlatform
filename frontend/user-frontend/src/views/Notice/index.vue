<template>
  <div class="notice-container">
    <van-nav-bar
      title="公告通知"
      left-arrow
      @click-left="onClickLeft"
    />

    <van-pull-refresh v-model="refreshing" @refresh="onRefresh">
      <van-list
        v-model:loading="loading"
        :finished="finished"
        finished-text="没有更多了"
        @load="onLoad"
      >
        <div class="notice-list" v-if="noticeList.length > 0">
          <div
            v-for="item in noticeList"
            :key="item.id"
            class="notice-item"
            @click="goDetail(item.id)"
          >
            <div class="notice-header">
              <van-tag v-if="item.type === 'important'" type="danger" size="small">重要</van-tag>
              <van-tag v-else type="primary" size="small">公告</van-tag>
              <span class="notice-title">{{ item.title }}</span>
            </div>
            <div class="notice-content text-ellipsis text-gray">{{ item.content }}</div>
            <div class="notice-footer">
              <span class="create-time text-gray">{{ item.created_at }}</span>
            </div>
          </div>
        </div>
        <van-empty v-else description="暂无公告" />
      </van-list>
    </van-pull-refresh>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

const refreshing = ref(false)
const loading = ref(false)
const finished = ref(false)

const noticeList = ref([
  {
    id: 1,
    title: '关于平台系统升级维护的通知',
    content: '为了给您提供更好的服务体验，我们计划于2026年5月10日凌晨2:00-6:00进行系统升级维护。维护期间，平台将暂停服务，给您带来的不便敬请谅解。',
    type: 'important',
    created_at: '2026-05-01 00:00'
  },
  {
    id: 2,
    title: '助农活动火热进行中',
    content: '为了助力乡村振兴，平台推出助农优惠活动，部分商品低至5折起，欢迎大家踊跃参与，共同助力农业发展！',
    type: 'normal',
    created_at: '2026-04-20 10:00'
  },
  {
    id: 3,
    title: '新用户注册福利公告',
    content: '新用户注册即送10元优惠券，首单满50元即可使用。活动截止日期：2026年6月30日，快来注册体验吧！',
    type: 'normal',
    created_at: '2026-04-15 09:00'
  }
])

const onClickLeft = () => {
  router.back()
}

const goDetail = (id) => {
  router.push(`/notice/${id}`)
}

const onLoad = () => {
  setTimeout(() => {
    if (refreshing.value) {
      refreshing.value = false
    }
    loading.value = false
    finished.value = true
  }, 1000)
}

const onRefresh = () => {
  onLoad()
}
</script>

<style lang="scss" scoped>
.notice-container {
  background-color: #f5f5f5;
  min-height: 100vh;
}

.notice-list {
  padding: 10px;
}

.notice-item {
  background-color: #fff;
  border-radius: 8px;
  padding: 15px;
  margin-bottom: 10px;
}

.notice-header {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
}

.notice-title {
  font-size: 15px;
  font-weight: 500;
  color: #333;
  margin-left: 8px;
  flex: 1;
}

.notice-content {
  font-size: 14px;
  line-height: 1.6;
  margin-bottom: 10px;
}

.notice-footer {
  padding-top: 10px;
  border-top: 1px solid #f5f5f5;
}

.create-time {
  font-size: 12px;
}

.text-gray {
  color: #999;
}
</style>
