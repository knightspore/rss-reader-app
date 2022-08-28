import "../styles/globals.css";
import type { AppProps } from "next/app";
import { useState } from "react";
import { QueryClient, QueryClientProvider } from "react-query";
import Head from "next/head";
import UserContextProvider from "../components/context/UserProvider";
import SubscriptionContextProvider from "../components/context/SubscriptionProvider";

function MyApp({ Component, pageProps: { session, ...pageProps } }: AppProps) {
  const [queryClient] = useState(() => new QueryClient());

  return (
    <>
      <Head>
        <title>RSS Reader</title>
        <link
          rel="shortcut icon"
          href="https://t3.gstatic.com/faviconV2?client=SOCIAL&type=FAVICON&fallback_opts=TYPE,SIZE,URL&url=http://ciaran.co.za&size=48"
          type="image/x-icon"
        />
      </Head>
      <QueryClientProvider client={queryClient}>
        <UserContextProvider>
          <SubscriptionContextProvider>
            <Component {...pageProps} />
          </SubscriptionContextProvider>
        </UserContextProvider>
      </QueryClientProvider>
    </>
  );
}

export default MyApp;
