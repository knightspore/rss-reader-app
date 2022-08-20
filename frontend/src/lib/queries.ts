import {
  ArticleEvent,
  SubscriptionEvent,
  UserEvent,
} from "../types/backend-module";

export const fetchReadingList = async (event: UserEvent) => {
  const res = await fetch("http://127.0.0.1:1337/api/readinglist/get", {
    method: "POST",
    body: JSON.stringify(event),
  });
  return res.json();
};

export const createSubscription = async (event: SubscriptionEvent) => {
  const res = await fetch("http://127.0.0.1:1337/api/subscription/create", {
    method: "POST",
    body: JSON.stringify(event),
  });
  return res.json();
};

export const fetchSubscription = async (event: UserEvent) => {
  const res = await fetch("http://127.0.0.1:1337/api/subscription/get", {
    method: "POST",
    body: JSON.stringify(event),
  });
  return res.json();
};

export const readArticle = async (event: ArticleEvent) => {
  const res = await fetch("http://127.0.0.1:1337/api/article/read", {
    method: "POST",
    body: JSON.stringify(event),
  });
  return res.json();
};
