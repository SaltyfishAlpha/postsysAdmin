<script setup>
import {computed, ref, watch} from "vue";
import http from "@/utils/request";
import myEnum from "@/main";
import {usePagination, useRequest} from "vue-request";
import axios from "axios";

const columns = [
  {
    name: '快递公司',
    dataIndex: 'company',
    key: 'company',
    width: 130,
  },
  {
    title: '是否到付',
    dataIndex: 'cash_on',
    key: 'cash_on',
    width: 100,
  },
  {
    title: 'pre_freight',
    dataIndex: 'pre_freight',
    key: 'pre_freight',
    width: 100,
  },
  {
    title: '快递单号',
    key: 'tracing_num',
    dataIndex: 'tracing_num',
    width: 200,
  },
  {
    title: 'registered',
    key: 'registered',
    dataIndex: 'registered',
    width: 100,
  },
  {
    title: '收件人电话号码',
    key: 'receiver_phone',
    dataIndex: 'receiver_phone',
    width: 150,
  },
  {
    title: '货架号',
    key: 'shelf_num',
    dataIndex: 'shelf_num',
    width: 200,
  },
  {
    title: '交付时间',
    key: 'deliver_time',
    dataIndex: 'deliver_time',
    width: 200,
  },
  {
    title: 'rejected',
    key: 'rejected',
    dataIndex: 'rejected',
    width: 100,
  },
  {},
  {
    title: '操作',
    key: 'action',
    fixed: 'right',
    width: 130,
  },
];

const testData = [
  {
    company: "char(50),",
    cash_on: false,
    pre_freight: false,
    tracing_num: "char(21)",
    registered: false,
    receiver_phone: "char(12)",
    shelf_num: "-1",
    deliver_time: "2023-01-08:03:18:00",
    rejected: false,
  },
];

const lists = ref(testData)
const loading = ref(false)
const take_object = (tracing_num, receiver_phone) => {
  console.log({tracing_num: tracing_num.value, receiver_phone: receiver_phone.value})
  loading.value = true
  http.get('/remove', {tracing_num: tracing_num.value, receiver_phone: receiver_phone.value}).then((res) => {
    console.log(res.data)
    // TODO: if(res.data.code == 200) delete
    lists.value = lists.value.filter(item => item.tracing_num !== tracing_num)
  }).catch(err => console.log(err))
};

//添加请求拦截器
http.interceptors.request.use(config => {
  // show loading
  loading.value = true
  return config;
}, err => {
  // hide loading
  loading.value = false
  return Promise.resolve(err);
});

//响应拦截器
http.interceptors.response.use(
    response => {
      // hide loading
      loading.value = false
      return response;
    },
    error => {
      // err message
      return Promise.reject(error);
    }
);

const datas = computed(() => {
  // if (dataSource === undefined) console.log('aaaaaaaaaaaaaa'); else console.log(dataSource);
  // return [dataSource]
  console.log(lists.value)
  return lists.value
})
const number = ref('')
const onSearch = () => {
  console.log(number.value)
  http.get('/query', {
    params: {
      receiver_phone: number.value
    }
  }).then((res) => {
    console.log(res.data)
    lists.value.push(res.data)
  })
}
</script>

<template>
  <a-input-search
      v-model:value="number.receiver_phone"
      placeholder="请输入手机号"
      style="width: 200px; margin: 0px 0px 20px 0px"
      @search="onSearch"
  />
  <a-table
      :columns="columns"
      :data-source="datas"
      :scroll="{ x: 1500 }"
      :loading="loading"
  >

    <template #headerCell="{ column }">
      <template v-for="name in column.name">
        <span>{{ name }}</span>
      </template>
    </template>

    <template #bodyCell="{ column, record }">

      <template v-if="column.key === 'cash_on'">
        <span>
          {{ record.cash_on ? '是' : '否' }}
        </span>
      </template>
      <template v-else-if="column.key === 'action'">
        <span>
          <a-popconfirm
              title="确认取走吗？"
              @confirm="take_object(record.tracing_num, record.receiver_phone)"
          >
            <a>取走</a>
          </a-popconfirm>
        </span>
      </template>

    </template>

  </a-table>

</template>

<style scoped>

</style>