// src/api/auth.js
import axios from "axios"; // or fetch, if you prefer
import api from "./api";

const API_URL = "http://localhost:8080/pizzas"; // adjust to your backend

export async function registerAPI({ username, email, password, role = "user" }) {
  const res = await axios.post(`${API_URL}/register`, {
    username,
    email,
    password,
    role,
  });
  return res.data;
}

export function updatePassword({ userId, oldPassword, newPassword, confirmPassword }) {
  return api
    .post("/pizzas/update-password", {
      userId,
      oldPassword,
      newPassword,
      confirmPassword,
    })
    .then((res) => res.data);
}