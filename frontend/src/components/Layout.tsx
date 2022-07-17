import { json } from "stream/consumers";
import type { Feed, PostContent } from "../types/server";

const Layout = ({
  content,
}: {
  content: any;
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
      {/* Main UI */}
      <div className="grid flex-grow grid-cols-6">
        <section id="recent-posts-feed" className="col-span-2 bg-slate-200"></section>
        <section id="read-post-area" className="col-span-4 bg-slate-50">
					{ content }
				</section>
      </div>
      {/* Foot */}
    </div>
  );
};

export default Layout;
