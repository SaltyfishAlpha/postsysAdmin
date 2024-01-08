<script setup>
import {onMounted, ref, watch} from 'vue';
import http from "@/utils/request.js";
import router from "@/router/index.js";
import { MailOutlined } from "@ant-design/icons-vue";

const collapsed = ref(false);
const selectedKeys = ref(['3']);

const options = {
  '0': 'ERROR',
  '1': '分配单号',
  '2': '快递入库',
  '3': '快递取出',
}
const links = {
  '0': 'noFound',
  '1': 'main',
  '2': 'upload',
  '3': 'find',
}

onMounted(()=>{
  router.push(`/${links[selectedKeys.value]}`)
})

watch(selectedKeys, (newValue, oldValue) => {
  // console.log(selectedKeys.value)
  router.push(`/${links[selectedKeys.value]}`)
})

</script>

<template>
  <a-layout style="min-height: 100vh">

    <a-layout-sider v-model:collapsed="collapsed" collapsible>
      <a-flex align="center" justify="center"><div class="logo"><MailOutlined style="font-size: 25px;margin: 5px"/><p style="margin: 5px">快递管理系统</p></div></a-flex>
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
  color: white;
  font-size: 20px;
  font-weight: 70;
  display: flex;
}

.site-layout .site-layout-background {
  background: #fff;
}
[data-theme='dark'] .site-layout .site-layout-background {
  background: #141414;
}
</style>