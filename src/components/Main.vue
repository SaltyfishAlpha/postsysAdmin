<script setup>
import {computed, onMounted, ref} from "vue";
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
    title: '寄件人姓名',
    dataIndex: 'sender_name',
    key: 'sender_name',
    width: 130,
  },
  {
    title: '寄件人电话号码',
    dataIndex: 'sender_phone',
    key: 'sender_phone',
    width: 150,
  },
  {
    title: '寄件人身份证号',
    key: 'sender_id',
    dataIndex: 'sender_id',
    width: 200,
  },
  {
    title: '收件人姓名',
    key: 'receiver_name',
    dataIndex: 'receiver_name',
    width: 130,
  },
  {
    title: '收件人电话号码',
    key: 'receiver_phone',
    dataIndex: 'receiver_phone',
    width: 150,
  },
  {
    title: '收件地址',
    key: 'receiver_addr',
    dataIndex: 'receiver_addr',
    width: 200,
  },
  {
    title: '物品信息',
    key: 'object_type',
    width: 100,
  },
  {
    title: '保价信息',
    key: 'parcel_insurance',
    width: 100,
  },
  {
    title: '是否到付',
    key: 'cash_on',
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
    company: "aaa",
    sender_name: "ax",
    sender_phone: "12312312312",
    sender_id: "abzdbz",
    receiver_name: "bx",
    receiver_phone: "11122233344",
    receiver_addr: "//",
    object_type: 0,
    object_weights: 20.0,
    parcel_insurance: false,
    object_value: 10.0,
    cash_on: true
  },
  {
    company: "aaa",
    sender_name: "ax",
    sender_phone: "12312312312",
    sender_id: "11111",
    receiver_name: "bx",
    receiver_phone: "11122233344",
    receiver_addr: "//",
    object_type: 0,
    object_weights: 10.0,
    parcel_insurance: true,
    object_value: 10.0,
    cash_on: true
  },
];

// const {
//   data: dataSource,
//   run,
//   loading,
// } = usePagination(()=>{http.get('/getinfo')})
const loading = ref(false)
const dataSource = ref([])
const getItems = ()=>{
  http.get('/getinfo').then((res)=>{
    dataSource.value = res.data.data
  })
}
onMounted(()=>{
  getItems()
})

const currentNum = ref('')
const trace_num = ref('')
const open = ref(false)
const editTrace = function(num) {
  currentNum.value = num
  open.value = true
}
const handleOk = () => {
  console.log({sender_id: currentNum.value, tracing_num: trace_num.value})
  loading.value = true
  http.post('/updateinfo', {sender_id: currentNum.value, tracing_num: trace_num.value}).then(()=>{
  //http.post('/updateinfo', {sender_id: '124124124413551241', tracing_num: trace_num.value}).then(()=>{
    console.log('success')
    open.value = false
    getItems()
  }).catch(err=>console.log(err))
};

const handleCancel = () => {
  currentNum.value = ''
  open.value = false;
};

const handleChange = () => {
  //TODO
  console.log('change')
  run()
}

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

</script>

<template>
  <a-table
      :columns="columns"
      :data-source="dataSource"
      :scroll="{ x: 1500 }"
      :loading="loading"
      @change="handleChange"
  >

    <template #headerCell="{ column }">
      <template v-for="name in column.name">
        <span>{{ name }}</span>
      </template>
    </template>

    <template #bodyCell="{ column, record }">

      <template v-if="column.key === 'object_type'">
        <span>
          <a-tag color="geekblue">
            {{ myEnum[record.object_type] }}/{{ record.object_weights }}kg
          </a-tag>
        </span>
      </template>
      <template v-else-if="column.key === 'parcel_insurance'">
        <span>
          {{ record.parcel_insurance ? record.object_value + '￥' : '否' }}
        </span>
      </template>
      <template v-else-if="column.key === 'cash_on'">
        <span>
          {{ record.cash_on ? '是' : '否' }}
        </span>
      </template>
      <template v-else-if="column.key === 'action'">
        <span>
          <a @click="editTrace(record.sender_id)">
            分配单号
          </a>
          <!--a-divider type="vertical"/>
          <a>Delete</a-->
        </span>
      </template>

    </template>

  </a-table>

  <a-modal v-model:open="open" title="填写单号" @ok="handleOk">
    <template #footer>
      <a-button key="back" @click="handleCancel">返回</a-button>
      <a-button key="submit" type="primary" :loading="loading" @click="handleOk">提交</a-button>
    </template>

    <a-form>
      <a-form-item
          label="订单号"
          :rules="[{ required: true, message: '请输入订单号' }]"
      >
        <a-input v-model:value="trace_num" />
      </a-form-item>
    </a-form>
  </a-modal>


</template>

<style scoped>

</style>