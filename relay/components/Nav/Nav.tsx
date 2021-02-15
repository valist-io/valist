import React, { FunctionComponent, useContext} from 'react';
import Link from 'next/link';
import LoginContext from '../LoginContext/LoginContext';

export const Nav:FunctionComponent<any> = () => {
    const login = useContext(LoginContext);

    return (
        <nav className="flex-shrink-0 bg-indigo-700">
            <div className="max-w-7xl mx-auto px-2 sm:px-4 lg:px-8">
                <div className="relative flex items-center justify-between h-16">
                    <div className="flex items-center px-2 lg:px-0 xl:w-64">
                        <Link href="/">
                            <div className="flex-shrink-0">
                                <img className="h-8 w-auto" src="/images/ValistLogo128.png" alt="Workflow logo" />
                            </div>
                        </Link>
                    </div>

                    <div className="flex lg:hidden">
                        <button className="inline-flex items-center justify-center p-2 rounded-md text-indigo-400 hover:text-white hover:bg-indigo-600 focus:outline-none focus:bg-indigo-600 focus:text-white transition duration-150 ease-in-out" aria-label="Main menu" aria-expanded="false">
                            <svg className="block h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M4 6h16M4 12h8m-8 6h16" />
                            </svg>
                            <svg className="hidden h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M6 18L18 6M6 6l12 12" />
                            </svg>
                        </button>
                    </div>
                    <div className="hidden lg:block lg:w-80">
                    <div className="flex items-center justify-end">
                        <div className="flex">
                        <a href="https://docs.valist.io" className="px-3 py-2 rounded-md text-sm leading-5 font-medium text-indigo-200 hover:text-white focus:outline-none focus:text-white focus:bg-indigo-600 transition duration-150 ease-in-out">Documentation</a>
                        <a href="mailto:support@valist.io?subject=Valist Support" className="px-3 py-2 rounded-md text-sm leading-5 font-medium text-indigo-200 hover:text-white focus:outline-none focus:text-white focus:bg-indigo-600 transition duration-150 ease-in-out">Support</a>
                            {login.loggedIn ? <button onClick={() => login.logOut()} type="button" className="inline-flex ml-2 items-center px-6 py-1.5 border border-transparent text-xs font-medium rounded-full shadow-sm text-indigo-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
                                Logout
                            </button> :
                            <button onClick={() => login.setShowLogin(true)} type="button" className="inline-flex ml-2 items-center px-6 py-1.5 border border-transparent text-xs font-medium rounded-full shadow-sm text-indigo-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
                                Login
                            </button>}
                        </div>

                        <div className="ml-4 relative flex-shrink-0">
                        <div className={"origin-top-right absolute right-0 mt-2 w-48 rounded-md shadow-lg " + "hidden"}>
                            <div className="py-1 rounded-md bg-white shadow-xs" role="menu" aria-orientation="vertical" aria-labelledby="user-menu">
                            <a href="#" className="block px-4 py-2 text-sm leading-5 text-gray-700 hover:bg-gray-100 focus:outline-none focus:bg-gray-100 transition duration-150 ease-in-out" role="menuitem">View Profile</a>
                            <a href="#" className="block px-4 py-2 text-sm leading-5 text-gray-700 hover:bg-gray-100 focus:outline-none focus:bg-gray-100 transition duration-150 ease-in-out" role="menuitem">Settings</a>
                            <a href="#" className="block px-4 py-2 text-sm leading-5 text-gray-700 hover:bg-gray-100 focus:outline-none focus:bg-gray-100 transition duration-150 ease-in-out" role="menuitem">Logout</a>
                            </div>
                        </div>
                        </div>
                    </div>
                    </div>
                </div>
                </div>
                <div className="hidden lg:hidden">
                <div className="px-2 pt-2 pb-3">
                    <a href="#" className="block px-3 py-2 rounded-md text-base font-medium text-white bg-indigo-800 focus:outline-none focus:text-indigo-100 focus:bg-indigo-600 transition duration-150 ease-in-out">Dashboard</a>
                    <a href="#" className="mt-1 block px-3 py-2 rounded-md text-base font-medium text-indigo-200 hover:text-indigo-100 hover:bg-indigo-600 focus:outline-none focus:text-white focus:bg-indigo-600 transition duration-150 ease-in-out">Support</a>
                </div>
            </div>
        </nav>
    );
}

export default Nav;
