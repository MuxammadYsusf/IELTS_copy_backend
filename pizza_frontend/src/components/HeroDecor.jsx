import React from "react";

/** Big blurred gradient blobs that hug the left/right edges */
export function SauceWaves() {
  return (
    <div className="pointer-events-none absolute inset-0 overflow-hidden">
      {/* LEFT blob */}
      <svg
        className="hidden md:block absolute -left-40 bottom-[-60px] w-[420px] opacity-[0.35]"
        viewBox="0 0 600 600"
        aria-hidden
      >
        <defs>
          <radialGradient id="sw-left" cx="45%" cy="55%" r="70%">
            <stop offset="0%" stopColor="#ff7a59" />
            <stop offset="60%" stopColor="#f1513d" />
            <stop offset="100%" stopColor="#d73b2e" />
          </radialGradient>
          <filter id="blur-40">
            <feGaussianBlur stdDeviation="40" />
          </filter>
        </defs>
        <circle cx="300" cy="320" r="240" fill="url(#sw-left)" filter="url(#blur-40)" />
      </svg>

      {/* RIGHT blob (mirrored hue) */}
      <svg
        className="hidden md:block absolute -right-40 bottom-[-60px] w-[420px] opacity-[0.35]"
        viewBox="0 0 600 600"
        aria-hidden
      >
        <defs>
          <radialGradient id="sw-right" cx="55%" cy="50%" r="70%">
            <stop offset="0%" stopColor="#ffb449" />
            <stop offset="60%" stopColor="#ff7a59" />
            <stop offset="100%" stopColor="#f1513d" />
          </radialGradient>
          <filter id="blur-40b">
            <feGaussianBlur stdDeviation="40" />
          </filter>
        </defs>
        <circle cx="300" cy="320" r="240" fill="url(#sw-right)" filter="url(#blur-40b)" />
      </svg>
    </div>
  );
}

/** Tiny floating “toppings” chips (pepperoni & basil) with gentle drift */
export function ToppingsConfetti() {
  const chip =
    "absolute rounded-full opacity-80 shadow-[0_2px_8px_rgba(0,0,0,.12)] will-change-transform";
  const leaf =
    "absolute rounded-[999px] opacity-75 shadow-[0_2px_8px_rgba(0,0,0,.12)] will-change-transform rotate-[15deg]";

  return (
    <div className="pointer-events-none absolute inset-0">
      {/* left side cluster */}
      <span className={`${chip} bg-[#b61f2e] w-4 h-4 left-10 md:left-20 bottom-24 animate-drift-slow`} />
      <span className={`${leaf} bg-[#2FA34F] w-3 h-5 left-24 md:left-32 bottom-40 animate-drift`} />
      <span className={`${chip} bg-[#b61f2e] w-3 h-3 left-36 bottom-28 animate-drift-fast`} />

      {/* right side cluster */}
      <span className={`${chip} bg-[#b61f2e] w-4 h-4 right-10 md:right-20 bottom-28 animate-drift`} />
      <span className={`${leaf} bg-[#2FA34F] w-3 h-5 right-28 md:right-36 bottom-44 animate-drift-slow`} />
      <span className={`${chip} bg-[#b61f2e] w-3 h-3 right-40 bottom-32 animate-drift-fast`} />
    </div>
  );
}