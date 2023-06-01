import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import Components from 'unplugin-vue-components/vite';
import { AntDesignVueResolver } from 'unplugin-vue-components/resolvers';
import Icons from 'unplugin-icons/vite'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    Components({
      resolvers: [
        AntDesignVueResolver({
          importStyle: false, // css in js
        }),
      ],
    }),,
    Icons({
      // experimental
      autoInstall: true,
    })
  ],
  server: {
    proxy: {
      '^/v1/.*/ws': {
        target: 'wss://q.vatprc.net',
        changeOrigin: true,
        ws: true,
      },
      '^/v1/.*': {
        target: 'https://q.vatprc.net',
        changeOrigin: true,
      },
    },
  },
})
