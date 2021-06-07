import React, { useContext } from "react";
import { userConstants } from "../constants/UserConstants";
import { UserContext } from "../contexts/UserContext";
import EditProfileInfoForm from "./EditProfileInfoForm";
import EditProfileSidebar from "./EditProfileSidebar";
import FailureAlert from "./FailureAlert";
import SuccessAlert from "./SuccessAlert";

const UserEditProfile = () => {
	const { userState, dispatch } = useContext(UserContext);

	return (
		<React.Fragment>
			<div className="row">
				<div className="col-12">
					<SuccessAlert
						hidden={!userState.editProfile.showSuccessMessage}
						header="Success"
						message={userState.editProfile.successMessage}
						handleCloseAlert={() => dispatch({ type: userConstants.UPDATE_USER_REQUEST })}
					/>
					<FailureAlert
						hidden={!userState.editProfile.showError}
						header="Error"
						message={userState.editProfile.errorMessage}
						handleCloseAlert={() => dispatch({ type: userConstants.UPDATE_USER_REQUEST })}
					/>
				</div>
			</div>

			<div className="row border" style={{ backgroundColor: "white" }}>
				<div className="col-3">
					<EditProfileSidebar />
				</div>
				<div className="col-9 border-left">
					<EditProfileInfoForm />
				</div>
			</div>
		</React.Fragment>
	);
};

export default UserEditProfile;
