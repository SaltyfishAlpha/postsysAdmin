import Mock from 'mockjs'

const data = {
    company: "aaa",
    sender_name: "ax",
    sender_phone: "12312312312",
    sender_id: "123123123412341234",
    receiver_name: "bx",
    receiver_phone: "11122233344",
    receiver_addr: "//",
    object_type: 0,
    object_weights: 20.1,
    parcel_insurance: false,
    object_value: 0.0,
    cash_on: false
};

Mock.setup({
    timeout: '1000-2000'
})

Mock.mock('/api/ping', 'get', function () {
    console.log('get111');
    return {}
})

Mock.mock('/api/getinfo', 'get', function () {
    console.log('get111')
    return data
})