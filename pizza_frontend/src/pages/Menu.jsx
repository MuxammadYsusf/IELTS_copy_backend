import React from "react";
import PizzaList from "../components/PizzaList";

export default function MenuPage() {
  return (
    <main className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
      <h1 className="py-8 text-3xl font-extrabold tracking-tight">Our Menu</h1>
      <PizzaList />
    </main>
  );
}