
const RecentPostsFeed = ({posts, focus, setFocus}) => {

  const handleClick = (event, url) => {
    setFocus(null)
    setFocus({user_id: '00001', read_post_url: url})
  }
  
  return posts.map(post => {
    const f = post.url === focus?.read_post_url
    return (
      <div
        key={post.url}
        onClick={() => handleClick(null, post.url)}
        className={`p-2 ${ f ? "bg-slate-300": "bg-slate-100"}`}
      >
        <h3 className="text-md">{post.title}</h3>
        <p className="text-sm italic font-bold text-slate-300">{post.domain}</p>
      </div>
    );
  });
};

export default RecentPostsFeed
