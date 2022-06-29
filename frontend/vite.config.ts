import { defineConfig } from 'vite'
import * as path from 'path'
import vue from '@vitejs/plugin-vue'
const resolve = path.resolve
// https://vitejs.dev/config/
export default defineConfig({
  server: {
    port: 8530,
  },
  build: {
    target: 'esnext',
  },
  plugins: [vue({ reactivityTransform: true })],
  resolve: {
    //导入时想要省略的扩展名列表。注意，不 建议忽略自定义导入类型的扩展名（例如：.vue），因为它会干扰 IDE 和类型支持。
    alias: [{ find: '@', replacement: resolve(__dirname, './src') }],
  },
})
