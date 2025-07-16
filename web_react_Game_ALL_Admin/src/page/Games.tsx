import React from "react";
import { Table, Typography, Tag, Space, Popconfirm, message } from "antd";
import MainLayout from "../layout/MainLayout";
import { EyeOutlined, EditOutlined, DeleteOutlined } from "@ant-design/icons";

const { Title } = Typography;

const dataSource = [
  {
    key: "1",
    name: "Assassin's Creed",
    publisher: "Ubisoft",
    genre: "Phiêu lưu",
    price: 790000,
    status: "active",
  },
  {
    key: "2",
    name: "Super Mario Odyssey",
    publisher: "Nintendo",
    genre: "Platform",
    price: 990000,
    status: "active",
  },
  {
    key: "3",
    name: "Half-Life 2",
    publisher: "Valve",
    genre: "Bắn súng",
    price: 350000,
    status: "inactive",
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
    title: "Tên game",
    dataIndex: "name",
    key: "name",
    render: (name: string) => <b>{name}</b>,
  },
  {
    title: "Nhà sản xuất",
    dataIndex: "publisher",
    key: "publisher",
    render: (publisher: string) => <Tag color="blue">{publisher}</Tag>,
  },
  {
    title: "Thể loại",
    dataIndex: "genre",
    key: "genre",
    render: (genre: string) => <Tag color="purple">{genre}</Tag>,
  },
  {
    title: "Giá",
    dataIndex: "price",
    key: "price",
    render: (price: number) => price.toLocaleString() + " ₫",
  },
  {
    title: "Trạng thái",
    dataIndex: "status",
    key: "status",
    render: (status: string) => (
      <Tag color={status === "active" ? "green" : "volcano"}>{status === "active" ? "Đang bán" : "Ngừng bán"}</Tag>
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
        <Popconfirm title="Bạn có chắc muốn xóa game này?" okText="Xóa" cancelText="Hủy" onConfirm={() => message.success('Đã xóa!')}>
          <a title="Xóa"><DeleteOutlined style={{ color: '#ff4d4f' }} /></a>
        </Popconfirm>
      </Space>
    ),
  },
];

const Games: React.FC = () => {
  return (
    <MainLayout>
      <div style={{ padding: 24 }}>
        <Title level={2} style={{ marginBottom: 24 }}>Quản lý trò chơi</Title>
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

export default Games; 