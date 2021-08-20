import React, { useState, useEffect } from 'react';

interface Props {
  threshold: number,
  voteThreshold: (threshold: string) => Promise<void>
}

const EditProjectThresholdForm: React.FC<Props> = (props: Props): JSX.Element => {
  const [threshold, setThreshold] = useState('0');

  useEffect(() => {
    setThreshold(`${props.threshold}`);
  }, [props.threshold]);

  return (
      <div className="px-4 py-5 sm:rounded-lg sm:p-6">
          <div className="md:grid md:grid-cols-3 md:gap-6">
              <div className="md:col-span-1">
                  <h3 className="text-lg font-medium leading-6 text-gray-900">Multi-factor</h3>
                  <p className="mt-1 text-sm leading-5 text-gray-500">
                      This sets the amount of votes required to add / remove keys and publish new releases.
                      Three or more members are required to enable multi-factor voting.
                  </p>
              </div>
              <div className="mt-5 md:mt-0 md:col-span-2">
                  <form className="grid grid-cols-1 gap-y-6 sm:grid-cols-2 sm:gap-x-8">
                      <div className="sm:col-span-2">
                          <label htmlFor="VoteThreshold" className="block text-sm font-medium
                          leading-5 text-gray-700">Voting Threshold</label>
                          <div className="mt-1 relative rounded-md shadow-sm">
                              <input value={threshold} type="number"
                              onChange={ (e) => setThreshold(e.target.value) }
                              required id="VoteThreshold" className="form-input block w-full sm:text-sm
                              sm:leading-5 transition ease-in-out duration-150" />
                          </div>
                      </div>
                      <div className="sm:col-span-2">
                      <span className="w-full inline-flex rounded-md shadow-sm">
                          <button onClick={() => props.voteThreshold(threshold)}
                          value="Submit" type="button"
                          className="w-full inline-flex items-center justify-center px-6 py-3 border
                          border-transparent text-base leading-6 font-medium rounded-md text-white
                          bg-indigo-600 hover:bg-indigo-500 focus:outline-none focus:border-indigo-700
                          focus:shadow-outline-indigo active:bg-indigo-700 transition ease-in-out duration-150">
                              Update Threshold
                          </button>
                      </span>
                      </div>
                  </form>
              </div>
          </div>
      </div>
  );
};

export default EditProjectThresholdForm;
