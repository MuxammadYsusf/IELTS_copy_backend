// src/components/PizzaList.jsx
import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import PizzaCard from "./PizzaCard";
import api from "../api/api";
import { isAdmin } from "../utils/auth";

export default function PizzaList() {
  const [pizzas, setPizzas] = useState([]);
  const [loading, setLoading] = useState(true);
  const [errStatus, setErrStatus] = useState(null);

  // create modal (kept from your previous version)
  const [open, setOpen] = useState(false);
  const [mode, setMode] = useState("pizza"); // 'pizza' | 'type'
  const [form, setForm] = useState({ name: "", price: "13", photo: "" });
  const [typeName, setTypeName] = useState("");

  const navigate = useNavigate();

  const load = async () => {
    try {
      setLoading(true);
      const res = await api.get("/pizzas/get", { headers: { Accept: "application/json" } });
      const raw = Array.isArray(res?.data?.pizzas)
        ? res.data.pizzas
        : Array.isArray(res?.data)
        ? res.data
        : [];

      const list = raw.map((p, i) => ({
        id: p.id ?? p.pizzaId ?? p.pizza_id ?? p.PizzaId ?? p.PizzaID ?? i + 1,
        name: p.name ?? p.pizza_name ?? p.Name ?? "Pizza",
        price: Number(p.price ?? p.cost ?? p.Price ?? 0),
        photo: p.photo ?? p.image ?? p.Photo ?? "",
      }));

      setPizzas(list);
      setErrStatus(null);
    } catch (e) {
      const status = e?.response?.status ?? "unknown";
      console.error("Failed to fetch pizzas:", status, e);
      setErrStatus(status);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    load();
  }, []);

  const openCreator = () => {
    setMode("pizza");
    setForm({ name: "", price: "13", photo: "" });
    setTypeName("");
    setOpen(true);
  };
  const closeCreator = () => setOpen(false);

  // Create handlers (kept same endpoints you shared)
  const submitPizza = async (e) => {
    e.preventDefault();
    if (!form.name.trim()) return alert("Name is required");
    if (!form.price || isNaN(Number(form.price))) return alert("Price must be a number");
    try {
      await api.post("/admin/pizzas/create", {
        name: form.name.trim(),
        price: Number(form.price),
        typeId: 1, // or expose a select if you want; backend create needs a type
        photo: form.photo.trim(),
      });
      closeCreator();
      load();
    } catch (err) {
      console.error(err);
      alert("Create failed");
    }
  };

  const submitType = async (e) => {
    e.preventDefault();
    if (!typeName.trim()) return alert("Type name required");
    try {
      await api.post("/admin/pizzas/create/type", { name: typeName.trim() });
      closeCreator();
      load();
    } catch (err) {
      console.error(err);
      alert("Create type failed");
    }
  };

  if (loading) return <p className="p-4">Loading pizzas…</p>;

  if (errStatus) {
    const needLogin = errStatus === 401 || errStatus === 403;
    return (
      <div className="flex flex-col items-center text-center py-12">
        <h2 className="text-2xl font-bold mb-2">
          {needLogin ? "Please log in to view our menu" : "Menu is temporarily unavailable"}
        </h2>
        <p className="text-gray-600 dark:text-gray-300 mb-6">
          {needLogin
            ? "Sign in to wake the oven and see all the delicious pizzas."
            : "We’re warming the oven. Please try again shortly."}
        </p>
        {needLogin ? (
          <button
            onClick={() => navigate("/auth/login")}
            className="px-5 py-2.5 rounded-xl bg-gradient-to-r from-red-500 to-orange-500 text-white font-semibold hover:scale-105 active:scale-95 transition-transform"
          >
            Unlock Menu
          </button>
        ) : (
          <button
            onClick={() => window.location.reload()}
            className="px-5 py-2.5 rounded-xl bg-gray-900 text-white font-semibold hover:scale-105 active:scale-95 transition-transform dark:bg-white dark:text-gray-900"
          >
            Retry
          </button>
        )}
      </div>
    );
  }

  return (
    <>
      <div className="grid gap-8 p-2 sm:grid-cols-2 lg:grid-cols-3">
        {pizzas.map((p) => (
          <PizzaCard key={p.id ?? p.name} pizza={p} onAfterChange={load} />
        ))}

        {/* Add tile LAST (admins only) */}
        {isAdmin() && (
          <button
            onClick={openCreator}
            className="group flex aspect-[4/3] w-full items-center justify-center
                       rounded-2xl border-2 border-dashed border-white/20
                       bg-white/5 text-white/90 hover:bg-white/10 transition shadow-inner"
            title="Add new…"
          >
            <div className="flex flex-col items-center gap-2">
              <div className="grid place-items-center h-14 w-14 rounded-xl bg-white text-black
                              text-3xl font-bold shadow group-hover:scale-105 transition">+</div>
              <span className="font-semibold">Add new…</span>
            </div>
          </button>
        )}
      </div>

      {/* Create modal */}
      {open && (
        <div
          className="fixed inset-0 z-50 grid place-items-center bg-black/70 p-4"
          onClick={closeCreator}
        >
          <div
            className="w-full max-w-lg rounded-2xl bg-white/90 dark:bg-zinc-900/90
                       backdrop-blur shadow-2xl ring-1 ring-black/5 dark:ring-white/10"
            onClick={(e) => e.stopPropagation()}
          >
            {/* Header tabs */}
            <div className="flex items-center justify-between px-5 pt-4">
              <div className="flex gap-2">
                <button
                  onClick={() => setMode("pizza")}
                  className={`px-3 py-1.5 rounded-lg text-sm font-medium ${
                    mode === "pizza"
                      ? "bg-red-500 text-white"
                      : "bg-black/5 dark:bg-white/10 text-gray-700 dark:text-gray-200"
                  }`}
                >
                  Create Pizza
                </button>
                <button
                  onClick={() => setMode("type")}
                  className={`px-3 py-1.5 rounded-lg text-sm font-medium ${
                    mode === "type"
                      ? "bg-red-500 text-white"
                      : "bg-black/5 dark:bg-white/10 text-gray-700 dark:text-gray-200"
                  }`}
                >
                  Create Type
                </button>
              </div>
              <button
                onClick={closeCreator}
                className="rounded-lg px-2 py-1 text-sm text-gray-500 hover:bg-black/5 dark:hover:bg-white/10"
              >
                Close
              </button>
            </div>

            {/* Body */}
            <div className="p-5">
              {mode === "pizza" ? (
                <form onSubmit={submitPizza} className="space-y-4">
                  <div>
                    <label className="block text-sm mb-1">Name</label>
                    <input
                      className="input"
                      value={form.name}
                      onChange={(e) => setForm({ ...form, name: e.target.value })}
                      placeholder="Margherita"
                    />
                  </div>
                  <div className="grid grid-cols-2 gap-4">
                    <div>
                      <label className="block text-sm mb-1">Price</label>
                      <input
                        className="input"
                        value={form.price}
                        onChange={(e) => setForm({ ...form, price: e.target.value })}
                        inputMode="decimal"
                        placeholder="13"
                      />
                    </div>
                    <div>
                      <label className="block text-sm mb-1">Image URL</label>
                      <input
                        className="input"
                        value={form.photo}
                        onChange={(e) => setForm({ ...form, photo: e.target.value })}
                        placeholder="https://…"
                      />
                    </div>
                  </div>

                  <div className="pt-2 flex justify-end gap-2">
                    <button
                      type="button"
                      onClick={closeCreator}
                      className="px-4 py-2 rounded-xl bg-black/5 dark:bg-white/10"
                    >
                      Cancel
                    </button>
                    <button
                      type="submit"
                      className="px-4 py-2 rounded-xl bg-gradient-to-r from-red-500 to-orange-500 text-white font-semibold
                                 hover:scale-[1.02] active:scale-[0.99] transition-transform"
                    >
                      Create Pizza
                    </button>
                  </div>
                </form>
              ) : (
                <form onSubmit={submitType} className="space-y-4">
                  <div>
                    <label className="block text-sm mb-1">Type name</label>
                    <input
                      className="input"
                      value={typeName}
                      onChange={(e) => setTypeName(e.target.value)}
                      placeholder="Veggie / Pepperoni / …"
                    />
                  </div>

                  <div className="pt-2 flex justify-end gap-2">
                    <button
                      type="button"
                      onClick={closeCreator}
                      className="px-4 py-2 rounded-xl bg-black/5 dark:bg-white/10"
                    >
                      Cancel
                    </button>
                    <button
                      type="submit"
                      className="px-4 py-2 rounded-xl bg-gradient-to-r from-red-500 to-orange-500 text-white font-semibold
                                 hover:scale-[1.02] active:scale-[0.99] transition-transform"
                    >
                      Create Type
                    </button>
                  </div>
                </form>
              )}
            </div>
          </div>
        </div>
      )}
    </>
  );
}