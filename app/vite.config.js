import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import{ fileURLToPath, URL } from 'node:url'
import { VitePWA } from 'vite-plugin-pwa'
import svgLoader from 'vite-svg-loader'

// https://vite.dev/config/
export default defineConfig({
  plugins: [
      vue(),
    svgLoader(),
    VitePWA({
      registerType: 'autoUpdate', // 自動更新（またはprompt）
      includeAssets: ['vite.svg',],
      manifest: {
        name: 'PlanArc',
        short_name: 'PlanArc',
        start_url: '/',
        display: 'standalone',
        background_color: '#ffffff',
        theme_color: '#42b883',
        icons: [
          {
            src: 'pwa-192x192.png',
            sizes: '192x192',
            type: 'image/png'
          },
          {
            src: 'pwa-512x512.png',
            sizes: '512x512',
            type: 'image/png'
          }
        ]
      }
    })
  ],

  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    }
  },
  server: {
    host: true, // 外部アクセス許可
    port: 5173,
    allowedHosts: ['planarc.kencode.tech']
  },
  preview: { port: 4173 }
})
