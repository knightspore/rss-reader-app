import Error from "next/error";
import { useQuery } from "react-query";
import { fetchReadingList } from "../lib/queries";
import { UserEvent } from "../types/backend-module";
import { Article } from "../types/backend-vo";
import ArticleCard from "./ArticleCard";

export default function RecentArticlesFeed({
  setFocus,
}: {
  setFocus: Function;
}) {
  const e: UserEvent = { id: "parabyl" };
  const { isLoading, error, data } = useQuery(["readingList"], () =>
    fetchReadingList(e)
  );

  if (isLoading || !data) {
    return <div className="p-4 text-center">✨ Loading Feeds...</div>;
  }

  if (error) {
    return <Error statusCode={500} />;
  }

  return data.map((article: Article) => {
      return <ArticleCard key={article.id} {...{ article, setFocus }} />;
    }
  );
}
