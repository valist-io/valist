const AddressIdenticon = ({address, height}: {address: string, height: number}) => (
  <img className={`mx-auto h-${height} w-${height} rounded-full`} src={`https://identicon-api.herokuapp.com/${address}/128?format=png`} alt="" />
);

export default AddressIdenticon;
