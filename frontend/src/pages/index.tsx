import type {NextPage} from 'next';
import {getSession} from 'next-auth/react';
import {useEffect, useState} from 'react';
import Layout from '../components/Layout';
import Post from '../components/Post';
import RecentPostsFeed from '../components/RecentPostFeed';
import {Event} from '../types/server';

const Home: NextPage = () => {

  const [focusArticle, setFocusArticle] = useState(null);

  const posts = [
    {
      title: 'Smart Contracts 101 for Web Developers',
      url: 'https://ciaran.co.za/smart-contracts-101-for-web-developers',
      domain: 'ciaran.co.za',
    },
    {
      title: 'The Case for Bad Coffee',
      url: 'https://www.seriouseats.com/the-case-for-bad-coffee',
      domain: 'seriouseats.com',
    },
  ];

  return (
    <Layout
      recentPosts={<RecentPostsFeed posts={posts} focus={focusArticle} setFocus={setFocusArticle} />}
    >
      <Post event={focusArticle} />
    </Layout>
  );
};

export default Home;

