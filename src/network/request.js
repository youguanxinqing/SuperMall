import axios from 'axios'

export function request(config) {
  const instance = axios.create({
    baseURL: "http://127.0.0.1:8081",
    timeout: 5000
  })

  // 1. 拦截请求
  instance.interceptors.request.use(
    config => config, 
    err => {
      console.log(err)
  })

  // 2. 拦截响应
  instance.interceptors.response.use(
    res => res.data, 
    err => {
      console.log(err)
  })

  return instance(config)
}