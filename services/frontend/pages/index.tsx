import Head from "next/head";
import Link from "next/link";

import { TChannelsData } from "./api/channels";

export type THomePage = {
  data: TChannelsData;
};

export default function HomePage({ data }: THomePage) {
  if (!data) return <div>Loading...</div>;

  const { channels } = data;

  return (
    <div>
      <Head>
        <title>Create Next App</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main>
        <h1>Welcome to ReactTV</h1>
        <div>
          <h1>Channels:</h1>
          <div>
            {channels.map((channel) => (
              <div>
                <Link href={`/ch/${channel.id}`}>{channel.name}</Link>
              </div>
            ))}
          </div>
        </div>
      </main>
    </div>
  );
}

export async function getServerSideProps() {
  const res = await fetch("http://localhost:3000/api/channels");
  const data = await res.json();

  // By returning { props: { posts } }, the Blog component
  // will receive `posts` as a prop at build time
  return {
    props: {
      data,
    },
  };
}
