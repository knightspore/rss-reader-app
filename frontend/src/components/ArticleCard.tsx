import Image from "next/image";
import { Article } from "../types/backend-vo";
import Icon from "./Icon";

export default function ArticleCard({ article }: { article: Article }) {
  return (
    <div key={article.url} className="px-2 py-4 rounded-md bg-slate-300 text-slate-900">
      <a href={article.url} target="_blank" rel="noreferrer noopener">
        <div className="flex items-center gap-2 mb-1 text-xs italic font-bold">
          <Icon src={article.icon}/>
          <span> 
            {article.parent}
          </span>
          <time className="text-slate-500">
            {new Date(article.pubDate).toLocaleDateString("en-ZA", {
              day: "numeric",
              month: "short",
              year: "2-digit",
            })}
          </time>
        </div>
        <div className="p-1">
          <h3 className="text-lg font-medium">{article.title}</h3>
        </div>
      </a>
    </div>
  );
}
