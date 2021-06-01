import { modalConstants } from "../constants/ModalConstants";
import { storyConstants } from "../constants/StoryConstants";

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
		default:
			return state;
	}
};
