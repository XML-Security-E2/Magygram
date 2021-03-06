import React, { createContext, useReducer } from "react";
import { profileSettingsReducer } from "../reducers/ProfileSettings";

export const ProfileSettingsContext = createContext();

const ProfileSettingsContextProvider = (props) => {
	const [profileSettingsState, profileSettingsDispatch] = useReducer(profileSettingsReducer, {
		activeSideBar: {
			showEditProfile: true,
			showVerifyAccount: false,
			showEditNotifications: false,
			showEditPrivacySettings: false,
			showInfluencerCampagn: false,
		},
		sendedVerifyRequest: false,
		sendRequest: {
			showError: false,
			errorMessage: "",
		},
		isUserVerified: false,
	});

	return <ProfileSettingsContext.Provider value={{ profileSettingsState, profileSettingsDispatch }}>{props.children}</ProfileSettingsContext.Provider>;
};

export default ProfileSettingsContextProvider;
