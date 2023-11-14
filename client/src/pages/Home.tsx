import Layout from "../components/Layout";
import { useState } from "react";

const Home = () => {
    const [input, setInput] = useState<string>("");
    const [distance, setDistance] = useState<number>(25);
    const [filter, setFilter] = useState<string>("")

    return (
        <Layout>
            <form 
                onSubmit={(e) => {console.log(input, filter, distance); e.preventDefault();}}
                className="h-full w-screen flex flex-col justify-center items-center space-y-2"
            >
                <div
                    className="w-1/3 h-16 place-self-center flex items-center shadow-inner rounded-full bg-neutral-100"
                >
                    <input 
                        type="search"
                        placeholder="search"
                        onChange={(e) => setInput(e.target.value)} 
                        className="h-full bg-transparent rounded-l-full flex-grow focus:outline-none pl-8 text-lg"
                    />
                    <button className="rounded-r-full h-full w-16 pl-4">
                        <svg className="stroke-neutral-400 hover:stroke-neutral-800 hover:ping w-6 h-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5}>
                            <path strokeLinecap="round" strokeLinejoin="round" d="M21 21l-5.197-5.197m0 0A7.5 7.5 0 105.196 5.196a7.5 7.5 0 0010.607 10.607z" />
                        </svg>
                    </button>
                </div>

                {/* filters */}
                <div className="w-1/3 h-12 flex items-center space-x-2 text-neutral-800">
                    <button 
                        onClick={(e) => {setFilter(`${filter =="attractions" ? "" : "attractions"}`); e.preventDefault();}}
                        className={`rounded-lg bg-neutral-100 shaddow-inner h-full w-1/6 ${filter == "attractions" ? "bg-neutral-300 text-neutral-100" : ""}`}
                    >
                        Attractions
                    </button>
                    <button 
                        onClick={(e) => {setFilter(`${filter =="restaurants" ? "" : "restaurants"}`); e.preventDefault();}}
                        className={`rounded-lg bg-neutral-100 shaddow-inner h-full w-1/6 ${filter == "restaurants" ? "bg-neutral-300 text-neutral-100" : ""}`}
                    >
                        Restaurants
                    </button>
                    <button 
                        onClick={(e) => {setFilter(`${filter =="hotels" ? "" : "hotels"}`); e.preventDefault();}}
                        className={`rounded-lg bg-neutral-100 shaddow-inner h-full w-1/6 ${filter == "hotels" ? "bg-neutral-300 text-neutral-100" : ""}`}
                    >
                        Hotels
                    </button>
                    <div className="rounded-lg bg-neutral-100 shaddow-inner h-full flex-grow flex items-center justify-center space-x-2">
                        <input 
                            type="range" 
                            min="1" 
                            max="50" 
                            value={distance} 
                            onChange={(e) => (setDistance(parseInt(e.target.value)))}
                        />
                        <span>
                            {distance} miles
                        </span>
                    </div>
                </div>
            </form>
        </Layout>
    )
}

export default Home;