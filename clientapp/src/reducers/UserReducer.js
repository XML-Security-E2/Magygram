import { userConstants } from "../constants/UserConstants";
import { userService } from "../services/UserService";



export const userReducer = (state, action) => {
	switch (action.type) {
		case userConstants.REGISTER_REQUEST:
			return {
				registrationError : { 
					showError: false,
					errorMessage: '' } 
			};
		case userConstants.REGISTER_SUCCESS:
			return {
				registrationError : { 
					showError: false,
					errorMessage: '' } 
			};
		case userConstants.REGISTER_FAILURE:
			return {
				registrationError : { 
					showError: true,
					errorMessage: action.errorMessage } 
			};
		case userConstants.REGISTER_VALIDATION_FAILURE:
			return {
				registrationError : { 
					showError: true,
					errorMessage: action.errorMessage } 
			};
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
			return {
				forgotPasswordLinkError : {
					showError: false,
					errorMessage: '',
					showSuccessMessage: false,
					emailAddress:''
				}
			}
		case userConstants.RESET_PASSWORD_LINK_SUCCESS:
			return {
				forgotPasswordLinkError : {
					showError: false,
					errorMessage: '',
					showSuccessMessage: true,
					emailAddress: action.emailAddress
				}
			}
		case userConstants.RESET_PASSWORD_LINK_FAILURE:
			return {
				forgotPasswordLinkError : {
					showError: true,
					errorMessage: action.errorMessage,
					showSuccessMessage: false,
					emailAddress:''
				}
			}

		case userConstants.RESET_PASSWORD_REQUEST:
			userService.resetPasswordRequest(action.resetPasswordRequest);
			return state;
		case userConstants.RESEND_ACTIVATION_LINK_REQUEST:
			userService.resendActivationLink(action.resendActivationLink);
			return state;
		case userConstants.BLOCKED_USER_EMAIL_REQUEST:
			return {
				blockedUser : {
					emailAddress: action.emailAddress
				}
			}
		default:
			return state;
	}
};
