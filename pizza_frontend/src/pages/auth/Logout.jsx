// src/pages/auth/Logout.jsx
import React from "react";
import api from "../../api/api";

export default function Logout() {
  const [msg, setMsg] = React.useState("");

  const doLogout = async () => {
    try {
      const token = localStorage.getItem("token");
      await api.post("/logout", null, {
        headers: token ? { Authorization: `Bearer ${token}` } : {},
      });
    } catch (e) {
      // ignore
    } finally {
      localStorage.removeItem("token");
      setMsg("Logged out.");
    }
  };

  return (
    <div className="space-y-3">
      <p>Click to log out from this device.</p>
      <button
        onClick={doLogout}
        className="w-full bg-gray-200 dark:bg-zinc-700 text-gray-900 dark:text-gray-100 py-2 rounded hover:bg-gray-300 dark:hover:bg-zinc-600 transition"
      >
        Logout
      </button>
      {msg && <p className="text-sm text-emerald-600">{msg}</p>}
    </div>
  );
}