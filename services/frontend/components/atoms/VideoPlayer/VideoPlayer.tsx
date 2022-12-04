import ReactPlayer from "react-player";

import styles from "./VideoPlayer.module.css";

export type TVideoPlayerProps = {
  url: string;
};

const VideoPlayer = ({ url }: TVideoPlayerProps) => {
  return (
    <div className={styles.VideoPlayer}>
      <ReactPlayer url={url} />
    </div>
  );
};

export default VideoPlayer;
