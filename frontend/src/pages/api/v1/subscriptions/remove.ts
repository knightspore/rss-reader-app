import axios, { AxiosResponse } from "axios";
import { NextApiRequest, NextApiResponse } from "next";
import { api } from "../../../../lib/constants";
import type { Event } from "../../../../types/server";

const handler = async (req: NextApiRequest, res: NextApiResponse) => {
  const SubscriptionEvent: Event = {
    user_id: "00001",
    feeds: [
      {
        url: "https://ciaran.co.za/rss.xml",
        title: "Ciaran's Blog",
      },
    ],
  };

  try {
    const server: AxiosResponse = await axios.post(
      api.subscriptions.remove,
      SubscriptionEvent
    );
    server.status == 200 ? res.send(200) : res.send(500);
  } catch (e) {
    console.warn("Error Removing Subscription");
    console.error(e);
  }
};

export default handler;
