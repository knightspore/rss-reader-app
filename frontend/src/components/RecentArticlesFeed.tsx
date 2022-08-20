import Error from "next/error";
import { useQuery } from "react-query";
import { fetchReadingList } from "../lib/queries";
import { Article } from "../types/backend-vo";
import ArticleCard from "./ArticleCard";

export default function RecentArticlesFeed() {
  const { isLoading, error, data } = useQuery(["readingList"], () =>
    fetchReadingList("parabyl")
  );

  if (isLoading) {
    return <div className="p-4 text-center">✨ Loading Feeds...</div>;
  }

  if (error | !data) {
    return <Error statusCode={500} />;
  }

  return data.map((article: Article) => {
    return <ArticleCard key={article.id} {...{ article }} />;
  });
}
