import ReactPlayer from "react-player";

import styles from "./VideoEmbed.module.css";

export type TVideoEmbedProps = {
  url: string;
  onEnded?: () => void;
};

const VideoEmbed = ({ url, onEnded = () => {} }: TVideoEmbedProps) => {
  return (
    <div className={styles.videoEmbed}>
      <ReactPlayer url={url} playing controls pip onEnded={onEnded} />
    </div>
  );
};

export default VideoEmbed;
