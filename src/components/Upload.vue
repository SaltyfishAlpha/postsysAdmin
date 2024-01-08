<script setup>
import {computed, onMounted, ref, watch} from "vue";
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
    title: '是否垫付运费',
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
    title: '是否为挂号信',
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

// const {
//   data: dataSource,
//   run,
//   loading,
// } = useRequest(()=>{http.get('/allocate').then(()=>{ console.log('update') })})
const dataSource = ref([])
const loading = ref(false)
const getItems = () => {
  http.get('/allocate').then((res) => {
    console.log("res.data.data")
    console.log(res.data.data)
    dataSource.value = res.data.data
  })
}
onMounted(()=>{
  getItems()
})

const currentNum = ref('')
const my_shelf_num = ref('')
const open = ref(false)
const editTrace = function (num) {
  currentNum.value = num
  open.value = true
}
const handleOk = () => {
  // console.log({tracing_num: currentNum.value, shelf_num: my_shelf_num.value})
  loading.value = true
  http.post('/allocated', {tracing_num: currentNum.value, shelf_num: my_shelf_num.value}).then((res) => {
    // http.post('/allocated', {tracing_num: '124124124413551241', shelf_num: my_shelf_num.value}).then((res)=>{
    console.log(res.data)
    open.value = false
    getItems()
  }).catch(err => console.log(err))
};

const handleCancel = () => {
  currentNum.value = ''
  open.value = false;
};

const handleChange = () => {
  //TODO
  console.log('change')
  getItems()
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

/////////////////////
//     POST       //
////////////////////

const visible_data = ref(false)
const postdata = ref({
  company: "",
  cash_on: false,
  pre_freight: false,
  tracing_num: "",
  registered: false,
  receiver_phone: "",
  rejected: false,
})
const showForm = () => {
  visible_data.value = true
}
const datahandleOk = () => {
  console.log(postdata.value)
  loading.value = true
  http.post('/upload', postdata).then((res) => {
    console.log(res.data)
    visible_data.value = false
    run()
  }).catch(err => console.log(err))
};
const datahandleCancel = () => {
  visible_data.value = false;
};
//for test
watch(postdata, (n, o) => {
  console.log(postdata.value)
})
</script>

<template>
  <a-button type="primary" style="margin: 0px 0px 20px 0px" @click="showForm">快递录入</a-button>
  <a-table
      :columns="columns"
      :data-source="datas"
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

      <template v-if="column.key === 'cash_on'">
        <span>
          {{ record.cash_on ? '是' : '否' }}
        </span>
      </template>
      <template v-else-if="column.key === 'action'">
        <span>
          <a @click="editTrace(record.tracing_num)">
            分配货架号
          </a>
        </span>
      </template>

    </template>

  </a-table>

  <a-modal v-model:open="open" title="填写货架号" @ok="handleOk">
    <template #footer>
      <a-button key="back" @click="handleCancel">返回</a-button>
      <a-button key="submit" type="primary" :loading="loading" @click="handleOk">提交</a-button>
    </template>

    <a-form>
      <a-form-item
          label="货架号"
          :rules="[{ required: true, message: '请输入货架号' }]"
      >
        <a-input v-model:value="my_shelf_num"/>
      </a-form-item>
    </a-form>
  </a-modal>

  <a-modal v-model:open="visible_data" title="录入信息" @ok="datahandleOk">
    <template #footer>
      <a-button key="back" @click="datahandleCancel">返回</a-button>
      <a-button key="submit" type="primary" :loading="loading" @click="datahandleOk">提交</a-button>
    </template>

    <a-form>
      <a-form-item
          label="快递公司"
          :rules="[{ required: true, message: '请输入' }]"
          :span="24"
      >
        <a-input v-model:checked="postdata.company"/>
      </a-form-item>
      <a-form-item
          label="是否到付"
          :rules="[{ required: true, message: '请输入' }]"
      >
        <a-checkbox v-model:checked="postdata.cash_on"/>
      </a-form-item>
      <a-form-item
          label="pre_freight"
          :rules="[{ required: true, message: '请输入' }]"
      >
        <a-checkbox v-model:checked="postdata.pre_freight"/>
      </a-form-item>
      <a-form-item
          label="单号"
          :rules="[{ required: true, message: '请输入' }]"
      >
        <a-input v-model:checked="postdata.tracing_num"/>
      </a-form-item>
      <a-form-item
          label="registered"
          :rules="[{ required: true, message: '请输入' }]"
      >
        <a-checkbox v-model:checked="postdata.registered"/>
      </a-form-item>
      <a-form-item
          label="收件人手机号"
          :rules="[{ required: true, message: '请输入' }]"
      >
        <a-input v-model:checked="postdata.receiver_phone"/>
      </a-form-item>
      <a-form-item
          label="rejected"
          :rules="[{ required: true, message: '请输入' }]"
      >
        <a-checkbox v-model:checked="postdata.rejected"/>
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<style scoped>

</style>