import axios from "axios";
const api = axios.create({
    baseURL: 'http://192.168.2.38:8000',
})
export default api;
