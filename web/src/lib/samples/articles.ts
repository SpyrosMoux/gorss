export interface Article {
  id: string;
  externalId: string;
  title: string;
  content: string;
  link: string;
  imageUrl: string;
  createdAt: string;
  updatedAt: string;
}

// Sample articles for development and testing
export const sampleArticles: Article[] = [
  {
    id: "1",
    externalId: "ext1",
    title: "Sample Article 1",
    content: "This is a sample article content.",
    link: "https://example.com/article1",
    imageUrl: "https://picsum.photos/600/400",
    createdAt: "2024-01-01T00:00:00Z",
    updatedAt: "2024-01-01T00:00:00Z"
  },
  {
    id: "2",
    externalId: "ext2",
    title: "Sample Article 2",
    content: "This is another sample article content.",
    link: "https://example.com/article2",
    imageUrl: "",
    createdAt: "2024-01-02T00:00:00Z",
    updatedAt: "2024-01-02T00:00:00Z"
  }
]; 