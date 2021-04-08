import { toSvg } from "jdenticon";

const AddressIdenticon = ({address, height}: {address: string, height: number}) => {
  const svgString = toSvg(address, 128);
  return (
    <img className={`mx-auto h-${height} w-${height} rounded-full`} src={`data:image/svg+xml;base64,${Buffer.from(svgString, 'utf8').toString('base64')}`} alt="" />
  )
};

export default AddressIdenticon;
