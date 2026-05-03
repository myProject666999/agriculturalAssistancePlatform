<template>
  <div class="forum-container">
    <van-nav-bar
      title="论坛交流"
      left-arrow
      @click-left="onClickLeft"
    >
      <template #right>
        <van-icon name="edit" size="20" @click="goAdd" />
      </template>
    </van-nav-bar>

    <van-pull-refresh v-model="refreshing" @refresh="onRefresh">
      <van-list
        v-model:loading="loading"
        :finished="finished"
        finished-text="没有更多了"
        @load="onLoad"
      >
        <div class="forum-list" v-if="forumList.length > 0">
          <div
            v-for="item in forumList"
            :key="item.id"
            class="forum-item"
            @click="goDetail(item.id)"
          >
            <div class="forum-header">
              <van-avatar :size="36" icon="user" />
              <div class="user-info">
                <div class="user-name">{{ item.user_name || '匿名用户' }}</div>
                <div class="create-time text-gray">{{ item.created_at }}</div>
              </div>
            </div>
            <div class="forum-content">
              <div class="forum-title">{{ item.title }}</div>
              <div class="forum-desc text-ellipsis text-gray">{{ item.content }}</div>
            </div>
            <div class="forum-footer">
              <span class="count-item text-gray">
                <van-icon name="eye-o" size="14" />
                {{ item.views }}
              </span>
              <span class="count-item text-gray">
                <van-icon name="chat-o" size="14" />
                {{ item.replies }}
              </span>
              <van-tag v-if="item.is_top" type="danger" size="small">置顶</van-tag>
              <van-tag v-else-if="item.is_essence" type="warning" size="small">精华</van-tag>
            </div>
          </div>
        </div>
        <van-empty v-else description="暂无帖子" />
      </van-list>
    </van-pull-refresh>

    <div class="fab" @click="goAdd">
      <van-icon name="plus" size="24" color="#fff" />
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

const refreshing = ref(false)
const loading = ref(false)
const finished = ref(false)

const forumList = ref([
  { id: 1, title: '分享一下我的种菜经验', content: '今年种了西红柿、黄瓜和茄子，收成还不错，跟大家分享一下经验...', user_name: '张三', views: 1256, replies: 23, created_at: '2026-05-01 10:30', is_top: true, is_essence: false },
  { id: 2, title: '求助！这个蔬菜是什么品种？', content: '在老家种了一些蔬菜，不知道是什么品种，有认识的吗？', user_name: '李四', views: 856, replies: 15, created_at: '2026-05-02 14:20', is_top: false, is_essence: true },
  { id: 3, title: '推荐几个适合家庭种植的蔬菜', content: '推荐几个好种又好吃的蔬菜品种，适合新手种植...', user_name: '王五', views: 2341, replies: 45, created_at: '2026-05-03 09:15', is_top: false, is_essence: false }
])

const onClickLeft = () => {
  router.back()
}

const goAdd = () => {
  const token = localStorage.getItem('token')
  if (token) {
    router.push('/forum/add')
  } else {
    router.push('/login')
  }
}

const goDetail = (id) => {
  router.push(`/forum/${id}`)
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
.forum-container {
  background-color: #f5f5f5;
  min-height: 100vh;
  padding-bottom: 20px;
}

.forum-list {
  padding: 10px;
}

.forum-item {
  background-color: #fff;
  border-radius: 8px;
  padding: 15px;
  margin-bottom: 10px;
}

.forum-header {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
}

.user-info {
  margin-left: 10px;
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

.forum-title {
  font-size: 16px;
  font-weight: 500;
  color: #333;
  margin-bottom: 8px;
}

.forum-desc {
  font-size: 14px;
  line-height: 1.5;
}

.forum-footer {
  display: flex;
  align-items: center;
  gap: 15px;
  margin-top: 12px;
  padding-top: 12px;
  border-top: 1px solid #f5f5f5;
}

.count-item {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 13px;
}

.fab {
  position: fixed;
  right: 20px;
  bottom: 80px;
  width: 50px;
  height: 50px;
  background: linear-gradient(135deg, #4caf50 0%, #8bc34a 100%);
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  box-shadow: 0 4px 15px rgba(76, 175, 80, 0.4);
  z-index: 100;
}

.text-gray {
  color: #999;
}
</style>
