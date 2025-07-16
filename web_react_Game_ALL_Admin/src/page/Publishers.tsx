import React from "react";
import { Table, Typography, Space, Popconfirm, message, Tag } from "antd";
import MainLayout from "../layout/MainLayout";
import { EyeOutlined, EditOutlined, DeleteOutlined } from "@ant-design/icons";

const { Title } = Typography;

const dataSource = [
  {
    key: "1",
    name: "Ubisoft",
    country: "Pháp",
    games: 12,
  },
  {
    key: "2",
    name: "Nintendo",
    country: "Nhật Bản",
    games: 20,
  },
  {
    key: "3",
    name: "Valve",
    country: "Mỹ",
    games: 8,
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
    title: "Tên nhà sản xuất",
    dataIndex: "name",
    key: "name",
    render: (name: string) => <b>{name}</b>,
  },
  {
    title: "Quốc gia",
    dataIndex: "country",
    key: "country",
    render: (country: string) => <Tag color="blue">{country}</Tag>,
  },
  {
    title: "Số lượng game",
    dataIndex: "games",
    key: "games",
    render: (games: number) => <span style={{ fontWeight: 500 }}>{games}</span>,
  },
  {
    title: "Thao tác",
    key: "action",
    width: 120,
    render: (_: any, record: any) => (
      <Space size="middle">
        <a title="Xem"><EyeOutlined style={{ color: '#1890ff' }} /></a>
        <a title="Sửa"><EditOutlined style={{ color: '#faad14' }} /></a>
        <Popconfirm title="Bạn có chắc muốn xóa nhà sản xuất này?" okText="Xóa" cancelText="Hủy" onConfirm={() => message.success('Đã xóa!')}>
          <a title="Xóa"><DeleteOutlined style={{ color: '#ff4d4f' }} /></a>
        </Popconfirm>
      </Space>
    ),
  },
];

const Publishers: React.FC = () => {
  return (
    <MainLayout>
      <div style={{ padding: 24 }}>
        <Title level={2} style={{ marginBottom: 24 }}>Quản lý nhà sản xuất</Title>
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

export default Publishers; 