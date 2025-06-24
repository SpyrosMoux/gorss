export interface ArticlesResponse {
  articles: Article[];
}

interface Article {
  id: string;
  title: string;
  content: string;
  link: string;
  feedId: string;
  imageUrl: string;
  date: string;
}
