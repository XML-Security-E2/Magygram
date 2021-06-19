import React, { useState, useContext } from "react";
import { postService } from "../services/PostService";
import { PostContext } from "../contexts/PostContext";

const ReportedPost = ({id}) => {

	const { state,dispatch } = useContext(PostContext);

    const handleViewPost = async (postId) => {
	
		await postService.findPostById(postId, dispatch);
	};

    return (
            <a onClick={() => handleViewPost(id)} class="link-primary">View post</a>
        
	);
    
};

export default ReportedPost;