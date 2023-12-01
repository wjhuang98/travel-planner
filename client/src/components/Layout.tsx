// import Navbar from "./navigation/Navbar";
import { ReactNode } from "react";

interface Props {
    children: ReactNode;
}

const Layout = ({children}: Props) => {
    return (
        <div className="h-screen w-screen flex flex-wrap justify-center">
            {/* <Navbar /> */}
            {children}
        </div>
    )
}

export default Layout;