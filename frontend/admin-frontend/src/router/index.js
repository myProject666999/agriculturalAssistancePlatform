import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from '@/store/user'

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/Login/index.vue'),
    meta: { title: '登录' }
  },
  {
    path: '/',
    redirect: '/dashboard',
    component: () => import('@/layout/index.vue'),
    meta: { requiresAuth: true },
    children: [
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('@/views/Dashboard/index.vue'),
        meta: { title: '数据统计' }
      },
      {
        path: 'orders',
        name: 'Orders',
        component: () => import('@/views/Order/index.vue'),
        meta: { title: '订单管理' }
      },
      {
        path: 'products',
        name: 'Products',
        component: () => import('@/views/Product/index.vue'),
        meta: { title: '商品管理' }
      },
      {
        path: 'categories',
        name: 'Categories',
        component: () => import('@/views/Category/index.vue'),
        meta: { title: '分类管理' }
      },
      {
        path: 'users',
        name: 'Users',
        component: () => import('@/views/User/index.vue'),
        meta: { title: '用户管理', role: 'admin' }
      },
      {
        path: 'merchants',
        name: 'Merchants',
        component: () => import('@/views/Merchant/index.vue'),
        meta: { title: '商家管理', role: 'admin' }
      },
      {
        path: 'admins',
        name: 'Admins',
        component: () => import('@/views/Admin/index.vue'),
        meta: { title: '管理员管理', role: 'admin' }
      },
      {
        path: 'forums',
        name: 'Forums',
        component: () => import('@/views/Forum/index.vue'),
        meta: { title: '论坛管理', role: 'admin' }
      },
      {
        path: 'banners',
        name: 'Banners',
        component: () => import('@/views/Banner/index.vue'),
        meta: { title: '轮播图管理', role: 'admin' }
      },
      {
        path: 'news',
        name: 'News',
        component: () => import('@/views/News/index.vue'),
        meta: { title: '资讯管理', role: 'admin' }
      },
      {
        path: 'messages',
        name: 'Messages',
        component: () => import('@/views/Message/index.vue'),
        meta: { title: '留言板管理', role: 'admin' }
      },
      {
        path: 'notices',
        name: 'Notices',
        component: () => import('@/views/Notice/index.vue'),
        meta: { title: '公告管理', role: 'admin' }
      }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  document.title = to.meta.title ? `${to.meta.title} - 助农平台管理系统` : '助农平台管理系统'
  
  const userStore = useUserStore()
  const token = userStore.token || localStorage.getItem('admin_token')
  const userInfo = userStore.userInfo

  if (to.meta.requiresAuth) {
    if (!token) {
      next('/login')
    } else {
      if (to.meta.role === 'admin' && userInfo && userInfo.role !== 'admin') {
        next('/dashboard')
      } else {
        next()
      }
    }
  } else {
    if (to.path === '/login' && token) {
      next('/dashboard')
    } else {
      next()
    }
  }
})

export default router
