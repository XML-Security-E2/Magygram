import { modalConstants } from "../constants/ModalConstants";
import { storyConstants } from "../constants/StoryConstants";

let storyCopy = {};
export const storyReducer = (state, action) => {
	switch (action.type) {
		case storyConstants.CREATE_STORY_REQUEST:
			return {
				...state,
				createStory: {
					showModal: false,
					showError: false,
					errorMessage: "",
					showSuccessMessage: false,
					successMessage: "",
				},
			};
		case storyConstants.CREATE_STORY_SUCCESS:
			return {
				...state,
				createStory: {
					showModal: false,
					showError: false,
					errorMessage: "",
					showSuccessMessage: true,
					successMessage: action.successMessage,
				},
			};
		case storyConstants.CREATE_STORY_FAILURE:
			return {
				...state,
				createStory: {
					showModal: false,
					showError: true,
					errorMessage: action.errorMessage,
					showSuccessMessage: false,
					successMessage: "",
				},
			};
		case modalConstants.OPEN_CREATE_STORY_MODAL:
			console.log(state);
			return {
				...state,
				createStory: {
					showModal: true,
					showError: false,
					errorMessage: "",
					showSuccessMessage: false,
					successMessage: "",
				},
			};
		case storyConstants.STORYLINE_STORY_REQUEST:
			return {
				...state,
				storyline: {
					stories: [],
				},
			};
		case modalConstants.SHOW_STORY_OPTIONS_MODAL:
			return {
				...state,
				editPost: {
					showModal: false,
					showError: false,
					errorMessage: "",
					showSuccessMessage: false,
					successMessage: "",
				},
				stories: action.stories,
				postOptions: {
					showModal: true,
				},
			};
		case modalConstants.HIDE_STORY_OPTIONS_MODAL:
			return {
				...state,
				editPost: {
					showModal: false,
					showError: false,
					errorMessage: "",
					showSuccessMessage: false,
					successMessage: "",
				},
				postOptions: {
					showModal: false,
				},
			};
		case storyConstants.STORYLINE_STORY_SUCCESS:
			return {
				...state,
				storyline: {
					stories: action.stories,
				},
			};
		case storyConstants.STORYLINE_STORY_FAILURE:
			return {
				...state,
				storyline: {
					stories: [],
				},
			};
		case modalConstants.SHOW_STORY_SLIDER_MODAL:
			return {
				...state,
				storySliderModal: {
					showModal: true,
					stories: createStories(action.stories),
					firstUnvisitedStory: action.stories.FirstUnvisitedStory,
					visited: action.visited,
					userId: action.userId,
				},
			};
		case modalConstants.HIDE_STORY_SLIDER_MODAL:
			return {
				...state,
				storySliderModal: {
					showModal: false,
					stories: action.stories,
					firstUnvisitedStory: 0,
				},
			};
		case storyConstants.PROFILE_HIGHLIGHTS_REQUEST:
			return {
				...state,
				profileHighlights: {
					highlights: [],
				},
			};
		case storyConstants.PROFILE_HIGHLIGHTS_SUCCESS:
			return {
				...state,
				profileHighlights: {
					highlights: action.highlights,
				},
			};
		case storyConstants.PROFILE_HIGHLIGHTS_FAILURE:
			return {
				...state,
				profileHighlights: {
					highlights: [],
				},
			};
		case storyConstants.USER_HIGHLIGHTS_STORY_REQUEST:
			storyCopy = { ...state };
			storyCopy.highlights.stories = [];
			return storyCopy;
		case storyConstants.USER_HIGHLIGHTS_STORY_SUCCESS:
			storyCopy = { ...state };
			storyCopy.highlights.stories = action.stories;
			return storyCopy;

		case storyConstants.USER_HIGHLIGHTS_STORY_FAILURE:
			storyCopy = { ...state };
			storyCopy.highlights.stories = [];
			return storyCopy;

		case modalConstants.SHOW_STORY_SELECT_HIGHLIGHTS_MODAL:
			storyCopy = { ...state };
			storyCopy.highlights.showModal = true;
			return storyCopy;

		case modalConstants.HIDE_STORY_SELECT_HIGHLIGHTS_MODAL:
			storyCopy = { ...state };
			storyCopy.highlights.showModal = false;
			storyCopy.highlights.showError = false;
			storyCopy.highlights.errorMessage = "";
			storyCopy.highlights.showHighlightsName = false;
			return storyCopy;

		case storyConstants.SHOW_HIGHLIGHTS_NAME_INPUT:
			storyCopy = { ...state };
			storyCopy.highlights.showHighlightsName = true;
			return storyCopy;

		case storyConstants.HIDE_HIGHLIGHTS_NAME_INPUT:
			storyCopy = { ...state };
			storyCopy.highlights.showHighlightsName = false;
			return storyCopy;

		case storyConstants.SHOW_HIGHLIGHTS_MODAL_ERROR_MESSAGE:
			storyCopy = { ...state };

			storyCopy.highlights.showError = true;
			storyCopy.highlights.errorMessage = action.errorMessage;
			return storyCopy;

		case storyConstants.HIDE_HIGHLIGHTS_MODAL_ERROR_MESSAGE:
			storyCopy = { ...state };
			storyCopy.highlights.showError = false;
			storyCopy.highlights.errorMessage = "";

			return storyCopy;

		case storyConstants.CREATE_HIGHLIGHTS_STORY_REQUEST:
			storyCopy = { ...state };
			storyCopy.highlights.showError = false;
			storyCopy.highlights.errorMessage = "";
			return storyCopy;

		case storyConstants.CREATE_HIGHLIGHTS_STORY_SUCCESS:
			return {
				...state,
				highlights: {
					showModal: false,
					showError: false,
					errorMessage: "",
					showHighlightsName: false,
					stories: [...state.highlights.stories],
				},
				profileHighlights: {
					highlights: [...state.profileHighlights.highlights, action.highlight],
				},
			};

		case storyConstants.CREATE_HIGHLIGHTS_STORY_FAILURE:
			storyCopy = { ...state };
			storyCopy.highlights.showError = true;
			storyCopy.highlights.errorMessage = action.errorMessage;

			return storyCopy;

		case storyConstants.FIND_HIGHLIGHT_BY_NAME_REQUEST:
			storyCopy = { ...state };
			storyCopy.highlightsSliderModal.showModal = false;
			storyCopy.highlightsSliderModal.highlights = [];
			return storyCopy;

		case storyConstants.FIND_HIGHLIGHT_BY_NAME_SUCCESS:
			storyCopy = { ...state };

			storyCopy.highlightsSliderModal.showModal = true;
			storyCopy.highlightsSliderModal.highlights = createHighlights(action.highlights, action.name);
			return storyCopy;

		case storyConstants.FIND_HIGHLIGHT_BY_NAME_FAILURE:
			storyCopy = { ...state };
			storyCopy.highlightsSliderModal.showModal = false;
			storyCopy.highlightsSliderModal.highlights = [];

			return storyCopy;

		case modalConstants.HIDE_STORY_SLIDER_HIGHLIGHTS_MODAL:
			storyCopy = { ...state };
			storyCopy.highlightsSliderModal.showModal = false;
			return storyCopy;
		case storyConstants.VISITED_STORY_SUCCESS: {
			return state;
		}
		case storyConstants.HAVE_LOGGED_USER_STORY_REQUEST:
			return {
				...state,
			};
		case storyConstants.HAVE_LOGGED_USER_STORY_SUCCESS:
			return {
				...state,
				iHaveAStory: action.haveStories,
			};
		case storyConstants.HAVE_LOGGED_USER_STORY_FAILURE: {
			return {
				...state,
				iHaveAStory: false,
			};
		}

		case storyConstants.REPORT_STORY_REQUEST:
			storyCopy = { ...state };
			storyCopy.storyReport.showError = false;
			storyCopy.storyReport.errorMessage = "";
			storyCopy.storyReport.showSuccessMessage = false;
			storyCopy.storyReport.successMessage = "";

			return storyCopy;
		case storyConstants.REPORT_STORY_SUCCESS:
			storyCopy = { ...state };
			storyCopy.storyReport.showError = false;
			storyCopy.storyReport.errorMessage = "";
			storyCopy.storyReport.showSuccessMessage = true;
			storyCopy.storyReport.successMessage = action.successMessage;

			return storyCopy;

		case storyConstants.REPORT_STORY_FAILURE:
			storyCopy = { ...state };
			storyCopy.storyReport.showError = true;
			storyCopy.storyReport.errorMessage = action.errorMessage;
			storyCopy.storyReport.showSuccessMessage = false;
			storyCopy.storyReport.successMessage = "";

			return storyCopy;
		default:
			return state;
	}
};

function createStories(stories) {
	var retVal = [];

	stories.Media.forEach((media) => {
		retVal.push({
			url: media.Url,
			header: {
				heading: stories.UserInfo.Username,
				profileImage: stories.UserInfo.ImageURL,
				storyId: media.StoryId,
			},
			type: media.MediaType === "VIDEO" ? "video" : "image",
			tags: media.Tags,
		});
	});

	return retVal;
}

function createHighlights(highlights, name) {
	var retVal = [];

	console.log(highlights);
	highlights.media.forEach((media) => {
		retVal.push({
			url: media.media.url,
			header: {
				heading: name,
				profileImage: highlights.url !== "" ? highlights.url : "assets/img/profile.jpg",
				storyId: media.id,
			},
			type: media.media.mediaType === "VIDEO" ? "video" : "image",
		});
	});
	console.log(retVal);
	return retVal;
}
