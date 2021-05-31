import Axios from "axios";
import { postConstants } from "../constants/PostConstants";
import { authHeader } from "../helpers/auth-header";

export const postService = {
	findPostsForTimeline,
};

async function findPostsForTimeline(dispatch) {
	dispatch(request());
	await Axios.get(`/api/posts`, { validateStatus: () => true, headers: {
		'Authorization': localStorage.getItem("accessToken")
	}	 })
		.then((res) => {
			if (res.status === 200) {
				
				dispatch(success(res.data));
			} else {
				failure()
			}
		})
		.catch((err) => {
			console.log(err);
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
