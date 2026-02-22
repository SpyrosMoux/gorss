import { ArticleCard } from "../../components/ArticleCard.jsx";
import useLatestArticles from "../../hooks/articles/useLatestArticles.js";

export const Home = () => {
  const { data: latestArticleData } = useLatestArticles();

  return (
    <>
      <h1 className="text-4xl text-center underline zag-offset font-bold py-4">
        Latest Articles
      </h1>

      <div className="flex flex-col gap-6 py-8">
        {latestArticleData?.articles?.map((article) => (
          <ArticleCard
            key={article?.id}
            title={article.title}
            date={article.date}
            hyperlink={article?.link}
          />
        ))}
      </div>
    </>
  );
};