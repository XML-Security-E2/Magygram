import Axios from "axios";
import { postConstants } from "../constants/PostConstants";

export const postService = {
	findPostsForTimeline,
};

async function findPostsForTimeline(dispatch) {
	dispatch(request());

	await Axios.get(`/api/posts`, { validateStatus: () => true })
		.then((res) => {
			if (res.status === 200) {
				dispatch(success(res.data));
			} else {
				failure()
			}
		})
		.catch((err) => {
			failure()
		});

	function request() {
		return { type: postConstants.TIMELINE_POSTS_REQUEST };
	}
	function success(data) {
		return { type: postConstants.TIMELINE_POSTS_SUCCESS, posts: data };
	}
	function failure() {
		return { type: postConstants.TIMELINE_POSTS_FAILURE };
	}
}
