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
	});

	return <StoryContext.Provider value={{ storyState, dispatch }}>{props.children}</StoryContext.Provider>;
};

export default StoryContextProvider;