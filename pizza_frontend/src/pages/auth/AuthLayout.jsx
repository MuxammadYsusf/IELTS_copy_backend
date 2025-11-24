// src/pages/auth/AuthLayout.jsx
import React from "react";
import { NavLink, Outlet } from "react-router-dom";

export default function AuthLayout() {
  const tab = "px-4 py-2 rounded-lg transition";
  const idle = "text-gray-700 hover:bg-gray-100 dark:text-gray-100/85 dark:hover:bg-white/5";
  const active = "bg-white text-red-600 shadow dark:bg-zinc-800 dark:text-red-400";
  return (
    <main className="max-w-md mx-auto p-6">
      <h1 className="text-2xl font-bold mb-4">Account</h1>
      <div className="mb-6 flex gap-2">
        <NavLink to="login"    className={({isActive}) => `${tab} ${isActive?active:idle}`}>Login</NavLink>
        <NavLink to="register" className={({isActive}) => `${tab} ${isActive?active:idle}`}>Register</NavLink>
        <NavLink to="logout"   className={({isActive}) => `${tab} ${isActive?active:idle}`}>Logout</NavLink>
      </div>
      <Outlet />
    </main>
  );
}

export const registerAPI = async ({ username, email, password, role = "user" }) => {
  const res = await api.post("/pizzas/register", { username, email, password, role });
  return res.data; // { message: "success" }
};