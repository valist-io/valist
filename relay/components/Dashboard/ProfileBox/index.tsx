import { useState, useEffect, useContext } from 'react';
import ValistContext from '../../ValistContext';
import AddressIdenticon from '../AddressIdenticon';

const ProfileBox = () => {
  const valist = useContext(ValistContext);
  const [profile, setProfile] = useState({ address: "0x0" });

  useEffect(() => {
    (async function() {
        if (valist) {
            setProfile({ address: valist.defaultAccount });
        }
    })()
  }, [valist]);

  return (
    <section aria-labelledby="profile-overview-title">
      <div className="rounded-lg bg-white overflow-hidden shadow">
        <h2 className="sr-only" id="profile-overview-title">Profile Overview</h2>
        <div className="bg-white p-6">
          <div className="sm:flex sm:items-center sm:justify-between">
            <div className="sm:flex sm:space-x-5">
              <div className="flex-shrink-0">
                <AddressIdenticon address={profile.address} height={20}/>
              </div>
              <div className="mt-4 text-center sm:mt-0 sm:pt-1 sm:text-left">
                <p className="text-sm font-medium text-gray-600">Welcome back,</p>
                <p className="text-xl font-bold text-gray-900 sm:text-2xl">{profile.address.replace(profile.address.substring(8,36), "...")}</p>
                <p className="text-sm font-medium text-gray-600">Software Developer</p>
              </div>
            </div>
            <div className="mt-5 flex justify-center sm:mt-0">
              <a href="#" className="flex justify-center items-center px-4 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50">
                Create Project
              </a>
            </div>
          </div>
        </div>
        <div className="border-t border-gray-200 bg-gray-50 grid grid-cols-1 divide-y divide-gray-200 sm:grid-cols-3 sm:divide-y-0 sm:divide-x">
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
        </div>
      </div>
    </section>
  )
}

export default ProfileBox;