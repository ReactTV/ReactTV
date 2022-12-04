import React from "react";
import { ComponentStory, ComponentMeta } from "@storybook/react";

import VideoPlayer from "./VideoPlayer";

export default {
  title: "Atoms/VideoPlayer",
  component: VideoPlayer,
  args: {
    url: [
      "https://www.youtube.com/watch?v=4c52GsNrJ_E",
      "https://www.dailymotion.com/video/x8ebz36",
      "https://www.twitch.tv/willneff",
    ],
  },
} as ComponentMeta<typeof VideoPlayer>;

const Template: ComponentStory<typeof VideoPlayer> = (args) => (
  <VideoPlayer {...args} />
);

export const Default = Template.bind({});
Default.args = {};