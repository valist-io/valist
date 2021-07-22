/* eslint-disable no-underscore-dangle */
import React, { useContext } from 'react';
import {
  VoteThresholdEvent, VoteKeyEvent, VoteReleaseEvent, Release, PendingRelease,
} from 'valist/dist/types';
import { ADD_KEY, REVOKE_KEY } from 'valist/dist/constants';
import ValistContext from '../Valist/ValistContext';

interface VotesProps {
  votes: any[],
  pendingKeys: string[],
  pendingThresholds: string[],
  pendingReleases?: PendingRelease[],
  grantKey: (key: string) => Promise<void>,
  revokeKey: (key: string) => Promise<void>,
  voteThreshold: (threshold: string) => Promise<void>,
  voteRelease?: (release: Release) => Promise<void>,
  clearPendingKey: (key: string, operation: string, index: number) => Promise<void>,
  clearPendingThreshold: (threshold: string, index: number) => Promise<void>,
  clearPendingRelease?: (release: Release, index: number) => Promise<void>,
}

interface VoteThresholdProps {
  event: VoteThresholdEvent,
  index: number,
  pending: string,
  clear: (threshold: string, index: number) => Promise<void>,
  vote: (threshold: string) => Promise<void>
}

interface VoteKeyProps {
  event: VoteKeyEvent,
  index: number,
  pending: string,
  clear: (key: string, operation: string, index: number) => Promise<void>,
  vote: (key: string) => Promise<void>
}

interface VoteReleaseProps {
  event: VoteReleaseEvent,
  index: number,
  pending: PendingRelease,
  tag: string,
  clear: (release: Release, index: number) => Promise<void>,
  vote: (release: Release) => Promise<void>
}

const VoteThreshold: React.FC<VoteThresholdProps> = (props: VoteThresholdProps): JSX.Element => {
  const threshold = props.event._pendingThreshold > props.event._threshold
    ? props.event._pendingThreshold
    : props.event._threshold;
  const completed = props.event._sigCount >= threshold;

  return (
    <tr>
      <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
        Set Threshold
      </td>
      <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500 font-mono">
        { props.event._pendingThreshold }
      </td>
      <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
        { props.event._sigCount } / { threshold }
      </td>
      <td className="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
        <a href="#" className="text-indigo-600 hover:text-indigo-900"
        onClick={() => (completed
          ? props.clear(props.event._pendingThreshold, props.index)
          : props.vote(props.event._pendingThreshold))}>
          { completed ? 'Clear' : 'Approve' }
        </a>
      </td>
    </tr>
  );
};

const VoteKey: React.FC<VoteKeyProps> = (props: VoteKeyProps): JSX.Element => {
  const completed = props.event._sigCount >= props.event._threshold;

  return (
    <tr>
      <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
        { props.event._operation === ADD_KEY && 'Add Key' }
        { props.event._operation === REVOKE_KEY && 'Revoke Key' }
      </td>
      <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500 font-mono">
        { props.event._key }
      </td>
      <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
        { props.event._sigCount } / { props.event._threshold }
      </td>
      <td className="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
        <a href="#" className="text-indigo-600 hover:text-indigo-900"
        onClick={() => (completed
          ? props.clear(props.event._key, props.event._operation, props.index)
          : props.vote(props.event._key)) }>
          { completed ? 'Clear' : 'Approve' }
        </a>
      </td>
    </tr>
  );
};

const VoteRelease: React.FC<VoteReleaseProps> = (props: VoteReleaseProps): JSX.Element => {
  const completed = props.event._sigCount >= props.event._threshold;

  return (
    <tr>
      <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
        Publish Release
      </td>
      <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500 font-mono">
        tag: { props.tag }<br />
        metaCID: {props.event._metaCID }<br />
        releaseCID: { props.event._releaseCID }
      </td>
      <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
        { props.event._sigCount } / { props.event._threshold }
      </td>
      <td className="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
        <a href="#" className="text-indigo-600 hover:text-indigo-900"
        onClick={() => (completed
          ? props.clear(props.pending, props.index)
          : props.vote(props.pending))}>
          { completed ? 'Clear' : 'Approve' }
        </a>
      </td>
    </tr>
  );
};

const Votes: React.FC<VotesProps> = (props: VotesProps): JSX.Element => {
  const valist = useContext(ValistContext);

  const getKeyVote = (pending: string, index: number): JSX.Element => {
    const vote = props.votes.find((v) => v.returnValues._key === pending);
    if (!vote) return <React.Fragment key={index} />;
    return <VoteKey key={index} index={index} pending={pending}
      event={vote.returnValues} clear={props.clearPendingKey} vote={props.grantKey} />;
  };

  const getThresholdVote = (pending: string, index: number): JSX.Element => {
    const vote = props.votes.find((v) => v.returnValues._pendingThreshold === pending);
    if (!vote) return <React.Fragment key={index} />;
    return <VoteThreshold key={index} index={index} pending={pending}
      event={vote.returnValues} clear={props.clearPendingThreshold} vote={props.voteThreshold} />;
  };

  const getReleaseVote = (pending: PendingRelease, index: number): JSX.Element => {
    const vote = props.votes.find((v) => v.returnValues._tag === valist.web3.utils.keccak256(pending.tag)
        && v.returnValues._releaseCID === pending.releaseCID
        && v.returnValues._metaCID === pending.metaCID);
    if (!vote || !props.clearPendingRelease || !props.voteRelease) return <React.Fragment key={index} />;
    return <VoteRelease key={index} index={index} pending={pending} tag={pending.tag}
      event={vote.returnValues} clear={props.clearPendingRelease} vote={props.voteRelease} />;
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
                  { props.pendingKeys.map(getKeyVote) }
                  { props.pendingThresholds.map(getThresholdVote) }
                  { props.pendingReleases && props.pendingReleases.map(getReleaseVote) }
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
  );
};

export default Votes;
