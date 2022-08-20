import { useState } from "react";
import AddSubscriptionForm from "./AddSubscriptionForm";
import ArticleContent from "./ArticleContent";

export default function Layout({
  subscriptions,
  articles,
  focus,
}: {
  subscriptions: any;
  articles: any;
  focus: string | null;
}) {
  return (
    <div className="flex flex-col w-screen h-screen text-slate-200 bg-slate-800">
      <header className="flex justify-between p-4 align-middle bg-slate-700">
        <div>
          <h1 className="font-bold">📄 RSS Reader</h1>
        </div>
        <AddSubscriptionForm />
      </header>
      <div className="grid flex-initial grid-cols-1 md:grid-cols-8 overflow-clip">
        <section className="col-span-1 p-2 md:col-span-1">
          <h2 className="mb-2 font-medium text-md">Subscriptions</h2>
          {subscriptions}
        </section>
        <section
          id="recent-posts-feed"
          className={`flex flex-col col-span-1 gap-2 p-2 overflow-y-scroll divide-y ${
            focus ? "md:col-span-3" : "md:col-span-7"
          } divide-slate-500 divide`}
        >
          <h2 className="mb-2 font-medium text-md">Reading List</h2>
          {articles}
        </section>
        {focus && <ArticleContent text={focus}/>}
      </div>
    </div>
  );
}
