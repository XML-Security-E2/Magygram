import { adsConstants } from "../constants/AdsConstants";
import Axios from "axios";
import { authHeader } from "../helpers/auth-header";

export const adsService = {
	getStoryCampaignStatistic,
	getPostCampaignStatistic,
};

async function getStoryCampaignStatistic(dispatch) {
	dispatch(request());
	await Axios.get(`/api/ads/campaign/story/statistic`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success(res.data));
			} else {
				dispatch(failure("Error while loading collections"));
			}
		})
		.catch((err) => {
			dispatch(failure());
		});

	function request() {
		return { type: adsConstants.SET_STORY_CAMPAIGN_STATISTIC_REQUEST };
	}

	function success(data) {
		return { type: adsConstants.SET_STORY_CAMPAIGN_STATISTIC_SUCCESS, campaigns: data };
	}
	function failure(message) {
		return { type: adsConstants.SET_STORY_CAMPAIGN_STATISTIC_FAILURE, errorMessage: message };
	}
}

async function getPostCampaignStatistic(dispatch) {
	dispatch(request());
	await Axios.get(`/api/ads/campaign/post/statistic`, { validateStatus: () => true, headers: authHeader() })
		.then((res) => {
			console.log(res.data);
			if (res.status === 200) {
				dispatch(success(res.data));
			} else {
				dispatch(failure("Error while loading collections"));
			}
		})
		.catch((err) => {
			dispatch(failure());
		});

	function request() {
		return { type: adsConstants.SET_POST_CAMPAIGN_STATISTIC_REQUEST };
	}

	function success(data) {
		return { type: adsConstants.SET_POST_CAMPAIGN_STATISTIC_SUCCESS, campaigns: data };
	}
	function failure(message) {
		return { type: adsConstants.SET_POST_CAMPAIGN_STATISTIC_FAILURE, errorMessage: message };
	}
}
