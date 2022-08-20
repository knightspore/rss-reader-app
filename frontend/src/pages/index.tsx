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


  return (
    <Layout
      subscriptions={<SubscriptionsList />}
      articles={<RecentArticlesFeed setFocus={setFocus} />}
			focus={focus}
    />
  );
};

export default Home;
