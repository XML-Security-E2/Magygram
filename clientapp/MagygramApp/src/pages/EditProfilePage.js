import React from "react";
import Header from "../components/Header";
import UserEditProfile from "../components/UserEditProfile";
import PostContextProvider from "../contexts/PostContext";
import UserContextProvider from "../contexts/UserContext";

const EditProfilePage = () => {
	return (
		<React.Fragment>
			<PostContextProvider>
				<Header />
			</PostContextProvider>
			<div className="mt-4">
				<div className="container d-flex justify-content-center ">
					<div className="col-9">
						<UserContextProvider>
							<UserEditProfile />
						</UserContextProvider>
					</div>
				</div>
			</div>
		</React.Fragment>
	);
};

export default EditProfilePage;
