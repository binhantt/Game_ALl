import React, { createContext } from "react";
import type { AuthContextProps } from "../type/auth";
import { loginService } from "../services/authService";

export const AuthContext = createContext<AuthContextProps>({
  login: async () => {},
  loading: false,
  error: "",
  success: "",
});

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [loading, setLoading] = React.useState(false);
  const [error, setError] = React.useState("");
  const [success, setSuccess] = React.useState("");

  const login = async (email: string, password: string) => {
    setLoading(true);
    setError("");
    setSuccess("");
    try {
      await loginService(email, password);
      setSuccess("Đăng nhập thành công!");
    } catch (err: any) {
      setError(err.message || "Đăng nhập thất bại!");
    } finally {
      setLoading(false);
    }
  };

  return (
    <AuthContext.Provider value={{ login, loading, error, success }}>
      {children}
    </AuthContext.Provider>
  );
}; 