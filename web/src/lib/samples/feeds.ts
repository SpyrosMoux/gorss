export interface Feed {
  id: string;
  externalId: string;
  name: string;
  link: string;
}

export const sampleFeeds: Feed[] = [
  {
    id: "1",
    externalId: "techcrunch",
    name: "TechCrunch",
    link: "https://techcrunch.com/feed/"
  },
  {
    id: "2",
    externalId: "theverge",
    name: "The Verge",
    link: "https://www.theverge.com/rss/index.xml"
  },
  {
    id: "3",
    externalId: "hackernews",
    name: "Hacker News",
    link: "https://news.ycombinator.com/rss"
  }
]; 