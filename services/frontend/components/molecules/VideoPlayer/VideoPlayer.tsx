import VideoEmbed from "../../atoms/VideoEmbed";

import styles from "./VideoPlayer.module.css";

export type TVideoPlayerProps = {
  url: string | string[];
};

const VideoPlayer = ({ url }: TVideoPlayerProps) => {
  return (
    <div className={styles.videoPlayer}>
      <VideoEmbed url={url} />
    </div>
  );
};

export default VideoPlayer;
