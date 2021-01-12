const AddressIdenticon = ({address}: {address: string}) => (
    <div className="flex items-center space-x-3">
        <div className="flex-shrink-0 h-12 w-12">
            <img className="h-12 w-12 rounded-full bg-black" src={`https://identicon-api.herokuapp.com/${address}/32?format=png`} alt="" />
        </div>
        <div className="space-y-1">
            <div className="text-sm leading-5 font-medium text-gray-900">{address.replace(address.substring(8,36), "...")}</div>
        </div>
    </div>
);

export default AddressIdenticon;
