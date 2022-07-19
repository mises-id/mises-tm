import axios, { AxiosRequestConfig } from 'axios'

export interface MyResponseType<T = any> {
  code: number
  message: string
  data: T
  pagination?: {
    page_num: number
    page_size: number
    total_page: number
    total_records: number
  }
}

const instance = axios.create({
  baseURL: process.env.VUE_APP_API_EXPLORER,
  timeout: 100000 // request timeout
})
const request = async <T = any>(config: AxiosRequestConfig): Promise<any> => {
  try {
    const { data } = await instance.request<MyResponseType<T>>(config)
    data.code === 0
      ? console.log(data.message) // 成功消息提示
      : console.error(data.message) // 失败消息提示
    if (data.pagination) {
      return data
    }
    return data.data
  } catch (err: any) {
    const message = err.message || 'Network error'
    if (err.response?.data) {
      return await Promise.reject(err.response.data.message)
    }
    return await Promise.reject(message)
  }
}

export default request
