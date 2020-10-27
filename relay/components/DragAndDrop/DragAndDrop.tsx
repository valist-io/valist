import React from 'react';

const DragAndDrop = (props: any) => {

    const { data, dispatch } = props;
    
    const handleDragEnter = (e: any) => {
        e.preventDefault();
        e.stopPropagation();
        dispatch({ type: 'SET_DROP_DEPTH', dropDepth: data.dropDepth + 1 });
    };
    const handleDragLeave = (e: any) => {
        e.preventDefault();
        e.stopPropagation();
        dispatch({ type: 'SET_DROP_DEPTH', dropDepth: data.dropDepth - 1 });
        if (data.dropDepth > 0) return
        dispatch({ type: 'SET_IN_DROP_ZONE', inDropZone: false })
    };
    const handleDragOver = (e: any) => {
        e.preventDefault();
        e.stopPropagation();
        e.dataTransfer.dropEffect = 'copy';
        dispatch({ type: 'SET_IN_DROP_ZONE', inDropZone: true });
    };
    const handleDrop = (e: any) => {
        e.preventDefault();
        e.stopPropagation();
        let files = [...e.dataTransfer.files];

        if (files && files.length > 0) {
            const existingFiles = data.fileList.map((f:any) => f.name)
            files = files.filter(f => !existingFiles.includes(f.name))
            
            dispatch({ type: 'ADD_FILE_TO_LIST', files });
            //e.dataTransfer.clearData();
            dispatch({ type: 'SET_DROP_DEPTH', dropDepth: 0 });
            dispatch({ type: 'SET_IN_DROP_ZONE', inDropZone: false });
        }
    };

    return (
        <div className="sm:col-span-2">
            <div className="border-4 border-dashed border-gray-200 rounded-lg h-96 relative h-32 ">
                <div onDrop={e => handleDrop(e)}
                    onDragOver={e => handleDragOver(e)}
                    onDragEnter={e => handleDragEnter(e)}
                    onDragLeave={e => handleDragLeave(e)}
                    className="absolute inset-0 flex items-center justify-center">
                    Drop Files Here
                </div> 
            </div>
        </div>
    )
};
export default DragAndDrop;