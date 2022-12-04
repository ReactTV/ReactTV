import React from "react";
import { ComponentStory, ComponentMeta } from "@storybook/react";

import VideoEmbed from "./VideoEmbed";

export default {
  title: "Atoms/VideoEmbed",
  component: VideoEmbed,
  args: {
    url: "https://www.twitch.tv/willneff",
  },
} as ComponentMeta<typeof VideoEmbed>;

const Template: ComponentStory<typeof VideoEmbed> = (args) => (
  <VideoEmbed {...args} />
);

export const Default = Template.bind({});
Default.args = {};
