import React from "react";
import { useNavigate, Link } from "react-router-dom";
import { registerAPI } from "../../api/auth";

export default function Register() {
  const [username, setUsername] = React.useState("");
  const [email, setEmail]     = React.useState("");
  const [password, setPassword] = React.useState("");
  const [showPwd, setShowPwd] = React.useState(false);
  const [loading, setLoading] = React.useState(false);
  const [err, setErr]         = React.useState("");
  const [ok, setOk]           = React.useState("");
  const navigate = useNavigate();

  const onSubmit = async (e) => {
    e.preventDefault();
    setErr(""); setOk("");

    // tiny client-side validation
    if (!username.trim() || !email.trim() || !password.trim()) {
      setErr("Please fill all fields.");
      return;
    }
    if (!/^\S+@\S+\.\S+$/.test(email)) {
      setErr("Please enter a valid email.");
      return;
    }

    try {
      setLoading(true);
      const data = await registerAPI({ username, email, password, role: "user" });
      setOk(data?.message || "Registered!");
      // Redirect to login after a short delay
      setTimeout(() => navigate("/auth/login"), 800);
    } catch (e2) {
      const msg =
        e2?.response?.data?.error ||
        e2?.response?.data?.message ||
        e2?.message ||
        "Registration failed";
      setErr(msg);
    } finally {
      setLoading(false);
    }
  };

  return (
    <main className="max-w-md mx-auto p-6">
      <h1 className="text-2xl font-bold mb-6">Create account</h1>

      <form onSubmit={onSubmit} className="space-y-4">
        {err && <div className="rounded-md bg-rose-50 border border-rose-200 text-rose-700 px-3 py-2">{err}</div>}
        {ok &&  <div className="rounded-md bg-emerald-50 border border-emerald-200 text-emerald-700 px-3 py-2">{ok}</div>}

        <div>
          <label className="block text-sm mb-1">Username</label>
          <input
            className="input"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            placeholder="e.g. johnny"
            autoComplete="username"
          />
        </div>

        <div>
          <label className="block text-sm mb-1">Email</label>
          <input
            className="input"
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="you@example.com"
            autoComplete="email"
          />
        </div>

        <div>
          <label className="block text-sm mb-1">Password</label>
          <div className="relative">
            <input
              className="input pr-10"
              type={showPwd ? "text" : "password"}
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              placeholder="••••••••"
              autoComplete="new-password"
            />
            <button
              type="button"
              onClick={() => setShowPwd((s) => !s)}
              className="absolute right-2 top-1/2 -translate-y-1/2 text-sm text-gray-500 hover:text-gray-800"
            >
              {showPwd ? "Hide" : "Show"}
            </button>
          </div>
        </div>

        <button
          type="submit"
          disabled={loading}
          className="w-full rounded-xl bg-red-500 text-white font-semibold px-4 py-2.5
                     transition-transform duration-200 hover:scale-105 hover:bg-red-600
                     disabled:opacity-60 disabled:hover:scale-100"
        >
          {loading ? "Creating account..." : "Create account"}
        </button>
      </form>

      <p className="mt-4 text-sm text-gray-600">
        Already have an account?{" "}
        <Link className="text-red-600 hover:underline" to="/auth/login">
          Log in
        </Link>
      </p>
    </main>
  );
}