import React from "react";
import { Table, Tag, Typography, Space, Popconfirm, message } from "antd";
import MainLayout from "../layout/MainLayout";
import { EyeOutlined, EditOutlined, DeleteOutlined } from "@ant-design/icons";

const { Title } = Typography;

const dataSource = [
  {
    key: "1",
    name: "Nguyễn Văn A",
    email: "a@gmail.com",
    role: "Admin",
    status: "active",
  },
  {
    key: "2",
    name: "Trần Thị B",
    email: "b@gmail.com",
    role: "Khách hàng",
    status: "inactive",
  },
  {
    key: "3",
    name: "Lê Văn C",
    email: "c@gmail.com",
    role: "Khách hàng",
    status: "active",
  },
];

const columns = [
  {
    title: "STT",
    dataIndex: "key",
    key: "key",
    width: 60,
    render: (text: string, record: any, idx: number) => idx + 1,
  },
  {
    title: "Tên",
    dataIndex: "name",
    key: "name",
  },
  {
    title: "Email",
    dataIndex: "email",
    key: "email",
  },
  {
    title: "Vai trò",
    dataIndex: "role",
    key: "role",
    render: (role: string) => (
      <Tag color={role === "Admin" ? "geekblue" : "green"}>{role}</Tag>
    ),
  },
  {
    title: "Trạng thái",
    dataIndex: "status",
    key: "status",
    render: (status: string) => (
      <Tag color={status === "active" ? "green" : "volcano"}>{status === "active" ? "Hoạt động" : "Khóa"}</Tag>
    ),
  },
  {
    title: "Thao tác",
    key: "action",
    width: 120,
    render: (_: any, record: any) => (
      <Space size="middle">
        <a title="Xem"><EyeOutlined style={{ color: '#1890ff' }} /></a>
        <a title="Sửa"><EditOutlined style={{ color: '#faad14' }} /></a>
        <Popconfirm title="Bạn có chắc muốn xóa người dùng này?" okText="Xóa" cancelText="Hủy" onConfirm={() => message.success('Đã xóa!')}>
          <a title="Xóa"><DeleteOutlined style={{ color: '#ff4d4f' }} /></a>
        </Popconfirm>
      </Space>
    ),
  },
];

const Users: React.FC = () => {
  return (
    <MainLayout>
      <div style={{ padding: 24 }}>
        <Title level={2} style={{ marginBottom: 24 }}>Quản lý người dùng</Title>
        <Table
          dataSource={dataSource}
          columns={columns}
          pagination={{ pageSize: 8 }}
          bordered
          rowKey="key"
        />
      </div>
    </MainLayout>
  );
};

export default Users; 