import React from "react";
import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";
import { AuthProvider } from "./context/AuthContext";
import LoginForm from "./page/LoginForm";
import Home from "./page/Home";

function App() {
  return (
    <AuthProvider>
      <Router>
        <Routes>
          <Route path="/login" element={<LoginForm />} />
          <Route path="/home" element={<Home />} />
          {/* <Route path="*" element={<Navigate to="/home" replace />} /> */}
        </Routes>
      </Router>
    </AuthProvider>
  );
}

export default App;
