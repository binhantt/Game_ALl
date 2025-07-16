import React from "react";
import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";
import { AuthProvider } from "./context/AuthContext";
import LoginForm from "./page/LoginForm";
import Home from "./page/Home";
import Dashboard from "./page/Dashboard";
import Users from "./page/Users";
import Publishers from "./page/Publishers";
import Reports from "./page/Reports";
import Games from "./page/Games";
import Genres from "./page/Genres";

function App() {
  return (
    <AuthProvider>
      <Router>
        <Routes>
        // ... existing code ..
            <Route path="/home" element={<Home />} />
            <Route path="/dashboard" element={<Dashboard />} />
            <Route path="/users" element={<Users />} />
            <Route path="/publishers" element={<Publishers />} />
            <Route path="/games" element={<Games />} />
            <Route path="/genres" element={<Genres />} />
            <Route path="/reports" element={<Reports />} />
// ... existing code ...
          {/* <Route path="*" element={<Navigate to="/home" replace />} /> */}
        </Routes>
      </Router>
    </AuthProvider>
  );
}

export default App;
