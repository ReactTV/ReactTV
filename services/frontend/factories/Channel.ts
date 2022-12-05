import { Factory } from "rosie";
import { randMovie, rand } from "@ngneat/falso";

const VIDEO_LIST = [
  "https://www.youtube.com/watch?v=Jk9gUzo2AV8",
  "https://www.youtube.com/watch?v=rZmnZI2KqVE",
  "https://www.youtube.com/watch?v=pR3SCgGohcw",
  "https://www.youtube.com/watch?v=jzQYMwPXgVU",
  "https://www.youtube.com/watch?v=CcijaBdQIFY",
  "https://www.youtube.com/watch?v=y7m3if-WCLQ",
  "https://www.youtube.com/watch?v=p1tQ-U4VNx0",
  "https://www.youtube.com/watch?v=E53K6wGzVRg",
  "https://www.youtube.com/watch?v=D6rlUL3LgCQ",
  "https://www.youtube.com/watch?v=5dW3ECb-31g",
  "https://www.youtube.com/watch?v=AWm9TGsqy7o",
  "https://www.youtube.com/watch?v=Nr0fLkMP260",
  "https://www.youtube.com/watch?v=Qf77tu2mLBE",
  "https://www.youtube.com/watch?v=BcnEP5MZRxw",
  "https://www.youtube.com/watch?v=Jk9gUzo2AV8",
  "https://www.youtube.com/watch?v=rZmnZI2KqVE",
  "https://www.youtube.com/watch?v=pR3SCgGohcw",
  "https://www.youtube.com/watch?v=jzQYMwPXgVU",
  "https://www.youtube.com/watch?v=CcijaBdQIFY",
  "https://www.youtube.com/watch?v=y7m3if-WCLQ",
  "https://www.youtube.com/watch?v=p1tQ-U4VNx0",
  "https://www.youtube.com/watch?v=E53K6wGzVRg",
  "https://www.youtube.com/watch?v=D6rlUL3LgCQ",
  "https://www.youtube.com/watch?v=5dW3ECb-31g",
  "https://www.youtube.com/watch?v=AWm9TGsqy7o",
  "https://www.youtube.com/watch?v=Nr0fLkMP260",
  "https://www.youtube.com/watch?v=Qf77tu2mLBE",
  "https://www.youtube.com/watch?v=BcnEP5MZRxw",
];

export default new Factory().sequence("id").attrs({
  name: () => randMovie(),
  playlist: () => rand(VIDEO_LIST, { length: 20 }),
});
