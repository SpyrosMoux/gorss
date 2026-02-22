import { useMemo } from "react";
import { DateTime } from "luxon";

export function ArticleCard({ article }) {
  const formattedDate = useMemo(() => {
    if (!article?.date) return "";

    const truncatedIso = article.date.split(".")[0] + "Z";
    return DateTime.fromISO(truncatedIso).toFormat("LLLL d, yyyy");
  }, [article?.date]);

  return (
    <a
      href={article.link}
      target="_blank"
      rel="noopener noreferrer"
      className={[
        "block",
        "zag-bg zag-text zag-border-b zag-transition",
        "border border-black p-6",
        "shadow-[3px_3px_0_0_black]",
        "hover:-translate-y-1 hover:shadow-[8px_8px_0_0_black]",
        "cursor-pointer select-none",
      ].join(" ")}
      aria-label={`Open article: ${article.title}`}
    >
      <h2 className="text-2xl font-bold leading-snug underline zag-offset">
        {article.title}
      </h2>
      <p className="zag-muted mt-2 font-mono text-sm">{formattedDate}</p>
    </a>
  );
}