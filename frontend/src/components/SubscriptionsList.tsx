import Error from "next/error";
import { useState } from "react";
import { useQuery } from "react-query";
import { fetchSubscription } from "../lib/queries";
import { UserEvent } from "../types/backend-module";
import { Subscription } from "../types/backend-vo";
import SubscriptionCard from "./SubscriptionCard";

export default function SubscriptionsList({
  filters,
  setFilters,
}: {
  filters: [];
  setFilters: Function;
}) {
  const e: UserEvent = { id: "parabyl" };

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

  return data
    .sort((a: Subscription, b: Subscription) => a.title.localeCompare(b.title))
    .map((sub: Subscription) => {
      // TODO: Fix broken icons
      return (
        sub.icon && (
          <SubscriptionCard key={sub.id} {...{ sub, filters, setFilters }} />
        )
      );
    });
}
