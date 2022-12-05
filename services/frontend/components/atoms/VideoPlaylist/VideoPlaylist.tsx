import VideoPlaylistItem from "../VideoPlaylistItem/VideoPlaylistItem";

import styles from "./VideoPlaylist.module.css";

export type TVideoPlaylistProps = {
  currentUrl?: string;
  playlistUrls: string[];
  onPlaylistClick: (url: string) => void;
};

const VideoPlaylist = ({
  currentUrl,
  playlistUrls,
  onPlaylistClick,
}: TVideoPlaylistProps) => {
  return (
    <div className={styles.videoPlaylist}>
      {playlistUrls.map((url, index) => (
        <VideoPlaylistItem
          isActive={url === currentUrl}
          key={url + index}
          url={url}
          onClick={onPlaylistClick}
        />
      ))}
    </div>
  );
};

export default VideoPlaylist;
