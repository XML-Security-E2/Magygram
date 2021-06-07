import React from "react";
import PostContextProvider from "../contexts/PostContext";
import UserContextProvider from "../contexts/UserContext";
import Header from "./Header";

const HeaderWrapper = () => {
	return (
		<React.Fragment>
			<PostContextProvider>
				<UserContextProvider>
					<Header />
				</UserContextProvider>
			</PostContextProvider>
		</React.Fragment>
	);
};

export default HeaderWrapper;
