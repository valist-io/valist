const copyToCB = (element: any) => {
  if (typeof document !== 'undefined') {
    const copyBoxElement = element.current;
    if (copyBoxElement) {
      copyBoxElement.contentEditable = true;
      copyBoxElement.focus();
      document.execCommand('selectAll');
      document.execCommand('copy');
      copyBoxElement.contentEditable = false;
      // @ts-ignore
      getSelection().empty();
    }
  }
};

export default copyToCB;
