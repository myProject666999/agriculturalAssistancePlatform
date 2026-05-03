<template>
  <div class="news-container">
    <van-nav-bar
      title="新闻资讯"
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
        <div class="news-list" v-if="newsList.length > 0">
          <div
            v-for="item in newsList"
            :key="item.id"
            class="news-item"
            @click="goDetail(item.id)"
          >
            <img :src="item.cover || getDefaultImage()" class="news-image" />
            <div class="news-content">
              <div class="news-title text-ellipsis-2">{{ item.title }}</div>
              <div class="news-desc text-ellipsis text-gray">{{ item.description }}</div>
              <div class="news-meta">
                <van-tag type="primary" size="small">{{ item.category }}</van-tag>
                <span class="views text-gray">
                  <van-icon name="eye-o" size="12" />
                  {{ item.views }}
                </span>
                <span class="create-time text-gray">{{ item.created_at }}</span>
              </div>
            </div>
          </div>
        </div>
        <van-empty v-else description="暂无新闻资讯" />
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

const getDefaultImage = () => {
  return 'https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=farm%20agriculture%20news&image_size=square'
}

const newsList = ref([
  { id: 1, title: '春季蔬菜种植技术指南：如何在家庭菜园种植有机蔬菜', description: '本文将详细介绍春季蔬菜种植的最佳时间、土壤准备、种子选择和日常管理...', category: '种植技术', views: 1256, created_at: '2026-05-01', cover: '' },
  { id: 2, title: '助农政策解读：2026年最新农业补贴申请指南', description: '了解2026年最新的农业补贴政策，包括申请条件、补贴标准和申请流程...', category: '政策解读', views: 856, created_at: '2026-04-28', cover: '' },
  { id: 3, title: '市场行情：本周蔬菜价格走势分析', description: '本周蔬菜市场价格分析，包括白菜、萝卜、番茄等常见蔬菜的价格变动...', category: '市场行情', views: 2341, created_at: '2026-04-25', cover: '' },
  { id: 4, title: '乡村振兴典型案例：某县如何通过电商助农实现增收', description: '本文将分享一个成功的乡村振兴案例，看看当地农民如何通过电商平台...', category: '成功案例', views: 1856, created_at: '2026-04-20', cover: '' }
])

const onClickLeft = () => {
  router.back()
}

const goDetail = (id) => {
  router.push(`/news/${id}`)
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
.news-container {
  background-color: #f5f5f5;
  min-height: 100vh;
}

.news-list {
  padding: 10px;
}

.news-item {
  display: flex;
  background-color: #fff;
  border-radius: 8px;
  padding: 12px;
  margin-bottom: 10px;
}

.news-image {
  width: 120px;
  height: 90px;
  border-radius: 4px;
  object-fit: cover;
  flex-shrink: 0;
}

.news-content {
  flex: 1;
  margin-left: 12px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.news-title {
  font-size: 15px;
  font-weight: 500;
  color: #333;
  line-height: 1.4;
}

.news-desc {
  font-size: 13px;
  margin-top: 6px;
}

.news-meta {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-top: 8px;
  font-size: 12px;
}

.views {
  display: flex;
  align-items: center;
  gap: 3px;
}

.text-gray {
  color: #999;
}
</style>
