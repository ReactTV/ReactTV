// import getHTML from "html-get";
// import metascraper from "metascraper";

// const getMetadata = async () => {
//   /**
//    * `browserless` will be passed to `html-get`
//    * as driver for getting the rendered HTML.
//    */
//   const browserless = require("browserless")();

//   const getContent = async (url: string) => {
//     // create a browser context inside the main Chromium process
//     const browserContext = browserless.createContext();
//     const promise = getHTML(url, { getBrowserless: () => browserContext });
//     // close browser resources before return the result
//     promise
//       .then(() => browserContext)
//       .then((browser: any) => browser.destroyContext());
//     return promise;
//   };
// };

// export default getMetadata;
