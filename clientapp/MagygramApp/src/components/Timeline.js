import React, {useContext} from "react";
import { PostContext } from "../contexts/PostContext";
import Post from "./Post";
import "react-image-gallery/styles/css/image-gallery.css";

const Timeline = () => {
	const { postState } = useContext(PostContext);

	return (
        <React.Fragment>
            {postState.timeline.posts.map((post) => {
						return <Post post={post}/>; })}
        </React.Fragment>
	);
};

export default Timeline;
