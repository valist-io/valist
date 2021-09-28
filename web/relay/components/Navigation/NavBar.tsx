import { useState, useContext } from 'react';
import Link from 'next/link';
import LoginContext from '../Login/LoginContext';
import AddressIdenticon from '../Identicons/AddressIdenticon';
import ValistContext from '../Valist/ValistContext';

export const Nav = () => {
  const login = useContext(LoginContext);
  const valist = useContext(ValistContext);
  const [menuVisible, setMenuVisible] = useState(false);

  return (
    <header className="pb-24 bg-gradient-to-r from-light-blue-800 to-violet-600">
      <div className="max-w-3xl mx-auto px-4 sm:px-6 lg:max-w-7xl lg:px-8">
        <div className="relative flex flex-wrap items-center justify-center lg:justify-between">

          <div className="w-full py-5">
            <div className="lg:grid lg:grid-cols-3 lg:gap-8 lg:items-center">
              <div className="hidden lg:block lg:col-span-2">
                <nav className="flex space-x-4">
                  <Link href="/">
                    <a className="text-white text-sm font-medium rounded-md bg-white
                    bg-opacity-0 px-3 py-2 hover:bg-opacity-10" aria-current="page">
                      Home
                    </a>
                  </Link>
                  <a href="https://docs.valist.io/" className="text-violet-100 text-sm
                  font-medium rounded-md bg-white bg-opacity-0 px-3 py-2 hover:bg-opacity-10">
                    Documentation
                  </a>

                  <a href="https://valist.io/discord" className="text-violet-100 text-sm
                  font-medium rounded-md bg-white bg-opacity-0 px-3 py-2 hover:bg-opacity-10">
                    Discord
                  </a>

                  <a href="mailto:support@valist.io" className="text-violet-100 text-sm
                  font-medium rounded-md bg-white bg-opacity-0 px-3 py-2 hover:bg-opacity-10">
                    Support
                  </a>
                </nav>
              </div>
              <div className={`hidden lg:block ${!login.loggedIn ? 'justify-self-end' : ''}`}>
                <div className="flex-grow">
                  {login.loggedIn
                    ? (<div onClick={() => setMenuVisible(!menuVisible)}
                        className="bg-white rounded-md flex items-center text-sm p-3 cursor-pointer"
                        id="user-menu" aria-expanded="false" aria-haspopup="true">
                          <span className="sr-only">Open user menu</span>
                            <AddressIdenticon address={valist.defaultAccount} height={25}/>
                            <span className="p-2 text-xs font-mono truncate">{valist.defaultAccount}</span>
                      </div>)
                    : (<button onClick={() => login.setShowLogin(true)} type="button"
                    className="inline-flex items-center px-6 py-1.5 border border-transparent
                    text-xs font-medium rounded-md shadow-sm text-indigo-700 bg-white hover:bg-gray-50
                    focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 cursor-pointer p-3 m-3">
                          Change RPC Provider (Read-Only)
                      </button>)}
                  {menuVisible && <div className="origin-top-right z-40 absolute right-0 mt-2 w-48 rounded-md
                                          shadow-lg py-1 bg-white ring-1 ring-black ring-opacity-5 focus:outline-none"
                                        role="menu" aria-orientation="vertical"
                                        aria-labelledby="user-menu">
                    <a onClick={() => { login.logOut(); setMenuVisible(!menuVisible); }}
                    className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 cursor-pointer"
                    role="menuitem">Switch RPC Provider</a>
                  </div>}
                </div>
              </div>
            </div>
          </div>

          <div className="absolute right-0 flex-shrink-0 lg:hidden">
            <button onClick={() => setMenuVisible(!menuVisible)} type="button"
            className="bg-transparent p-2 rounded-md inline-flex items-center
            justify-center text-violet-200 hover:text-white hover:bg-white
            hover:bg-opacity-10 focus:outline-none focus:ring-2 focus:ring-white"
            aria-expanded="false">
              <span className="sr-only">Open main menu</span>

              <svg className="block h-6 w-6" xmlns="http://www.w3.org/2000/svg"
              fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M4 6h16M4 12h16M4 18h16" />
              </svg>

              <svg className="hidden h-6 w-6" xmlns="http://www.w3.org/2000/svg"
              fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
        </div>
      </div>

      { menuVisible && <div className="lg:hidden">
        <div className="z-20 fixed inset-0 bg-black bg-opacity-25" aria-hidden="true"></div>
        <div className="z-30 absolute top-0 inset-x-0 max-w-3xl mx-auto w-full p-2 transition transform origin-top">
          <div className="rounded-lg shadow-lg ring-1 ring-black ring-opacity-5 bg-white divide-y divide-gray-200">
            <div className="pt-3 pb-2">
              <div className="flex items-center justify-between px-4">
                <div>
                  <img className="h-8 w-auto"
                    src="/images/ValistLogo128.png" alt="Logo" />
                </div>
                <div className="-mr-2">
                  <button onClick={() => setMenuVisible(!menuVisible)}
                    type="button" className="bg-white rounded-md p-2 inline-flex
                    items-center justify-center text-gray-400 hover:text-gray-500
                    hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-inset
                    focus:ring-violet-500">
                    <span className="sr-only">Close menu</span>
                    <svg className="h-6 w-6" xmlns="http://www.w3.org/2000/svg"
                      fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M6 18L18 6M6 6l12 12" />
                    </svg>
                  </button>
                </div>
              </div>
              <div className="mt-3 px-2 space-y-1">
                <Link href="/dashboard">
                  <a className="block rounded-md px-3 py-2 text-base text-gray-900 font-medium
                  hover:bg-gray-100 hover:text-gray-800">Home</a>
                </Link>
                <a href="https://docs.valist.io/" className="block rounded-md px-3 py-2 text-base
                  text-gray-900 font-medium hover:bg-gray-100 hover:text-gray-800">Documentation</a>

                <a href="https://valist.io/discord" className="block rounded-md px-3 py-2 text-base
                  text-gray-900 font-medium hover:bg-gray-100 hover:text-gray-800">Discord</a>

                <a href="mailto:support@valist.io" className="block rounded-md px-3 py-2 text-base
                  text-gray-900 font-medium hover:bg-gray-100 hover:text-gray-800">Support</a>
              </div>
            </div>
            <div className="pt-4 pb-2">
              <div className="flex items-center px-5">
                <div className="flex-shrink-0">
                  <AddressIdenticon address={valist.defaultAccount} height={8}/>
                </div>
                <div className="ml-3 min-w-0 flex-1">
                  <div className="text-base font-xs text-gray-800 font-mono">{valist.defaultAccount}</div>
                </div>
              </div>
              <div className="mt-3 px-2 space-y-1">
                <a href="#" className="block rounded-md px-3 py-2 text-base text-gray-900
                font-medium hover:bg-gray-100 hover:text-gray-800">Sign out</a>
              </div>
            </div>
          </div>
        </div>
      </div>}
    </header>
  );
};

export default Nav;
