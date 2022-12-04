import React from "react";
import { ComponentStory, ComponentMeta } from "@storybook/react";

import VideoPlaylist from "./VideoPlaylist";

export default {
  title: "Atoms/VideoPlaylist",
  component: VideoPlaylist,
  args: {
    url: [
      "https://www.youtube.com/watch?v=4c52GsNrJ_E",
      "https://www.dailymotion.com/video/x8ebz36",
      "https://www.twitch.tv/willneff",
    ],
  },
} as ComponentMeta<typeof VideoPlaylist>;

const Template: ComponentStory<typeof VideoPlaylist> = (args) => (
  <VideoPlaylist {...args} />
);

export const Default = Template.bind({});
Default.args = {};
