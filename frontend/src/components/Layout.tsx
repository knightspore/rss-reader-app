const Layout = ({
  focusedContent,
  recentPosts,
  ...props
}: {
  focusedContent: React.ReactNode;
  recentPosts: any;
  props: { children: any };
}) => {
  return (
    <div className="flex flex-col w-screen h-screen text-slate-900">
      <header className="flex justify-between p-4 align-middle bg-slate-300">
        <div>
          <h1>RSS Reader</h1>
        </div>
        <div>
          <button className="px-2 py-1 text-xs font-bold uppercase transition-all rounded-sm hover:bg-slate-600 duration-250 bg-slate-900 text-slate-50">
            Add Subscription
          </button>
        </div>
      </header>
      <div className="grid flex-grow grid-cols-6">
        <section
          id="recent-posts-feed"
          className="flex flex-col col-span-2 gap-4 p-4 bg-slate-200"
        >
          {recentPosts}
        </section>
        <section id="read-post-area" className="col-span-4 bg-slate-50">
        </section>
      </div>
    </div>
  );
};

export default Layout;
