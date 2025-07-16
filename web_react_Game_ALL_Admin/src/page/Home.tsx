import React from "react";
import { Typography, Card, Row, Col } from "antd";
import MainLayout from "../layout/MainLayout";
import {
  DashboardOutlined,
  UserOutlined,
  AppstoreOutlined,
  BarChartOutlined,
} from "@ant-design/icons";

const { Title, Paragraph } = Typography;

const cardData = [
  {
    title: "Dashboard",
    description: "Tổng quan hoạt động shop bán game.",
    bg: "linear-gradient(135deg, #1890ff 60%, #40a9ff 100%)", // Xanh dương
    color: "#fff",
    icon: <DashboardOutlined style={{ fontSize: 36, color: '#1890ff' }} />,
  },
  {
    title: "Người dùng",
    description: "Quản lý tài khoản và thông tin khách hàng.",
    bg: "linear-gradient(135deg, #52c41a 60%, #73d13d 100%)", // Xanh lá
    color: "#fff",
    icon: <UserOutlined style={{ fontSize: 36, color: '#52c41a' }} />,
  },
  {
    title: "Trò chơi",
    description: "Quản lý danh mục và chi tiết game bán ra.",
    bg: "linear-gradient(135deg, #722ed1 60%, #b37feb 100%)", // Tím
    color: "#fff",
    icon: <AppstoreOutlined style={{ fontSize: 36, color: '#722ed1' }} />,
  },
  {
    title: "Báo cáo",
    description: "Xem báo cáo, thống kê doanh thu và hiệu quả bán hàng.",
    bg: "linear-gradient(135deg, #fa8c16 60%, #faad14 100%)", // Cam
    color: "#fff",
    icon: <BarChartOutlined style={{ fontSize: 36, color: '#fa8c16' }} />,
  },
];

const cardStyle = {
  borderRadius: 24,
  minHeight: 200,
  maxWidth: 350,
  width: "100%",
  boxShadow: "0 8px 32px rgba(0,0,0,0.10)",
  display: "flex",
  flexDirection: "column" as const,
  alignItems: "center" as const,
  justifyContent: "center" as const,
  transition: "transform 0.2s, box-shadow 0.2s",
  cursor: "pointer",
};

const Home: React.FC = () => {
  return (
    <MainLayout>
      <div
        style={{
          minHeight: "calc(100vh - 180px)",
          display: "flex",
          flexDirection: "column",
          justifyContent: "center",
        }}
      >
        <Title level={2} style={{ textAlign: "center", marginBottom: 8 }}>
          Chào mừng đến với <span style={{ color: "#1890ff" }}>Game Admin</span>
        </Title>
        <Paragraph style={{ textAlign: "center", marginBottom: 40, fontSize: 18 }}>
          Quản lý shop bán game, sản phẩm, đơn hàng và khách hàng một cách dễ dàng, hiệu quả.
        </Paragraph>
        <Row gutter={[40, 40]} justify="center" align="middle">
          {cardData.map((card, idx) => (
            <Col xs={24} sm={12} md={8} lg={6} key={card.title} style={{ display: "flex", justifyContent: "center" }}>
              <Card
                bordered={false}
                style={{
                  ...cardStyle,
                  background: card.bg,
                  color: card.color,
                }}
                bodyStyle={{
                  textAlign: "center",
                  padding: 32,
                  color: card.color,
                  fontSize: 20,
                  fontWeight: 500,
                }}
                headStyle={{
                  background: "transparent",
                  borderBottom: "none",
                  textAlign: "center",
                  color: card.color,
                  fontSize: 24,
                  fontWeight: 700,
                }}
                hoverable
                onMouseEnter={e => {
                  (e.currentTarget as HTMLDivElement).style.transform = "translateY(-8px) scale(1.03)";
                  (e.currentTarget as HTMLDivElement).style.boxShadow = "0 16px 48px rgba(0,0,0,0.15)";
                }}
                onMouseLeave={e => {
                  (e.currentTarget as HTMLDivElement).style.transform = "";
                  (e.currentTarget as HTMLDivElement).style.boxShadow = cardStyle.boxShadow;
                }}
                title={<div style={{ display: 'flex', alignItems: 'center', justifyContent: 'center', gap: 12 }}>{card.icon}<span>{card.title}</span></div>}
              >
                {card.description}
              </Card>
            </Col>
          ))}
        </Row>
      </div>
    </MainLayout>
  );
};

export default Home; 