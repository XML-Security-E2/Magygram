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
		case storyConstants.STORYLINE_STORY_REQUEST:
			return {
				...state,
				storyline: {		
					stories : []
				}
			};
		case storyConstants.STORYLINE_STORY_SUCCESS:
			return {
				...state,
				storyline: {		
					stories : action.stories
				}
			};
		case storyConstants.STORYLINE_STORY_FAILURE:
			return {
				...state,
				storyline: {		
					stories : []
				}
			};
		case modalConstants.SHOW_STORY_SLIDER_MODAL:
			return {
				...state,
				storySliderModal: {
					showModal: true,
					stories: createStories(action.stories),
					firstUnvisitedStory: action.stories.FirstUnvisitedStory
				}
			};
		case modalConstants.HIDE_STORY_SLIDER_MODAL:
			return {
				...state,
				storySliderModal: {
					showModal: false,
					stories: action.stories,
					firstUnvisitedStory: 0,
				}
			};
		default:
			return state;
	}
};

function createStories(stories){
	var retVal =[]

	stories.Media.forEach(media =>{
		retVal.push({
			url:media.Url,
			header: {
				heading: stories.UserInfo.Username,
				profileImage: stories.UserInfo.ImageURL,
				storyId: media.StoryId
			},
			type: media.MediaType==='VIDEO'?'video':'image',
		})
	})

	return retVal;
}