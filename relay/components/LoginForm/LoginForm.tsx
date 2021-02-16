import React, { useEffect, useRef } from 'react';

const LoginForm = ({ handleLogin, setEmail, setShowLogin }: { handleLogin: any, setEmail: any, setShowLogin: any }) => {
    const element = useRef<HTMLDivElement>(null);

    const loginMetaMask = async () => {
        await handleLogin("metaMask");
        setShowLogin(false);
    }

    const handleClick = (e: any) => {
        if (element.current && e.target && element.current.contains(e.target)) return;

        setShowLogin(false);
    };

    useEffect(() => {
        document.addEventListener("mousedown", handleClick);
    
        return () => {
            document.removeEventListener("mousedown", handleClick);
        };
    }, []);

    return (
        <div className="fixed top-0 left-0 z-50 w-screen h-screen flex items-center justify-center" style={{background: "rgba(0, 0, 0, 0.3)"}}>
            <div ref={element} className="bg-white border py-2 px-5 rounded-lg flex items-center flex-col">
                <div className="text-gray-500 font-light mt-2 text-center">
                    <div className="bg-gray-50 flex flex-col justify-center py-28 sm:px-6 lg:px-8">
                        <div className="sm:mx-auto sm:w-full sm:max-w-md">
                            <img className="mx-auto h-12 w-auto" src="/images/ValistLogo128.png" alt="Valist Logo" />
                            <h2 className="mt-6 text-center text-3xl leading-9 font-extrabold text-gray-900">
                            </h2>
                        </div>
                        <div className="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
                            <div className="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
                                <div>
                                    <label htmlFor="email" className="block text-sm font-medium leading-5 text-gray-700">
                                        Email address
                                    </label>
                                    <div className="mt-1 rounded-md shadow-sm">
                                        <input id="email" type="email" onChange={(e) => setEmail(e.target.value)} onKeyDown={(e) => { if (e.key === 'Enter') handleLogin("magic") }} required className="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md placeholder-gray-400 focus:outline-none focus:shadow-outline-blue focus:border-blue-300 transition duration-150 ease-in-out sm:text-sm sm:leading-5" />
                                    </div>
                                </div>
                                <div className="mt-6">
                                    <span className="block w-full rounded-md shadow-sm">
                                        <button type="submit" onClick={ async () => { await handleLogin("magic"); setShowLogin(false); }} className="w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-500 focus:outline-none focus:border-indigo-700 focus:shadow-outline-indigo active:bg-indigo-700 transition duration-150 ease-in-out">
                                            Send magic link
                                        </button>
                                    </span>
                                </div>
                                <div className="mt-6">
                                    <div className="relative">
                                        <div className="absolute inset-0 flex items-center">
                                            <div className="w-full border-t border-gray-300" />
                                        </div>
                                        <div className="relative flex justify-center text-sm leading-5">
                                            <span className="px-2 bg-white text-gray-500">
                                                Or continue with
                                            </span>
                                        </div>
                                    </div>
                                    <div className="mt-6 grid grid-cols-3 gap-3">
                                        <div>
                                            <span className="w-full inline-flex rounded-md shadow-sm">
                                                <button onClick={async () => await loginMetaMask()} type="button" className="w-full inline-flex justify-center py-2 px-4 border border-gray-300 rounded-md bg-white text-sm leading-5 font-medium text-gray-500 hover:text-gray-400 focus:outline-none focus:border-blue-300 focus:shadow-outline-blue transition duration-150 ease-in-out" aria-label="Sign in with GitHub">
                                                    <img width="85px" src="/images/metamask.svg"/>
                                                </button>
                                            </span>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default LoginForm;
