import styles from "./VideoPlaylist.module.css";

export type TVideoPlaylistProps = {
  url: string | string[];
  onPlaylistClick: (url: string) => void;
};

const VideoPlaylist = ({ url, onPlaylistClick }: TVideoPlaylistProps) => {
  const playlistUrls: string[] = Array.isArray(url) ? url : [url];

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
