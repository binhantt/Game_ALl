import React from "react";
import { Form, Input } from "antd";

interface CustomInputProps {
  label: string;
  name: string;
  type?: string;
  placeholder?: string;
  rules?: any[];
  autoComplete?: string;
}

const CustomInput: React.FC<CustomInputProps> = ({
  label,
  name,
  type = "text",
  placeholder,
  rules,
  autoComplete,
}) => (
  <Form.Item label={label} name={name} rules={rules}>
    {type === "password" ? (
      <Input.Password placeholder={placeholder} autoComplete={autoComplete} />
    ) : (
      <Input placeholder={placeholder} autoComplete={autoComplete} />
    )}
  </Form.Item>
);

export default CustomInput; 