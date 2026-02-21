import { NavLink } from "react-router-dom";

export default function Anchor({ url, children, ...props }) {

  const isExternal = /^https?:\/\//.test(url);

  if (isExternal) {
    return (
      <a
        href={url}
        target="_blank"
        rel="noreferrer"
        className="zag-text zag-transition hover:underline underline-offset-4"
        {...props}
      >
        {children}
      </a>
    );
  }

  return (
    <NavLink
      to={url}
      className={({ isActive }) =>
        [
          "zag-text zag-transition hover:underline underline-offset-4",
          isActive ? "underline" : "",
        ].join(" ")
      }
      {...props}
    >
      {children}
    </NavLink>
  );
}