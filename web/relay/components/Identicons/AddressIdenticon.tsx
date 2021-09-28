import { toSvg } from 'jdenticon';

const AddressIdenticon = ({ address, height }: { address: string, height: number }): JSX.Element => {
  const svgString = toSvg(address, 128);
  return (
    <img height={height} width={height} className='mx-auto rounded-full'
      src={`data:image/svg+xml;base64,${Buffer.from(svgString, 'utf8').toString('base64')}`} alt="" />
  );
};

export default AddressIdenticon;
