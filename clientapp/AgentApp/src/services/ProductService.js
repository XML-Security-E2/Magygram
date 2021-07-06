import Axios from "axios";
import { productConstants } from "../constants/ProductConstants";
import { authHeader } from "../helpers/auth-header";

export const productService = {
	findAllProducts,
	createProduct,
	findById,
	updateProductInfo,
	updateProductImage,
	deleteProduct,
	createCampaign,
	getCampaignStatistics,
	getGeneratedCampaignStatisticsReports,
};

async function getGeneratedCampaignStatisticsReports(dispatch) {
	dispatch(request());

	await Axios.get(`/api/products/campaign/reports`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res);
			if (res.status === 200) {
				dispatch(success(res.data));
			} else {
				dispatch(failure("Error while fetching data"));
			}
		})
		.catch((err) => {
			console.log(err);
			dispatch(failure("Error"));
		});

	function request() {
		return { type: productConstants.SET_CAMPAIGN_REPORTS_REQUEST };
	}
	function success(data) {
		return { type: productConstants.SET_CAMPAIGN_REPORTS_SUCCESS, reports: data };
	}
	function failure(message) {
		return { type: productConstants.SET_CAMPAIGN_REPORTS_FAILURE, errorMessage: message };
	}
}

async function findAllProducts(dispatch) {
	dispatch(request());

	await Axios.get(`/api/products`, { validateStatus: () => true })
		.then((res) => {
			console.log(res);
			if (res.status === 200) {
				dispatch(success(res.data));
			} else {
				dispatch(failure("Error while fetching data"));
			}
		})
		.catch((err) => {
			console.log(err);
			dispatch(failure("Error"));
		});

	function request() {
		return { type: productConstants.SET_PRODUCTS_REQUEST };
	}
	function success(data) {
		return { type: productConstants.SET_PRODUCTS_SUCCESS, products: data };
	}
	function failure(message) {
		return { type: productConstants.SET_PRODUCTS_FAILURE, errorMessage: message };
	}
}

async function getCampaignStatistics(dispatch) {
	dispatch(request());

	await Axios.get(`/api/products/campaign/statistics`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res);
			if (res.status === 200) {
				dispatch(success(res.data));
			} else {
				dispatch(failure());
			}
		})
		.catch((err) => {
			console.log(err);
			dispatch(failure());
		});

	function request() {
		return { type: productConstants.SET_CAMPAIGNS_STATS_REQUEST };
	}
	function success(data) {
		return { type: productConstants.SET_CAMPAIGNS_STATS_SUCCESS, report: data };
	}
	function failure() {
		return { type: productConstants.SET_CAMPAIGNS_STATS_FAILURE };
	}
}

function createCampaign(campaignDTO, dispatch) {
	dispatch(request());

	Axios.post(`/api/products/campaign`, campaignDTO, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res);
			if (res.status === 201) {
				dispatch(success("Campaign successfully created"));
			} else {
				dispatch(failure("Error while creating campaign"));
			}
		})
		.catch((err) => {
			dispatch(failure("Error"));
		});

	function request() {
		return { type: productConstants.CREATE_CAMPAIGN_REQUEST };
	}
	function success(message) {
		return { type: productConstants.CREATE_CAMPAIGN_SUCCESS, successMessage: message };
	}
	function failure(message) {
		return { type: productConstants.CREATE_CAMPAIGN_FAILURE, errorMessage: message };
	}
}

function createProduct(productDTO, dispatch) {
	let formData = new FormData();
	formData.append("image", productDTO.image);
	formData.append("name", productDTO.name);
	formData.append("price", productDTO.price);

	dispatch(request());

	Axios.post(`/api/products`, formData, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res);
			if (res.status === 201) {
				dispatch(success(res.data));
			} else {
				dispatch(failure("Error while creating product"));
			}
		})
		.catch((err) => {
			console.log(err);
			dispatch(failure("Error"));
		});

	function request() {
		return { type: productConstants.CREATE_PRODUCT_REQUEST };
	}
	function success(data) {
		return { type: productConstants.CREATE_PRODUCT_SUCCESS, product: data };
	}
	function failure(message) {
		return { type: productConstants.CREATE_PRODUCT_FAILURE, errorMessage: message };
	}
}

function updateProductInfo(productDTO, productId, dispatch) {
	dispatch(request());

	Axios.put(`/api/products/${productId}`, productDTO, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res);
			if (res.status === 200) {
				dispatch(success(res.data));
			} else {
				dispatch(failure("Error while editing product"));
			}
		})
		.catch((err) => {
			console.log(err);
			dispatch(failure("Error"));
		});

	function request() {
		return { type: productConstants.EDIT_PRODUCT_REQUEST };
	}
	function success(data) {
		return { type: productConstants.EDIT_PRODUCT_SUCCESS, product: data };
	}
	function failure(message) {
		return { type: productConstants.EDIT_PRODUCT_FAILURE, errorMessage: message };
	}
}

function updateProductImage(image, productId, dispatch) {
	let formData = new FormData();
	formData.append("image", image, "img");

	dispatch(request());
	Axios.put(`/api/products/${productId}/image`, formData, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res);
			if (res.status === 200) {
				dispatch(success(res.data, productId));
			} else {
				dispatch(failure("Error while editing product's image"));
			}
		})
		.catch((err) => {
			console.log(err);
			dispatch(failure("Error"));
		});

	function request() {
		return { type: productConstants.EDIT_PRODUCT_IMAGE_REQUEST };
	}
	function success(data, productId) {
		return { type: productConstants.EDIT_PRODUCT_IMAGE_SUCCESS, imageUrl: data, productId };
	}
	function failure(message) {
		return { type: productConstants.EDIT_PRODUCT_IMAGE_FAILURE, errorMessage: message };
	}
}

function deleteProduct(productId, dispatch) {
	dispatch(request());

	Axios.delete(`/api/products/${productId}`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res);
			if (res.status === 200) {
				dispatch(success(productId));
			} else {
				dispatch(failure("Error while editing product's image"));
			}
		})
		.catch((err) => {
			console.log(err);
			dispatch(failure("Error"));
		});

	function request() {
		return { type: productConstants.DELETE_PRODUCT_REQUEST };
	}
	function success(productId) {
		return { type: productConstants.DELETE_PRODUCT_SUCCESS, productId };
	}
	function failure(message) {
		return { type: productConstants.DELETE_PRODUCT_FAILURE, errorMessage: message };
	}
}

async function findById(id, dispatch) {
	dispatch(request());

	await Axios.get(`/api/products/${id}`, { validateStatus: () => true })
		.then((res) => {
			console.log(res);
			if (res.status === 200) {
				dispatch(success(res.data));
			} else {
				dispatch(failure("Error while fetching data"));
			}
		})
		.catch((err) => {
			console.log(err);
			dispatch(failure("Error"));
		});

	function request() {
		return { type: productConstants.FIND_BY_ID_PRODUCT_REQUEST };
	}
	function success(data) {
		return { type: productConstants.FIND_BY_ID_PRODUCT_SUCCESS, product: data };
	}
	function failure(message) {
		return { type: productConstants.FIND_BY_ID_PRODUCT_FAILURE, errorMessage: message };
	}
}
