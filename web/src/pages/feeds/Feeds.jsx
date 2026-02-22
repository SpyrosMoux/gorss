import { Link } from "react-router-dom";
import { useFeeds }  from "../../hooks/feeds/useFeeds";

export function Feeds() {
  const { data: feeds, isLoading, isError, error } = useFeeds(); 
  if (isLoading) return <div className="p-4">Loading feeds…</div>;
  if (isError) return <div className="p-4">Error: {error?.message ?? "Failed to load feeds"}</div>;
  if (!feeds?.length) return <div className="p-4">No feeds found.</div>;

  return (
    <div className="p-4">
      <h1 className="text-xl font-semibold mb-4">Feeds</h1>

      <ul className="space-y-2">
        {feeds.map((feed) => (
          <li key={feed.id} className="border rounded p-3">
            <div className="font-medium">{feed.title ?? feed.name ?? "Untitled feed"}</div>

            <Link className="underline text-sm" to={`/feeds/${feed.id}`}>
              View articles
            </Link>
          </li>
        ))}
      </ul>
    </div>
  );
}