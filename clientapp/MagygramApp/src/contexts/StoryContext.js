import React, { createContext, useReducer } from "react";
import { storyReducer } from "../reducers/StoryReducer";

export const StoryContext = createContext();

const StoryContextProvider = (props) => {
	const [storyState, dispatch] = useReducer(storyReducer, {
		createStory: {
			showModal: false,
			showError: false,
			errorMessage: "",
			showSuccessMessage: false,
			successMessage: "",
		},
		createAgentStory: {
			showModal: false,
			showError: false,
			errorMessage: "",
			showSuccessMessage: false,
			successMessage: "",
		},
		storyline: {
			stories: [],
		},
		userInfo: {},
		postOptions: {
			showModal: false,
		},
		searchInfluencer: {
			storyId: "", 
		},
		campaignOptions: {
			showModal: false,
		},
		highlights: {
			showModal: false,
			showError: false,
			errorMessage: "",
			showHighlightsName: false,
			stories: [],
		},
		highlightsSliderModal: {
			showModal: false,
			highlights: [],
		},
		profileHighlights: {
			highlights: [],
		},
		storySliderModal: {
			showModal: false,
			stories: [],
			firstUnvisitedStory: 0,
			visited: false,
			userId: "",
		},
		agentCampaignStoryModal: {
			showModal: false,
			stories: "",
			storyId: "",
		},
		agentCampaignStoryOptionModal: {
			showModal: false,
			showError: false,
			errorMessage: "",
			showSuccessMessage: false,
			successMessage: "",
			campaign: {
				minAge: "",
				maxAge: "",
				minDisplays: "",
				gender: "ANY",
				frequency: "",
				startDate: new Date(),
				endDate: new Date(new Date().getTime() + 24 * 60 * 60 * 1000),
			},
		},
		agentCampaignStories: [],
		storyId: "",
		iHaveAStory: false,
		storyReport: {
			showError: false,
			errorMessage: "",
			showSuccessMessage: false,
			successMessage: "",
		},
	});

	return <StoryContext.Provider value={{ storyState, dispatch }}>{props.children}</StoryContext.Provider>;
};

export default StoryContextProvider;
