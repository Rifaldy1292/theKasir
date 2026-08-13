import { useRuntimeConfig } from '#app'

export const useApi = () => {
  const config = useRuntimeConfig()
  
  // Create a base custom fetcher
  const apiFetch = $fetch.create({
    baseURL: config.public.apiBase || 'http://localhost:8080/api/v1',
    onRequest({ request, options }) {
      // In a real app, attach JWT token from cookies/localStorage
      const token = useCookie('auth_token').value
      if (token) {
        options.headers = options.headers || {}
        // @ts-ignore
        options.headers.Authorization = `Bearer ${token}`
      }
    },
    onResponseError({ response }) {
      // Handle global errors, e.g., redirect to login on 401
      if (response.status === 401) {
        useCookie('auth_token').value = null
        navigateTo('/login')
      }
    }
  })

  return apiFetch
}
