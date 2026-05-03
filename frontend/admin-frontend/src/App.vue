<template>
  <router-view />
</template>

<script setup>
import { onMounted } from 'vue'
import { useUserStore } from './store/user'
import { useRoute, useRouter } from 'vue-router'

const userStore = useUserStore()
const route = useRoute()
const router = useRouter()

onMounted(() => {
  const token = localStorage.getItem('admin_token')
  if (token) {
    userStore.setToken(token)
    userStore.getUserInfo().catch(() => {
      localStorage.removeItem('admin_token')
      userStore.setToken('')
      userStore.setUserInfo(null)
      if (route.meta.requiresAuth) {
        router.push('/login')
      }
    })
  }
})
</script>

<style lang="scss">
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  background-color: #f0f2f5;
  color: #333;
}

a {
  text-decoration: none;
  color: inherit;
}

img {
  max-width: 100%;
  display: block;
}

#app {
  width: 100%;
  min-height: 100vh;
}

.layout-container {
  display: flex;
  height: 100vh;
  overflow: hidden;
}

.layout-sidebar {
  width: 220px;
  background: linear-gradient(180deg, #2e7d32 0%, #1b5e20 100%);
  transition: width 0.3s;
  overflow: hidden;

  .sidebar-header {
    height: 60px;
    display: flex;
    align-items: center;
    justify-content: center;
    color: #fff;
    font-size: 18px;
    font-weight: bold;
    border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  }

  .sidebar-menu {
    padding: 10px 0;
  }
}

.layout-main {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.layout-header {
  height: 60px;
  background: #fff;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.1);
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 20px;
  z-index: 100;

  .header-left {
    display: flex;
    align-items: center;

    .header-title {
      font-size: 18px;
      font-weight: 500;
      color: #333;
    }
  }

  .header-right {
    display: flex;
    align-items: center;

    .user-info {
      display: flex;
      align-items: center;
      cursor: pointer;
    }
  }
}

.layout-content {
  flex: 1;
  padding: 20px;
  overflow-y: auto;
  background-color: #f0f2f5;
}

.page-container {
  background: #fff;
  border-radius: 4px;
  padding: 20px;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;

  .page-title {
    font-size: 18px;
    font-weight: 500;
    color: #333;
  }
}

.el-menu {
  border-right: none;
  background: transparent;
}

.el-menu-item,
.el-sub-menu__title {
  color: rgba(255, 255, 255, 0.8) !important;

  &:hover {
    background-color: rgba(255, 255, 255, 0.1) !important;
  }
}

.el-menu-item.is-active {
  background-color: rgba(255, 255, 255, 0.2) !important;
  color: #fff !important;
}

.el-sub-menu.is-active > .el-sub-menu__title {
  color: #fff !important;
}

.el-table {
  margin-top: 15px;
}

.pagination-container {
  display: flex;
  justify-content: flex-end;
  margin-top: 20px;
}

.search-form {
  background: #f5f7fa;
  padding: 20px;
  border-radius: 4px;
  margin-bottom: 20px;
}

.operate-btn {
  margin-right: 10px;
}

.status-tag {
  &.status-active {
    background-color: #e1f3d8;
    color: #67c23a;
  }

  &.status-inactive {
    background-color: #fde2e2;
    color: #f56c6c;
  }

  &.status-pending {
    background-color: #fdf6ec;
    color: #e6a23c;
  }
}

.chart-container {
  width: 100%;
  height: 300px;
}

.stats-card {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: #fff;
  border-radius: 8px;
  padding: 20px;
  margin-bottom: 20px;

  &.card-green {
    background: linear-gradient(135deg, #11998e 0%, #38ef7d 100%);
  }

  &.card-blue {
    background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);
  }

  &.card-orange {
    background: linear-gradient(135deg, #fa709a 0%, #fee140 100%);
  }

  .card-title {
    font-size: 14px;
    opacity: 0.9;
    margin-bottom: 10px;
  }

  .card-value {
    font-size: 32px;
    font-weight: bold;
  }
}

.dialog-footer {
  text-align: right;
}
</style>
