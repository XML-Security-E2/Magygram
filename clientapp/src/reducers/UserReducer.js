import { userConstants } from "../constants/UserConstants";
import { userService } from "../services/UserService";



export const userReducer = (state, action) => {
	switch (action.type) {
		case userConstants.REGISTER_REQUEST:
			userService.register(action.user);
			return state;
		case userConstants.REGISTER_SUCCESS:
			return state;
		case userConstants.REGISTER_FAILURE:
			return state;
		case userConstants.LOGIN_REQUEST:
		    return {
				loginError : { 
					showError: false,
					errorMessage: '' } 
			};
		case userConstants.LOGIN_FAILURE:
			return {
				loginError : { 
					showError: true, 
					errorMessage: 'Sorry, your email or password was incorrect. Please double-check your password.' 
				} 
			};
		case userConstants.LOGIN_SUCCESS:
			return {
				loginError : { 
					showError: false, 
					errorMessage: '' 
				} 
			};
		case userConstants.RESET_PASSWORD_LINK_REQUEST:
			userService.resetPasswordLinkRequest(action.resetPasswordLinkRequest);
			return state;
		case userConstants.RESET_PASSWORD_REQUEST:
			userService.resetPasswordRequest(action.resetPasswordRequest);
			return state;
		case userConstants.RESEND_ACTIVATION_LINK_REQUEST:
			userService.resendActivationLink(action.resendActivationLink);
			return state;
		default:
			return state;
	}
};
