// src/pages/History.jsx
import React from "react";
import api from "../api/api"; // used to preload pizza catalog for names/photos
import { fetchHistory, fetchHistoryItem } from "../api/history";

/* --------- icons (no emoji) --------- */
const CartIcon = ({ className = "h-5 w-5" }) => (
  <svg viewBox="0 0 24 24" fill="none" stroke="currentColor"
       strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round"
       className={className}>
    <circle cx="9" cy="21" r="1.8" />
    <circle cx="18" cy="21" r="1.8" />
    <path d="M1 1h3l2.6 13.2a2 2 0 0 0 2 1.6h8.8a2 2 0 0 0 2-1.5L22 6H6" />
  </svg>
);

const CalendarIcon = ({ className = "h-4 w-4" }) => (
  <svg viewBox="0 0 24 24" fill="none" stroke="currentColor"
       strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round"
       className={className}>
    <rect x="3" y="4" width="18" height="18" rx="2" />
    <path d="M16 2v4M8 2v4M3 10h18" />
  </svg>
);

/* --------- utils --------- */
const formatDate = (d) => {
  if (!d) return "";
  if (typeof d === "object" && typeof d.seconds === "number") {
    return new Date(d.seconds * 1000).toLocaleString();
  }
  const t = new Date(d);
  return Number.isNaN(t.getTime()) ? String(d) : t.toLocaleString();
};

const money = (n) => `$${(Number(n) || 0).toFixed(2)}`;

/* --------- status pill (colors per request) --------- */
const StatusPill = ({ status }) => {
  const s = (status || "").toLowerCase().trim();
  let cls = "bg-gray-200 text-gray-800 dark:bg-zinc-800 dark:text-zinc-200";
  let text = status || "completed";

  if (s === "completed") {
    cls = "bg-emerald-100 text-emerald-800 dark:bg-emerald-900/30 dark:text-emerald-200";
    text = "completed";
  } else if (s === "in progress" || s === "progress" || s === "active") {
    cls = "bg-amber-100 text-amber-800 dark:bg-amber-900/30 dark:text-amber-200";
    text = "in progress";
  } else if (s === "canceled" || s === "cancelled") {
    cls = "bg-rose-100 text-rose-800 dark:bg-rose-900/30 dark:text-rose-200";
    text = "canceled";
  }

  return (
    <span className={`inline-flex items-center gap-1.5 rounded-full px-2.5 py-1 text-xs font-semibold ${cls}`}>
      <span className="h-1.5 w-1.5 rounded-full bg-current/70" />
      {text}
    </span>
  );
};

const ViewBtn = ({ onClick }) => (
  <button
    onClick={onClick}
    className="px-3 py-1.5 rounded-xl bg-red-600 text-white shadow hover:bg-red-700 hover:shadow-md hover:scale-[1.02] active:scale-95 transition"
  >
    View items
  </button>
);

