export interface Article {
  id: string;
  title: string;
  content: string;
  link: string;
  feedId: string;
  date: string;
}

// Sample articles for development and testing
export const sampleArticles: Article[] = [
  {
    id: "1",
    title: "Sample Article 1",
    content: "<div>Sample content 1</div>",
    link: "https://example.com/1",
    feedId: "feed-1",
    date: "2024-04-07T22:22:47.645525+03:00"
  },
  {
    id: "2",
    title: "Sample Article 2",
    content: "<div>Sample content 2</div>",
    link: "https://example.com/2",
    feedId: "feed-2",
    date: "2024-04-07T22:22:47.645525+03:00"
  }
]; 