// src/api/history.js
import api from "./api";

/** Fetch user's cart/order history list */
export async function fetchHistory() {
  const res = await api.get("/pizzas/history");
  return res.data;
}

/** Fetch details for one history entry by id (currently backend expects ORDER id) */
export async function fetchHistoryItem(id) {
  const res = await api.get(`/pizzas/history/${id}`);
  return res.data;
}