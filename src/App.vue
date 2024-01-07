<script setup>
import { ref } from 'vue';
import http from "@/utils/request.js";
const collapsed = ref(false);
const selectedKeys = ref(['1']);

const options = {
  '0': 'ERROR',
  '1': '分配单号',
  '2': '快递入库',
  '3': '衣物',
  '4': '数码产品',
  '5': '药品',
  '6': '生鲜',
  '7': '易碎品',
  '8': '其他',
}

</script>

<template>
  <a-layout style="min-height: 100vh">

    <a-layout-sider v-model:collapsed="collapsed" collapsible>
      <div class="logo"></div>
      <a-menu v-model:selectedKeys="selectedKeys" theme="dark" mode="inline">
        <a-menu-item key="1">
          <pie-chart-outlined />
          <span>{{ options['1'] }}</span>
        </a-menu-item>
        <a-menu-item key="2">
          <desktop-outlined />
          <span>{{ options['2'] }}</span>
        </a-menu-item>
        <a-menu-item key="3">
          <desktop-outlined />
          <span>{{ options['3'] }}</span>
        </a-menu-item>
        <a-menu-item key="4">
          <desktop-outlined />
          <span>{{ options['4'] }}</span>
        </a-menu-item>
      </a-menu>
    </a-layout-sider>

    <a-layout>

      <a-layout-content style="margin: 50px">
        <router-view></router-view>
      </a-layout-content>

      <a-layout-footer style="text-align: center">
        <a @click="()=>http.get('/ping')">SaltyfishAlpha Design ©2023 Created by zjh/cd/thk/zyf</a>
      </a-layout-footer>
    </a-layout>

  </a-layout>
</template>

<style scoped>
.logo {
  height: 32px;
  margin: 16px;
  background: rgba(255, 255, 255, 0.3);
}

.site-layout .site-layout-background {
  background: #fff;
}
[data-theme='dark'] .site-layout .site-layout-background {
  background: #141414;
}
</style>