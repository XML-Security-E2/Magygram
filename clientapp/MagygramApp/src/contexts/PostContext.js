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
			posts: [],
		},
		addToFavouritesModal: {
			renderCollectionSwitch: false,
			showModal: false,
			selectedPostId: "",
		},
		userCollections: {
			showError: false,
			errorMessage: "",
			showSuccessMessage: false,
			successMessage: "",
			collections: [],
		},
		loggedUserInfo: {
			Id: "063d2368-9035-4184-8b0f-bdb255e6a492",
			ImageURL: "assets/images/profiles/profile-1.jpg",
			Username: "nikolakolovic",
		},
		postLikedBy: {
			showModal: false,
			likedBy: [],
		},
		postDislikes: {
			showModal: false,
			dislikes: [],
		},
		viewPostModal: {
			showModal: false,
			post: {
				Id: "",
				Description: "",
				Location: "",
				ContentType: "",
				Tags: null,
				HashTags: null,
				Media: [{}],
				UserInfo: {},
				LikedBy: [{}],
				DislikedBy: [{}],
				Comments: [{}],
				Liked: false,
				Disliked: false,
			},
		},
	});

	return <PostContext.Provider value={{ postState, dispatch }}>{props.children}</PostContext.Provider>;
};

export default PostContextProvider;
