export const apiUrl: string = "http://192.168.0.109:1337";

export const api = {
  subscriptions: {
    add: apiUrl + "/subscriptions/add",
    remove: apiUrl + "/subscriptions/remove",
  },
  posts: {
    read: apiUrl + "/posts/read",
  },
};
