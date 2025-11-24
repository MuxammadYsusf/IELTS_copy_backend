import React from "react";
import { Link, NavLink, useLocation } from "react-router-dom";
import ThemeToggle from "./ThemeToggle";
import { useCart } from "../store/cart";
import { isAdmin } from "../utils/auth";

export default function Navbar() {
  const count = useCart((s) => s.count());
  const hasNew = useCart((s) => s.hasNew);
  const markSeen = useCart((s) => s.markSeen);
  const { pathname } = useLocation();
  const onCartPage = pathname.startsWith("/cart");

  React.useEffect(() => {
    if (onCartPage) markSeen();
  }, [onCartPage, markSeen]);

  const base  = "px-3 py-2 rounded transition";
  const idle  = "text-gray-700 hover:bg-gray-100 dark:text-gray-100/85 dark:hover:bg-white/5";
  const active= "text-red-500 dark:text-red-400";

  return (
    <nav className="bg-white dark:bg-zinc-900 shadow-sm">
      <div className="max-w-7xl mx-auto px-4 py-3 flex items-center justify-between">
        <Link to="/" className="text-2xl font-bold">PizzaTime <span aria-hidden>🍕</span></Link>

        <div className="flex items-center gap-4">
          <NavLink to="/" className={({isActive}) => `${base} ${isActive?active:idle}`}>Home</NavLink>
          <NavLink to="/menu" className={({isActive}) => `${base} ${isActive?active:idle}`}>Menu</NavLink>

          {/* Cart (with badge that only shows when something NEW was added) */}
          <NavLink
            to="/cart"
            onClick={() => markSeen()}
            className={({isActive}) => `${base} ${isActive?active:idle} relative`}
          >
            Cart
            {!onCartPage && hasNew && count > 0 && (
              <span
                className="absolute -top-2 -right-2 inline-flex min-w-5 items-center justify-center
                           rounded-full bg-rose-600 px-1.5 text-xs font-bold text-white"
                aria-live="polite"
              >
                {count}
              </span>
            )}
          </NavLink>

          <NavLink to="/order" className={({isActive}) => `${base} ${isActive?active:idle}`}>Order</NavLink>
          <NavLink to="/history" className={({isActive}) => `${base} ${isActive?active:idle}`}>History</NavLink>
          <NavLink to="/auth" className={({isActive}) => `${base} ${isActive?active:idle}`}>Auth</NavLink>

          {/* Profile button (circle with person icon) */}
          <NavLink
            to="/profile"
            className={({isActive}) =>
              `flex items-center justify-center h-9 w-9 rounded-full transition 
              ${isActive ? "bg-red-500 text-white" : "bg-gray-200 text-gray-600 hover:bg-gray-300"}`
            }
          >
            {/* Inline SVG user icon (no external lib needed) */}
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="currentColor"
              viewBox="0 0 24 24"
              className="h-5 w-5"
            >
              <path d="M12 12c2.7 0 4.8-2.1 4.8-4.8S14.7 2.4 12 2.4 7.2 4.5 7.2 7.2 9.3 12 12 12zm0 2.4c-3.2 0-9.6 1.6-9.6 4.8v2.4h19.2v-2.4c0-3.2-6.4-4.8-9.6-4.8z" />
            </svg>
          </NavLink>

          <ThemeToggle />
        </div>
      </div>
    </nav>
  );
}