import { defineConfig } from 'vite'
import preact from '@preact/preset-vite'
import tailwindcss from '@tailwindcss/vite';

//@ts-ignore
import path from 'path';
//@ts-ignore
import { fileURLToPath } from 'url';

//@ts-ignore
const __dirname = path.dirname(fileURLToPath(import.meta.url));

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    preact(),
    tailwindcss()
  ],
  resolve: {
    alias: {
      '@wails': path.resolve("./src/lib/wailsjs/"),
      '@components': path.resolve("./src/components/"),
      '@utils': path.resolve("./src/utils/"),
    },
  },
})
