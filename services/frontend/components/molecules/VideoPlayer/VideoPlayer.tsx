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

  const nextVideo = () => {
    const currentIndex = playlistUrls.indexOf(currentUrl);
    const nextUrl = playlistUrls[currentIndex + 1] || playlistUrls[0];
    setCurrentUrl(nextUrl);
  };

  return (
    <div className={styles.videoPlayer}>
      <VideoEmbed url={currentUrl} onEnded={nextVideo} />
      <VideoPlaylist
        playlistUrls={playlistUrls}
        currentUrl={currentUrl}
        onPlaylistClick={onChangeVideo}
      />
    </div>
  );
};

export default VideoPlayer;
