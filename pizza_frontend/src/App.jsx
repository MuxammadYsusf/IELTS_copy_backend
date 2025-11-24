// src/App.jsx
import React from "react";
import { BrowserRouter, Routes, Route, Navigate, useLocation } from "react-router-dom";
import Navbar from "./components/Navbar";

import Home from "./pages/Home";
import Menu from "./pages/Menu";
import Cart from "./pages/Cart";
import Order from "./pages/Order";
import History from "./pages/History";
import Profile from "./pages/Profile";

// Auth pages
import AuthLayout from "./pages/auth/AuthLayout";
import Login from "./pages/auth/Login";
import Register from "./pages/auth/Register";
import Logout from "./pages/auth/Logout";

// If you DO have src/utils/auth.js, keep this import.
// If not yet, temporarily comment the admin route below OR create the file.
import { isAdmin } from "./utils/auth";

// ---- Simple auth gate (localStorage token) ----
function RequireAuth({ children }) {
  const location = useLocation();
  const token = localStorage.getItem("token");
  if (!token) {
    return <Navigate to="/auth/login" replace state={{ from: location }} />;
  }
  return children;
}

export default function App() {
  return (
    <BrowserRouter>
      <Navbar />

      <Routes>
        <Route index element={<Home />} />
        <Route path="/menu" element={<Menu />} />

        {/* Protected pages */}
        <Route
          path="/cart"
          element={
            <RequireAuth>
              <Cart />
            </RequireAuth>
          }
        />
        <Route
          path="/order"
          element={
            <RequireAuth>
              <Order />
            </RequireAuth>
          }
        />
        <Route
          path="/history"
          element={
            <RequireAuth>
              <History />
            </RequireAuth>
          }
        />

        {/* Auth group */}
        <Route path="/auth" element={<AuthLayout />}>
          <Route index element={<Navigate to="login" replace />} />
          <Route path="login" element={<Login />} />
          <Route path="register" element={<Register />} />
          <Route path="logout" element={<Logout />} />
        </Route>

    <Route path="/profile" element={<Profile />} />

        {/* Admin route — only enable if both utils/auth.js and pages/admin/AdminLayout.jsx exist */}
        {/* 
        <Route
          path="/admin/*"
          element={
            isAdmin() ? (
              // replace with your real admin layout component
              <div className="p-6">Admin area placeholder</div>
            ) : (
              <Navigate to="/" replace />
            )
          }
        />
        */}

        {/* 404 fallback */}
        <Route path="*" element={<Navigate to="/" replace />} />
      </Routes>
    </BrowserRouter>
  );
}