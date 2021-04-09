import Link from 'next/link';
import React from 'react';
import ActivityFeed from './ActivityList';

export const ActivityBox = (): JSX.Element => (
  <section aria-labelledby="announcements-title">
    <div className="rounded-lg bg-white overflow-hidden shadow p-6">
        <h2 className="text-base font-medium text-gray-900">Activity</h2>
        <div className="flow-root mt-6 overflow-hidden max-h-72">
          <ActivityFeed />
        </div>
        <div className="mt-6">
          <Link href="/activity">
            <a className="w-full flex justify-center items-center px-4 py-2 border border-gray-300 shadow-sm
                          text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50">
              View all
            </a>
          </Link>
        </div>
    </div>
  </section>
);

export default ActivityBox;
