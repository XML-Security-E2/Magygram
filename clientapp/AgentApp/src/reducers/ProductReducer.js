import { modalConstants } from "../constants/ModalConstants";
import { productConstants } from "../constants/ProductConstants";

var prodCpy = {};

export const productReducer = (state, action) => {
	switch (action.type) {
		case productConstants.SET_PRODUCTS_REQUEST:
			return {
				...state,
				listProducts: {
					showError: false,
					errorMessage: "",
					products: [],
				},
			};
		case productConstants.SET_PRODUCTS_SUCCESS:
			return {
				...state,
				listProducts: {
					showError: false,
					errorMessage: "",
					products: action.products,
				},
			};
		case productConstants.SET_PRODUCTS_FAILURE:
			return {
				...state,
				listProducts: {
					showError: true,
					errorMessage: action.errorMessage,
					products: [],
				},
			};
		case productConstants.CREATE_PRODUCT_MODAL_HIDE_ERROR:
			prodCpy = { ...state };

			prodCpy.createProduct.showErrorMessage = false;
			prodCpy.createProduct.errorMessage = "";
			return prodCpy;

		case modalConstants.SHOW_CREATE_PRODUCT_MODAL:
			prodCpy = { ...state };
			prodCpy.createProduct.showModal = true;
			return prodCpy;

		case modalConstants.HIDE_CREATE_PRODUCT_MODAL:
			prodCpy = { ...state };
			prodCpy.createProduct.showModal = false;
			prodCpy.createProduct.imageSelected = false;
			prodCpy.createProduct.showedImage = "";
			return prodCpy;

		case productConstants.CREATE_PRODUCT_IMAGE_SELECTED:
			prodCpy = { ...state };
			prodCpy.createProduct.imageSelected = true;
			prodCpy.createProduct.showedImage = action.showedImage;
			return prodCpy;

		case productConstants.CREATE_PRODUCT_IMAGE_DESELECTED:
			prodCpy = { ...state };
			prodCpy.createProduct.imageSelected = false;
			prodCpy.createProduct.showedImage = "";
			return prodCpy;

		case productConstants.CREATE_PRODUCT_REQUEST:
			prodCpy = { ...state };
			prodCpy.createProduct.showErrorMessage = false;
			prodCpy.createProduct.errorMessage = "";
			return prodCpy;

		case productConstants.CREATE_PRODUCT_SUCCESS:
			prodCpy = { ...state };

			if (prodCpy.listProducts.products.find((post) => post.id === action.product.id) === undefined) {
				prodCpy.listProducts.products.push(action.product);
			}

			prodCpy.createProduct.showModal = false;
			prodCpy.createProduct.imageSelected = false;
			prodCpy.createProduct.showedImage = "";
			return prodCpy;

		case productConstants.CREATE_PRODUCT_FAILURE:
			prodCpy = { ...state };
			prodCpy.createProduct.showErrorMessage = true;
			prodCpy.createProduct.errorMessage = action.errorMessage;
			return prodCpy;

		case modalConstants.SHOW_EDIT_PRODUCT_MODAL:
			prodCpy = { ...state };

			prodCpy.updateProduct.showModal = true;
			return prodCpy;

		case modalConstants.HIDE_EDIT_PRODUCT_MODAL:
			prodCpy = { ...state };
			prodCpy.updateProduct.showModal = false;
			prodCpy.updateProduct.imageSelected = false;
			prodCpy.updateProduct.showedImage = "";
			return prodCpy;

		case productConstants.FIND_BY_ID_PRODUCT_REQUEST:
			return state;

		case productConstants.FIND_BY_ID_PRODUCT_SUCCESS:
			prodCpy = { ...state };

			prodCpy.updateProduct.product = action.product;
			prodCpy.updateProduct.imageSelected = true;
			prodCpy.updateProduct.showedImage = action.product.imageUrl;
			return prodCpy;

		case productConstants.FIND_BY_ID_PRODUCT_FAILURE:
			return state;

		case productConstants.EDIT_PRODUCT_IMAGE_DESELECTED:
			prodCpy = { ...state };
			prodCpy.updateProduct.imageSelected = false;
			prodCpy.updateProduct.showedImage = "";

			return prodCpy;

		case productConstants.EDIT_PRODUCT_IMAGE_SELECTED:
			prodCpy = { ...state };

			prodCpy.updateProduct.imageSelected = true;
			prodCpy.updateProduct.showedImage = action.showedImage;
			return prodCpy;

		case productConstants.EDIT_PRODUCT_MODAL_HIDE_ERROR:
			prodCpy = { ...state };

			prodCpy.updateProduct.showErrorMessage = false;
			prodCpy.updateProduct.errorMessage = "";
			return prodCpy;

		case productConstants.EDIT_PRODUCT_REQUEST:
			prodCpy = { ...state };
			prodCpy.updateProduct.showErrorMessage = false;
			prodCpy.updateProduct.errorMessage = "";
			return prodCpy;

		case productConstants.EDIT_PRODUCT_SUCCESS:
			prodCpy = { ...state };

			let prdIdx = prodCpy.listProducts.products.findIndex((post) => post.id === action.product.id);
			prodCpy.listProducts.products[prdIdx] = action.product;

			prodCpy.updateProduct.showModal = false;
			prodCpy.updateProduct.imageSelected = false;
			prodCpy.updateProduct.showedImage = "";
			return prodCpy;

		case productConstants.EDIT_PRODUCT_FAILURE:
			prodCpy = { ...state };
			prodCpy.updateProduct.showErrorMessage = true;
			prodCpy.updateProduct.errorMessage = action.errorMessage;
			return prodCpy;

		case productConstants.EDIT_PRODUCT_IMAGE_REQUEST:
			prodCpy = { ...state };
			prodCpy.updateProduct.showErrorMessage = false;
			prodCpy.updateProduct.errorMessage = "";
			return prodCpy;

		case productConstants.EDIT_PRODUCT_IMAGE_SUCCESS:
			prodCpy = { ...state };

			let prdcIdx = prodCpy.listProducts.products.findIndex((post) => post.id === action.productId);
			prodCpy.listProducts.products[prdcIdx].imageUrl = action.imageUrl;

			prodCpy.updateProduct.showModal = false;
			prodCpy.updateProduct.imageSelected = false;
			prodCpy.updateProduct.showedImage = "";
			return prodCpy;

		case productConstants.EDIT_PRODUCT_IMAGE_FAILURE:
			prodCpy = { ...state };
			prodCpy.updateProduct.showErrorMessage = true;
			prodCpy.updateProduct.errorMessage = action.errorMessage;
			return prodCpy;

		case modalConstants.SHOW_OPTIONS_MODAL:
			prodCpy = { ...state };
			prodCpy.optionsModal.showModal = true;
			prodCpy.optionsModal.productId = action.productId;
			return prodCpy;

		case modalConstants.HIDE_OPTIONS_MODAL:
			prodCpy = { ...state };
			prodCpy.optionsModal.showModal = false;
			prodCpy.optionsModal.productId = "";
			return prodCpy;

		case productConstants.DELETE_PRODUCT_REQUEST:
			return state;

		case productConstants.DELETE_PRODUCT_SUCCESS:
			let cp = state.listProducts.products.filter((post) => post.id !== action.productId);

			return {
				...state,
				listProducts: {
					showError: false,
					errorMessage: "",
					products: cp,
				},
			};

		case productConstants.DELETE_PRODUCT_FAILURE:
			return state;

		case productConstants.CREATE_CAMPAIGN_REQUEST:
			return {
				...state,
				createCampaign: {
					showError: false,
					errorMessage: "",
					showSuccessMessage: false,
					successMessage: "",
				},
			};

		case productConstants.CREATE_CAMPAIGN_SUCCESS:
			prodCpy = { ...state };
			prodCpy.createCampaign.showError = false;
			prodCpy.createCampaign.errorMessage = "";
			prodCpy.createCampaign.showSuccessMessage = true;
			prodCpy.createCampaign.successMessage = action.successMessage;

			return prodCpy;

		case productConstants.CREATE_CAMPAIGN_FAILURE:
			prodCpy = { ...state };
			prodCpy.createCampaign.showError = true;
			prodCpy.createCampaign.errorMessage = action.errorMessage;
			prodCpy.createCampaign.showSuccessMessage = false;
			prodCpy.createCampaign.successMessage = "";
			return prodCpy;

		case productConstants.SET_CAMPAIGNS_STATS_REQUEST:
			return {
				...state,
				campaigns: [],
			};

		case productConstants.SET_CAMPAIGNS_STATS_SUCCESS:
			return {
				...state,
				campaigns: action.campaigns,
			};

		case productConstants.SET_CAMPAIGNS_STATS_FAILURE:
			return state;

		default:
			return state;
	}
};
