import React, { useContext } from "react";
import { Form, Button, Alert } from "antd";
import { AuthContext } from "../context/AuthContext";
import CustomInput from "../components/iu/CustomInput";

const LoginForm: React.FC = () => {
  const { login, loading, error, success } = useContext(AuthContext);
  const [form] = Form.useForm();

  const handleFinish = (values: { email: string; password: string }) => {
    login(values.email, values.password);
  };

  return (
    <div style={{ maxWidth: 400, margin: "40px auto", padding: 24, border: "1px solid #eee", borderRadius: 8 }}>
      <h2>Đăng nhập</h2>
      <Form
        form={form}
        layout="vertical"
        onFinish={handleFinish}
        autoComplete="off"
      >
        <CustomInput
          label="Email"
          name="email"
          placeholder="Email"
          autoComplete="username"
          rules={[{ required: true, message: "Vui lòng nhập email!" }]}
        />
        <CustomInput
          label="Mật khẩu"
          name="password"
          type="password"
          placeholder="Mật khẩu"
          autoComplete="current-password"
          rules={[{ required: true, message: "Vui lòng nhập mật khẩu!" }]}
        />
        <Form.Item>
          <Button type="primary" htmlType="submit" loading={loading} style={{ width: "100%" }}>
            Đăng nhập
          </Button>
        </Form.Item>
      </Form>
      {error && <Alert type="error" message={error} showIcon style={{ marginTop: 16 }} />}
      {success && <Alert type="success" message={success} showIcon style={{ marginTop: 16 }} />}
    </div>
  );
};

export default LoginForm; 