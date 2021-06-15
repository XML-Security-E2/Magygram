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
		editPost: {
			showModal: false,
			post: {
				id: "",
				location: "",
				tags: [],
				description: "",
				media: [],
			},
			showError: false,
			errorMessage: "",
			showSuccessMessage: false,
			successMessage: "",
		},
		postOptions: {
			showModal: false,
		},
		timeline: {
			posts: [],
		},
		guestTimeline: {
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
		postLikedBy: {
			showModal: false,
			likedBy: [],
		},
		postDislikes: {
			showModal: false,
			dislikes: [],
		},
		userProfileContent: {
			showError: false,
			errorMessage: "",
			showPosts: false,
			showCollections: false,
			showUnauthorizedErrorMessage: false,
			showCollectionPosts: false,
			selectedCollectionName: "",
			posts: [],
			collections: [],
			collectionPosts: [],
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
				Favourites: false,
			},
		},
		viewPostModalForGuest: {
			showModal: false,
			post: {
				Id: "",
				Description: "",
				Location: "",
				Media: [{}],
				UserInfo: {},
			},
		},
		userLikedPosts: null

	});

	return <PostContext.Provider value={{ postState, dispatch }}>{props.children}</PostContext.Provider>;
};

export default PostContextProvider;
