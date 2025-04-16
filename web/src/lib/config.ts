const API_BASE_URL = import.meta.env.PUBLIC_API_BASE_URL || 'http://localhost:8080';

export const API_ENDPOINTS = {
  articles: {
    latest: `${API_BASE_URL}/api/articles/latest`,
    byFeed: (feedId: string) => `${API_BASE_URL}/api/articles/${feedId}`,
  },
  feeds: {
    list: `${API_BASE_URL}/api/feeds`,
    add: `${API_BASE_URL}/api/feeds`,
    delete: (feedId: string) => `${API_BASE_URL}/api/feeds/${feedId}`,
  },
}; 