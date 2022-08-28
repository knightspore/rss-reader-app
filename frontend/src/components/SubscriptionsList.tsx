import Error from "next/error";
import { Dispatch, SetStateAction, useState } from "react";
import { useQuery } from "react-query";
import { fetchSubscription } from "../lib/queries";
import { SubscriptionEvent, UserEvent } from "../types/backend-server";
import { Subscription } from "../types/backend-vo";
import SubscriptionCard from "./SubscriptionCard";

export default function SubscriptionsList({
  subscriptions,
  filters,
  setFilters,
}: {
  subscriptions: [] | [string];
  filters: [] | [string];
  setFilters: Dispatch<SetStateAction<[] | [string]>>;
}) {

  const e: SubscriptionEvent = {
    id: "https://ciaran.co.za/rss.xml",
    title: "Ciaran's Blog",
    url: "https://ciaran.co.za/rss.xml",
    userId: "parabyl",
  };

  const { isLoading, error, data } = useQuery(
    ["subscriptions"],
    () => fetchSubscription(e),
    {
      notifyOnChangeProps: ["data", "error"],
    }
  );

  if (isLoading || !data) {
    return <div className="p-4 text-center">✨ Loading Subscriptions...</div>;
  }

  if (error) {
    return <Error statusCode={500} />;
  }

  console.log(data)

  return <SubscriptionCard key={data.id} sub={data} {...{filters, setFilters}}/>

  // return data
  //   .sort((a: Subscription, b: Subscription) => a.title.localeCompare(b.title))
  //   .map((sub: Subscription) => {
  //     // TODO: Fix broken icons
  //     return (
  //       sub.icon && (
  //         <SubscriptionCard key={sub.id} {...{ sub, filters, setFilters }} />
  //       )
  //     );
  //   });
}
