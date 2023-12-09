// import Layout from "../components/Layout";
import { useState } from "react";
import { GoogleLogin } from "@react-oauth/google";

const Home = () => {
  const [input, setInput] = useState<string>("");
  const [distance, setDistance] = useState<number>(25);
  const [filter, setFilter] = useState<string>("");
  const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);

  async function fetchBackendData(): Promise<any> {
    // const queryParams = {
    //     location: input,
    //     distance: distance,
    //     filter: filter,
    // };

    // const queryString = Object.keys(queryParams)
    // .map(key => `${encodeURIComponent(key)}=${encodeURIComponent(queryParams[key])}`)
    // .join('&');
    const queryString =
      "location=" +
      encodeURIComponent(input) +
      "&distance=" +
      encodeURIComponent(distance) +
      "&filter=" +
      encodeURIComponent(filter);

    try {
      const response = await fetch(
        `http://localhost:8080/api/search?${queryString}`
      );
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
      const data = await response.json();
      // Process your data
      return data; // Return the data for further processing
    } catch (error) {
      console.error("Error fetching data: ", error);
      // Handle the error
      throw error; // Re-throw the error if you want to handle it at a higher level
    }
  }

  const handleLoginSuccess = (credentialResponse) => {
    console.log(credentialResponse);

    // Update the authentication status
    setIsLoggedIn(true);
  };

  const handleLoginError = () => {
    console.log("Login Failed");
  };

  return (
    <div className="w-full h-screen pt-4 bg-neutral-900">
      <div className="flex justify-center w-full">
        <form
          onSubmit={(e) => {
            fetchBackendData();
            e.preventDefault();
          }}
          className="h-full w-2/5 flex flex-col justify-center items-center space-y-2"
        >
          <h1 className="text-4xl font-semibold bg-gradient-to-r from-cyan-500 to-blue-500 text-transparent bg-clip-text">
            Travel Planner
          </h1>
          <div className="w-full h-16 place-self-center flex items-center shadow-inner rounded-full bg-neutral-700">
            <input
              type="text"
              placeholder="Search Destination"
              onChange={(e) => setInput(e.target.value)}
              className="h-full bg-transparent rounded-l-full flex-grow focus:outline-none pl-8 text-lg"
            />
            <button className="rounded-r-full h-full w-16 pl-4">
              <svg
                className="stroke-neutral-400 hover:stroke-neutral-200 hover:ping w-6 h-6"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                strokeWidth={1.5}
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  d="M21 21l-5.197-5.197m0 0A7.5 7.5 0 105.196 5.196a7.5 7.5 0 0010.607 10.607z"
                />
              </svg>
            </button>
          </div>

          <div className="w-full h-full flex justify-center items-center space-x-2 text-neutral-200">
            {isLoggedIn ? (
              // Render Home content for authenticated user
              <div className="w-full h-full text-center">
                <h1 className="text-4xl font-semibold bg-gradient-to-r from-cyan-500 to-blue-500 text-transparent bg-clip-text">
                  Plan your Dream Vacation!
                </h1>
              </div>
            ) : (
              // Render Login component for non-authenticated user
              <GoogleLogin
                onSuccess={handleLoginSuccess}
                onError={handleLoginError}
              />
            )}
          </div>

          <div className="w-full h-12 flex items-center space-x-2 text-neutral-200">
            <button
              onClick={(e) => {
                setFilter(`${filter == "attractions" ? "" : "attractions"}`);
                e.preventDefault();
              }}
              className={`rounded-lg bg-neutral-700 shaddow-lg h-full w-1/5 ${
                filter == "attractions"
                  ? "bg-gradient-to-r from-cyan-500 to-blue-500 text-neutral-100 hover:from-cyan-700 hover:to-blue-700"
                  : "hover:bg-neutral-800"
              }`}
            >
              Attractions
            </button>
            <button
              onClick={(e) => {
                setFilter(`${filter == "restaurants" ? "" : "restaurants"}`);
                e.preventDefault();
              }}
              className={`rounded-lg bg-neutral-700 shaddow-lg h-full w-1/5 ${
                filter == "restaurants"
                  ? "bg-gradient-to-r from-cyan-500 to-blue-500 text-neutral-100 hover:from-cyan-700 hover:to-blue-700"
                  : "hover:bg-neutral-800"
              }`}
            >
              Restaurants
            </button>
            <button
              onClick={(e) => {
                setFilter(`${filter == "hotels" ? "" : "hotels"}`);
                e.preventDefault();
              }}
              className={`rounded-lg bg-neutral-700 shaddow-lg h-full w-1/5 ${
                filter == "hotels"
                  ? "bg-gradient-to-r from-cyan-500 to-blue-500 text-neutral-100 hover:from-cyan-700 hover:to-blue-700"
                  : "hover:bg-neutral-800"
              }`}
            >
              Hotels
            </button>
            <div className="rounded-lg bg-neutral-700 shaddow-lg h-full flex-grow flex items-center justify-center space-x-2">
              <input
                type="range"
                min="1"
                max="50"
                value={distance}
                onChange={(e) => setDistance(parseInt(e.target.value))}
              />
              <span>{distance} miles</span>
            </div>
          </div>
        </form>
      </div>
    </div>
  );
};

export default Home;
