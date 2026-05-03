<template>
  <div class="layout-container">
    <aside class="layout-sidebar">
      <div class="sidebar-header">
        <el-icon :size="20"><Guide /></el-icon>
        <span class="logo-text">助农平台</span>
      </div>
      <div class="sidebar-menu">
        <el-menu
          :default-active="activeMenu"
          router
          background-color="transparent"
          text-color="rgba(255, 255, 255, 0.8)"
          active-text-color="#fff"
        >
          <el-menu-item index="/dashboard">
            <el-icon><DataAnalysis /></el-icon>
            <span>数据统计</span>
          </el-menu-item>
          <el-menu-item index="/orders">
            <el-icon><Document /></el-icon>
            <span>订单管理</span>
          </el-menu-item>
          <el-menu-item index="/products">
            <el-icon><Goods /></el-icon>
            <span>商品管理</span>
          </el-menu-item>
          <el-menu-item index="/categories">
            <el-icon><Menu /></el-icon>
            <span>分类管理</span>
          </el-menu-item>
          
          <template v-if="userStore.isAdmin()">
            <el-sub-menu index="user-manage">
              <template #title>
                <el-icon><User /></el-icon>
                <span>用户管理</span>
              </template>
              <el-menu-item index="/users">用户列表</el-menu-item>
              <el-menu-item index="/merchants">商家列表</el-menu-item>
              <el-menu-item index="/admins">管理员列表</el-menu-item>
            </el-sub-menu>
            
            <el-sub-menu index="content-manage">
              <template #title>
                <el-icon><Picture /></el-icon>
                <span>内容管理</span>
              </template>
              <el-menu-item index="/forums">论坛管理</el-menu-item>
              <el-menu-item index="/banners">轮播图管理</el-menu-item>
              <el-menu-item index="/news">资讯管理</el-menu-item>
              <el-menu-item index="/messages">留言板管理</el-menu-item>
              <el-menu-item index="/notices">公告管理</el-menu-item>
            </el-sub-menu>
          </template>
        </el-menu>
      </div>
    </aside>

    <div class="layout-main">
      <header class="layout-header">
        <div class="header-left">
          <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="{ path: '/dashboard' }">首页</el-breadcrumb-item>
            <el-breadcrumb-item v-if="currentRoute.meta && currentRoute.meta.title">
              {{ currentRoute.meta.title }}
            </el-breadcrumb-item>
          </el-breadcrumb>
        </div>
        <div class="header-right">
          <el-dropdown @command="handleCommand">
            <div class="user-info">
              <el-avatar :size="32" icon="UserFilled" />
              <span class="user-name">{{ userStore.userInfo?.username || userStore.userInfo?.nickname }}</span>
              <el-tag :type="userStore.isAdmin() ? 'success' : 'primary'" size="small" effect="dark">
                {{ userStore.isAdmin() ? '管理员' : '商家' }}
              </el-tag>
            </div>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="profile">个人信息</el-dropdown-item>
                <el-dropdown-item command="logout" divided>退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </header>

      <main class="layout-content">
        <router-view />
      </main>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useUserStore } from '@/store/user'
import { ElMessageBox } from 'element-plus'
import {
  Guide,
  DataAnalysis,
  Document,
  Goods,
  Menu,
  User,
  Picture
} from '@element-plus/icons-vue'

const route = useRoute()
const router = useRouter()
const userStore = useUserStore()

const currentRoute = computed(() => route)
const activeMenu = computed(() => route.path)

const handleCommand = (command) => {
  if (command === 'logout') {
    ElMessageBox.confirm('确定要退出登录吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      userStore.logout()
      router.push('/login')
    }).catch(() => {})
  }
}
</script>
