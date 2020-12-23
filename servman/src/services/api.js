import axios from "axios";
const api = axios.create({
    baseURL: 'https://803d12034421.ngrok.io',
})

export default api;
