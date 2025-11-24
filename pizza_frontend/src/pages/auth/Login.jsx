import React, { useState } from "react";
import api from "../../api/api";
import { useNavigate } from "react-router-dom";

export default function Login() { // <-- default here
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const handleLogin = async (e) => {
    e.preventDefault();
    try {
      const res = await api.post("/pizzas/login", { username, password });
      localStorage.setItem("token", res.data.token);
      navigate("/menu");
    } catch (err) {
      console.error("Login failed", err);
      alert("Invalid credentials");
    }
  };

return (
  <div className="max-w-sm mx-auto mt-10 bg-white dark:bg-zinc-900 p-6 shadow-md rounded">
    <h2 className="text-gray-800 dark:text-gray-100 font-bold">Login</h2>
    <form onSubmit={handleLogin} className="space-y-4">
      <input
        type="text"
        placeholder="Username"
        value={username}
        onChange={(e) => setUsername(e.target.value)}
        className="w-full rounded border px-3 py-2 
                   bg-white text-gray-900
                   dark:bg-zinc-800 dark:text-gray-100 
                   dark:border-zinc-700"
      />
      <input
        type="password"
        placeholder="Password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        className="w-full rounded border px-3 py-2 
                   bg-white text-gray-900
                   dark:bg-zinc-800 dark:text-gray-100 
                   dark:border-zinc-700"
      />
      <button
        type="submit"
        className="w-full bg-red-500 text-white py-2 rounded hover:bg-red-600"
      >
        Login
      </button>
    </form>
  </div>
);
}