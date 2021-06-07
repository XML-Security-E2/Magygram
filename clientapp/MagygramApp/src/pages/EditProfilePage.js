import React from "react";
import HeaderWrapper from "../components/HeaderWrapper";
import UserEditProfile from "../components/UserEditProfile";
import UserContextProvider from "../contexts/UserContext";

const EditProfilePage = () => {
	return (
		<React.Fragment>
			<HeaderWrapper />
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
