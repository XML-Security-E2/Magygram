import React, { useState, useContext } from "react";
import { storyService } from "../services/StoryService";
import { StoryContext } from "../contexts/StoryContext";

const ReportedPost = ({storyId,userId}) => {

	const { state,dispatch } = useContext(StoryContext);

    const handleViewPost = async (storyId) => {
	
		await storyService.findStoryById(storyId, userId, dispatch);
	};

	const handleVisitProfile = (userId) => {
		window.location = "#/profile?userId=" + userId;
	}

    return (
		<div>
			<div>
				<a onClick={() => handleViewPost(storyId)} class="link-primary">View story</a>
			</div>
			<div>
				<a onClick={() => handleVisitProfile(userId)} class="link-primary">View profile</a>
			</div>
		</div>
        
	);
    
};

export default ReportedPost;