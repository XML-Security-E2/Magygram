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
		createAgentPost: {
			showError: false,
			errorMessage: "",
			showSuccessMessage: false,
			successMessage: "",
		},
		searchInfluencer: {
			post: {
				id: "",
				userId: "",
				location: "",
				tags: [],
				description: "",
				media: [],
			},
		},
		editPost: {
			showModal: false,
			post: {
				id: "",
				userId: "",
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
		campaignOptions: {
			showModal: false,
		},
		postOptions: {
			showModal: false,
		},
		timeline: {
			posts: [],
		},
		postDetailsPage: {
			post: {
				Id: "",
				Description: "",
				Location: "",
				ContentType: "",
				Tags: null,
				HashTags: null,
				Media: [],
				UserInfo: {},
				LikedBy: [],
				DislikedBy: [],
				Comments: [],
				Liked: false,
				Disliked: false,
				Favourites: false,
			},
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
		viewAgentCampaignPostModal: {
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
		agentCampaignPostOptionModal: {
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
		agentCampaignPosts: [],
		userLikedPosts: null,
		userDislikedPosts: null,
		postReport: {
			showError: false,
			errorMessage: "",
			showSuccessMessage: false,
			successMessage: "",
		},
	});

	return <PostContext.Provider value={{ postState, dispatch }}>{props.children}</PostContext.Provider>;
};

export default PostContextProvider;
