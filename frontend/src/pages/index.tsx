import axios from "axios";
import type { NextPage } from "next";
import { useEffect, useState } from "react";
import { useQuery } from "react-query";
import Layout from "../components/Layout";
import { Event } from "../types/server";

const Home: NextPage = () => {
  return (
    <Layout content={<Post />} />
  );
};

  const event: Event = {
    user_id: "00001",
    read_post_url:
      "https://ciaran.co.za/smart-contracts-101-for-web-developers",
  };

function Post() {

  const { isLoading, error, data } = useQuery(['readPostContent'], async () => {
    await axios.post("/api/v1/posts/read", event).then((res) => {
      return res
    })
  })

  if (isLoading) return  <p>Loading...</p>

  if (error) return <p>An error has occured: {error.message}</p>

  console.log(data)

  return (
    <div>
      {data.title}
    </div>
  )
}

export default Home;
