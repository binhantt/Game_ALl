
// # import axios
// # export axios sau config
// # config
// # axiosInstance luôn gửi token
// # axiosInstance luôn kiểm tra được 401 và redirect về /login
// # axiosInstance sẽ trả về lỗi api
// axiosInstance.js
import axios from 'axios';
 const baseURL = 'http://localhost:3000/api' 
 
// Hàm lấy token từ localStorage (hoặc nơi bạn lưu)
const getToken = () => {
  return localStorage.getItem('token');
};

// Tạo instance
const axiosInstance = axios.create({
  baseURL: baseURL, // ✅ thay bằng API của bạn
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
});

// ✅ Add token vào headers mỗi request
axiosInstance.interceptors.request.use(
  (config) => {
    const token = getToken();
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`;
    }
    return config;
  },
  (error) => Promise.reject(error)
);

// ✅ Xử lý lỗi 401 → redirect login
axiosInstance.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response && error.response.status === 401) {
      // Xoá token nếu cần
      localStorage.removeItem('token');
      // Redirect tới login
      window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);

export default axiosInstance;
