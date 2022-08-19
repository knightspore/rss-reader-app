import type { Article } from "../types/backend-vo";

const RecentArticlesFeed = ({posts, focus, setFocus}:{posts: Article[] | null, focus: Article | null, setFocus: Function}) => {

  const handleClick = (url: string) => {
    console.log("Should focus article")
  }
  
  return posts ? posts.map(post => {
    const f = post.url === focus?.url
    return (
      <div
        key={post.url}
        onClick={() => handleClick(post.url)}
        className={`p-2 ${ f ? "bg-slate-300": "bg-slate-100"}`}
      >
        <h3 className="text-md">{post.title}</h3>
        <p className="text-sm italic font-bold text-slate-300">{post.parent}</p>
      </div>
    );
  }) : <>
  Loading...
  </>
};

export default RecentArticlesFeed
