import React from "react";
import { Typography, Card, Row, Col } from "antd";
import MainLayout from "../layout/MainLayout";

const { Title, Paragraph } = Typography;

const cardData = [
  {
    title: "Người dùng",
    description: "Quản lý tài khoản người dùng.",
    bg: "linear-gradient(135deg, #1890ff 60%, #40a9ff 100%)",
    color: "#fff",
    fontSize: 24,
  },
  {
    title: "Trò chơi",
    description: "Quản lý danh sách trò chơi.",
    bg: "linear-gradient(135deg, #52c41a 60%, #73d13d 100%)",
    color: "#fff",
    fontSize: 28,
  },
  {
    title: "Báo cáo",
    description: "Xem báo cáo và thống kê.",
    bg: "linear-gradient(135deg, #faad14 60%, #ffc53d 100%)",
    color: "#222",
    fontSize: 32,
  },
];

const cardStyle = {
  borderRadius: 24,
  minHeight: 220,
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
          Quản lý hệ thống game của bạn một cách dễ dàng và hiệu quả.
        </Paragraph>
        <Row gutter={[40, 40]} justify="center" align="middle">
          {cardData.map((card, idx) => (
            <Col xs={24} sm={12} md={8} key={card.title} style={{ display: "flex", justifyContent: "center" }}>
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
                  fontSize: card.fontSize,
                  fontWeight: 500,
                }}
                headStyle={{
                  background: "transparent",
                  borderBottom: "none",
                  textAlign: "center",
                  color: card.color,
                  fontSize: card.fontSize + 2,
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
                title={card.title}
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