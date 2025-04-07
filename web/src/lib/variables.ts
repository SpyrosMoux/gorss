// Set any item to undefined to remove it from the site or to use the default value

export const GLOBAL = {
  // Site metadata
  username: "GoRSS",
  rootUrl: "https://gorss.spyrosmoux.com",
  shortDescription: "Modern RSS Aggregator",
  longDescription: "A modern RSS aggregator built with Go and Astro, bringing your favorite content together in one place",

  // Social media links
  githubProfile: "https://github.com/SpyrosMoux/gorss",
  twitterProfile: undefined,
  linkedinProfile: undefined,

  // Common text names used throughout the site
  articlesName: "Latest Posts",
  projectsName: "Feeds",
  viewAll: "View All",

  // Common descriptions used throughout the site
  noArticles: "No posts available yet.",
  noProjects: "No feeds available yet.",

  // Blog metadata
  blogTitle: "Latest Updates",
  blogShortDescription: "Stay up to date with your favorite content.",
  blogLongDescription:
    "Aggregated content from your favorite RSS feeds in one convenient place.",

  // Project metadata
  projectTitle: "Your Feeds",
  projectShortDescription:
    "Manage and organize your RSS feeds",
  projectLongDescription:
    "Add, remove, and organize your RSS feeds to create your personalized content stream.",

  // Profile image
  profileImage: undefined,

  // Menu items
  menu: {
    home: "/",
    feeds: "/feeds",
  },
};
