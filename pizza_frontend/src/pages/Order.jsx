// src/pages/Order.jsx
import React from "react";
import { Link, useNavigate } from "react-router-dom";
import { useCart } from "../store/cart";
import api from "../api/api";
import { getCartAPI } from "../api/cart";

const StatusBadge = ({ status }) => {
  const map = {
    idle:     { text: "Review your order", cls: "bg-gray-100 text-gray-800" },
    paying:   { text: "Confirming payment…", cls: "bg-amber-100 text-amber-800" },
    queued:   { text: "Queued", cls: "bg-blue-100 text-blue-800" },
    cooking:  { text: "Cooking", cls: "bg-indigo-100 text-indigo-800" },
    ontheway: { text: "On the way", cls: "bg-sky-100 text-sky-800" },
    delivered:{ text: "Delivered", cls: "bg-emerald-100 text-emerald-800" },
    canceled: { text: "Canceled", cls: "bg-rose-100 text-rose-800" },
  };
  const { text, cls } = map[status] ?? map.idle;
  return (
    <span className={`rounded-full px-3 py-1 text-sm font-semibold ${cls}`}>
      {text}
    </span>
  );
};

export default function OrderPage() {
  const navigate = useNavigate();

  // cart store
  const items = useCart((s) => s.items);
  const total = useCart((s) => s.total());
  const clear = useCart((s) => s.clear);

  // local state
  const [address, setAddress] = React.useState("");
  const [card, setCard] = React.useState("");
  const [errors, setErrors] = React.useState({});
  const [status, setStatus] = React.useState("idle");
  const [submitting, setSubmitting] = React.useState(false);

  // backend cartId
  const [cartId, setCartId] = React.useState(null);
  React.useEffect(() => {
    (async () => {
      try {
        const data = await getCartAPI();
        const id = data?.cartId ?? data?.id ?? null;
        setCartId(id);
      } catch (e) {
        console.error("failed to load cartId", e);
      }
    })();
  }, []);


  // validators
  const validate = () => {
    const e = {};
    if (!address.trim()) e.address = "Address is required";
    const digits = card.replace(/\s+/g, "");
    if (digits.length < 12) e.card = "Enter a valid card number";
    setErrors(e);
    return Object.keys(e).length === 0;
  };

  // user clicked "Order"
  const handleOrder = (e) => {
    e.preventDefault();
    if (!validate()) return;
    setSubmitting(true);
    setStatus("paying");
  };

const buildOrderItemsPayload = () => {
  const distinct = new Map(); // pizzaId -> totalQty

  items.forEach((it) => {
    const pizzaId = it.id ?? it.pizzaId ?? it.pizza_id;
    if (!pizzaId) return;
    const qty = Number(it.qty ?? it.quantity ?? 1);
    distinct.set(pizzaId, (distinct.get(pizzaId) ?? 0) + qty);
  });

  const orderItems = Array.from(distinct.entries()).map(([pizzaId, quantity]) => ({
    pizzaId,
    quantity,
  }));

  return {
    items: orderItems,
    limit: orderItems.length, // dynamic: each new distinct pizza
  };
};

  // user clicked "Confirm payment"
const handleConfirm = async () => {
  try {
    // Build the body your backend expects for OrderItem
    const { items: orderItems, limit } = buildOrderItemsPayload();

    const res = await api.post("/pizzas/order", {
      items: orderItems,
      limit,
    });

    if (res.data?.message === "already ordered") {
      setStatus("canceled"); // or show a toast
    } else {
      setStatus("queued");
      await runTracking();
    }
  } catch (e) {
    console.error("order failed", e);
    setStatus("canceled");
  } finally {
    setSubmitting(false);
  }
};

  // fake backend tracking updates
  const runTracking = async () => {
    const seq = ["queued", "cooking", "ontheway", "delivered"];
    for (const step of seq) {
      await new Promise((r) => setTimeout(r, 1200));
      setStatus(step);
    }
    await clear().catch(() => {});
  };

  // empty state
  if (!items || items.length === 0) {
    return (
      <main className="max-w-3xl mx-auto p-6">
        <div className="rounded-2xl border border-black/10 dark:border-white/10 bg-white/70 dark:bg-zinc-900/60 p-10 text-center shadow">
          <div className="text-6xl mb-4">🍽️</div>
          <h1 className="text-2xl font-bold mb-2">Nothing to order yet</h1>
          <p className="text-gray-600 dark:text-gray-300 mb-6">
            Pick your favorites from the menu and they’ll appear here.
          </p>
          <Link
            to="/menu"
            className="px-5 py-2.5 rounded-xl bg-red-500 text-white font-semibold
                       hover:scale-110 active:scale-95 transition transform-gpu
                       shadow-lg hover:shadow-xl"
          >
            Browse Menu
          </Link>
        </div>
      </main>
    );
  }

  return (
    <main className="max-w-5xl mx-auto p-6 space-y-6">
      <header className="flex items-center justify-between">
        <h1 className="text-3xl font-extrabold tracking-tight">Order</h1>
        <StatusBadge status={status} />
      </header>

      <section className="grid lg:grid-cols-2 gap-6">
        {/* Summary */}
        <div className="rounded-2xl border border-black/10 dark:border-white/10 bg-white/70 dark:bg-zinc-900/60 p-6 shadow">
          <h2 className="text-xl font-semibold mb-4">Summary</h2>
          <ul className="space-y-3">
            {items.map((it) => {
              const qty = it.qty ?? 1;
              const price = Number(it.price) || 0;
              return (
                <li
                  key={it.id ?? it.name}
                  className="flex items-center justify-between"
                >
                  <div className="flex items-center gap-3">
                    <img
                      src={it.photo}
                      alt=""
                      className="h-12 w-12 rounded object-cover"
                    />
                    <div>
                      <div className="font-medium">{it.name}</div>
                      <div className="text-sm text-gray-500 dark:text-gray-400">
                        x{qty}
                      </div>
                    </div>
                  </div>
                  <div className="font-semibold">
                    ${(qty * price).toFixed(2)}
                  </div>
                </li>
              );
            })}
          </ul>
          <div className="mt-4 border-t border-black/10 dark:border-white/10 pt-4 flex justify-between">
            <span className="font-semibold">Total</span>
            <span className="text-lg font-extrabold">
              ${total.toFixed(2)}
            </span>
          </div>
        </div>

        {/* Details */}
        <form
          onSubmit={handleOrder}
          className="rounded-2xl border border-black/10 dark:border-white/10 bg-white/70 dark:bg-zinc-900/60 p-6 shadow space-y-4"
        >
          <h2 className="text-xl font-semibold">Delivery & Payment</h2>

          <div>
            <label className="block text-sm mb-1">Address</label>
            <input
              className="input"
              placeholder="Street, house, apartment"
              value={address}
              onChange={(e) => setAddress(e.target.value)}
            />
            {errors.address && (
              <p className="text-sm text-rose-600 mt-1">{errors.address}</p>
            )}
          </div>

          <div>
            <label className="block text-sm mb-1">Card number</label>
            <input
              className="input"
              placeholder="1234 5678 9012 3456"
              value={card}
              onChange={(e) => setCard(e.target.value)}
            />
            {errors.card && (
              <p className="text-sm text-rose-600 mt-1">{errors.card}</p>
            )}
          </div>

          {status === "idle" && (
            <button
              type="submit"
              disabled={submitting}
              className="w-full px-4 py-2 rounded-xl bg-emerald-600 text-white font-semibold
                         hover:scale-[1.03] active:scale-95 transition transform-gpu disabled:opacity-60"
            >
              Order
            </button>
          )}

          {status === "paying" && (
            <button
              type="button"
              onClick={handleConfirm}
              className="w-full px-4 py-2 rounded-xl bg-indigo-600 text-white font-semibold
                         hover:scale-[1.03] active:scale-95 transition transform-gpu"
            >
              Confirm payment
            </button>
          )}
        </form>
      </section>
    </main>
  );
}