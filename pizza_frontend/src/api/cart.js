// src/api/cart.js
import api from "./api";

// Add or increase quantity by 1
export const addToCartAPI = async ({ pizzaId, pizzaTypeId = 1, qty = 1 }) => {
  const res = await api.post("/pizzas/cart", {
    items: [{ pizzaId, pizzaTypeId, quantity: qty }],
  });
  return res.data;
};

// Decrease quantity by 1 (backend uses pizzaId + quantity)
export const decreaseCartAPI = async ({ pizzaId, qty = 1 }) => {
  const res = await api.put("/pizzas/decrease", {
    pizzaId,
    quantity: qty,
  });
  return res.data;
};

// Current cart
export const getCartAPI = async () => {
  const res = await api.get("/pizzas/cart");
  return res.data; // expect { cartItems: [...] , totalCost?: number } OR { items: [...] }
};

// Remove a single pizza line
export const removeFromCartAPI = async ({ pizzaId }) => {
  const res = await api.delete(`/pizzas/cart/${pizzaId}`);
  return res.data;
};

// Clear all
export const clearCartAPI = async () => {
  const res = await api.delete("/pizzas/cart");
  return res.data;
};