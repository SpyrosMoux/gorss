import React from "react";
import { ROUTES, GITHUB_PROFILE } from "../../../config/routes";
import Anchor from "../../../components/Anchor";

const NAV_ITEMS = [
  { label: "Home", to: ROUTES.home },
  { label: "Feeds", to: ROUTES.feeds },
];

function ThemeToggle({ onClick }) {
  return (
    <button
      type="button"
      className="zag-text zag-transition font-mono text-sm underline underline-offset-4"
      aria-label="Toggle theme"
      onClick={onClick}
    >
      theme
    </button>
  );
}

export default function Header() {
  const [isOpen, setIsOpen] = React.useState(false);
  const [isMobile, setIsMobile] = React.useState(() =>
    window.matchMedia("(max-width: 640px)").matches
  );

  React.useEffect(() => {
    const mq = window.matchMedia("(max-width: 640px)");

    const sync = () => {
      setIsMobile(mq.matches);
      // leaving mobile -> force nav hidden and reset open state
      if (!mq.matches) setIsOpen(false);
    };

    sync();
    mq.addEventListener?.("change", sync);
    window.addEventListener("resize", sync);

    return () => {
      mq.removeEventListener?.("change", sync);
      window.removeEventListener("resize", sync);
    };
  }, []);

  return (
    <header className="zag-bg zag-border-b zag-transition sticky top-0 w-full z-10">
      {/* Mobile top bar with hamburger */}
      <div className="zag-bg zag-transition sm:hidden relative z-50 py-4 flex items-center">
        <button
          type="button"
          className="px-4"
          aria-label="Toggle navigation menu"
          onClick={() => setIsOpen((v) => !v)}
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="32"
            height="32"
            viewBox="0 0 512 512"
            className="zag-fill zag-transition"
          >
            <path d="M80 96h352v32H80zm0 144h352v32H80zm0 144h352v32H80z" />
          </svg>
        </button>
      </div>

      {/* Nav */}
      <nav
        className={[
          "zag-bg zag-border-b zag-transition fixed sm:relative inset-x-0 top-0 h-auto",
          "sm:px-4 flex justify-between flex-col gap-8 py-4 text-xl",
          "sm:flex-row max-w-2xl mx-auto sm:pt-4 sm:border-none",
        ].join(" ")}
        style={{
          transform: isMobile
            ? isOpen
              ? "translateY(0)"
              : "translateY(-100%)"
            : "translateY(0)",
        }}
      >
        <div className="flex items-center gap-2 px-4 sm:px-0">
          <span className="text-2xl font-bold">GoRSS</span>
        </div>

        <div className="flex flex-col font-mono font-medium gap-4 sm:flex-row px-4 sm:px-0 mt-16 sm:mt-0">
          {NAV_ITEMS.map((item) => (
            <Anchor key={item.to} url={item.to} onClick={() => setIsOpen(false)}>
              {item.label}
            </Anchor>
          ))}
        </div>

        <div className="flex gap-4 justify-between px-4 sm:px-0">
          {GITHUB_PROFILE ? (
            <Anchor
              url={GITHUB_PROFILE}
              aria-label="GoRSS GitHub Repository"
              onClick={() => setIsOpen(false)}
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="32"
                height="32"
                viewBox="0 0 24 24"
                className="zag-fill zag-transition"
              >
                <path d="M12 2A10 10 0 0 0 2 12c0 4.42 2.87 8.17 6.84 9.5c.5.08.66-.23.66-.5v-1.69c-2.77.6-3.36-1.34-3.36-1.34c-.46-1.16-1.11-1.47-1.11-1.47c-.91-.62.07-.6.07-.6c1 .07 1.53 1.03 1.53 1.03c.87 1.52 2.34 1.07 2.91.83c.09-.65.35-1.09.63-1.34c-2.22-.25-4.55-1.11-4.55-4.92c0-1.11.38-2 1.03-2.71c-.1-.25-.45-1.29.1-2.64c0 0 .84-.27 2.75 1.02c.79-.22 1.65-.33 2.5-.33s1.71.11 2.5.33c1.91-1.29 2.75-1.02 2.75-1.02c.55 1.35.2 2.39.1 2.64c.65.71 1.03 1.6 1.03 2.71c0 3.82-2.34 4.66-4.57 4.91c.36.31.69.92.69 1.85V21c0 .27.16.59.67.5C19.14 20.16 22 16.42 22 12A10 10 0 0 0 12 2" />
              </svg>
            </Anchor>
          ) : null}

          {/* Theme toggle not implemented yet: this only closes the mobile menu */}
          <ThemeToggle onClick={() => setIsOpen(false)} />
        </div>
      </nav>
    </header>
  );
}