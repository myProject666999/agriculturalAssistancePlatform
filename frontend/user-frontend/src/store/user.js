import { defineStore } from 'pinia'
import { ref } from 'vue'
import { get, post } from '@/api/request'

export const useUserStore = defineStore('user', () => {
  const token = ref(localStorage.getItem('token') || '')
  const userInfo = ref(JSON.parse(localStorage.getItem('userInfo') || 'null'))

  const setToken = (newToken) => {
    token.value = newToken
    if (newToken) {
      localStorage.setItem('token', newToken)
    } else {
      localStorage.removeItem('token')
    }
  }

  const setUserInfo = (info) => {
    userInfo.value = info
    if (info) {
      localStorage.setItem('userInfo', JSON.stringify(info))
    } else {
      localStorage.removeItem('userInfo')
    }
  }

  const login = async (username, password) => {
    const res = await post('/api/auth/user/login', { username, password })
    if (res.code === 200) {
      setToken(res.data.token)
      setUserInfo(res.data.user_info)
    }
    return res
  }

  const register = async (username, password, nickname, phone) => {
    const res = await post('/api/auth/user/register', { username, password, nickname, phone })
    return res
  }

  const getUserInfo = async () => {
    const res = await get('/api/user/info')
    if (res.code === 200) {
      setUserInfo(res.data)
    }
    return res
  }

  const updateProfile = async (data) => {
    const res = await post('/api/user/profile', data)
    if (res.code === 200) {
      const infoRes = await getUserInfo()
      if (infoRes.code === 200) {
        setUserInfo(infoRes.data)
      }
    }
    return res
  }

  const updatePassword = async (oldPassword, newPassword) => {
    const res = await post('/api/user/password', { old_password: oldPassword, new_password: newPassword })
    return res
  }

  const logout = () => {
    setToken('')
    setUserInfo(null)
  }

  return {
    token,
    userInfo,
    setToken,
    setUserInfo,
    login,
    register,
    getUserInfo,
    updateProfile,
    updatePassword,
    logout
  }
})
