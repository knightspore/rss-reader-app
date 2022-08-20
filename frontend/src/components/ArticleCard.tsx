import { Article } from "../types/backend-vo";
import Icon from "./Icon";

export default function ArticleCard({
  article,
  setFocus,
}: {
  article: Article;
  setFocus: Function;
}) {

  return (
    <div
      key={article.url}
      className="px-2 py-4 text-slate-200"
      onClick={() => setFocus(article.id)}
    >
      <div className="flex items-center gap-2 mb-1 text-sm italic font-bold text-slate-300">
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
        <h3 className="text-lg font-medium">{article.title}</h3>
      </div>
    </div>
  );
}
