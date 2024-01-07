import axios from 'axios';

const http = axios.create({
    baseURL: '/api',
    timeout:5000, //超时时长
    headers: {
        "Content-Type" : "multipart/form-data",
        //"Content-Type" : "application/json"
    }
})

export default http;