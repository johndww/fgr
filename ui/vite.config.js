// vite.config.js
import vue from '@vitejs/plugin-vue';

export default {
    plugins: [
        vue(),
    ],

    server: {
        host: 'localhost',
        port: 8080,
    },
}