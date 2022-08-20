import { SubscriptionEvent } from "../types/backend-module";
import { Subscription } from "../types/backend-vo";

export const fetchReadingList = async (id: string) => {
  const res = await fetch("http://127.0.0.1:1337/api/readinglist/get", {
    method: "POST",
    body: JSON.stringify({ id }),
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

export const fetchSubscription = async (id: string) => {
  const res = await fetch("http://127.0.0.1:1337/api/subscription/get", {
    method: "POST",
    body: JSON.stringify({ id }),
  });
  return res.json();
};
