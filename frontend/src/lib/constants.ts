export const apiUrl: string = "http://localhost:8080";

export const api = {
  subscriptions: {
    add: apiUrl + "/subscriptions/add",
    remove: apiUrl + "/subscriptions/remove",
  },
  posts: {
    read: apiUrl + "/posts/read",
  },
};
