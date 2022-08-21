import type { NextPage } from "next";
import { useEffect, useState } from "react";
import { useQuery } from "react-query";
import Layout from "../components/Layout";
import RecentArticlesFeed from "../components/RecentArticlesFeed";
import SubscriptionsList from "../components/SubscriptionsList";
import { readArticle } from "../lib/queries";
import { ArticleEvent } from "../types/backend-module";

const Home: NextPage = () => {

  const [focus, setFocus] = useState<string | null>(null);
  const [text, setText] = useState<string|null>(null)

  useEffect(() => {
    if (focus != null) {
      const e: ArticleEvent = {
        userId: "parabyl",
        url: focus,
      }

      readArticle(e).then(text => setText(text))
    }
  }, [focus, text])

  return (
    <Layout
      subscriptions={<SubscriptionsList />}
      articles={<RecentArticlesFeed setFocus={setFocus} />}
			text={text}
    />
  );
};

export default Home;
