import React, { useContext, useEffect } from "react";
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
import EditNotificationsSidebar from "./EditNotificationsSidebar";
import EditPrivacySettingsSidebar from "./EditPrivacySettingsSidebar";
import EditNotificationsForm from "./EditNotificationsForm";
import EditPrivacySettingsForm from "./EditPrivacySettingsForm";
import { userService } from "../services/UserService";
import { requestsService } from "../services/RequestsService";

const UserEditProfile = () => {
	const { userState, dispatch } = useContext(UserContext);
	const { profileSettingsState, profileSettingsDispatch } = useContext(ProfileSettingsContext);

	const handleEditProfile = () => {
		profileSettingsDispatch({ type: profileSettingsConstants.SHOW_EDIT_PROFILE_PAGE });
	};

	const handleVerifyAccount = () => {
		profileSettingsDispatch({ type: profileSettingsConstants.SHOW_VERIFY_ACCOUNT_PAGE });
	};

	const handleEditNotifications = () => {
		profileSettingsDispatch({ type: profileSettingsConstants.SHOW_EDIT_NOTIFICATIONS_PAGE });
	};

	const handleEditPrivacySettings = () => {
		profileSettingsDispatch({ type: profileSettingsConstants.SHOW_EDIT_PRIVACY_SETTINGS_PAGE });
	};

	useEffect(() => {
		userService.IsUserVerified(profileSettingsDispatch);
		requestsService.hasUserPendingRequest(profileSettingsDispatch);
	}, [profileSettingsDispatch]);

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
					<EditNotificationsSidebar show={profileSettingsState.activeSideBar.showEditNotifications} handleEditNotifications={handleEditNotifications} />
					<EditPrivacySettingsSidebar show={profileSettingsState.activeSideBar.showEditPrivacySettings} handleEditPrivacySettings={handleEditPrivacySettings} />
					<VerifyAccoundSidebar show={profileSettingsState.activeSideBar.showVerifyAccount} handleVerifyAccount={handleVerifyAccount} />
				</div>
				<div className="col-9 border-left" style={{ minHeight: "600px" }}>
					<EditProfileInfoForm show={profileSettingsState.activeSideBar.showEditProfile} />
					<EditNotificationsForm show={profileSettingsState.activeSideBar.showEditNotifications} />
					<VerifyAccontInfoForm show={profileSettingsState.activeSideBar.showVerifyAccount} />
					<EditPrivacySettingsForm show={profileSettingsState.activeSideBar.showEditPrivacySettings} />
				</div>
			</div>
		</React.Fragment>
	);
};

export default UserEditProfile;
