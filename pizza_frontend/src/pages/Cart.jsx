import React, { useEffect } from "react";
import { Link, useNavigate } from "react-router-dom";
import { useCart } from "../store/cart";
import { FiTrash2, FiCheckCircle } from "react-icons/fi";

// AnimatedNumber (unchanged)
function AnimatedNumber({ value, format = (v) => v.toString(), className = "" }) {
  const [display, setDisplay] = React.useState(value);
  const [animating, setAnimating] = React.useState(false);

  React.useEffect(() => {
    if (value === display) return;
    setAnimating(true);
    const t = setTimeout(() => {
      setDisplay(value);
      setAnimating(false);
    }, 150);
    return () => clearTimeout(t);
  }, [value, display]);

  return (
    <span
      className={[
        "inline-block transition-all duration-300 ease-out will-change-transform",
        animating
          ? value > display
            ? "translate-y-[-8px] opacity-0 text-emerald-600"
            : "translate-y-[8px] opacity-0 text-rose-600"
          : "translate-y-0 opacity-100 text-inherit",
        className,
      ].join(" ")}
    >
      {format(display)}
    </span>
  );
}

export default function Cart() {
  const navigate = useNavigate();
  const items  = useCart((s) => s.items);
  const add    = useCart((s) => s.add);
  const dec    = useCart((s) => s.dec);
  const remove = useCart((s) => s.remove);
  const clear  = useCart((s) => s.clear);
  const total  = useCart((s) => s.total());
  const load   = useCart((s) => s.load);
  const markSeen = useCart((s) => s.markSeen); // 👈

  useEffect(() => {
    // load cart and clear the "new items" badge because the user is looking at it
    load().catch(console.error);
    markSeen();
  }, [load, markSeen]);

  if (!items || items.length === 0) {
    return (
      <main className="max-w-4xl mx-auto p-6">
        <div className="flex items-center justify-between mb-4">
          <h1 className="text-2xl font-bold">Your Cart</h1>
        </div>
        <div className="rounded-2xl border border-gray-200 dark:border-white/10 bg-white/70 dark:bg-zinc-900/60 p-8 text-center shadow">
          <h2 className="text-xl font-semibold mb-2">Your cart is empty</h2>
          <p className="text-gray-600 dark:text-gray-300 mb-6">
            Add some delicious pizzas to get started.
          </p>
          <Link
            to="/menu"
            className="inline-flex items-center justify-center rounded-xl bg-rose-600 px-5 py-2.5 
                       text-white font-medium shadow transition-transform duration-200 ease-in-out
                       hover:scale-110 hover:bg-rose-700"
          >
            Browse Menu
          </Link>
        </div>
      </main>
    );
  }

  return (
    <main className="max-w-4xl mx-auto p-6">
      <div className="flex items-center justify-between mb-4">
        <h1 className="text-2xl font-bold">Your Cart</h1>
        <button
          onClick={() => clear().catch(console.error)}
          aria-label="Clear cart"
          className="inline-flex h-10 w-10 items-center justify-center rounded-md bg-rose-600 text-white shadow-sm 
                     transition-transform duration-200 ease-in-out
                     hover:scale-110 hover:bg-rose-700 active:scale-95 focus:outline-none focus:ring-2 focus:ring-rose-400"
          title="Clear cart"
        >
          <FiTrash2 className="text-lg" />
        </button>
      </div>

      <ul className="space-y-4">
        {items.map((it, idx) => {
          const idOrName = it.id ?? it.name;
          const reactKey = it.id ?? it.name ?? `row-${idx}`;
          const qty = it.qty ?? 1;
          const price = Number(it.price) || 0;
          const lineTotal = price * qty;

          return (
            <li
              key={reactKey}
              className="flex items-center gap-4 rounded-xl bg-white/80 dark:bg-zinc-900/70
                         ring-1 ring-black/5 dark:ring-white/10 shadow p-4"
            >
              <img
                src={it.photo}
                alt={it.name}
                className="h-16 w-16 rounded object-cover"
                onError={(e) => (e.currentTarget.style.visibility = "hidden")}
              />
              <div className="flex-1">
                <div className="font-semibold text-gray-900 dark:text-gray-100">{it.name}</div>
                <div className="text-sm text-gray-500 dark:text-gray-400">
                  ${price.toFixed(2)} each
                </div>
              </div>

              <div className="flex items-center gap-2">
                <button
                  onClick={() => dec(idOrName).catch(console.error)}
                  className="px-2 py-1 rounded bg-gray-200 dark:bg-zinc-700
                             hover:bg-gray-300 dark:hover:bg-zinc-600 transition"
                  aria-label="Decrease quantity"
                >
                  −
                </button>
                <span className="w-8 text-center">{qty}</span>
                <button
                  onClick={() => add(it).catch(console.error)}
                  className="px-2 py-1 rounded bg-gray-200 dark:bg-zinc-700
                             hover:bg-gray-300 dark:hover:bg-zinc-600 transition"
                  aria-label="Increase quantity"
                >
                  +
                </button>
              </div>

              <div className="w-24 text-right font-semibold">
                <AnimatedNumber value={lineTotal} format={(v) => `$${v.toFixed(2)}`} />
              </div>

              <button
                onClick={() => remove(idOrName).catch(console.error)}
                className="text-rose-600 hover:underline"
              >
                Remove
              </button>
            </li>
          );
        })}
      </ul>

      <div className="mt-6 flex items-center justify-between">
        <div className="text-xl font-bold">
          Total:{" "}
          <AnimatedNumber
            value={total}
            format={(v) => `$${v.toFixed(2)}`}
            className="ml-1"
          />
        </div>
        <button
          onClick={() => navigate("/order")}
          className="inline-flex items-center gap-2 px-4 py-2 rounded-md bg-emerald-600 text-white 
                     shadow-sm transition-transform duration-200 ease-in-out
                     hover:scale-110 hover:bg-emerald-700 active:scale-100 
                     focus:outline-none focus:ring-2 focus:ring-emerald-400"
        >
          <FiCheckCircle className="text-lg" />
          Checkout
        </button>
      </div>
    </main>
  );
}