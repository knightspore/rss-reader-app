import { useState, useEffect, Dispatch, SetStateAction } from "react";
import { Subscription } from "../types/backend-vo";
import Icon from "./Icon";

export default function SubscriptionCard({
  sub,
  filters,
  setFilters,
}: {
  sub: Subscription;
  filters: [] | [string];
  setFilters: any;
}) {
  const [active, setActive] = useState(true);

  useEffect(() => {
    if (filters.includes(sub.url)) {
      setActive(false);
    } else {
      setActive(true);
    }
  }, [filters, sub.url]);

  const handleClick = () => {
    if (active) {
      if (filters.length === 0) {
        setFilters([sub.title])
      } else {
      setFilters([sub.title, ...filters]);
      }
    } else if (!active) {
      setFilters(
        filters.map((f) => {
          return f != sub.title && f;
        })
      );
    }
  };
  return (
    <div className="flex cursor-pointer" onClick={handleClick}>
      <span
        className={`flex items-center p-px px-2 font-bold border-2 rounded-full gap-2 line-clamp border-slate-700/30 ${
          active ? "bg-slate-700/30" : "opacity-30"
        } hover:bg-slate-700 transition-all duration-150`}
      >
        <Icon src={sub.icon} />
        <p>{sub.title}</p>
      </span>
    </div>
  );
}
