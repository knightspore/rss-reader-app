import axios, { AxiosResponse } from "axios";
import { NextApiRequest, NextApiResponse } from "next";
import { api } from "../../../../lib/constants";
import type { Event } from "../../../../types/server";

const handler = async (req: NextApiRequest, res: NextApiResponse) => {
  const event: Event = req.body;

  try {
    const server: AxiosResponse = await axios.post(api.posts.read, event);
    server.status == 200 ? res.json(server.data) : res.send(500);
  } catch (e) {
    console.warn("Error Getting Read Post Text");
    console.error(e.message);
  }
};

export default handler;
