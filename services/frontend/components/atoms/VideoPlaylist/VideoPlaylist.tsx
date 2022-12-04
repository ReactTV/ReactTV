import styles from "./VideoPlaylist.module.css";

export type TVideoPlaylistProps = {
  playlistUrls: string[];
  onPlaylistClick: (url: string) => void;
};

const VideoPlaylist = ({
  playlistUrls,
  onPlaylistClick,
}: TVideoPlaylistProps) => {
  return (
    <div className={styles.videoPlaylist}>
      {playlistUrls.map((url, index) => (
        <div key={url + index} onClick={() => onPlaylistClick(url)}>
          {url}
        </div>
      ))}
    </div>
  );
};

export default VideoPlaylist;
