import React, { createContext, useReducer } from "react";
import { adsReducer } from "../reducers/AdsReducer";

export const AdsContext = createContext();

const AdsContextProvider = (props) => {
	const [adsState, dispatch] = useReducer(adsReducer, {
		storyCampaigns: [],
		postCampaigns: [],
	});

	return <AdsContext.Provider value={{ adsState, dispatch }}>{props.children}</AdsContext.Provider>;
};

export default AdsContextProvider;
