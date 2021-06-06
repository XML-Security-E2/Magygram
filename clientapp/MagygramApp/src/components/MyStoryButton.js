
import React,{ useContext } from "react";
import { UserContext } from "../contexts/UserContext";
import { StoryContext } from "../contexts/StoryContext";
import { getUserInfo } from "../helpers/auth-header";

const MyStoryButton = ({openStorySlider, hideButton}) => {
	const { userState } = useContext(UserContext);
	const imgStyle = {"transform":"scale(1.5)","width":"100%","position":"absolute","left":"0"};
	const { storyState } = useContext(StoryContext);

	return (
		<li hidden={!storyState.iHaveAStory} class="list-inline-item">
			<button onClick={()=>openStorySlider(getUserInfo().Id)} className="btn p-0 m-0">
				<div className="d-flex flex-column align-items-center">
					<div className="btn-secondary rounded-circle overflow-hidden d-flex justify-content-center align-items-center border border-danger story-profile-photo-my-story">
						<img src={getUserInfo().ImageURL} alt="..." style={imgStyle}/>
					</div>
					<small>My stories</small>
				</div>
			</button>
		</li>
	);
};

export default MyStoryButton;
