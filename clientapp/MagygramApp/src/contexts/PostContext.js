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
		timeline: {
			posts: []
		},
		loggedUserInfo: {
			Id: "063d2368-9035-4184-8b0f-bdb255e6a492",
			ImageURL: "",
			Username: "nikolakolovic",
		},
		postLikedBy:{
			showModal: false,
			likedBy: []
		},
		postDislikes:{
			showModal: false,
			dislikes: []
		}
	});

	return <PostContext.Provider value={{ postState, dispatch }}>{props.children}</PostContext.Provider>;
};

export default PostContextProvider;
