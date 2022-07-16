import type { NextPage } from "next";
import { Event } from "../types/server";

const Home: NextPage = () => {
  const event: Event = {
    user_id: "00001",
    feeds: [
      {
        url: "https://ciaran.co.za/rss.xml",
        title: "Ciaran's RSS Feed",
      },
    ],
  };

  return (
    <>
      Hello
      {event.feeds.map((feed) => {
        return (
          <div key={feed.url}>
            <a href={"#"+feed.url}>{feed.title}</a>
          </div>
        );
      })}
    </>
  );
};

export default Home;
