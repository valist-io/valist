import React, { useContext } from 'react';
import ValistContext from '../Valist/ValistContext';
import { ADD_KEY, REVOKE_KEY, ROTATE_KEY } from 'valist/dist/constants';

const OrgVotes = ({ orgName, orgVotes }: { orgName: string, orgVotes: any[] }) => {
  const valist = useContext(ValistContext);

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
                      Expiration
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
                  { orgVotes && orgVotes.map((vote, index) => (
                    <tr key={index}>
                      <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                        { vote.returnValues._operation === ADD_KEY && 'Add Key' }
                        { vote.returnValues._operation === REVOKE_KEY && 'Revoke Key' }
                        { vote.returnValues._operation === ROTATE_KEY && 'Rotate Key' }
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500 font-mono">
                        { vote.returnValues._key }
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                        TODO!!!
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                        { vote.returnValues._sigCount } / { vote.returnValues._threshold }
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                        <a href="#" className="text-indigo-600 hover:text-indigo-900">Approve</a>
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

export default OrgVotes;
