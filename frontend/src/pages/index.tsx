import type {NextPage} from 'next';
import {useState} from 'react';
import Layout from '../components/Layout';
import Article from '../components/Article';
import RecentArticlesFeed from '../components/RecentArticlesFeed';

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
      recentPosts={<RecentArticlesFeed posts={posts} focus={focusArticle} setFocus={setFocusArticle} />}
    >
      <Article event={focusArticle} />
    </Layout>
  );
};

export default Home;

