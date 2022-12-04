import React from "react";
import { ComponentStory, ComponentMeta } from "@storybook/react";

import VideoPlayer from "./VideoPlayer";

export default {
  title: "Atoms/VideoPlayer",
  component: VideoPlayer,
  args: {
    url: "https://www.youtube.com/watch?v=ysz5S6PUM-U",
  },
} as ComponentMeta<typeof VideoPlayer>;

const Template: ComponentStory<typeof VideoPlayer> = (args) => (
  <VideoPlayer {...args} />
);

export const Default = Template.bind({});
Default.args = {};
