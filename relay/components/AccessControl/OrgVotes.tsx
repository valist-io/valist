/* eslint-disable no-underscore-dangle */
import React from 'react';
import { VoteKeyEvent, VoteThresholdEvent } from 'valist/dist/types';
import { ADD_KEY, REVOKE_KEY, ROTATE_KEY } from 'valist/dist/constants';

const OrgVotes = ({
  orgKeyVotes,
  orgThresholdVotes,
  voteAdmin,
  revokeAdmin,
  voteThreshold,
}: {
  orgKeyVotes: VoteKeyEvent[],
  orgThresholdVotes: VoteThresholdEvent[],
  voteAdmin: (key: string) => Promise<void>,
  revokeAdmin: (key: string) => Promise<void>,
  voteThreshold: (threshold: number) => Promise<void>
}) => (
      <div className="flex flex-col w-full">
        <div className="-my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
          <div className="py-2 align-middle inline-block min-w-full sm:px-6 lg:px-8">
            <div className="shadow overflow-hidden border-b border-gray-200 sm:rounded-lg">
              <table className="min-w-full divide-y divide-gray-200">
                <thead className="bg-gray-50">
                  <tr>
                    <th scope="col"
                    className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Description
                    </th>
                    <th scope="col"
                    className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Info
                    </th>
                    <th scope="col"
                    className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Votes
                    </th>
                    <th scope="col" className="relative px-6 py-3">
                      <span className="sr-only">Edit</span>
                    </th>
                  </tr>
                </thead>
                <tbody className="bg-white divide-y divide-gray-200">
                  { orgKeyVotes && orgKeyVotes.filter((vote) => vote._operation !== ROTATE_KEY).map((vote, index) => (
                    <tr key={index}>
                      <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                        { vote._operation === ADD_KEY && 'Add Key' }
                        { vote._operation === REVOKE_KEY && 'Revoke Key' }
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500 font-mono">
                        { vote._key }
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                        { vote._sigCount } / { vote._threshold }
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                        <a href="#" className="text-indigo-600 hover:text-indigo-900"
                        onClick={() => {
                          if (vote._operation === ADD_KEY) {
                            voteAdmin(vote._key);
                          } else if (vote._operation === REVOKE_KEY) {
                            revokeAdmin(vote._key);
                          }
                        }}>
                          Approve
                        </a>
                      </td>
                    </tr>
                  ))}
                  { orgThresholdVotes && orgThresholdVotes.map((vote, index) => (
                    <tr key={index}>
                      <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                        Set Threshold
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500 font-mono">
                        { vote._pendingThreshold }
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                        { vote._sigCount } / { vote._threshold }
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                        <a href="#" className="text-indigo-600 hover:text-indigo-900"
                        onClick={() => voteThreshold(parseInt(vote._pendingThreshold, 10)) }>
                          Approve
                        </a>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
);

export default OrgVotes;
