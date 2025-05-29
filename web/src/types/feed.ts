export interface GetAllFeedsResponse {
  feeds: Feed[];
}

interface Feed {
  id: string;
  name: string;
  link: string;
  createdAt: string;
  updatedAt: string;
}
