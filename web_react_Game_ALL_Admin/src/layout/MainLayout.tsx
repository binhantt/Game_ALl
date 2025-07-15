import React from "react";
import { Layout, Avatar, Typography } from "antd";
import SiderBar from "./SiderBar";

const { Header, Content, Footer } = Layout;
const { Title } = Typography;

const MainLayout: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  return (
    <Layout style={{ minHeight: "100vh" }}>
      <SiderBar />
      <Layout>
        <Header style={{ display: "flex", alignItems: "center", justifyContent: "space-between", background: "#fff", boxShadow: "0 2px 8px #f0f1f2" }}>
          <div style={{ display: "flex", alignItems: "center" }}>
            <Avatar style={{ backgroundColor: "#1890ff", marginRight: 12 }}>A</Avatar>
            <Title level={4} style={{ margin: 0 }}>Game Admin</Title>
          </div>
          <div>
            {/* User/account area, can add dropdown or user info here */}
            <Avatar style={{ backgroundColor: "#87d068" }} icon={"U"} />
          </div>
        </Header>
        <Content style={{ margin: "24px 16px 0", padding: 24, background: "#fff", borderRadius: 8, minHeight: 360 }}>
          {children}
        </Content>
        <Footer style={{ textAlign: "center" }}>Game Admin ©2024 Created by You</Footer>
      </Layout>
    </Layout>
  );
};

export default MainLayout; 