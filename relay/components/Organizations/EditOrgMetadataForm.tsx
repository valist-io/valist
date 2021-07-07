import React, { useState, useEffect, useContext } from 'react';
import ValistContext from '../Valist/ValistContext';
import { OrgMeta } from '../../../lib/dist/types';

const EditOrgMetadataForm = ({ orgName, orgMeta }: { orgName: string, orgMeta: OrgMeta }) => {
  const valist = useContext(ValistContext);

  const [name, setName] = useState('');
  const [description, setDescription] = useState('');

  const updateOrgMeta = async () => {
    try {
      await valist.setOrgMeta(orgName, { name, description }, valist.defaultAccount);
    } catch (e) { console.log(e); }
  };

  useEffect(() => {
    setName(orgMeta.name);
    setDescription(orgMeta.description);
  }, [valist, orgName]);

  return (
    <div className="px-4 py-5 sm:rounded-lg sm:p-6">
        <div className="md:grid md:grid-cols-3 md:gap-6">
            <div className="md:col-span-1">
                <h3 className="text-lg font-medium leading-6 text-gray-900">Metadata</h3>
                <p className="mt-1 text-sm leading-5 text-gray-500">
                    This information will be displayed publicly so be careful what you share.
                </p>
            </div>
            <div className="mt-5 md:mt-0 md:col-span-2">
                <form className="grid grid-cols-1 gap-y-6 sm:grid-cols-2 sm:gap-x-8">
                    <div className="sm:col-span-2">
                        <label htmlFor="OrgFullName" className="block text-sm font-medium
                        leading-5 text-gray-700">Full Name</label>
                        <div className="mt-1 relative rounded-md shadow-sm">
                            <input value={name} type="text" onChange={
                              (e) => setName(e.target.value)
                            }
                            required id="OrgFullName" className="form-input block w-full sm:text-sm
                            sm:leading-5 transition ease-in-out duration-150" />
                        </div>
                    </div>
                    <div className="sm:col-span-2">
                        <label htmlFor="OrgDescription" className="block text-sm font-medium leading-5
                        text-gray-700">Description</label>
                        <div className="mt-1 relative rounded-md shadow-sm">
                            <textarea value={description} onChange={(e) => setDescription(e.target.value)}
                            required id="OrgDescription" className="h-20 form-input block w-full sm:text-sm
                            sm:leading-5 transition ease-in-out duration-150" />
                        </div>
                    </div>
                    <div className="sm:col-span-2">
                    <span className="w-full inline-flex rounded-md shadow-sm">
                        <button onClick={updateOrgMeta} value="Submit" type="button" className="w-full inline-flex
                        items-center justify-center px-6 py-3 border border-transparent text-base leading-6
                        font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-500 focus:outline-none
                        focus:border-indigo-700 focus:shadow-outline-indigo active:bg-indigo-700 transition
                        ease-in-out duration-150">
                            Update Metadata
                        </button>
                    </span>
                    </div>
                </form>
            </div>
        </div>
    </div>
  );
};

export default EditOrgMetadataForm;
