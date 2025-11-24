import React, { useEffect, useState } from "react";
import { PiMoonStarsBold, PiSunBold } from "react-icons/pi";

export default function ThemeToggle() {
  const [dark, setDark] = useState(() => localStorage.getItem("theme") === "dark");

  useEffect(() => {
    const root = document.documentElement;
    if (dark) {
      root.classList.add("dark");
      localStorage.setItem("theme", "dark");
    } else {
      root.classList.remove("dark");
      localStorage.setItem("theme", "light");
    }
  }, [dark]);

  return (
    <button
      aria-pressed={dark}
      onClick={() => setDark(d => !d)}
      className="group relative inline-flex items-center gap-3 rounded-full border border-gray-200 dark:border-white/10
                 bg-white/70 dark:bg-zinc-800/60 px-4 py-2 text-sm font-medium text-gray-800 dark:text-gray-100
                 shadow-sm backdrop-blur active:scale-[0.98] transition"
      title={dark ? "Switch to light mode" : "Switch to dark mode"}
    >
      {/* Track */}
      <div
        className={`relative h-7 w-20 overflow-hidden rounded-full transition-colors duration-700 ease-in-out
                    ${dark ? "bg-gradient-to-r from-violet-700/40 to-fuchsia-600/40"
                           : "bg-gradient-to-r from-amber-300/50 to-yellow-400/50"}`}
      >
        {/* Knob = Sun/Moon */}
        <div
          className={`absolute top-1/2 -translate-y-1/2 h-6 w-6 rounded-full shadow-md
                      flex items-center justify-center transition-[transform,background-color,box-shadow]
                      duration-700 ease-in-out
                      ${dark
                        ? "translate-x-[56px] bg-zinc-200 shadow-[0_0_18px_rgba(168,85,247,0.45)]"
                        : "translate-x-[6px] bg-yellow-300 shadow-[0_0_18px_rgba(255,215,0,0.55)]"}`}
        >
          {/* Icon inside the knob */}
          <span
            className={`transition-all duration-700 ease-in-out
                        ${dark ? "opacity-100 rotate-0" : "opacity-0 rotate-45"} text-violet-600`}
          >
            <PiMoonStarsBold size={16} />
          </span>
          <span
            className={`absolute transition-all duration-700 ease-in-out
                        ${dark ? "opacity-0 -rotate-45" : "opacity-100 rotate-0"} text-yellow-700`}
          >
            <PiSunBold size={16} />
          </span>

          {/* Soft “rays”/“glow” using a pseudo-ring */}
          <div
            className={`absolute inset-0 rounded-full -z-10 transition duration-700 ease-in-out
                        ${dark ? "ring-4 ring-violet-400/20" : "ring-4 ring-yellow-300/25"}`}
          />
        </div>
      </div>

      {/* Optional label */}
      <span className="min-w-[44px] text-center select-none">{dark ? "Dark" : "Light"}</span>
    </button>
  );
}