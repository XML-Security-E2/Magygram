import React, { useState, useContext } from "react";
import { storyService } from "../services/StoryService";
import { StoryContext } from "../contexts/StoryContext";

const ReportedPost = ({userId}) => {

	const { state,dispatch } = useContext(StoryContext);

    const handleViewPost = async (userId) => {
	
		await storyService.GetStoriesForUser(userId, 0, dispatch);
	};

    return (
            <a onClick={() => handleViewPost(userId)} class="link-primary">View story</a>
        
	);
    
};

export default ReportedPost;