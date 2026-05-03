import { defineStore } from 'pinia'
import { ref } from 'vue'
import { get, post } from '@/api/request'

export const useUserStore = defineStore('user', () => {
  const token = ref(localStorage.getItem('admin_token') || '')
  const userInfo = ref(JSON.parse(localStorage.getItem('admin_userInfo') || 'null'))

  const setToken = (newToken) => {
    token.value = newToken
    if (newToken) {
      localStorage.setItem('admin_token', newToken)
    } else {
      localStorage.removeItem('admin_token')
    }
  }

  const setUserInfo = (info) => {
    userInfo.value = info
    if (info) {
      localStorage.setItem('admin_userInfo', JSON.stringify(info))
    } else {
      localStorage.removeItem('admin_userInfo')
    }
  }

  const login = async (username, password, role) => {
    let url = '/api/auth/admin/login'
    if (role === 'merchant') {
      url = '/api/auth/merchant/login'
    }
    
    const res = await post(url, { username, password })
    if (res.code === 200) {
      setToken(res.data.token)
      setUserInfo(res.data.user_info)
    }
    return res
  }

  const getUserInfo = async () => {
    const res = await get('/api/user/info')
    if (res.code === 200) {
      setUserInfo(res.data)
    }
    return res
  }

  const logout = () => {
    setToken('')
    setUserInfo(null)
  }

  const isAdmin = () => {
    return userInfo.value?.role === 'admin'
  }

  const isMerchant = () => {
    return userInfo.value?.role === 'merchant'
  }

  return {
    token,
    userInfo,
    setToken,
    setUserInfo,
    login,
    getUserInfo,
    logout,
    isAdmin,
    isMerchant
  }
})
