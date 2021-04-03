import React, { useContext, useEffect } from 'react';
import Layout from '../../../components/Dashboard/Layout';
import ProjectList from '../../../components/List/ProjectList';
import ProfileBox from '../../../components/Dashboard/ProfileBox';
import { useRouter } from 'next/router';

export default function Dashboard(){
  const router = useRouter();
  const { orgName } = router.query;

  return(
    <Layout>
        <main className="-mt-24 pb-8">
          <div className="max-w-3xl mx-auto px-4 sm:px-6 lg:max-w-7xl lg:px-8">
            <h1 className="sr-only">Profile</h1>
            <div className="grid grid-cols-1 gap-4 items-start lg:grid-cols-3 lg:gap-8">
              <div className="grid grid-cols-1 gap-4 lg:col-span-2">
                <ProfileBox />
                <ProjectList orgName={`${orgName}`}/>
              </div>
              <div className="grid grid-cols-1 gap-4">
                <section aria-labelledby="announcements-title">
                  <div className="rounded-lg bg-white overflow-hidden shadow">
                    <div className="p-6">
                      <h2 className="text-base font-medium text-gray-900" id="announcements-title">Activity</h2>
                      <div className="flow-root mt-6">
                        <ul className="-my-5 divide-y divide-gray-200">

                          <li className="py-5">
                            <div className="relative focus-within:ring-2 focus-within:ring-cyan-500">
                              <h3 className="text-sm font-semibold text-gray-800">
                                <a href="#" className="hover:underline focus:outline-none">
                                  <span className="absolute inset-0" aria-hidden="true"></span>
                                  Valist organization created.
                                </a>
                              </h3>
                              <p className="mt-1 text-sm text-gray-600 line-clamp-2">
                                Cum qui rem deleniti. Suscipit in dolor veritatis sequi aut. Vero ut earum quis deleniti. Ut a sunt eum cum ut repudiandae possimus. Nihil ex tempora neque cum consectetur dolores.
                              </p>
                            </div>
                          </li>

                        </ul>
                      </div>
                      <div className="mt-6">
                        <a href="#" className="w-full flex justify-center items-center px-4 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50">
                          View all
                        </a>
                      </div>
                    </div>
                  </div>
                </section>

                
              </div>
            </div>
          </div>
        </main>
    </Layout>
  )
}
