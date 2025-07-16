import React from "react";
import { Row, Col, Card, Typography } from "antd";
import MainLayout from "../layout/MainLayout";
import {
  UserOutlined,
  AppstoreOutlined,
  DollarOutlined,
  ShoppingCartOutlined,
} from "@ant-design/icons";

const { Title } = Typography;

const stats = [
  {
    title: "Người dùng",
    value: 1280,
    icon: <UserOutlined style={{ fontSize: 32, color: "#52c41a" }} />,
    bg: "linear-gradient(135deg, #eafff1 60%, #b7eb8f 100%)",
  },
  {
    title: "Trò chơi",
    value: 320,
    icon: <AppstoreOutlined style={{ fontSize: 32, color: "#722ed1" }} />,
    bg: "linear-gradient(135deg, #f9f0ff 60%, #d3adf7 100%)",
  },
  {
    title: "Doanh thu",
    value: "$12,500",
    icon: <DollarOutlined style={{ fontSize: 32, color: "#faad14" }} />,
    bg: "linear-gradient(135deg, #fffbe6 60%, #ffe58f 100%)",
  },
  {
    title: "Đơn hàng",
    value: 540,
    icon: <ShoppingCartOutlined style={{ fontSize: 32, color: "#1890ff" }} />,
    bg: "linear-gradient(135deg, #e6f7ff 60%, #91d5ff 100%)",
  },
];

const Dashboard: React.FC = () => {
  return (
    <MainLayout>
      <div style={{ padding: 24 }}>
        <Title level={2} style={{ marginBottom: 32, textAlign: "center" }}>
          Tổng quan hệ thống
        </Title>
        <Row gutter={[32, 32]} justify="center">
          {stats.map((stat) => (
            <Col xs={24} sm={12} md={12} lg={6} key={stat.title} style={{ display: "flex", justifyContent: "center" }}>
              <Card
                bordered={false}
                style={{
                  width: "100%",
                  minWidth: 220,
                  minHeight: 140,
                  borderRadius: 18,
                  background: stat.bg,
                  boxShadow: "0 4px 24px rgba(0,0,0,0.07)",
                  display: "flex",
                  flexDirection: "column",
                  alignItems: "center",
                  justifyContent: "center",
                }}
                bodyStyle={{ padding: 24, textAlign: "center" }}
              >
                <div style={{ marginBottom: 12 }}>{stat.icon}</div>
                <div style={{ fontSize: 28, fontWeight: 700, marginBottom: 4 }}>{stat.value}</div>
                <div style={{ fontSize: 16, color: "#555" }}>{stat.title}</div>
              </Card>
            </Col>
          ))}
        </Row>
      </div>
    </MainLayout>
  );
};

export default Dashboard; 