/* --------- main page --------- */
export default function History() {
  const [orders, setOrders] = React.useState([]);
  const [loading, setLoading] = React.useState(true);
  const [err, setErr] = React.useState(null);

  // pizza catalog lookup (id -> {name, photo})
  const [catalog, setCatalog] = React.useState({});

  // modal
  const [modalOpen, setModalOpen] = React.useState(false);
  const [modalTitle, setModalTitle] = React.useState("");
  const [items, setItems] = React.useState([]);
  const [itemsLoading, setItemsLoading] = React.useState(false);
  const [itemsErr, setItemsErr] = React.useState(null);

  // load history + catalog
  React.useEffect(() => {
    (async () => {
      try {
        setLoading(true);
        setErr(null);

        // 1) history
        const h = await fetchHistory();
        const arr =
          h?.CartHistory ?? h?.cartHistory ?? h?.history ?? (Array.isArray(h) ? h : []);
const normalized = arr.map((it, idx) => ({
  cartId: it.cartId ?? it.CartId ?? it.cart_id ?? it.id ?? idx + 1,
  status:
    (it.status ?? it.Status) ??
    ((it.isActive ?? it.IsActive) ? "in progress" : "completed"),
  date: it.date ?? it.Date ?? null,
}));

// Make a per-user sequence number based on date (oldest → newest)
const sorted = [...normalized].sort((a, b) => {
  const ta = new Date(a.date).getTime() || 0;
  const tb = new Date(b.date).getTime() || 0;
  return ta - tb;
});
const seqByCartId = new Map(sorted.map((o, i) => [o.cartId, i + 1]));
const withNumbers = normalized.map((o) => ({
  ...o,
  orderNo: seqByCartId.get(o.cartId) || 1,
}));

setOrders(withNumbers);

        // 2) pizzas for name/photo lookup
        const res = await api.get("/pizzas/get");
        const raw = Array.isArray(res?.data?.pizzas)
          ? res.data.pizzas
          : Array.isArray(res?.data)
          ? res.data
          : [];
        const map = {};
        raw.forEach((p) => {
          const id = p.id ?? p.pizzaId ?? p.pizza_id;
          if (!id) return;
          map[id] = {
            name: p.name ?? p.pizza_name ?? "Pizza",
            photo: p.photo ?? p.image ?? "",
            price: Number(p.price ?? p.cost ?? 0),
          };
        });
        setCatalog(map);
      } catch (e) {
        console.error("history load failed", e);
        setErr(e?.response?.status ?? "unknown");
      } finally {
        setLoading(false);
      }
    })();
  }, []);

  const openItems = async (order) => {
    setModalOpen(true);
    setModalTitle(`Order #${order.orderNo} items`);
    setItems([]);
    setItemsErr(null);
    setItemsLoading(true);
    try {
      const data = await fetchHistoryItem(order.cartId);
      // backend may return one object or an array
      const list =
        data?.CartHistory ??
        data?.cartHistory ??
        (Array.isArray(data) ? data : [data]).filter(Boolean);

      const norm = list.map((x) => ({
        pizzaId: x?.PizzaId ?? x?.pizzaId ?? null,
        quantity: Number(x?.Quantity ?? x?.quantity ?? 0),
        cost: Number(x?.Cost ?? x?.cost ?? 0),
      }));
      setItems(norm);
    } catch (e) {
      console.error("items load failed", e);
      setItemsErr(e?.response?.status ?? "unknown");
    } finally {
      setItemsLoading(false);
    }
  };

  const closeItems = () => setModalOpen(false);

  /* ---------- render ---------- */
  if (loading) {
    return <main className="max-w-6xl mx-auto p-6">Loading history…</main>;
  }

  if (err) {
    const needAuth = err === 401 || err === 403;
    return (
      <main className="max-w-6xl mx-auto p-6">
        <div className="rounded-3xl border border-black/10 dark:border-white/10 bg-white/70 dark:bg-zinc-900/60 p-10 text-center shadow">
          <h1 className="text-2xl font-bold mb-2">
            {needAuth ? "Please log in to view your order history" : "History is unavailable"}
          </h1>
          <p className="text-gray-600 dark:text-gray-300">
            {needAuth ? "Sign in and try again." : "Please try again later."}
          </p>
        </div>
      </main>
    );
  }

  return (
    <main className="max-w-6xl mx-auto p-6 space-y-6">
      <header className="flex items-center justify-between">
        <h1 className="text-3xl font-extrabold tracking-tight">Order history</h1>
      </header>

      {/* Card grid */}
      <div className="grid gap-5 sm:grid-cols-2">
        {orders.length === 0 ? (
          <div className="col-span-full rounded-3xl bg-white/60 dark:bg-zinc-900/60 p-10 text-center">
            No orders yet.
          </div>
        ) : (
          orders.map((o) => (
            <div
              key={`${o.cartId}-${formatDate(o.date)}`}
              className="rounded-3xl bg-white/60 dark:bg-zinc-900/60 backdrop-blur-xl
                         border border-black/10 dark:border-white/10 p-5 shadow-lg
                         hover:shadow-2xl transition"
            >
              <div className="flex items-start justify-between">
                <div className="space-y-1.5">
                  <div className="flex items-center gap-2 text-xl font-bold">
                    <CartIcon className="h-5 w-5" />
                   <span>Order #{o.orderNo}</span>
                  </div>
                  <div className="flex items-center gap-2 text-sm text-gray-600 dark:text-gray-400">
                    <CalendarIcon className="h-4 w-4" />
                    <span>{formatDate(o.date)}</span>
                  </div>
                </div>
                <StatusPill status={(o.status || "").toString()} />
              </div>

              <div className="mt-5 flex items-center justify-between">
                <span className="text-sm text-gray-500 dark:text-gray-400">
                  View all items in this order
                </span>
                <ViewBtn onClick={() => openItems(o)} />
              </div>
            </div>
          ))
        )}
      </div>

      {/* Items modal */}
      {modalOpen && (
        <div
          className="fixed inset-0 z-50 grid place-items-center bg-black/70 p-4"
          onClick={closeItems}
        >
          <div
            onClick={(e) => e.stopPropagation()}
            className="w-full max-w-2xl rounded-3xl bg-white/85 dark:bg-zinc-900/85 backdrop-blur-xl
                       ring-1 ring-black/10 dark:ring-white/10 shadow-2xl overflow-hidden"
          >
            <div className="flex items-center justify-between px-6 py-4">
              <h3 className="text-xl font-semibold">{modalTitle}</h3>
              <button
                onClick={closeItems}
                className="px-3 py-1.5 rounded-lg text-sm text-gray-600 hover:bg-black/5 dark:text-gray-300 dark:hover:bg-white/10"
              >
                Close
              </button>
            </div>

            <div className="px-6 pb-6">
              {itemsLoading ? (
                <div className="py-10 text-center">Loading…</div>
              ) : itemsErr ? (
                <div className="py-10 text-center text-rose-600">
                  Failed to load (status {itemsErr})
                </div>
              ) : items.length === 0 ? (
                <div className="py-10 text-center text-gray-600 dark:text-gray-300">
                  No items in this order.
                </div>
              ) : (
                <>
                  <div className="overflow-x-auto rounded-2xl ring-1 ring-black/5 dark:ring-white/10">
                    <table className="min-w-full text-sm bg-white/60 dark:bg-zinc-900/60">
                      <thead className="text-left text-gray-600 dark:text-gray-300">
                        <tr>
                          <th className="px-4 py-3">Item</th>
                          <th className="px-4 py-3">Quantity</th>
                          <th className="px-4 py-3 text-right">Cost</th>
                        </tr>
                      </thead>
                      <tbody>
                        {items.map((row, idx) => {
                          const meta = catalog[row.pizzaId] || {};
                          return (
                            <tr
                              key={`${row.pizzaId}-${idx}`}
                              className="border-top border-black/5 dark:border-white/10"
                              style={{ borderTopWidth: 1 }}
                            >
                              <td className="px-4 py-3">
                                <div className="flex items-center gap-3">
                                  {meta.photo ? (
                                    <img
                                      src={meta.photo}
                                      alt={meta.name || `#${row.pizzaId}`}
                                      className="h-10 w-10 rounded-lg object-cover ring-1 ring-black/10 dark:ring-white/10"
                                    />
                                  ) : (
                                    <div className="h-10 w-10 rounded-lg bg-gray-200 dark:bg-zinc-800 grid place-items-center text-xs text-gray-500">
                                      #{row.pizzaId}
                                    </div>
                                  )}
                                  <div className="leading-tight">
                                    <div className="font-medium">
                                      {meta.name || `Pizza #${row.pizzaId}`}
                                    </div>
                                    <div className="text-xs text-gray-500">
                                      ID: {row.pizzaId}
                                    </div>
                                  </div>
                                </div>
                              </td>
                              <td className="px-4 py-3">
                                <span className="inline-block rounded-full bg-gray-100 dark:bg-zinc-800 px-2.5 py-1 text-xs font-semibold">
                                  x{row.quantity}
                                </span>
                              </td>
                              <td className="px-4 py-3 text-right font-semibold">
                                {money(row.cost)}
                              </td>
                            </tr>
                          );
                        })}
                      </tbody>
                    </table>
                  </div>

                  {/* footer total */}
                  <div className="mt-4 flex items-center justify-end gap-6">
                    <div className="text-sm text-gray-600 dark:text-gray-300">Total</div>
                    <div className="text-lg font-extrabold">
                      {money(items.reduce((s, r) => s + Number(r.cost || 0), 0))}
                    </div>
                  </div>
                </>
              )}
            </div>
          </div>
        </div>
      )}
    </main>
  );
}