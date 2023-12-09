interface Props {
    name: string;
    address: string;
    rating: number;
    url: string;
    photos: string[];
}

const Place = (props:Props) => {
    return (
        <li className="w-full flex bg-neutral-700 shadow-md rounded-lg">
                        {/* <div>{props.photos.map(i => <img src={i}/>)}</div> */}
            <div className="aspect-square w-40 object-cover">
                <img src={props.photos[0]}/>
            </div>
            <div className="flex flex-col">
                <h2 className="">{props.name}</h2>  
                <h3>{props.address}</h3>
                <h3>{props.rating}</h3>
                {/* <h3>{props.url}</h3> */}
            </div>


        </li>
    )
}

export default Place;