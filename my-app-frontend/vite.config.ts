import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
// @ts-ignore
export default defineConfig({
  mode: "production",
  base: "/",
  plugins: [vue()],
  server: {
    host: "0.0.0.0"
  }
})
