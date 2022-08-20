import { useState } from "react";
import AddSubscriptionForm from "./AddSubscriptionForm";

export default function Layout({ subscriptions, articles }: { subscriptions: any, articles: any }) {
  const [open, setOpen] = useState(false);

  return (
    <div className="flex flex-col w-screen h-screen text-slate-200 bg-slate-800">
      <header className="flex justify-between p-4 align-middle bg-slate-700">
        <div>
          <h1 className="font-bold">📄 RSS Reader</h1>
        </div>
        <AddSubscriptionForm {...{ open, setOpen }} />
      </header>
      <div className="grid flex-initial grid-cols-1 md:grid-cols-7 overflow-clip">
        <section className="col-span-1 p-2 md:col-span-2">
          <h2 className="mb-2 font-medium text-md">Subscriptions</h2>
          {subscriptions}
        </section>
        <section
          id="recent-posts-feed"
          className="flex flex-col col-span-1 gap-2 p-2 overflow-y-scroll md:col-span-5"
        >
          <h2 className="mb-2 font-medium text-md">Reading List</h2>
          {articles}
        </section>
      </div>
    </div>
  );
}
