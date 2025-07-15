export interface AuthContextProps {
  login: (email: string, password: string) => Promise<void>;
  loading: boolean;
  error: string;
  success: string;
} 