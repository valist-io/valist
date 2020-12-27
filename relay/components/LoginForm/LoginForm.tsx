import React, { useEffect, useState } from 'react';

import { Magic } from 'magic-sdk';

function LoginForm({ magic, setLoggedIn }: { magic: Magic, setLoggedIn: any }) {

    const [email, setEmail] = useState("");
    const [checking, setChecking] = useState(true);

    useEffect(() => {
        (async function () {
            if (magic) {
                setLoggedIn(await magic.user.isLoggedIn(), setChecking(false));
            }
        })();
    }, [magic]);

    const handleLogin = async () => {
        await magic.auth.loginWithMagicLink({ email });
        setLoggedIn(await magic.user.isLoggedIn());
    }

    return (
        <div className="min-h-screen bg-gray-50 flex flex-col justify-center py-12 sm:px-6 lg:px-8">
            <div className="sm:mx-auto sm:w-full sm:max-w-md">
                <img className="mx-auto h-12 w-auto" src="/images/ValistLogo128.png" alt="Valist Logo" />
                <h2 className="mt-6 text-center text-3xl leading-9 font-extrabold text-gray-900">
                    {checking ? `Checking for Magic session...` : `Sign in, or sign up`}
                </h2>
            </div>
            { !checking && <div className="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
                <div className="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
                    <div>
                        <label htmlFor="email" className="block text-sm font-medium leading-5 text-gray-700">
                            Email address
                            </label>
                        <div className="mt-1 rounded-md shadow-sm">
                            <input id="email" type="email" onChange={(e) => setEmail(e.target.value)} onKeyDown={(e) => { if (e.key === 'Enter') handleLogin() }} required className="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md placeholder-gray-400 focus:outline-none focus:shadow-outline-blue focus:border-blue-300 transition duration-150 ease-in-out sm:text-sm sm:leading-5" />
                        </div>
                    </div>
                    <div className="mt-6">
                        <span className="block w-full rounded-md shadow-sm">
                            <button type="submit" onClick={handleLogin} className="w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-500 focus:outline-none focus:border-indigo-700 focus:shadow-outline-indigo active:bg-indigo-700 transition duration-150 ease-in-out">
                                Send magic link
                                </button>
                        </span>
                    </div>
                    {/* <div className="mt-6">
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
                                    <button type="button" className="w-full inline-flex justify-center py-2 px-4 border border-gray-300 rounded-md bg-white text-sm leading-5 font-medium text-gray-500 hover:text-gray-400 focus:outline-none focus:border-blue-300 focus:shadow-outline-blue transition duration-150 ease-in-out" aria-label="Sign in with GitHub">
                                        <svg className="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
                                            <path fillRule="evenodd" d="M10 0C4.477 0 0 4.484 0 10.017c0 4.425 2.865 8.18 6.839 9.504.5.092.682-.217.682-.483 0-.237-.008-.868-.013-1.703-2.782.605-3.369-1.343-3.369-1.343-.454-1.158-1.11-1.466-1.11-1.466-.908-.62.069-.608.069-.608 1.003.07 1.531 1.032 1.531 1.032.892 1.53 2.341 1.088 2.91.832.092-.647.35-1.088.636-1.338-2.22-.253-4.555-1.113-4.555-4.951 0-1.093.39-1.988 1.029-2.688-.103-.253-.446-1.272.098-2.65 0 0 .84-.27 2.75 1.026A9.564 9.564 0 0110 4.844c.85.004 1.705.115 2.504.337 1.909-1.296 2.747-1.027 2.747-1.027.546 1.379.203 2.398.1 2.651.64.7 1.028 1.595 1.028 2.688 0 3.848-2.339 4.695-4.566 4.942.359.31.678.921.678 1.856 0 1.338-.012 2.419-.012 2.747 0 .268.18.58.688.482A10.019 10.019 0 0020 10.017C20 4.484 15.522 0 10 0z" clipRule="evenodd" />
                                        </svg>
                                    </button>
                                </span>
                            </div>
                        </div>
                    </div> */}
                </div>
            </div> }
        </div>
    )
}

export default LoginForm;
