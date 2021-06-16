import { modalConstants } from "../constants/ModalConstants";
import { orderConstants } from "../constants/OrderConstants";

var ordCpy = {};

export const orderReducer = (state, action) => {
	switch (action.type) {
		case orderConstants.ADD_PRODUCT_TO_ORDER:
			ordCpy = { ...state };
			if (ordCpy.shoppingCart.items.find((item) => item.id === action.item.id) === undefined) {
				ordCpy.shoppingCart.items.push(action.item);
			}
			return ordCpy;

		case orderConstants.SET_PRODUCT_COUNT_TO_ORDER:
			ordCpy = { ...state };
			let prdIdx = ordCpy.shoppingCart.items.findIndex((item) => item.id === action.id);
			ordCpy.shoppingCart.items[prdIdx].count = action.count;
			return ordCpy;

		case orderConstants.REMOVE_PRODUCT_FROM_ORDER:
			ordCpy = { ...state };
			ordCpy.shoppingCart.items = ordCpy.shoppingCart.items.filter((item) => item.id !== action.id);
			return ordCpy;

		case modalConstants.SHOW_ORDER_CHECKOUT_MODAL:
			return {
				...state,
				orderCheckout: {
					showModal: true,
					showErrorMessage: false,
					errorMessage: "",
					showSuccessMessage: false,
					successMessage: "",
				},
			};

		case modalConstants.HIDE_ORDER_CHECKOUT_MODAL:
			return {
				...state,
				orderCheckout: {
					showModal: false,
					showErrorMessage: false,
					errorMessage: "",
					showSuccessMessage: false,
					successMessage: "",
				},
			};

		case orderConstants.CHECKOUT_MODAL_HIDE_MESSAGE:
			ordCpy = { ...state };
			ordCpy.orderCheckout.showErrorMessage = false;
			ordCpy.orderCheckout.errorMessage = "";
			ordCpy.orderCheckout.showSuccessMessage = false;
			ordCpy.orderCheckout.successMessage = "";

			return ordCpy;

		case orderConstants.CREATE_ORDER_REQUEST:
			ordCpy = { ...state };
			ordCpy.orderCheckout.showErrorMessage = false;
			ordCpy.orderCheckout.errorMessage = "";
			ordCpy.orderCheckout.showSuccessMessage = false;
			ordCpy.orderCheckout.successMessage = "";

			return ordCpy;

		case orderConstants.CREATE_ORDER_SUCCESS:
			ordCpy = { ...state };
			ordCpy.orderCheckout.showErrorMessage = false;
			ordCpy.orderCheckout.errorMessage = "";
			ordCpy.orderCheckout.showSuccessMessage = true;
			ordCpy.orderCheckout.successMessage = action.successMessage;

			return ordCpy;

		case orderConstants.CREATE_ORDER_FAILURE:
			ordCpy = { ...state };
			ordCpy.orderCheckout.showErrorMessage = true;
			ordCpy.orderCheckout.errorMessage = action.errorMessage;
			ordCpy.orderCheckout.showSuccessMessage = false;
			ordCpy.orderCheckout.successMessage = "";

			return ordCpy;
		default:
			return state;
	}
};
