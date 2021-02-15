import React from 'react';
import Valist, { Web3Providers } from 'valist';
import getConfig from "next/config";

const { publicRuntimeConfig } = getConfig();

export default React.createContext<Valist>(new Valist({ web3Provider: new Web3Providers.HttpProvider(publicRuntimeConfig.WEB3_PROVIDER), metaTx: false } ));
