import { useParams } from "react-router-dom";
import { ArticleCard } from "../../components/ArticleCard.jsx";
import { useArticlesByFeed } from "../../hooks/articles/useArticlesByFeed";

export function FeedArticles() {
  const { feedId } = useParams();
  const { data: articles, isLoading, isError, error } = useArticlesByFeed(feedId);

  if (isLoading) return <div className="p-4">Loading articles…</div>;
  if (isError) return <div className="p-4">Error: {error?.message ?? "Failed to load articles"}</div>;
  if (!articles?.length) return <div className="p-4">No articles found.</div>;

  return (
    <div className="p-4">
      <h1 className="text-xl font-semibold mb-4">Feed Articles</h1>

      <div className="flex flex-col gap-6 py-8">
        {articles.map((article) => (
          <ArticleCard key={article.id} article={article} />
        ))}
      </div>
    </div>
  );
}