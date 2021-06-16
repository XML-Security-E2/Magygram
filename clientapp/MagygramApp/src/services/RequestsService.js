import Axios from "axios";
import { profileSettingsConstants } from "../constants/ProfileSettingsConstants";
import { authHeader } from "../helpers/auth-header";

export const requestsService = {
	createVerificationRequest
};

function createVerificationRequest(formData,dispatch){
    dispatch(request());

	Axios.post(`/api/requests`, formData, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			if (res.status === 200) {
				dispatch(success())
			} else {
				dispatch(failure("Sorry, we have some internal problem."));
			}
		})
		.catch((err) => console.error(err));

    function request() {
		return { type: profileSettingsConstants.CREATE_VERIFICATION_REQUEST_REQUEST };
	}
	function success() {
		return { type: profileSettingsConstants.CREATE_VERIFICATION_REQUEST_SUCCESS };
	}
	function failure(error) {
		return { type: profileSettingsConstants.CREATE_VERIFICATION_REQUEST_FAILURE, error };
	}
}

