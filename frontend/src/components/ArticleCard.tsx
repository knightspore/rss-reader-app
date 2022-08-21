import { useState } from "react";
import { Article } from "../types/backend-vo";
import Icon from "./Icon";

export default function ArticleCard({
  article,
  setFocus,
}: {
  article: Article;
  setFocus: Function;
}) {
  const handleClick = () => {
    setFocus(article.id);
    article.read = true;
  };

  return (
    <div
      key={article.url}
      className={`pb-2 text-slate-200 ${article.read && "opacity-50"}`}
      onClick={handleClick}
    >
      <div className="flex items-center mb-1 text-sm italic font-medium gap-2 text-slate-300">
        <Icon src={article.icon} />
        <a href={article.url} target="_blank" rel="noreferrer noopener">
          <span>{article.parent}</span>
        </a>
        <time className="text-slate-500">
          {new Date(article.pubDate).toLocaleDateString("en-ZA", {
            day: "numeric",
            month: "short",
            year: "2-digit",
          })}
        </time>
      </div>
      <div className="py-1">
        <h3 className="font-medium text-md">{article.title}</h3>
      </div>
    </div>
  );
}
