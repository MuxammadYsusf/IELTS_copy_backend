import React from "react";
import { isAdmin } from "../utils/auth";

export default function AddPizzaTile({ onAdd }) {
  if (!isAdmin()) return null;
  return (
    <button
      onClick={onAdd}
      className="group flex aspect-[4/3] w-full items-center justify-center
                 rounded-2xl border-2 border-dashed border-white/20
                 bg-white/5 text-white/90 hover:bg-white/10 transition
                 shadow-inner"
      title="Add new pizza"
    >
      <div className="flex flex-col items-center gap-2">
        <div className="grid place-items-center h-14 w-14 rounded-xl bg-white text-black
                        text-3xl font-bold shadow group-hover:scale-105 transition">
          +
        </div>
        <span className="font-semibold">Add New Pizza</span>
      </div>
    </button>
  );
}