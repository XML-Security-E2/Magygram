import { userConstants } from "../constants/UserConstants";
import { userService } from "../services/UserService";

export const userReducer = (state, action) => {
	switch (action.type) {
		case userConstants.REGISTER_REQUEST:
			userService.register(action.user);
			return state;
		case userConstants.LOGIN_REQUEST:
			userService.login(action.loginRequest);
			return state;
		default:
			return state;
	}
};
