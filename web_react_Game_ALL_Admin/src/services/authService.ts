import axios from "../config/axiosInstance";

export async function loginService(email: string, password: string) {
  try {
    const res = await axios.post("/api/login", { email, password });
    // Có thể lưu token/res ở đây nếu cần
    return res.data;
  } catch (err: any) {
    throw new Error(err.response?.data?.message || "Đăng nhập thất bại!");
  }
} 