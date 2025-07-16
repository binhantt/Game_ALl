import React from "react";
import { Row, Col, Card, Typography, Table, Tag } from "antd";
import MainLayout from "../layout/MainLayout";

const { Title } = Typography;

const stats = [
  { title: "Tổng doanh thu", value: "₫25,000,000", color: "#faad14" },
  { title: "Số đơn hàng", value: 320, color: "#1890ff" },
  { title: "Game bán chạy", value: 12, color: "#722ed1" },
  { title: "Người dùng mới", value: 45, color: "#52c41a" },
];

const orderData = [
  { key: 1, code: "ORD001", user: "Nguyễn Văn A", game: "Assassin's Creed", price: 790000, status: "Hoàn thành" },
  { key: 2, code: "ORD002", user: "Trần Thị B", game: "Super Mario Odyssey", price: 990000, status: "Đang xử lý" },
  { key: 3, code: "ORD003", user: "Lê Văn C", game: "Half-Life 2", price: 350000, status: "Đã hủy" },
];

const orderColumns = [
  { title: "Mã đơn", dataIndex: "code", key: "code" },
  { title: "Khách hàng", dataIndex: "user", key: "user" },
  { title: "Game", dataIndex: "game", key: "game" },
  { title: "Giá", dataIndex: "price", key: "price", render: (p: number) => p.toLocaleString() + " ₫" },
  { title: "Trạng thái", dataIndex: "status", key: "status", render: (s: string) => {
    let color = s === "Hoàn thành" ? "green" : s === "Đang xử lý" ? "blue" : "volcano";
    return <Tag color={color}>{s}</Tag>;
  } },
];

const Reports: React.FC = () => {
  return (
    <MainLayout>
      <div style={{ padding: 24 }}>
        <Title level={2} style={{ marginBottom: 32 }}>Thống kê & Báo cáo</Title>
        <Row gutter={[32, 32]} justify="center" style={{ marginBottom: 32 }}>
          {stats.map(stat => (
            <Col xs={24} sm={12} md={12} lg={6} key={stat.title} style={{ display: "flex", justifyContent: "center" }}>
              <Card bordered={false} style={{ minWidth: 200, minHeight: 120, borderRadius: 18, background: stat.color + '22', boxShadow: "0 4px 24px rgba(0,0,0,0.07)", display: "flex", flexDirection: "column", alignItems: "center", justifyContent: "center" }}>
                <div style={{ fontSize: 28, fontWeight: 700, marginBottom: 4, color: stat.color }}>{stat.value}</div>
                <div style={{ fontSize: 16, color: "#555" }}>{stat.title}</div>
              </Card>
            </Col>
          ))}
        </Row>
        <Card bordered={false} style={{ borderRadius: 18 }}>
          <Title level={4} style={{ marginBottom: 16 }}>Đơn hàng gần đây</Title>
          <Table dataSource={orderData} columns={orderColumns} pagination={false} rowKey="key" />
        </Card>
      </div>
    </MainLayout>
  );
};

export default Reports; 