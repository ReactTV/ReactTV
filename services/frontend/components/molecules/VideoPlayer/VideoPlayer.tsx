import { useState } from "react";

import VideoEmbed from "../../atoms/VideoEmbed";
import VideoPlaylist from "../../atoms/VideoPlaylist";

import styles from "./VideoPlayer.module.css";

export type TVideoPlayerProps = {
  url: string | string[];
};

const VideoPlayer = ({ url }: TVideoPlayerProps) => {
  const playlistUrls: string[] = Array.isArray(url) ? url : [url];

  const [currentUrl, setCurrentUrl] = useState(playlistUrls[0]);

  const onChangeVideo = (nextUrl: string) => {
    setCurrentUrl(nextUrl);
  };

  return (
    <div className={styles.videoPlayer}>
      <VideoEmbed url={currentUrl} />
      <VideoPlaylist url={playlistUrls} onPlaylistClick={onChangeVideo} />
    </div>
  );
};

export default VideoPlayer;
