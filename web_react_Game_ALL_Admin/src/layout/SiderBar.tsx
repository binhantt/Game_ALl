import React, { useState } from "react";
import { Layout, Menu } from "antd";
import {
  DashboardOutlined,
  UserOutlined,
  AppstoreOutlined,
  BarChartOutlined,
  TagsOutlined,
} from "@ant-design/icons";
import { useNavigate } from "react-router-dom";

const { Sider } = Layout;

const SiderBar: React.FC = () => {  
  const [collapsed, setCollapsed] = useState(false);
  const navigate = useNavigate();

  return (
    <Sider collapsible collapsed={collapsed} onCollapse={setCollapsed} breakpoint="lg">
      <div style={{ height: 64, display: "flex", alignItems: "center", justifyContent: "center", color: "#fff", fontWeight: "bold", fontSize: 20 }}>
        {collapsed ? "GA" : "Game Admin"}
      </div>
      <Menu
        theme="dark"
        mode="inline"
        defaultSelectedKeys={["dashboard"]}
        onClick={({ key }) => navigate(`/${key}`)}
      >
        <Menu.Item key="dashboard" icon={<DashboardOutlined />}>Dashboard</Menu.Item>
        <Menu.Item key="users" icon={<UserOutlined />}>Người dùng</Menu.Item>
        <Menu.Item key="games" icon={<AppstoreOutlined />}>Trò chơi</Menu.Item>
        <Menu.Item key="genres" icon={<TagsOutlined />}>Thể loại</Menu.Item>
        <Menu.Item key="reports" icon={<BarChartOutlined />}>Thống kê</Menu.Item>
      </Menu>
    </Sider>
  );
};

export default SiderBar; 