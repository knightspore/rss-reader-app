import type { NextPage } from "next";
import { useEffect, useState } from "react";
import Layout from "../components/Layout";
import RecentArticlesFeed from "../components/RecentArticlesFeed";
import SubscriptionsList from "../components/SubscriptionsList";
import { readArticle } from "../lib/queries";
import { ArticleEvent } from "../types/backend-module";

const Home: NextPage = () => {
  const [loading, setLoading] = useState<boolean>(false);
  const [focus, setFocus] = useState<string | null>(null);
  const [text, setText] = useState<string | null>(null);
  const [filters, setFilters] = useState<[string]|[]>([]);

  useEffect(() => {
    if (focus != null) {
      if (!text) {
        setLoading(true);
      }
      const e: ArticleEvent = {
        userId: "parabyl",
        url: focus,
      };

      readArticle(e).then((text) => {
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
    <Layout
      subscriptions={<SubscriptionsList {...{ filters, setFilters }} />}
      articles={<RecentArticlesFeed {...{ setFocus, filters }} />}
      {...{ closeArticle, text }}
    />
  );
};

export default Home;
