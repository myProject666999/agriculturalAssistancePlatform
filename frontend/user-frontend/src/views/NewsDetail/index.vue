<template>
  <div class="news-detail-container">
    <van-nav-bar
      title="资讯详情"
      left-arrow
      @click-left="onClickLeft"
    />

    <div class="article-content" v-if="article">
      <h1 class="article-title">{{ article.title }}</h1>
      <div class="article-meta">
        <van-tag type="primary" size="small">{{ article.category }}</van-tag>
        <span class="views text-gray">
          <van-icon name="eye-o" size="12" />
          {{ article.views }} 阅读
        </span>
        <span class="create-time text-gray">{{ article.created_at }}</span>
      </div>
      <div class="article-body">
        <img :src="article.cover || getDefaultImage()" class="article-image" v-if="article.cover" />
        <div class="article-text">
          <p>{{ article.content }}</p>
          <p>农业是国民经济的基础，发展现代农业是乡村振兴的重要支撑。近年来，随着科技的发展和政策的扶持，我国农业现代化水平不断提高，农民收入持续增长。</p>
          <p>助农平台致力于搭建农产品产销对接的桥梁，帮助农民拓宽销售渠道，让优质农产品能够直达消费者手中。我们相信，通过互联网+农业的模式，能够有效解决农产品销售难的问题，促进农民增收致富。</p>
          <p>如果您有任何问题或建议，欢迎通过留言板与我们联系。</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()
const route = useRoute()

const getDefaultImage = () => {
  return 'https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=farm%20agriculture%20news&image_size=square'
}

const article = ref({
  id: 1,
  title: '春季蔬菜种植技术指南：如何在家庭菜园种植有机蔬菜',
  category: '种植技术',
  views: 1256,
  created_at: '2026-05-01',
  cover: '',
  content: '春季是蔬菜种植的黄金季节，掌握正确的种植技术是获得丰收的关键。本文将为您详细介绍春季蔬菜种植的各项技术要点。'
})

const onClickLeft = () => {
  router.back()
}
</script>

<style lang="scss" scoped>
.news-detail-container {
  background-color: #fff;
  min-height: 100vh;
}

.article-content {
  padding: 15px;
}

.article-title {
  font-size: 20px;
  font-weight: bold;
  color: #333;
  line-height: 1.5;
  margin-bottom: 15px;
}

.article-meta {
  display: flex;
  align-items: center;
  gap: 12px;
  padding-bottom: 15px;
  border-bottom: 1px solid #f5f5f5;
  margin-bottom: 15px;
  font-size: 12px;
}

.views {
  display: flex;
  align-items: center;
  gap: 3px;
}

.article-image {
  width: 100%;
  height: 200px;
  object-fit: cover;
  border-radius: 8px;
  margin-bottom: 15px;
}

.article-text {
  p {
    font-size: 15px;
    color: #333;
    line-height: 1.8;
    margin-bottom: 15px;
    text-indent: 2em;
  }
}

.text-gray {
  color: #999;
}
</style>
