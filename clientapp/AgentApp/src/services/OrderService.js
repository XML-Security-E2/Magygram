import Axios from "axios";
import { orderConstants } from "../constants/OrderConstants";
import { authHeader } from "../helpers/auth-header";

export const orderService = {
	createOrder,
};

function createOrder(orderDTO, dispatch) {
	dispatch(request());

	if (orderDTO.items.length > 0) {
		Axios.post(`/api/orders`, orderDTO, { validateStatus: () => true, headers: authHeader() })
			.then((res) => {
				console.log(res);
				if (res.status === 201) {
					dispatch(success("Order successfully created"));
				} else {
					dispatch(failure("Error while creating order"));
				}
			})
			.catch((err) => {
				console.log(err);
				dispatch(failure("Error"));
			});
	} else {
		dispatch(failure("At least one product must be selected"));
	}
	function request() {
		return { type: orderConstants.CREATE_ORDER_REQUEST };
	}
	function success(successMessage) {
		return { type: orderConstants.CREATE_ORDER_SUCCESS, successMessage };
	}
	function failure(message) {
		return { type: orderConstants.CREATE_ORDER_FAILURE, errorMessage: message };
	}
}
