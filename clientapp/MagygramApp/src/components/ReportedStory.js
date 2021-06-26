import React, { useState, useContext } from "react";
import { storyService } from "../services/StoryService";
import { StoryContext } from "../contexts/StoryContext";

const ReportedPost = ({storyId}) => {

	const { state,dispatch } = useContext(StoryContext);

    const handleViewPost = async (storyId) => {
	
		await storyService.findStoryById(storyId, dispatch);
	};

    return (
            <a onClick={() => handleViewPost(storyId)} class="link-primary">View story</a>
        
	);
    
};

export default ReportedPost;