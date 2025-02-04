import { defineConfig } from "vite";
import solid from "vite-plugin-solid";

export default defineConfig({
  plugins: [solid()],
  build: {
    outDir: "../assets/web",
    emptyOutDir: true,
  },
  base: "/web",
  server: {
    port: 3000,
    open: true,
    proxy: {
      "/api": {
        target: "http://localhost:8080",
        changeOrigin: true,
      },
      "/media": {
        target: "http://localhost:8080",
        changeOrigin: true,
      },
      "/thumb": {
        target: "http://localhost:8080",
        changeOrigin: true,
      },
    },
  },
});
