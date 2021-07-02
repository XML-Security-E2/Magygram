import React, { useState, useContext } from "react";
import { storyService } from "../services/StoryService";
import { StoryContext } from "../contexts/StoryContext";

const ReportedPost = ({storyId}) => {

	const { state,dispatch } = useContext(StoryContext);

    const handleViewPost = async (storyId) => {
		await storyService.findStoryById(storyId,"", dispatch);
	};

	const handleVisitProfile = (userId) => {
		window.location = "#/profile?userId=" + userId;
	}

    return (
			<div>
				<a onClick={() => handleViewPost(storyId)} class="link-primary">View story</a>
			</div>
			
        
	);
    
};

export default ReportedPost;