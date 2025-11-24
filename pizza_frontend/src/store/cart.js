import { create } from "zustand";
import {
  addToCartAPI,
  decreaseCartAPI,
  getCartAPI,
  removeFromCartAPI,
  clearCartAPI,
} from "../api/cart";

// Normalize backend item -> frontend
function normalizeItem(it) {
  return {
    id: it.pizzaId ?? it.pizza_id ?? it.id,
    name: it.name ?? it.pizza_name ?? it.Name,
    price: Number(it.cost ?? it.price ?? it.Price ?? 0),
    photo: it.photo ?? it.image ?? it.Photo ?? "",
    qty: Number(it.quantity ?? it.qty ?? it.Quantity ?? 1),
    typeId: it.pizzaTypeId ?? it.pizza_type_id ?? it.typeId ?? 1,
  };
}

export const useCart = create((set, get) => ({
  items: [],
  serverTotal: 0,

  // 🔔 show a badge when new items were added
  hasNew: false,

  // Fetch current cart from backend
  load: async () => {
    const data = await getCartAPI();
    const raw = data?.cartItems ?? data?.items ?? data ?? [];
    const items = Array.isArray(raw) ? raw.map(normalizeItem) : [];
    const total = Number(data?.totalCost ?? 0);
    set({ items, serverTotal: total });
  },

  // Mark notification as seen (called when user opens the cart)
  markSeen: () => set({ hasNew: false }),

  // Add or increase
  add: async (pizza) => {
    if (!pizza?.id) return;
    await addToCartAPI({ pizzaId: pizza.id, pizzaTypeId: pizza.typeId ?? 1, qty: 1 });
    await get().load();
    set({ hasNew: true }); // 🔔 set badge after successful add
  },

  // Decrease by 1
  dec: async (idOrName) => {
    const found = get().items.find((x) => (x.id ?? x.name) === idOrName);
    if (!found) return;
    await decreaseCartAPI({ pizzaId: found.id, qty: 1 });
    await get().load();
  },

  // Remove a single line
  remove: async (idOrName) => {
    const found = get().items.find((x) => (x.id ?? x.name) === idOrName);
    if (!found) return;
    await removeFromCartAPI({ pizzaId: found.id });
    await get().load();
  },

  // Clear all
  clear: async () => {
    await clearCartAPI();
    await get().load();
    set({ hasNew: false }); // no badge if cart is empty
  },

  // Derived
  count: () => get().items.reduce((n, x) => n + (x.qty ?? 1), 0),
  total: () => get().items.reduce((s, x) => s + (Number(x.price) || 0) * (x.qty ?? 1), 0),
}));