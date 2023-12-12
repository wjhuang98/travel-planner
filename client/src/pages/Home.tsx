import { useState } from "react";
import Place from "../components/Place";
import { GoogleLogin, googleLogout } from "@react-oauth/google";

const Home = () => {
  const [input, setInput] = useState<string>("");
  const [distance, setDistance] = useState<number>(25);
  const [filter, setFilter] = useState<string>("");
  const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);

  interface apiResponse {
    name: string;
    address: string;
    rating: number;
    url: string;
    photos: string[];
  }

  const [data, setData] = useState<apiResponse[]>([]);

  async function fetchBackendData() {
    const queryString = "location=" + encodeURIComponent(input) + "&radius=" + encodeURIComponent(distance) + "&filter=" + encodeURIComponent(filter);

    try {
      const response = await fetch(`http://localhost:8080/api/search?${queryString}`);
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
      // const data = await response.json();
      setData(await response.json());
      console.log(data);
    } catch (error) {
      console.error('Error fetching data: ', error);
      // Handle the error
      throw error; // Re-throw the error if you want to handle it at a higher level
    }
  }

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const handleLoginSuccess = (credentialResponse: any) => {
    console.log(credentialResponse);

    // Update the authentication status
    setIsLoggedIn(true);
  };

  const handleLogout = async () => {
    try {
      await googleLogout();
      setIsLoggedIn(false);
    } catch (error) {
      console.log("Login Failed", error);
    }
  };

  const handleLoginError = () => {
    console.log("Login Failed");
  };

  return (
    <div className="w-full min-h-screen pt-4 bg-neutral-900">
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
          {isLoggedIn && (
            <div className="w-full h-16 place-self-center flex items-center shadow-inner rounded-full bg-neutral-700">
              <input
                type="text"
                placeholder="Search Destination"
                onChange={(e) => setInput(e.target.value)}
                className="h-full bg-transparent rounded-l-full flex-grow focus:outline-none pl-8 text-lg text-neutral-200"
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
          )}

          <div className="w-full h-full flex justify-center items-center space-x-2 text-neutral-200">
            {isLoggedIn ? (
              // Render Home content for authenticated user
              <div className="w-full h-full text-center">
                <h1 className="text-4xl font-semibold bg-gradient-to-r from-cyan-500 to-blue-500 text-transparent bg-clip-text pb-1">
                  Plan your Dream Vacation!
                </h1>
                <button onClick={handleLogout}>Logout</button>
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
              className={`rounded-lg bg-neutral-700 shaddow-lg h-full w-1/5 ${filter == "attractions"
                ? "bg-gradient-to-r from-cyan-500 to-blue-500 text-neutral-100 hover:from-cyan-700 hover:to-blue-700"
                : "hover:bg-neutral-800"
                }`}
            >
              <svg
                className="stroke-neutral-400 hover:stroke-neutral-200 hover:ping w-full h-6"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                strokeWidth={1.5}
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  d="M17.218,2.268L2.477,8.388C2.13,8.535,2.164,9.05,2.542,9.134L9.33,10.67l1.535,6.787c0.083,0.377,0.602,0.415,0.745,0.065l6.123-14.74C17.866,2.46,17.539,2.134,17.218,2.268 M3.92,8.641l11.772-4.89L9.535,9.909L3.92,8.641z M11.358,16.078l-1.268-5.613l6.157-6.157L11.358,16.078z"
                />
              </svg>
              Attractions
            </button>
            <button
              onClick={(e) => {
                setFilter(`${filter == "restaurants" ? "" : "restaurants"}`);
                e.preventDefault();
              }}
              className={`rounded-lg bg-neutral-700 shaddow-lg h-full w-1/5 ${filter == "restaurants"
                ? "bg-gradient-to-r from-cyan-500 to-blue-500 text-neutral-100 hover:from-cyan-700 hover:to-blue-700"
                : "hover:bg-neutral-800"
                }`}
            >
              <svg
                className="stroke-neutral-400 hover:stroke-neutral-200 hover:ping w-full h-6"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                strokeWidth={1.5}
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  d="M5.229,6.531H4.362c-0.239,0-0.434,0.193-0.434,0.434c0,0.239,0.194,0.434,0.434,0.434h0.868c0.24,0,0.434-0.194,0.434-0.434C5.663,6.724,5.469,6.531,5.229,6.531 M10,6.531c-1.916,0-3.47,1.554-3.47,3.47c0,1.916,1.554,3.47,3.47,3.47c1.916,0,3.47-1.555,3.47-3.47C13.47,8.084,11.916,6.531,10,6.531 M11.4,11.447c-0.071,0.164-0.169,0.299-0.294,0.406c-0.124,0.109-0.27,0.191-0.437,0.248c-0.167,0.057-0.298,0.09-0.492,0.098v0.402h-0.35v-0.402c-0.21-0.004-0.352-0.039-0.527-0.1c-0.175-0.064-0.324-0.154-0.449-0.27c-0.124-0.115-0.221-0.258-0.288-0.428c-0.068-0.17-0.1-0.363-0.096-0.583h0.664c-0.004,0.259,0.052,0.464,0.169,0.613c0.116,0.15,0.259,0.229,0.527,0.236v-1.427c-0.159-0.043-0.268-0.095-0.425-0.156c-0.157-0.061-0.299-0.139-0.425-0.235C8.852,9.752,8.75,9.631,8.672,9.486C8.594,9.34,8.556,9.16,8.556,8.944c0-0.189,0.036-0.355,0.108-0.498c0.072-0.144,0.169-0.264,0.292-0.36c0.122-0.097,0.263-0.17,0.422-0.221c0.159-0.052,0.277-0.077,0.451-0.077V7.401h0.35v0.387c0.174,0,0.29,0.023,0.445,0.071c0.155,0.047,0.29,0.118,0.404,0.212c0.115,0.095,0.206,0.215,0.274,0.359c0.067,0.146,0.103,0.315,0.103,0.508H10.74c-0.007-0.201-0.06-0.354-0.154-0.46c-0.096-0.106-0.199-0.159-0.408-0.159v1.244c0.174,0.047,0.296,0.102,0.462,0.165c0.167,0.063,0.314,0.144,0.443,0.241c0.128,0.099,0.23,0.221,0.309,0.366c0.077,0.146,0.116,0.324,0.116,0.536C11.509,11.092,11.473,11.283,11.4,11.447 M18.675,4.795H1.326c-0.479,0-0.868,0.389-0.868,0.868v8.674c0,0.479,0.389,0.867,0.868,0.867h17.349c0.479,0,0.867-0.389,0.867-0.867V5.664C19.542,5.184,19.153,4.795,18.675,4.795M1.76,5.664c0.24,0,0.434,0.193,0.434,0.434C2.193,6.336,2,6.531,1.76,6.531S1.326,6.336,1.326,6.097C1.326,5.857,1.52,5.664,1.76,5.664 M1.76,14.338c-0.24,0-0.434-0.195-0.434-0.434c0-0.24,0.194-0.434,0.434-0.434s0.434,0.193,0.434,0.434C2.193,14.143,2,14.338,1.76,14.338 M18.241,14.338c-0.24,0-0.435-0.195-0.435-0.434c0-0.24,0.194-0.434,0.435-0.434c0.239,0,0.434,0.193,0.434,0.434C18.675,14.143,18.48,14.338,18.241,14.338 M18.675,12.682c-0.137-0.049-0.281-0.08-0.434-0.08c-0.719,0-1.302,0.584-1.302,1.303c0,0.152,0.031,0.297,0.08,0.434H2.981c0.048-0.137,0.08-0.281,0.08-0.434c0-0.719-0.583-1.303-1.301-1.303c-0.153,0-0.297,0.031-0.434,0.08V7.318c0.136,0.049,0.28,0.08,0.434,0.08c0.718,0,1.301-0.583,1.301-1.301c0-0.153-0.032-0.298-0.08-0.434H17.02c-0.049,0.136-0.08,0.28-0.08,0.434c0,0.718,0.583,1.301,1.302,1.301c0.152,0,0.297-0.031,0.434-0.08V12.682z M18.241,6.531c-0.24,0-0.435-0.194-0.435-0.434c0-0.24,0.194-0.434,0.435-0.434c0.239,0,0.434,0.193,0.434,0.434C18.675,6.336,18.48,6.531,18.241,6.531 M9.22,8.896c0,0.095,0.019,0.175,0.058,0.242c0.039,0.066,0.088,0.124,0.148,0.171c0.061,0.047,0.13,0.086,0.21,0.115c0.079,0.028,0.11,0.055,0.192,0.073V8.319c-0.21,0-0.322,0.044-0.437,0.132C9.277,8.54,9.22,8.688,9.22,8.896 M15.639,12.602h-0.868c-0.239,0-0.434,0.195-0.434,0.434c0,0.24,0.194,0.436,0.434,0.436h0.868c0.24,0,0.434-0.195,0.434-0.436C16.072,12.797,15.879,12.602,15.639,12.602 M10.621,10.5c-0.068-0.052-0.145-0.093-0.23-0.124c-0.086-0.031-0.123-0.06-0.212-0.082v1.374c0.209-0.016,0.332-0.076,0.465-0.186c0.134-0.107,0.201-0.281,0.201-0.516c0-0.11-0.02-0.202-0.062-0.277C10.743,10.615,10.688,10.551,10.621,10.5"
                />
              </svg>
              Restaurants
            </button>
            <button
              onClick={(e) => {
                setFilter(`${filter == "hotels" ? "" : "hotels"}`);
                e.preventDefault();
              }}
              className={`rounded-lg bg-neutral-700 shaddow-lg h-full w-1/5 ${filter == "hotels"
                ? "bg-gradient-to-r from-cyan-500 to-blue-500 text-neutral-100 hover:from-cyan-700 hover:to-blue-700"
                : "hover:bg-neutral-800"
                }`}
            >
              <svg
                className="stroke-neutral-400 hover:stroke-neutral-200 hover:ping w-full h-6"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                strokeWidth={1.5}
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  d="M18.121,9.88l-7.832-7.836c-0.155-0.158-0.428-0.155-0.584,0L1.842,9.913c-0.262,0.263-0.073,0.705,0.292,0.705h2.069v7.042c0,0.227,0.187,0.414,0.414,0.414h3.725c0.228,0,0.414-0.188,0.414-0.414v-3.313h2.483v3.313c0,0.227,0.187,0.414,0.413,0.414h3.726c0.229,0,0.414-0.188,0.414-0.414v-7.042h2.068h0.004C18.331,10.617,18.389,10.146,18.121,9.88 M14.963,17.245h-2.896v-3.313c0-0.229-0.186-0.415-0.414-0.415H8.342c-0.228,0-0.414,0.187-0.414,0.415v3.313H5.032v-6.628h9.931V17.245z M3.133,9.79l6.864-6.868l6.867,6.868H3.133z"
                />
              </svg>
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
      <div className="w-full flex justify-center pt-4">
        <ul className="w-1/2 flex flex-col space-y-2 text-neutral-200 pb-4">
          {data.map(Place)}
        </ul>
      </div>
    </div>
  );
};
export default Home;
