import { AppProps } from 'next/app';
import getConfig from "next/config";
import React, { useEffect, useState } from 'react';

import Valist from 'valist';
import ValistContext from '../components/Valist/ValistContext';
import LoginContext from '../components/Login/LoginContext';
import getProviders from '../utils/providers';

import LoadingDialog from '../components/LoadingDialog/LoadingDialog';
import LoginForm from '../components/Login/LoginForm';
import { Magic } from 'magic-sdk';

import '../styles/main.css';

type ProviderType = "magic" | "metaMask" | "readOnly";

const { publicRuntimeConfig } = getConfig();

function App({ Component, pageProps }: AppProps) {
  const [valist, setValist] = useState<Valist>();
  const [email, setEmail] = useState("");
  const [showLogin, setShowLogin] = useState(false);
  const [loggedIn, setLoggedIn] = useState(false);
  const [magic, setMagic] = useState<Magic | undefined>();
  const [account, setAccount] =useState();

  const loginObject = {
    setShowLogin: setShowLogin,
    loggedIn: loggedIn,
    logOut: async () => {
      window.localStorage.clear();
      setShowLogin(false);
      setLoggedIn(false);
      if (magic) {
        magic.user.logout();
        setMagic(undefined);
      }
      await handleLogin("readOnly");
    }
  }

  const handleLogin = async (providerType: ProviderType) => {
    let providers = getProviders(setMagic, setLoggedIn, email);
    let provider;

    try {
      provider = await providers[providerType]();
      window.localStorage.setItem("loginType", providerType);
    } catch (e) {
      console.log("Could not set provider, falling back to readOnly", e);
      provider = await providers["readOnly"]();
      window.localStorage.setItem("loginType", "readOnly");
    }

    try {
      const valist = new Valist({
        web3Provider: provider,
        metaTx: provider == publicRuntimeConfig.WEB3_PROVIDER ? false : publicRuntimeConfig.METATX_ENABLED
      });

      await valist.connect();
      setValist(valist);
      console.log("Current Account: ", valist.defaultAccount);
      (window as any).valist = valist;
    } catch (e) {
      console.error("Could not initialize Valist object", e);

      try {
        await handleLogin("readOnly");
      } catch (e) {
        console.error("Critical error, could not login with desired method or readOnly", e);
      }
    }
  }

  useEffect(() => {
    (async function () {
      // check login type on first load and set provider to previous state
      const loginType = window.localStorage.getItem("loginType");

      if (loginType === "readOnly" || loginType === "magic" || loginType === "metaMask") {
        await handleLogin(loginType);
      } else {
        await handleLogin("readOnly");
        window.localStorage.setItem("loginType", "readOnly");
      }
    })();

    (window as any).ethereum.on('accountsChanged', function (accounts: any) {
      setAccount(accounts[0]);
    });
  }, [account]);

  return (
    <LoginContext.Provider value={loginObject}>
      { valist ?
        <ValistContext.Provider value={valist}>
          <Component loggedIn={loggedIn} setShowLogin={setShowLogin} {...pageProps} />
          { showLogin && <LoginForm setShowLogin={setShowLogin} handleLogin={handleLogin} setEmail={setEmail} /> }
        </ValistContext.Provider>
        :
        <LoadingDialog>Loading...</LoadingDialog>}
    </LoginContext.Provider>
  )
}

export default App;