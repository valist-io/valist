const setLoading = (text: string): string => {
  if (text === 'Loading') {
    return 'loading';
  }
  return ' ';
};

export default setLoading;
