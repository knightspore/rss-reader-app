import axios from "axios";
import { useState } from "react";
import { useQuery } from "react-query";

export default function Article ({ event }) {

  const [content, setContent] = useState(null);

  const { isLoading, error } = useQuery(["readPostContent"], () => {
    axios.post("/api/v1/posts/read", event).then((res) => {
      setContent(res.data);
    });
  });

  if (isLoading) return <p>Loading...</p>;

  if (error) return <p>An error has occured: {error.message}</p>;

  return (
    content && (
      <section id="now-reading" className="p-4">
        <div id="site-imprint" className="flex items-center gap-4">
          <img
            className="w-4 h-4"
            src={content.favicon}
            alt={content.title + " favicon"}
          />
          <p>
            {content.byline} | {content.site}
          </p>
        </div>
        <div className="p-4 mx-auto prose prose-slate">
          <h2>{content.title}</h2>
          <article>{content.text}</article>
        </div>
      </section>
    )
  );
};