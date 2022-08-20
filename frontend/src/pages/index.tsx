import type { NextPage } from "next";
import Layout from "../components/Layout";
import RecentArticlesFeed from "../components/RecentArticlesFeed";
import SubscriptionsList from "../components/SubscriptionsList";

const Home: NextPage = () => {
  return (
    <Layout articles={<RecentArticlesFeed/>} subscriptions={<SubscriptionsList/>} />
  );
};

export default Home;
