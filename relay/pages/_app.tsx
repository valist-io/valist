import type { AppProps /*, AppContext */ } from 'next/app'
import Web3Modal from 'web3modal'
import React, { useState, useEffect } from 'react';
import Valist from 'valist'
import '../styles/main.css';


function App({ Component, pageProps }: AppProps) {

  const [valistWeb3, setValistWeb3] = useState(null)
  const providerOptions = {}

    useEffect(() => {
      async function connectWeb3() {
          try {
            const web3Modal = new Web3Modal({
              cacheProvider: true, // optional
              providerOptions // required
            });
          
            const provider = await web3Modal.connect();
              const valist = new Valist(provider)
              await valist.connect();
              console.log(await valist.getOrganization("Akashic tech"))
              setValistWeb3(valist)
              
          } catch (error) {
              alert(
                  `Failed to load web3, accounts, or contract. Check console for details.`
              )
              console.log(error)
          }
      }
      connectWeb3()
    }, [])

  return <Component {...pageProps} valist={valistWeb3} />
}

// Only uncomment this method if you have blocking data requirements for
// every single page in your application. This disables the ability to
// perform automatic static optimization, causing every page in your app to
// be server-side rendered.
//
// MyApp.getInitialProps = async (appContext: AppContext) => {
//   // calls page's `getInitialProps` and fills `appProps.pageProps`
//   const appProps = await App.getInitialProps(appContext);

//   return { ...appProps }
// }

export default App
