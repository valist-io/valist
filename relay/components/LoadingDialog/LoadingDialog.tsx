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
                <div className="text-gray-500 font-light mt-2 text-center">
                { props.children || "Please Wait..." }
                </div>
            </div>
        </div>
    )
}

export default LoadingDialog;
