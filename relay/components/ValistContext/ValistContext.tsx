import React from 'react';
import Valist, { Web3Providers } from 'valist';

export default React.createContext<Valist>(new Valist({ web3Provider: new Web3Providers.HttpProvider("https://cloudflare-eth.com") } ));
