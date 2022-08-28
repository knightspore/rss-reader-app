import type { NextPage } from "next";
import { useContext, useEffect, useState } from "react";
import AddSubscriptionForm from "../components/AddSubscriptionForm";
import ArticleContent from "../components/ArticleContent";
import { UserContext } from "../components/context/UserProvider";
import SubscriptionsList from "../components/SubscriptionsList";
import { readArticle } from "../lib/fetch";
import { ArticleEvent } from "../types/backend-server";

const Home: NextPage = () => {

  const [loading, setLoading] = useState<boolean>(false);
  const [showSubs, setShowSubs] = useState(true);
  const [filters, setFilters] = useState<string[] | []>([]);
  const [focus, setFocus] = useState<string | null>(null);
  const [text, setText] = useState<string | null>(null);

  const subscriptions = useContext(UserContext)?.subscriptions || undefined;

  useEffect(() => {
    if (focus != null) {
      if (!text) {
        setLoading(true);
      }
      const ae: ArticleEvent[] = [
        {
          userId: "parabyl",
          url: focus,
          id: "",
          parent: "",
        },
      ];

      readArticle(ae).then((text) => {
        setText(text);
        setLoading(false);
      });
    }
  }, [focus, text]);

  useEffect(() => {
    if (loading == true) {
      setText(`\nLoading Article...`);
    }
  }, [loading]);

  function closeArticle() {
    setFocus(null);
    setText(null);
  }

  return (
    <div className="flex flex-col w-screen h-screen text-slate-200 bg-gradient-to-b from-slate-900 to-gray-900">
      <header className="flex justify-between p-2 align-middle">
        <div>
          <h1 className="font-bold">📄 Reader</h1>
        </div>
        <AddSubscriptionForm />
      </header>
      <div className="grid flex-initial grid-cols-1 md:grid-cols-8 overflow-clip">
        <section className="col-span-2 p-2">
          <h2
            onClick={() => setShowSubs(!showSubs)}
            className="mb-2 font-medium select-none text-md"
          >
            Subscriptions
          </h2>
          <div className="flex flex-row flex-wrap gap-2">
            {showSubs && subscriptions && (
              <SubscriptionsList {...{ filters, setFilters, subscriptions }} />
            )}
          </div>
        </section>
        <section
          className={`cols-span-1 gap-2 p-2 overflow-y-scroll ${
            text ? "md:col-span-3" : "md:col-span-6"
          }`}
        >
          <h2 className="mb-2 font-medium text-md">Reading List</h2>
          <div className="flex flex-col gap-2">{/* {articles} */}</div>
        </section>
        {text && <ArticleContent {...{ text, closeArticle }} />}
      </div>
    </div>
  );
};

export default Home;
