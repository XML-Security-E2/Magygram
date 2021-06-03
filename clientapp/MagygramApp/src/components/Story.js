import React from "react";;

const Story = ({story,openStorySlider}) => {
	const imgStyle = {"transform":"scale(1.5)","width":"100%","position":"absolute","left":"0"};
	const visitedStoryClassName= "rounded-circle overflow-hidden d-flex justify-content-center align-items-center border border-danger story-profile-photo-visited";
	const notVisitedStoryClassName= "rounded-circle overflow-hidden d-flex justify-content-center align-items-center border border-danger story-profile-photo";
	
	return (
        <React.Fragment>
            <li class="list-inline-item">
				<button onClick={() => openStorySlider(story.UserInfo.Id)} className="btn p-0 m-0">
					<div className="d-flex flex-column align-items-center">
						<div className={story.Visited ? visitedStoryClassName : notVisitedStoryClassName}>
							<img src={story.UserInfo.ImageURL} alt="..." style={imgStyle}/>
						</div>
						<small>{story.UserInfo.Username}</small>
					</div>
				</button>
			</li>
        </React.Fragment>
	);
};

export default Story;
