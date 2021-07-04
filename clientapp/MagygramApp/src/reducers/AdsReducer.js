import { adsConstants } from "../constants/AdsConstants";

export const adsReducer = (state, action) => {
	switch (action.type) {
		case adsConstants.SET_POST_CAMPAIGN_STATISTIC_REQUEST:
			return state;
		case adsConstants.SET_POST_CAMPAIGN_STATISTIC_SUCCESS:
			return {
				...state,
				postCampaigns: action.campaigns,
			};
		case adsConstants.SET_POST_CAMPAIGN_STATISTIC_FAILURE:
			return state;

		case adsConstants.SET_STORY_CAMPAIGN_STATISTIC_REQUEST:
			return state;
		case adsConstants.SET_STORY_CAMPAIGN_STATISTIC_SUCCESS:
			return {
				...state,
				storyCampaigns: action.campaigns,
			};
		case adsConstants.SET_STORY_CAMPAIGN_STATISTIC_FAILURE:
			return state;
		default:
			return state;
	}
};
