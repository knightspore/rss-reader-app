import Error from "next/error";
import { useQuery } from "react-query";
import { fetchSubscription } from "../lib/queries";
import { Subscription } from "../types/backend-vo";
import Icon from "./Icon";

export default function SubscriptionsList() {
  const { isLoading, error, data } = useQuery(["subscriptions"], () =>
    fetchSubscription("parabyl")
  );

  if (isLoading) {
    return <div className="p-4 text-center">✨ Loading Subscriptions...</div>;
  }

  if (error | !data) {
    return <Error statusCode={500} />;
  }

  return data.map((sub: Subscription) => {
    return <div key={sub.id} className="flex items-center gap-2"><Icon src={sub["channel>icon"]} />{sub.title}</div>;
  });
}
