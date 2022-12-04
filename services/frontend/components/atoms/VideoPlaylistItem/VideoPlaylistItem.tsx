import Image from "next/image";
import clsx from "clsx";
import { parser } from "html-metadata-parser";
import { useState, useEffect } from "react";

import styles from "./VideoPlaylistItem.module.css";

export type TVideoPlaylistItemProps = {
  isActive: boolean;
  url: string;
  onClick: (url: string) => void;
};

const VideoPlaylistItem = ({
  isActive,
  url,
  onClick,
}: TVideoPlaylistItemProps) => {
  const [metadata, setMetadata]: any = useState();

  useEffect(() => {
    if (metadata) return;

    parser(url).then((result) => {
      setMetadata(result.og);
    });
  });

  if (!metadata) {
    return <div>Loading...</div>;
  }

  return (
    <div
      className={clsx(styles.videoPlaylistItem, {
        [styles.isActive]: isActive,
      })}
      onClick={() => onClick(url)}
    >
      <Image
        className={styles.image}
        src={metadata.image}
        alt={metadata.title}
        width={50}
        height={50}
      />
      <div className={styles.textWrapper}>
        <div className={styles.title}>{metadata.title}</div>
        <div className={styles.sitename}>{metadata.site_name}</div>
      </div>
    </div>
  );
};

export default VideoPlaylistItem;
