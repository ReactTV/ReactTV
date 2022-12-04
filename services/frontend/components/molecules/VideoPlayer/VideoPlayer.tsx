import { useState } from "react";

import VideoEmbed from "../../atoms/VideoEmbed";
import VideoPlaylist from "../../atoms/VideoPlaylist";

import styles from "./VideoPlayer.module.css";

export type TVideoPlayerProps = {
  playlistUrls: string[];
};

const VideoPlayer = ({ playlistUrls }: TVideoPlayerProps) => {
  const [currentUrl, setCurrentUrl] = useState(playlistUrls[0]);

  const onChangeVideo = (nextUrl: string) => {
    setCurrentUrl(nextUrl);
  };

  return (
    <div className={styles.videoPlayer}>
      <VideoEmbed url={currentUrl} />
      <VideoPlaylist
        playlistUrls={playlistUrls}
        onPlaylistClick={onChangeVideo}
      />
    </div>
  );
};

export default VideoPlayer;
