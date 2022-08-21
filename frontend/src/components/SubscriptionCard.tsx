import { Subscription } from "../types/backend-vo";
import Icon from "./Icon";

export default function SubscriptionCard({ sub }: { sub: Subscription }) {
  return (
    <div className="flex">
      <span className="flex items-center p-px px-2 text-sm font-bold border-2 rounded-full line-clamp border-slate-700 bg-slate-700 gap-2">
        <Icon src={sub["channel>icon"]} />
        <p>{sub.title}</p>
      </span>
    </div>
  );
}
