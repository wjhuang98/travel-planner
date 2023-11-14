import { Link } from "react-router-dom";

const Navbar = () => {
    return (
        <nav className="absolute w-full h-12 flex justify-center items-center shadow-sm space-x-10 px-10">
            <Link to={"/"} className="text-neutral-800 hover:text-neutral-600">Home</Link>
            <div className="flex-grow"></div>
            <Link to={"/about"} className="text-neutral-800 hover:text-neutral-600">Login</Link>
            <Link to={"/login"} className="text-neutral-800 hover:text-neutral-600">About</Link>
        </nav>
    )
}

export default Navbar;