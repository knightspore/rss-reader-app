import type {NextPage} from 'next';
import {useEffect, useState} from 'react';
import Layout from '../components/Layout';
import Article from '../components/Article';
import RecentArticlesFeed from '../components/RecentArticlesFeed';
import { useQuery } from 'react-query';
import axios from 'axios';
import { UserEvent } from '../types/backend-module';

const Home: NextPage = () => {

  const [content, setContent] = useState(null)
  const [focusArticle, setFocusArticle] = useState(null);

  const result = fetch("http://127.0.0.1:1337/api/readinglist/get", {
    body: '{ "id": "parabyl" }',
    mode: 'no-cors',
    headers: {
      "Content-Type": "application/json",
      "Accept": "application/json",
    },
    method: "POST"
  }).then(res => res.json()).then(data => {
    setContent(data)
    console.log(content)
  })


  return (
    <Layout
      recentPosts={<RecentArticlesFeed posts={content} focus={focusArticle} setFocus={setFocusArticle} />}
    >
      <Article event={focusArticle} />
    </Layout>
  );
};

export default Home;

