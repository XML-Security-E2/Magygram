import React, { useContext, useEffect } from "react";
import { PostContext } from "../contexts/PostContext";
import Post from "./Post";
import "react-image-gallery/styles/css/image-gallery.css";
import { postService } from "../services/PostService";
import { useLocation } from 'react-router-dom';

const SearchedPostTimeline = (props) => {
	const { postState, dispatch } = useContext(PostContext);
	const value = props.id;
    const location = useLocation();

	useEffect(() => {
        if(location.pathname.startsWith('/search/hashtag/')){
            const getPostsHandler = async () => {
                await postService.findPostsForUserByHashtag(value, dispatch);
            };
            getPostsHandler();
        }else if(location.pathname.startsWith('/search/location/')){
            const getPostsHandler = async () => {
                await postService.findPostsForUserByLocation(value, dispatch);
            };
            getPostsHandler();
        }
	}, [value]);

	return (
		<React.Fragment>
			{postState.timeline.posts.map((post) => {
				console.log(post);
				return <Post post={post} />;
			})}
		</React.Fragment>
	);
};

export default SearchedPostTimeline;
