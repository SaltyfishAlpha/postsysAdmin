# postsys

This template should help get you started developing with Vue 3 in Vite.

## Recommended IDE Setup

[VSCode](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) (and disable Vetur) + [TypeScript Vue Plugin (Volar)](https://marketplace.visualstudio.com/items?itemName=Vue.vscode-typescript-vue-plugin).

## Customize configuration

See [Vite Configuration Reference](https://vitejs.dev/config/).

## Project Setup

```sh
npm install
```

### Compile and Hot-Reload for Development

```sh
npm run dev
```

### Compile and Minify for Production

```sh
npm run build
```

### 前后端分离配置

请修改`vite.config.js`中`proxy`的值



### 后端运行方式
cd到backend文件夹，请确保初始化了golang和mysql环境，安装好golang包依赖。
后端默认运行在http://127.0.0.1:1926上，您可以在backend/app/init.go的StartServer函数中修改这一端口号。

运行后端的方法为
```bash
cd backend && go run main.go 
```
