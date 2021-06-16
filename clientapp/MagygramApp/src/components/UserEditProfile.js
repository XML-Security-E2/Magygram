import React, { useContext } from "react";
import { userConstants } from "../constants/UserConstants";
import { ProfileSettingsContext } from "../contexts/ProfileSettingsContext";
import { UserContext } from "../contexts/UserContext";
import EditProfileInfoForm from "./EditProfileInfoForm";
import EditProfileSidebar from "./EditProfileSidebar";
import FailureAlert from "./FailureAlert";
import SuccessAlert from "./SuccessAlert";
import VerifyAccoundSidebar from "./VerifyAccountSidebar";
import { profileSettingsConstants } from "../constants/ProfileSettingsConstants";
import VerifyAccontInfoForm from "./VerifyAccountInfoForm";

const UserEditProfile = () => {
	const { userState, dispatch } = useContext(UserContext);
	const { profileSettingsState, profileSettingsDispatch } = useContext(ProfileSettingsContext);

	const handleEditProfile= () =>{
		profileSettingsDispatch({ type: profileSettingsConstants.SHOW_EDIT_PROFILE_PAGE })
	}

	const handleVerifyAccount =()=>{
		profileSettingsDispatch({ type: profileSettingsConstants.SHOW_VERIFY_ACCOUNT_PAGE })
	}

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
					<EditProfileSidebar show={profileSettingsState.activeSideBar.showEditProfile} handleEditProfile={handleEditProfile} />
					<VerifyAccoundSidebar show={profileSettingsState.activeSideBar.showVerifyAccount} handleVerifyAccount={handleVerifyAccount}/>
				</div>
				<div className="col-9 border-left" style={{minHeight: "600px"}}>
					<EditProfileInfoForm show={profileSettingsState.activeSideBar.showEditProfile} />
					<VerifyAccontInfoForm show={profileSettingsState.activeSideBar.showVerifyAccount}/ >
				</div>
			</div>
		</React.Fragment>
	);
};

export default UserEditProfile;
