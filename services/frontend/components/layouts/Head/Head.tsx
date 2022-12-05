import Head from "next/head";

export type THeadComponentProps = {
  pageTitle: string;
};

const HeadComponent = ({ pageTitle = "ReactTV" }) => {
  return (
    <Head>
      <title>{pageTitle}</title>
      <meta name="description" content="Your source for content" />
      <link rel="icon" href="/favicon.ico" />
    </Head>
  );
};

export default HeadComponent;
