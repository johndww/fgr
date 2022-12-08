// vite.config.js
import vue from '@vitejs/plugin-vue';
import { defineConfig } from 'vite'
import path from 'path'

export default defineConfig({
    plugins: [
        vue(),
    ],

    server: {
        host: 'localhost',
        port: 8080,
    },
    resolve: {
        alias: {
            '@': path.resolve(__dirname, './src'),
            "assets": path.resolve(__dirname, "src/assets"),
            "~assets": path.resolve(__dirname, "/src/assets"),
        }
    }
})