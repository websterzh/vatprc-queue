import axios from 'axios';


const api = axios.create({
  baseURL: `${window.location.protocol}//${window.location.host}`,
  timeout: 10000,
  headers: {
    "Accept": "application/json",
    "Content-Type": "application/json",
  },
});

api.interceptors.response.use(response => {
  return response.data;
});

export default api;