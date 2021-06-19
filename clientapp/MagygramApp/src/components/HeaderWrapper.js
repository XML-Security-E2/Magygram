import React from "react";
import NotificationContextProvider from "../contexts/NotificationContext";
import PostContextProvider from "../contexts/PostContext";
import UserContextProvider from "../contexts/UserContext";
import { hasRoles } from "../helpers/auth-header";
import GuestHeader from "./GuestHeader";
import Header from "./Header";

const HeaderWrapper = () => {
	return (
		<React.Fragment>
			<PostContextProvider>
				<NotificationContextProvider>
					<UserContextProvider>{localStorage.getItem("userId") !== null || hasRoles(["admin"]) ? <Header /> : <GuestHeader />}</UserContextProvider>
				</NotificationContextProvider>
			</PostContextProvider>
		</React.Fragment>
	);
};

export default HeaderWrapper;
