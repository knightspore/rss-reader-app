import type { NextPage } from "next";
import { useEffect, useState } from "react";
import { useQuery } from "react-query";
import Layout from "../components/Layout";
import RecentArticlesFeed from "../components/RecentArticlesFeed";
import SubscriptionsList from "../components/SubscriptionsList";
import { fetchUser, readArticle } from "../lib/queries";
import { ArticleEvent } from "../types/backend-server";
import { UserEvent } from "../types/backend-server";

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
      const ae: ArticleEvent = {
        userId: "parabyl",
        url: focus,
        id: "",
        parent: "",
      };

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

  const e: UserEvent = {id: "parabyl"};

  const { isLoading, error, data } = useQuery(
    ["user"],
    () => fetchUser(e),
    {
      notifyOnChangeProps: ['data', 'error']
    }
  )

  if (isLoading || !data) {
    <Layout
      subscriptions={"Loading..."}
      articles={"Loading..."}
      {...{ closeArticle, text }}
    />
  }

  if (error) {
    <Layout
      subscriptions={"Server Error"}
      articles={"Server Error"}
      {...{ closeArticle, text }}
    />
  }

  return data && (
    <Layout
      subscriptions={<SubscriptionsList subscriptions={data.subscriptions} {...{ filters, setFilters }} />}
      articles={<RecentArticlesFeed {...{ setFocus, filters }} />}
      {...{ closeArticle, text }}
    />
  );
};

export default Home;
