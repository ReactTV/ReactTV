import ReactPlayer, { ReactPlayerProps } from "react-player";

import styles from "./VideoEmbed.module.css";

export type TVideoEmbedProps = ReactPlayerProps & {
  url: string;
};

const VideoEmbed = ({ url, ...otherProps }: TVideoEmbedProps) => {
  return (
    <div className={styles.videoEmbed}>
      <ReactPlayer url={url} playing controls pip {...otherProps} />
    </div>
  );
};

export default VideoEmbed;
