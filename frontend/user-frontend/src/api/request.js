import axios from 'axios'
import { showToast } from 'vant'
import { useUserStore } from '@/store/user'
import router from '@/router'

const request = axios.create({
  baseURL: '',
  timeout: 10000
})

request.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

request.interceptors.response.use(
  (response) => {
    const res = response.data
    if (res.code === 200) {
      return res
    } else if (res.code === 401) {
      showToast('登录已过期，请重新登录')
      const userStore = useUserStore()
      userStore.logout()
      router.push('/login')
      return Promise.reject(new Error(res.message || '登录已过期'))
    } else {
      showToast(res.message || '请求失败')
      return Promise.reject(new Error(res.message || '请求失败'))
    }
  },
  (error) => {
    let message = '网络错误'
    if (error.response) {
      switch (error.response.status) {
        case 400:
          message = '请求错误'
          break
        case 401:
          message = '未授权，请登录'
          break
        case 403:
          message = '拒绝访问'
          break
        case 404:
          message = '请求地址不存在'
          break
        case 500:
          message = '服务器内部错误'
          break
        default:
          message = `连接错误 ${error.response.status}`
      }
    }
    showToast(message)
    return Promise.reject(error)
  }
)

export const get = (url, params) => {
  return request.get(url, { params })
}

export const post = (url, data) => {
  return request.post(url, data)
}

export const put = (url, data) => {
  return request.put(url, data)
}

export const del = (url, data) => {
  return request.delete(url, { data })
}

export default request
