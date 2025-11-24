import React from "react";
import { Link } from "react-router-dom";
import { SauceWaves, ToppingsConfetti } from "../components/HeroDecor";

export default function Home() {
  return (
    <section className="relative overflow-hidden bg-gradient-to-br from-red-600 to-orange-500 text-white">
      {/* soft vignette to deepen the center */}
      <div className="pointer-events-none absolute inset-0 opacity-[0.12] mix-blend-multiply">
        <div className="w-full h-full bg-[radial-gradient(130%_90%_at_50%_20%,transparent,rgba(0,0,0,0.45))]" />
      </div>

      {/* new decorative layer */}
      <SauceWaves />
      <ToppingsConfetti />

      {/* content */}
      <div className="relative max-w-7xl mx-auto px-4 py-20 sm:py-24 text-center">
        <h1 className="text-4xl sm:text-5xl lg:text-6xl font-extrabold mb-4 drop-shadow-[0_2px_6px_rgba(0,0,0,.25)]">
          Welcome to PizzaTime 🍕
        </h1>
        <p className="text-base sm:text-lg mb-8 opacity-95">
          Hot, fresh, and straight from the oven to your door
        </p>
        
        <Link
            to="/menu"
            className="inline-block px-6 py-3 bg-white text-red-600 rounded-full font-semibold
                 shadow-lg transition-transform duration-200 ease-in-out
                 hover:scale-110 hover:bg-red-50"
      >
        Order Now
      </Link>
      </div>
    </section>
  );
}