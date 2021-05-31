import React, { createContext, useReducer } from "react";
import { postReducer } from "../reducers/PostReducer";

export const PostContext = createContext();

const PostContextProvider = (props) => {
	const [postState, dispatch] = useReducer(postReducer, {
		createPost: {
			showError: false,
			errorMessage: "",
			showSuccessMessage: false,
			successMessage: "",
		},
	});

	return <PostContext.Provider value={{ postState, dispatch }}>{props.children}</PostContext.Provider>;
};

export default PostContextProvider;
