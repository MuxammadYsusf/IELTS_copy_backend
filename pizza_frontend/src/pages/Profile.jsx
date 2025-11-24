// src/pages/Profile.jsx
import React from "react";
import api from "../api/api";
import { updatePassword } from "../api/auth";
import { UserIcon } from "@heroicons/react/24/solid";

export default function Profile() {
  const [user, setUser] = React.useState(null);
  const [err, setErr] = React.useState(null);

  // change-password state
  const [oldPwd, setOldPwd] = React.useState("");
  const [newPwd, setNewPwd] = React.useState("");
  const [confirmPwd, setConfirmPwd] = React.useState("");
  const [pwdMsg, setPwdMsg] = React.useState(null);
  const [pwdLoading, setPwdLoading] = React.useState(false);

  // load profile
  React.useEffect(() => {
    (async () => {
      try {
        const res = await api.get("/pizzas/get/user-data");
        setUser({
          id: res.data?.id ?? res.data?.Id,
          username: res.data?.username ?? res.data?.Username,
          role: res.data?.role ?? res.data?.Role,
        });
      } catch (e) {
        console.error(e);
        setErr("Failed to load profile");
      }
    })();
  }, []);

  // submit password change
  const onChangePassword = async (e) => {
    e.preventDefault();
    setPwdMsg(null);

    if (!user?.id) return setPwdMsg("User not loaded.");
    if (!oldPwd || !newPwd || !confirmPwd) return setPwdMsg("Fill all fields.");
    if (newPwd !== confirmPwd) return setPwdMsg("New passwords do not match.");

    try {
      setPwdLoading(true);
      const resp = await updatePassword({
        userId: user.id,
        oldPassword: oldPwd,
        newPassword: newPwd,
        confirmPassword: confirmPwd,
      });
      setPwdMsg(resp?.message || "Password updated");
      setOldPwd("");
      setNewPwd("");
      setConfirmPwd("");
    } catch (err) {
      console.error(err);
      setPwdMsg(err?.response?.data?.error || "Update failed");
    } finally {
      setPwdLoading(false);
    }
  };

  if (err) return <main className="p-6">{err}</main>;
  if (!user) return <main className="p-6">Loading…</main>;

  return (
    <main className="max-w-4xl mx-auto p-6 space-y-6">
      {/* Profile card */}
      <div className="flex items-center gap-8 rounded-3xl bg-white/70 dark:bg-zinc-900/70 p-8 shadow-lg">
        {/* Avatar */}
        <div className="flex-shrink-0">
          <div className="h-24 w-24 rounded-full bg-gray-200 dark:bg-zinc-800 grid place-items-center ring-1 ring-black/10 dark:ring-white/10">
            <UserIcon className="h-12 w-12 text-gray-500" />
          </div>
        </div>

        {/* Data */}
        <div className="space-y-2">
          <h1 className="text-2xl font-bold">{user.username}</h1>
          <p className="text-gray-600 dark:text-gray-300">
            Role: <span className="font-medium">{user.role}</span>
          </p>
          <p className="text-gray-500 text-sm">User ID: {user.id}</p>
        </div>
      </div>

      {/* Change password */}
      <form
        onSubmit={onChangePassword}
        className="rounded-3xl bg-white/70 dark:bg-zinc-900/70 p-6 shadow-lg border border-black/10 dark:border-white/10 space-y-3"
      >
        <h3 className="font-semibold text-lg">Change password</h3>

        <input
          className="input w-full"
          type="password"
          placeholder="Current password"
          value={oldPwd}
          onChange={(e) => setOldPwd(e.target.value)}
        />
        <input
          className="input w-full"
          type="password"
          placeholder="New password"
          value={newPwd}
          onChange={(e) => setNewPwd(e.target.value)}
        />
        <input
          className="input w-full"
          type="password"
          placeholder="Confirm new password"
          value={confirmPwd}
          onChange={(e) => setConfirmPwd(e.target.value)}
        />

        {pwdMsg && <div className="text-sm text-gray-700 dark:text-gray-300">{pwdMsg}</div>}

        <button
          type="submit"
          disabled={pwdLoading}
          className="px-4 py-2 rounded-xl bg-red-500 text-white font-semibold disabled:opacity-60 hover:scale-[1.02] active:scale-95 transition"
        >
          {pwdLoading ? "Updating…" : "Update password"}
        </button>
      </form>
    </main>
  );
}