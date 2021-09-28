import React, { useEffect, useRef } from 'react';

const LoginForm = ({ handleLogin, setShowLogin }: { handleLogin: any, setEmail: any, setShowLogin: any }) => {
  const element = useRef<HTMLDivElement>(null);

  const loginMetaMask = async () => {
    await handleLogin('metaMask');
    setShowLogin(false);
  };

  const loginWalletConnect = async () => {
    await handleLogin('walletConnect');
    setShowLogin(false);
  };

  const loginReadOnly = async () => {
    await handleLogin('readOnly');
    setShowLogin(false);
  };

  const handleClick = (e: any) => {
    if (element.current && e.target && element.current.contains(e.target)) return;

    setShowLogin(false);
  };

  useEffect(() => {
    document.addEventListener('mousedown', handleClick);

    return () => {
      document.removeEventListener('mousedown', handleClick);
    };
  }, []);

  return (
        <div className="fixed top-0 left-0 z-50 w-screen h-screen flex items-center
        justify-center" style={{ background: 'rgba(0, 0, 0, 0.3)' }}>
            <div ref={element} className="bg-gray-50 border py-2 px-5 rounded-lg flex items-center flex-col">
                <div className="text-gray-500 font-light mt-2 text-center">
                    <div className="bg-gray-50 flex flex-col justify-center py-28 sm:px-6 lg:px-8">
                        <div className="sm:mx-auto sm:w-full sm:max-w-md">
                            <button onClick={async () => loginReadOnly()} type="button"
                                className="bg-white inline-flex justify-center py-4 px-4 border
                                rounded-md text-sm leading-5 font-medium
                                text-gray-500 hover:text-gray-400 focus:outline-none
                                focus:border-blue-300 focus:shadow-outline-blue transition
                                duration-150 ease-in-out" aria-label="Sign in with GitHub">
                                <img width="85px" src="/images/ValistLogo128.png" alt="Valist Logo" />
                            </button>
                            <h2 className="mt-6 text-center text-xl leading-9 text-gray-900">
                                Valist (Ready-Only)
                            </h2>
                        </div>
                        <div className="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
                            <div className="py-8 px-4 sm:rounded-lg sm:px-10">
                                <div className="mt-6">
                                    <div className="mt-6 grid grid-cols-3 gap-3 content-center">
                                        <div>
                                            <span className="w-full inline-flex rounded-md shadow-sm">
                                                <button onClick={async () => loginMetaMask()} type="button"
                                                  className="bg-white w-full inline-flex justify-center py-2 px-4 border
                                                   rounded-md text-sm leading-5 font-medium
                                                  text-gray-500 hover:text-gray-400 focus:outline-none
                                                  focus:border-blue-300 focus:shadow-outline-blue transition
                                                  duration-150 ease-in-out" aria-label="Sign in with GitHub">
                                                    <img width="85px" src="/images/metamask.svg"/>
                                                </button>
                                            </span>
                                            <h2 className="mt-6 text-center text-xl leading-9 text-gray-900">
                                                MetaMask
                                            </h2>
                                        </div>
                                        <div>
                                        </div>
                                        <div>
                                            <span className="w-full inline-flex rounded-md shadow-sm">
                                                <button onClick={async () => loginWalletConnect()} type="button"
                                                  className="bg-white w-full inline-flex justify-center py-2 px-4 border
                                                rounded-md text-sm leading-5 font-medium
                                                  text-gray-500 hover:text-gray-400 focus:outline-none
                                                  focus:border-blue-300 focus:shadow-outline-blue transition
                                                  duration-150 ease-in-out" aria-label="Sign in with GitHub">
                                                    <img width="85px" src="/images/walletConnect.jpeg"/>
                                                </button>
                                            </span>
                                            <h2 className="mt-6 text-center text-xl leading-9 text-gray-900">
                                                W. Connect
                                            </h2>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
  );
};

export default LoginForm;
