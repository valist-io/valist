import type { AppProps /*, AppContext */ } from 'next/app'
import React, { FunctionComponent, useState, useEffect  } from 'react';
import getValistContract from 'valist/dist/getValistContract';
import getValistOrganizationContract from 'valist/dist/getValistOrganizationContract'
import { Web3 } from 'valist';
import Web3Modal from "web3modal";

import '../styles/main.css';

function App({ Component, pageProps }: AppProps) {

  const providerOptions = {
    /* See Provider Options Section */
  };

  const [contract, setContract] = useState(null)
  const [provider, setProvider] = useState(null)
  const [web3, setWeb3] = useState(null)
  const [contractOutput, setContractOutput] = useState(null)
  const [organizationContractOutput, setOrganizationContractOutput] = useState(null)

    useEffect(() => {
      async function connectWeb3() {
          try {
              const web3Modal = new Web3Modal({
                  cacheProvider: true, // optional
                  providerOptions // required
              });
              
              const provider = await web3Modal.connect();
              const web3 = Web3(provider);
              const contract = await getValistContract(web3);
              const organization = await contract.methods.orgs("Akashic tech").call()
              const organizationContract = await getValistOrganizationContract(web3, organization)
              console.log( await organizationContract.methods.orgMeta.call())
              console.log(contract)
              
              setContract(contract)
              setProvider(provider)
              setWeb3(web3)

              setContractOutput(organization)
              // setOrganizationContractOutput(orgOutput)
              
          } catch (error) {
              alert(
                  `Failed to load web3, accounts, or contract. Check console for details.`
              )
              console.log(error)
          }
      }
      connectWeb3()
    }, [])

  return <Component {...pageProps} contract={contract} />
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
