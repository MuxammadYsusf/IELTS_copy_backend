// src/components/PizzaCard.jsx
import React, { useState } from "react";
import { FiShoppingCart, FiX, FiSettings } from "react-icons/fi";
import { useCart } from "../store/cart";
import api from "../api/api";
import { isAdmin } from "../utils/auth";

export default function PizzaCard({ pizza, onAfterChange }) {
  const [openPreview, setOpenPreview] = useState(false);
  const [adding, setAdding] = useState(false);

  // Admin UI state
  const [menuOpen, setMenuOpen] = useState(false);
  const [editOpen, setEditOpen] = useState(false);
  const [deleteOpen, setDeleteOpen] = useState(false);

  const [edit, setEdit] = useState({
    name: pizza?.name ?? "",
    price:
      typeof pizza?.price === "number"
        ? String(pizza.price)
        : String(parseFloat(pizza?.price ?? 0) || 0),
    photo: pizza?.photo ?? "",
  });

  const add = useCart((s) => s.add);

  const priceNumber =
    typeof pizza?.price === "number" ? pizza.price : parseFloat(pizza?.price || 0);

  // ----- Cart -----
  const handleAdd = async (e) => {
    e.stopPropagation();
    if (!pizza?.id) return alert("This item is missing an ID.");
    try {
      setAdding(true);
      await add(pizza);
    } catch (err) {
      console.error("Failed to add to cart:", err);
      alert("Could not add to cart. Try again.");
    } finally {
      setAdding(false);
    }
  };

  // ----- Admin actions -----
  const openEdit = (e) => {
    e.stopPropagation();
    setEdit({
      name: pizza?.name ?? "",
      price:
        typeof pizza?.price === "number"
          ? String(pizza.price)
          : String(parseFloat(pizza?.price ?? 0) || 0),
      photo: pizza?.photo ?? "",
    });
    setEditOpen(true);
    setMenuOpen(false);
  };

  const submitEdit = async (e) => {
    e.preventDefault();
    try {
      await api.put("/admin/pizzas/update", {
        id: pizza.id,
        name: edit.name.trim(),
        price: Number(edit.price),
        photo: edit.photo.trim(),
      });
      setEditOpen(false);
      onAfterChange?.();
    } catch (err) {
      console.error(err);
      alert("Update failed");
    }
  };

  const openDelete = (e) => {
    e.stopPropagation();
    setMenuOpen(false);
    setDeleteOpen(true);
  };
  const closeDelete = () => setDeleteOpen(false);

  const doDelete = async () => {
    try {
      await api.delete(`/admin/pizzas/delete/${pizza.id}`);
      setDeleteOpen(false);
      onAfterChange?.();
    } catch (err) {
      console.error(err);
      alert("Delete failed");
    }
  };

  return (
    <>
      <div
        className="group relative rounded-2xl overflow-hidden
                   ring-1 ring-black/10 dark:ring-white/10
                   bg-white/30 dark:bg-zinc-800/40 backdrop-blur-md
                   transition shadow-md hover:shadow-2xl"
      >
        {/* Admin gear (only for admins) */}
        {isAdmin() && (
          <div className="absolute left-3 top-3 z-20">
            <button
              onClick={(e) => {
                e.stopPropagation();
                setMenuOpen((s) => !s);
              }}
              className="p-2 rounded-xl bg-white/15 dark:bg-zinc-800/40
                         text-white hover:text-gray-200 backdrop-blur-md
                         shadow-md transition-transform hover:scale-110"
              title="Admin controls"
            >
              <FiSettings size={20} />
            </button>

            {menuOpen && (
              <div
                onClick={(e) => e.stopPropagation()}
                className="mt-2 w-40 overflow-hidden rounded-xl
                           bg-white/90 dark:bg-zinc-900/90 backdrop-blur
                           ring-1 ring-black/5 dark:ring-white/10 shadow-2xl"
              >
                <button
                  onClick={openEdit}
                  className="w-full px-3 py-2 text-left text-sm font-medium
                             hover:bg-black/[.06] dark:hover:bg-white/[.06] transition"
                >
                  ✏️ Update
                </button>
                <div className="mx-3 my-1 h-px bg-black/5 dark:bg-white/10" />
                <button
                  onClick={openDelete}
                  className="w-full px-3 py-2 text-left text-sm font-medium text-rose-600
                             hover:bg-rose-50/60 dark:hover:bg-rose-400/10 transition"
                >
                  🗑️ Delete
                </button>
              </div>
            )}
          </div>
        )}

        {/* Clickable area → preview */}
        <div
          role="button"
          tabIndex={0}
          onClick={() => setOpenPreview(true)}
          onKeyDown={(e) => e.key === "Enter" && setOpenPreview(true)}
          className="relative block w-full aspect-[4/3] overflow-hidden cursor-zoom-in"
          aria-label={`View ${pizza.name} larger`}
          title="View larger"
        >
          <img
            src={pizza.photo}
            alt={pizza.name}
            loading="lazy"
            className="h-full w-full object-cover transition-transform duration-300 group-hover:scale-[1.04]"
            onError={(e) => {
              e.currentTarget.src =
                "data:image/svg+xml;utf8," +
                encodeURIComponent(
                  `<svg xmlns='http://www.w3.org/2000/svg' width='600' height='450'><rect width='100%' height='100%' fill='#eee'/><text x='50%' y='50%' dominant-baseline='middle' text-anchor='middle' fill='#999' font-family='sans-serif'>No image</text></svg>`
                );
            }}
          />

          {/* Price badge */}
          <span
            className="absolute right-3 top-3 rounded-full px-3 py-1 text-sm font-semibold
                       bg-emerald-100/90 text-emerald-700 shadow
                       dark:bg-emerald-400/20 dark:text-emerald-200"
          >
            ${Number.isFinite(priceNumber) ? priceNumber.toFixed(2) : "0.00"}
          </span>

          {/* Footer gradient + CTA */}
          <div
            className="absolute inset-x-0 bottom-0 p-4
                       bg-gradient-to-t from-black/70 via-black/30 to-transparent
                       text-white"
          >
            <h3 className="mb-1 text-base font-semibold line-clamp-1">
              {pizza.name}
            </h3>
            <p className="mb-3 text-sm text-white/80">Freshly baked with love ❤️</p>

            <button
              onClick={handleAdd}
              disabled={adding}
              className="inline-flex items-center gap-2 rounded-xl bg-rose-500/95 px-4 py-2
                         text-sm font-medium text-white shadow 
                         transition-transform duration-200 ease-in-out
                         hover:scale-110 hover:bg-rose-600 active:scale-100
                         disabled:opacity-70 disabled:cursor-not-allowed"
            >
              <FiShoppingCart className="text-base" />
              {adding ? "Adding..." : "Add to Cart"}
            </button>
          </div>
        </div>
      </div>

      {/* Fullscreen preview modal */}
      {openPreview && (
        <div
          className="fixed inset-0 z-50 flex items-center justify-center bg-black/80 p-4"
          onClick={() => setOpenPreview(false)}
        >
          <div
            className="relative max-h-[90vh] w-full max-w-5xl"
            onClick={(e) => e.stopPropagation()}
          >
            <img
              src={pizza.photo}
              alt={pizza.name}
              className="h-full w-full rounded-2xl object-contain shadow-2xl"
            />
            <button
              onClick={() => setOpenPreview(false)}
              className="absolute right-3 top-3 rounded-full bg-white/90 p-2 text-gray-800 shadow
                         hover:bg-white dark:bg-zinc-800 dark:text-gray-100 dark:hover:bg-zinc-700"
              title="Close"
            >
              <FiX size={18} />
            </button>
          </div>
        </div>
      )}

      {/* Admin Update modal */}
      {editOpen && (
        <div
          className="fixed inset-0 z-50 grid place-items-center bg-black/70 p-4"
          onClick={() => setEditOpen(false)}
        >
          <div
            onClick={(e) => e.stopPropagation()}
            className="w-full max-w-lg rounded-2xl bg-white/90 dark:bg-zinc-900/90
                       backdrop-blur shadow-2xl ring-1 ring-black/5 dark:ring-white/10"
          >
            <div className="flex items-center justify-between px-5 pt-4">
              <h3 className="text-lg font-semibold">Update Pizza</h3>
              <button
                onClick={() => setEditOpen(false)}
                className="rounded-lg px-2 py-1 text-sm text-gray-500 hover:bg-black/5 dark:hover:bg-white/10"
              >
                Close
              </button>
            </div>

            <form onSubmit={submitEdit} className="p-5 space-y-4">
              <div className="grid grid-cols-2 gap-4">
                <div>
                  <label className="block text-sm mb-1">Name</label>
                  <input
                    className="input"
                    value={edit.name}
                    onChange={(e) => setEdit({ ...edit, name: e.target.value })}
                    placeholder="Margherita"
                  />
                </div>
                <div>
                  <label className="block text-sm mb-1">Price</label>
                  <input
                    className="input"
                    value={edit.price}
                    onChange={(e) => setEdit({ ...edit, price: e.target.value })}
                    inputMode="decimal"
                    placeholder="13"
                  />
                </div>
              </div>

              <div>
                <label className="block text-sm mb-1">Image URL</label>
                <input
                  className="input"
                  value={edit.photo}
                  onChange={(e) => setEdit({ ...edit, photo: e.target.value })}
                  placeholder="https://…"
                />
              </div>

              {edit.photo && (
                <div className="rounded-xl overflow-hidden ring-1 ring-black/5 dark:ring-white/10">
                  <img
                    src={edit.photo}
                    alt="Preview"
                    className="w-full h-40 object-cover"
                    onError={(e) => (e.currentTarget.style.display = "none")}
                  />
                </div>
              )}

              <div className="pt-2 flex justify-end gap-2">
                <button
                  type="button"
                  onClick={() => setEditOpen(false)}
                  className="px-4 py-2 rounded-xl bg-black/5 dark:bg-white/10"
                >
                  Cancel
                </button>
                <button
                  type="submit"
                  className="px-4 py-2 rounded-xl bg-gradient-to-r from-red-500 to-orange-500 text-white font-semibold
                             hover:scale-[1.02] active:scale-[0.99] transition-transform"
                >
                  Save Changes
                </button>
              </div>
            </form>
          </div>
        </div>
      )}

      {/* Admin Delete modal */}
      {deleteOpen && (
        <div
          className="fixed inset-0 z-50 grid place-items-center bg-black/70 p-4"
          onClick={closeDelete}
        >
          <div
            onClick={(e) => e.stopPropagation()}
            className="w-full max-w-md rounded-2xl bg-white/90 dark:bg-zinc-900/90
                       backdrop-blur shadow-2xl ring-1 ring-black/5 dark:ring-white/10"
          >
            <div className="px-5 pt-5">
              <h3 className="text-lg font-semibold">Delete pizza</h3>
              <p className="mt-2 text-sm text-gray-600 dark:text-gray-300">
                Are you sure you want to delete{" "}
                <span className="font-medium">{pizza.name}</span>? This action cannot be undone.
              </p>
            </div>

            <div className="p-5 flex justify-end gap-2">
              <button
                onClick={closeDelete}
                className="px-4 py-2 rounded-xl bg-black/5 dark:bg-white/10"
              >
                Cancel
              </button>
              <button
                onClick={doDelete}
                className="px-4 py-2 rounded-xl bg-rose-600 text-white font-semibold
                           hover:bg-rose-700 transition"
              >
                Delete
              </button>
            </div>
          </div>
        </div>
      )}
    </>
  );
}