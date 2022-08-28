import {
  ArticleEvent,
  SubscriptionEvent,
  UserEvent,
} from "../types/backend-server";

const baseUrl = "127.0.0.1";
const port = "1337";

export const fetchReadingList = async (event: UserEvent) => {
  const res = await fetch(`http://${baseUrl}:${port}/api/readinglist/get`, {
    method: "POST",
    body: JSON.stringify(event),
  });
  return res.json();
};

export const createSubscription = async (event: SubscriptionEvent) => {
  const res = await fetch(`http://${baseUrl}:${port}/api/subscription/create`, {
    method: "POST",
    body: JSON.stringify(event),
  });
  return res.json();
};

export const fetchSubscription = async (event: SubscriptionEvent) => {
  const res = await fetch(`http://${baseUrl}:${port}/api/subscription/get`, {
    method: "POST",
    body: JSON.stringify(event),
  });
  return res.json();
};

export const readArticle = async (event: ArticleEvent) => {
  const res = await fetch(`http://${baseUrl}:${port}/api/article/read`, {
    method: "POST",
    body: JSON.stringify(event),
  });
  return res.json();
};

export const fetchUser = async (event: UserEvent) => {
  const res = await fetch(`http://${baseUrl}:${port}/api/user/get`, {
    method: "POST",
    body: JSON.stringify(event),
  });
  return res.json();
};