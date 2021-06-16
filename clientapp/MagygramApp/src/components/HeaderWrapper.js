import React from "react";
import NotificationContextProvider from "../contexts/NotificationContext";
import PostContextProvider from "../contexts/PostContext";
import UserContextProvider from "../contexts/UserContext";
import GuestHeader from "./GuestHeader";
import Header from "./Header";

const HeaderWrapper = () => {
	return (
		<React.Fragment>
			<PostContextProvider>
				<NotificationContextProvider>
					<UserContextProvider>{localStorage.getItem("userId") !== null ? <Header /> : <GuestHeader />}</UserContextProvider>
				</NotificationContextProvider>
			</PostContextProvider>
		</React.Fragment>
	);
};

export default HeaderWrapper;
