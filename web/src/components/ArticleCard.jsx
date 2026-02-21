import { useMemo } from "react";
import { DateTime } from "luxon";

export function ArticleCard({ title, date, hyperlink }) {
  const formattedDate = useMemo(() => {
    const truncatedIso = date.split(".")[0] + "Z";
    return DateTime.fromISO(truncatedIso).toFormat("LLLL d, yyyy");
  }, [date]);

  const handleClick = () => {
    window.open(hyperlink, "_blank", "noopener,noreferrer");
  };

  return (
    <article
      className={[
        "zag-bg zag-text zag-border-b zag-transition",
        "border border-black p-6",
        "shadow-[3px_3px_0_0_black]",
        "hover:-translate-y-1 hover:shadow-[8px_8px_0_0_black]",
        "cursor-pointer select-none",
      ].join(" ")}
      onClick={handleClick}
      role="link"
      tabIndex={0}
      onKeyDown={(e) => {
        if (e.key === "Enter" || e.key === " ") handleClick();
      }}
      aria-label={`Open article: ${title}`}
    >
      <h2 className="text-2xl font-bold leading-snug underline zag-offset">
        {title}
      </h2>
      <p className="zag-muted mt-2 font-mono text-sm">{formattedDate}</p>
    </article>
  );
}