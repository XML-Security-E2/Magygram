import React from "react";
import PostContextProvider from "../contexts/PostContext";
import UserContextProvider from "../contexts/UserContext";
import GuestHeader from "./GuestHeader";
import Header from "./Header";

const HeaderWrapper = () => {
	return (
		<React.Fragment>
			<PostContextProvider>
				<UserContextProvider>{localStorage.getItem("userId") !== null ? <Header /> : <GuestHeader />}</UserContextProvider>
			</PostContextProvider>
		</React.Fragment>
	);
};

export default HeaderWrapper;
