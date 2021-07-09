/* eslint-disable no-underscore-dangle */
import React from 'react';
import {
  VoteKeyEvent, VoteThresholdEvent, VoteReleaseEvent, Release,
} from 'valist/dist/types';
import { ADD_KEY, REVOKE_KEY } from 'valist/dist/constants';

interface VotesProps {
  keyVotes: VoteKeyEvent[],
  thresholdVotes: VoteThresholdEvent[],
  releaseVotes?: VoteReleaseEvent[],
  grantKey: (key: string) => Promise<void>,
  revokeKey: (key: string) => Promise<void>,
  voteThreshold: (threshold: string) => Promise<void>,
  voteRelease?: (release: Release) => Promise<void>,
  clearPendingKey: (key: string, operation: string, index: number) => Promise<void>,
  clearPendingThreshold: (threshold: string, index: number) => Promise<void>,
  clearPendingRelease?: (release: Release, index: number) => Promise<void>
}

const Votes: React.FC<VotesProps> = (props: VotesProps): JSX.Element => {
  const getVoteThreshold = (vote: VoteThresholdEvent) => {
    if (vote._pendingThreshold > vote._threshold) {
      return vote._pendingThreshold;
    }
    return vote._threshold;
  };

  return (
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
                  { props.keyVotes && props.keyVotes.map((vote, index) => (
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
                          if (vote._sigCount >= vote._threshold) {
                            props.clearPendingKey(vote._key, vote._operation, vote.index);
                          } else if (vote._operation === ADD_KEY) {
                            props.grantKey(vote._key);
                          } else if (vote._operation === REVOKE_KEY) {
                            props.revokeKey(vote._key);
                          }
                        }}>
                          { vote._sigCount >= vote._threshold ? 'Clear' : 'Approve' }
                        </a>
                      </td>
                    </tr>
                  ))}
                  { props.thresholdVotes && props.thresholdVotes.map((vote, index) => (
                    <tr key={index}>
                      <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                        Set Threshold
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500 font-mono">
                        { vote._pendingThreshold }
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                        { vote._sigCount } / { getVoteThreshold(vote) }
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                        <a href="#" className="text-indigo-600 hover:text-indigo-900"
                        onClick={() => {
                          if (vote._sigCount >= getVoteThreshold(vote)) {
                            props.clearPendingThreshold(vote._pendingThreshold, vote.index);
                          } else {
                            props.voteThreshold(vote._pendingThreshold);
                          }
                        }}>
                          { vote._sigCount >= getVoteThreshold(vote) ? 'Clear' : 'Approve' }
                        </a>
                      </td>
                    </tr>
                  ))}
                  { props.releaseVotes && props.releaseVotes.map((vote, index) => (
                    <tr key={index}>
                      <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                        Publish Release
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500 font-mono">
                        tag={ vote._tag }
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                        { vote._sigCount } / { vote._threshold }
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                        <a href="#" className="text-indigo-600 hover:text-indigo-900"
                        onClick={() => {
                          if (vote._sigCount >= vote._threshold && props.clearPendingRelease) {
                            props.clearPendingRelease(vote.release, vote.index);
                          } else if (vote._sigCount < vote._threshold && props.voteRelease) {
                            props.voteRelease(vote.release);
                          }
                        }}>
                          { vote._sigCount >= vote._threshold ? 'Clear' : 'Approve' }
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
};

export default Votes;
