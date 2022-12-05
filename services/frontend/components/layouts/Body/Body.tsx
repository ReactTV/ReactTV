export type TBodyProps = {
  children: React.ReactNode;
};

const Body = ({ children }: TBodyProps) => {
  return <div>{children}</div>;
};

export default Body;
