import { TChannelsData } from "../../api/channels";
import VideoPlayer from "../../../components/molecules/VideoPlayer";

export type TChannelPage = {
  id: number;
  data: TChannelsData;
};

export default function ChannelPage({ id, data }: TChannelPage) {
  const channel =
    data.channels.find((channel) => channel.id === id) || data.channels[0];

  return (
    <div>
      <main>
        <h1>{channel.name}</h1>
        <VideoPlayer playlistUrls={channel.playlist} />
      </main>
    </div>
  );
}

export async function getServerSideProps({ params }: any) {
  const res = await fetch("http://localhost:3000/api/channels");
  const data = await res.json();

  // By returning { props: { posts } }, the Blog component
  // will receive `posts` as a prop at build time
  return {
    props: {
      data,
      id: parseInt(params.id),
    },
  };
}
