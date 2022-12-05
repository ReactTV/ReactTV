// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { NextApiRequest, NextApiResponse } from "next";
import Channel from "../../factories/Channel";

const TEMPORARY_CHANNELS = Channel.buildList(20);

export type TChannel = {
  id: number;
  name: string;
  playlist: string[];
};

export type TChannelsData = {
  channels: TChannel[];
};

export default function handler(
  req: NextApiRequest,
  res: NextApiResponse<TChannelsData>
) {
  res.status(200).json({ channels: TEMPORARY_CHANNELS });
}
