import { Box, Typography } from "@mui/material";
import { ArticleCard } from "../../components/ArticleCard.tsx";
import useLatestArticles from "../../hooks/article/useLatestArticles.ts";

export const NewsFeed = () => {
  const { data: latestArticleData } = useLatestArticles();

  return (
    <>
      <Typography
        variant={"h4"}
        textAlign={"center"}
        sx={{ textDecoration: "underline" }}
        py={4}
      >
        Latest Articles
      </Typography>
      <Box display={"flex"} flexDirection={"column"} gap={3} py={8}>
        {latestArticleData?.articles?.map((article) => (
          <ArticleCard
            key={article?.id}
            title={article.title}
            date={article.date}
            hyperlink={article?.link}
          />
        ))}
      </Box>
    </>
  );
};
