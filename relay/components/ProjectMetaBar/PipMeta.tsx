import React, { useEffect, useState } from 'react';

const PipMeta = (orgName: string = "organization", repoName: string = "repo") => {

    const [origin, setOrigin] = useState("https://app.valist.io");
    useEffect(() => {
        setOrigin(window.location.origin);
    });

    return (
        <div>
            <div className="pl-6 lg:w-80">
                <div className="pt-6 pb-2">
                    <h1 className="flex-1 text-lg leading-7 font-medium">Pip Install From Url</h1>
                </div>
                <div className="border-2 border-solid border-black-200 rounded-lg p-2">
                    pip install {origin}/api/{orgName}/{repoName}/latest
                </div>
            </div>
        </div>
    );
}

export default PipMeta;
