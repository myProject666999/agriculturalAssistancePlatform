import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from '@/store/user'

const routes = [
  {
    path: '/',
    redirect: '/home'
  },
  {
    path: '/home',
    name: 'Home',
    component: () => import('@/views/Home/index.vue'),
    meta: { title: '首页', keepAlive: true }
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/Login/index.vue'),
    meta: { title: '登录' }
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import('@/views/Register/index.vue'),
    meta: { title: '注册' }
  },
  {
    path: '/category',
    name: 'Category',
    component: () => import('@/views/Category/index.vue'),
    meta: { title: '分类', keepAlive: true }
  },
  {
    path: '/cart',
    name: 'Cart',
    component: () => import('@/views/Cart/index.vue'),
    meta: { title: '购物车', requiresAuth: true }
  },
  {
    path: '/user',
    name: 'User',
    component: () => import('@/views/User/index.vue'),
    meta: { title: '我的', requiresAuth: true, keepAlive: true }
  },
  {
    path: '/product/:id',
    name: 'ProductDetail',
    component: () => import('@/views/Product/index.vue'),
    meta: { title: '商品详情' }
  },
  {
    path: '/product-list',
    name: 'ProductList',
    component: () => import('@/views/ProductList/index.vue'),
    meta: { title: '商品列表' }
  },
  {
    path: '/order',
    name: 'OrderList',
    component: () => import('@/views/Order/index.vue'),
    meta: { title: '订单列表', requiresAuth: true }
  },
  {
    path: '/order/:id',
    name: 'OrderDetail',
    component: () => import('@/views/OrderDetail/index.vue'),
    meta: { title: '订单详情', requiresAuth: true }
  },
  {
    path: '/order-confirm',
    name: 'OrderConfirm',
    component: () => import('@/views/OrderConfirm/index.vue'),
    meta: { title: '确认订单', requiresAuth: true }
  },
  {
    path: '/address',
    name: 'AddressList',
    component: () => import('@/views/Address/index.vue'),
    meta: { title: '收货地址', requiresAuth: true }
  },
  {
    path: '/address/add',
    name: 'AddressAdd',
    component: () => import('@/views/AddressEdit/index.vue'),
    meta: { title: '添加地址', requiresAuth: true }
  },
  {
    path: '/address/edit/:id',
    name: 'AddressEdit',
    component: () => import('@/views/AddressEdit/index.vue'),
    meta: { title: '编辑地址', requiresAuth: true }
  },
  {
    path: '/favorite',
    name: 'Favorite',
    component: () => import('@/views/Favorite/index.vue'),
    meta: { title: '我的收藏', requiresAuth: true }
  },
  {
    path: '/news',
    name: 'NewsList',
    component: () => import('@/views/News/index.vue'),
    meta: { title: '新闻资讯' }
  },
  {
    path: '/news/:id',
    name: 'NewsDetail',
    component: () => import('@/views/NewsDetail/index.vue'),
    meta: { title: '资讯详情' }
  },
  {
    path: '/forum',
    name: 'ForumList',
    component: () => import('@/views/Forum/index.vue'),
    meta: { title: '论坛' }
  },
  {
    path: '/forum/:id',
    name: 'ForumDetail',
    component: () => import('@/views/ForumDetail/index.vue'),
    meta: { title: '帖子详情' }
  },
  {
    path: '/forum/add',
    name: 'ForumAdd',
    component: () => import('@/views/ForumAdd/index.vue'),
    meta: { title: '发布帖子', requiresAuth: true }
  },
  {
    path: '/profile',
    name: 'Profile',
    component: () => import('@/views/Profile/index.vue'),
    meta: { title: '个人信息', requiresAuth: true }
  },
  {
    path: '/password',
    name: 'Password',
    component: () => import('@/views/Password/index.vue'),
    meta: { title: '修改密码', requiresAuth: true }
  },
  {
    path: '/recharge',
    name: 'Recharge',
    component: () => import('@/views/Recharge/index.vue'),
    meta: { title: '充值', requiresAuth: true }
  },
  {
    path: '/message',
    name: 'Message',
    component: () => import('@/views/Message/index.vue'),
    meta: { title: '留言板', requiresAuth: true }
  },
  {
    path: '/notice',
    name: 'NoticeList',
    component: () => import('@/views/Notice/index.vue'),
    meta: { title: '公告' }
  },
  {
    path: '/notice/:id',
    name: 'NoticeDetail',
    component: () => import('@/views/NoticeDetail/index.vue'),
    meta: { title: '公告详情' }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  document.title = to.meta.title || '助农平台'
  
  const userStore = useUserStore()
  const token = userStore.token || localStorage.getItem('token')

  if (to.meta.requiresAuth && !token) {
    next('/login')
  } else {
    next()
  }
})

export default router
