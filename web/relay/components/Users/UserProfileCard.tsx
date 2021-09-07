import React, { useState, useEffect, useContext } from 'react';
import Link from 'next/link';
import ValistContext from '../Valist/ValistContext';
import LoginContext from '../Login/LoginContext';
import AddressIdenticon from '../Identicons/AddressIdenticon';

const ProfileBox = () => {
  const valist = useContext(ValistContext);
  const login = useContext(LoginContext);
  const [profile, setProfile] = useState({ address: '0x0' });

  useEffect(() => {
    (async () => {
      if (valist) {
        setProfile({ address: valist.defaultAccount });
      }
    })();
  }, [valist]);

  if (login.loggedIn) {
    return (
      <section aria-labelledby="profile-overview-title">
        <div className="rounded-lg bg-white overflow-hidden shadow">
          <h2 className="sr-only" id="profile-overview-title">Profile Overview</h2>
          <div className="bg-white p-6">
            <div className="sm:flex sm:flex-wrap sm:items-center sm:justify-between">
              <div className="sm:flex sm:items-center">
                <div className="pr-4">
                  <AddressIdenticon address={profile.address} height={20}/>
                </div>
                <div>
                  <p className="text-md font-medium text-gray-600">Welcome back! Your current account is:</p>
                  <p className="text-sm font-medium text-gray-900 sm:text-sm pt-2 font-mono">
                    {profile.address}
                  </p>
                </div>
              </div>
              <div className="mt-5 flex justify-center sm:mt-0">
                <Link href="/v/create">
                  <a href="#" className={`flex justify-center items-center px-4 py-2
                  border border-transparent text-sm leading-5 font-medium rounded-md text-white
                  bg-indigo-600 hover:bg-indigo-500 focus:outline-none focus:border-indigo-700
                  focus:shadow-outline-indigo active:bg-indigo-700 transition ease-in-out duration-150`}>
                    Create Organization
                  </a>
                </Link>
              </div>
            </div>
          </div>
          { /* <div className="border-t
          border-gray-200 bg-gray-50 grid
          grid-cols-1 divide-y divide-gray-200 sm:grid-cols-3 sm:divide-y-0 sm:divide-x">
            <div className="px-6 py-5 text-sm font-medium text-center">
              <span className="text-gray-900">2 </span>
              <span className="text-gray-600">Organizations</span>
            </div>

            <div className="px-6 py-5 text-sm font-medium text-center">
              <span className="text-gray-900">12 </span>
              <span className="text-gray-600">Projects</span>
            </div>

            <div className="px-6 py-5 text-sm font-medium text-center">
              <span className="text-gray-900">104 </span>
              <span className="text-gray-600">Releases</span>
            </div>
          </div> */}
        </div>
      </section>
    );
  }

  return (
    <section aria-labelledby="profile-overview-title">
      <div className="rounded-lg bg-white overflow-hidden shadow">
        <div className="bg-white p-6">
          <div className="sm:flex sm:items-center sm:justify-between">
          <div className="sm:flex sm:space-x-5">
            <div className="flex-shrink-0">
              <img className="h-15" src="/images/ValistLogo128.png" />
            </div>
            <div className="mt-4 text-center sm:mt-0 sm:pt-1 sm:text-left">
              <p className="text-3xl font-bold text-gray-900 sm:text-2xl">
                Welcome to Valist
              </p>
              <p className="text-md font-medium text-gray-600">Login to securely distribute software.</p>
            </div>
          </div>
          </div>
        </div>
      </div>
    </section>
  );
};

export default ProfileBox;
