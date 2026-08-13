import fs from 'node:fs'

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  future: {
    compatibilityVersion: 4,
  },
  experimental: {
    appManifest: false
  },
  devtools: { enabled: true },
  modules: ['@nuxtjs/tailwindcss', 'shadcn-nuxt', '@pinia/nuxt'],
  shadcn: {
    prefix: '',
    componentDir: '@/components/ui'
  },
  vite: {
    vue: {
      script: {
        fs: {
          fileExists(file: string) {
            return fs.existsSync(file)
          },
          readFile(file: string) {
            return fs.readFileSync(file, 'utf-8')
          }
        }
      }
    }
  }
})
