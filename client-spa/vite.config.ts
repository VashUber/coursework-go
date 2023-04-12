import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import { resolve, dirname } from "path";
import { fileURLToPath } from "url";

const __dirname = dirname(fileURLToPath(import.meta.url));

export default defineConfig({
  plugins: [vue()],

  resolve: {
    alias: {
      "~": resolve(__dirname, "./src"),
      types: resolve(__dirname, "./src/shared-types"),
    },
  },
  server: {
    port: 4020,
    strictPort: true,
  },
});
