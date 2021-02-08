import React, { FunctionComponent } from 'react';
// import styles from './LoadingDialog.module.css';

export const LoadingDialog:FunctionComponent<any> = (props: any) => {
    return (
        <div className="fixed top-0 left-0 z-50 w-screen h-screen flex items-center justify-center" style={{background: "rgba(0, 0, 0, 0.3)"}}>
            <div className="bg-white border py-2 px-5 rounded-lg flex items-center flex-col">
                {/* <div className={`${styles.loaderDots} block relative w-20 h-5 mt-2`}>
                    <div className="absolute top-0 mt-1 w-3 h-3 rounded-full bg-green-500"></div>
                    <div className="absolute top-0 mt-1 w-3 h-3 rounded-full bg-green-500"></div>
                    <div className="absolute top-0 mt-1 w-3 h-3 rounded-full bg-green-500"></div>
                    <div className="absolute top-0 mt-1 w-3 h-3 rounded-full bg-green-500"></div>
                </div> */}
                <svg xmlns="http://www.w3.org/2000/svg" xmlnsXlink="http://www.w3.org/1999/xlink" style={{margin: 'auto', display: 'block'}} width="200px" height="200px" viewBox="0 0 100 100" preserveAspectRatio="xMidYMid">
                    <circle cx={50} cy={50} r={32} strokeWidth={8} stroke="#5145cd" strokeDasharray="50.26548245743669 50.26548245743669" fill="none" strokeLinecap="round">
                        <animateTransform attributeName="transform" type="rotate" dur="1s" repeatCount="indefinite" keyTimes="0;1" values="0 50 50;360 50 50" />
                    </circle>
                    <circle cx={50} cy={50} r={23} strokeWidth={8} stroke="#000000" strokeDasharray="36.12831551628262 36.12831551628262" strokeDashoffset="36.12831551628262" fill="none" strokeLinecap="round">
                        <animateTransform attributeName="transform" type="rotate" dur="1s" repeatCount="indefinite" keyTimes="0;1" values="0 50 50;-360 50 50" />
                    </circle>
                </svg>
                <div className="text-gray-500 font-light mt-2 text-center">
                { props.children || "Please Wait..." }
                </div>
            </div>
        </div>
    )
}

export default LoadingDialog;
