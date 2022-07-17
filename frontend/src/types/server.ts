/* Do not change, this code is generated from Golang structs */

export interface Feed {
  url: string;
  title: string;
}
export interface Event {
  user_id: string;
  feeds?: Feed[];
  read_post_url?: string;
}
export interface PostContent {
  title: string;
  byline: string;
  site: string;
  text: string;
  favicon: string;
  url: string;
}
