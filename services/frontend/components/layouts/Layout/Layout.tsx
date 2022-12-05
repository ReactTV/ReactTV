import Body from "../Body";
import Footer from "../Footer";
import Header from "../Header";
import Head from "../Head";

export type TLayoutProps = {
  children: React.ReactNode;
};

const Layout = ({ children }: TLayoutProps) => {
  return (
    <div>
      <Head />
      <Header />
      <Body>{children}</Body>
      <Footer />
    </div>
  );
};

export default Layout;
