import Axios from "axios";
import { profileSettingsConstants } from "../constants/ProfileSettingsConstants";
import { authHeader } from "../helpers/auth-header";
import { adminConstants } from "../constants/AdminConstants";

export const requestsService = {
	createVerificationRequest,
	getAllPendingVerificationRequest,
	approveVerificationRequest,
	rejectVerificationRequest,
	hasUserPendingRequest,
};

function createVerificationRequest(formData,dispatch){
    dispatch(request());

	Axios.post(`/api/requests/verification`, formData, { validateStatus: () => true, headers: authHeader() })
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

async function getAllPendingVerificationRequest(dispatch) {
	dispatch(request());
	await Axios.get(`/api/requests/verification`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success(res.data));
			} else {
				failure();
			}
		})
		.catch((err) => {
			failure();
		});

	function request() {
		return { type: adminConstants.GET_PENDING_VERIFICATION_REQUEST_REQUEST };
	}

	function success(data) {
		return { type: adminConstants.GET_PENDING_VERIFICATION_REQUEST_SUCCESS, requests: data };
	}
	function failure() {
		return { type: adminConstants.GET_PENDING_VERIFICATION_REQUEST_FAILURE };
	}
}

function approveVerificationRequest(requestId, dispatch) {
	dispatch(request());

	Axios.put(`/api/requests/verification/${requestId}/approve`, {}, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res);
			if (res.status === 200) {
				dispatch(success(requestId,"Verification request has been approved"));
			} else {
				dispatch(failure(res.data.message));
			}
		})
		.catch((err) => {
			dispatch(failure("Error"));
		});

	function request() {
		return { type: adminConstants.APPROVE_VERIFICATION_REQUEST_REQUEST };
	}
	function success(requestId,successMessage) {
		return { type: adminConstants.APPROVE_VERIFICATION_REQUEST_SUCCESS, requestId,successMessage };
	}
	function failure(message) {
		return { type: adminConstants.APPROVE_VERIFICATION_REQUEST_FAILURE, errorMessage: message };
	}
}

function rejectVerificationRequest(requestId, dispatch) {
	dispatch(request());

	Axios.put(`/api/requests/verification/${requestId}/reject`, {}, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res);
			if (res.status === 200) {
				dispatch(success(requestId,"Verification request has been rejected"));
			} else {
				dispatch(failure(res.data.message));
			}
		})
		.catch((err) => {
			dispatch(failure("Error"));
		});

	function request() {
		return { type: adminConstants.REJECT_VERIFICATION_REQUEST_REQUEST };
	}
	function success(requestId,successMessage) {
		return { type: adminConstants.REJECT_VERIFICATION_REQUEST_SUCCESS, requestId,successMessage };
	}
	function failure(message) {
		return { type: adminConstants.REJECT_VERIFICATION_REQUEST_FAILURE, errorMessage: message };
	}
}

function hasUserPendingRequest(dispatch) {
	dispatch(request());

	Axios.get(`/api/requests/verification/has-pending-request`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			if (res.status === 200) {
				dispatch(success(res.data));
			} else {
				dispatch(failure("Usled internog problema trenutno nije moguce logovanje"));
			}
		})
		.catch((err) => console.error(err));

	function request() {
		return { type: profileSettingsConstants.CHECK_IF_USER_HAS_PENDING_REQUEST_REQUEST };
	}
	function success(result) {
		return { type: profileSettingsConstants.CHECK_IF_USER_HAS_PENDING_REQUEST_SUCCESS, result };
	}
	function failure(error) {
		return { type: profileSettingsConstants.CHECK_IF_USER_HAS_PENDING_REQUEST_FAILURE, error };
	}
}