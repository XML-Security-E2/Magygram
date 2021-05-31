import React, { createContext, useReducer } from "react";
import { postReducer } from "../reducers/PostReducer";

export const PostContext = createContext();

const PostContextProvider = (props) => {
	const [postState, dispatch] = useReducer(postReducer, {
		timeline: {
			posts: [{
				Description : "Deskripcija prvog posta",
    			Location : "Kraljevo",
    			PostType : "REGULAR",
				Media: [
					{ Url :"assets/images/posts/post-1.jpg", MediaType: "IMAGE" },
					{ Url: "http://lorempixel.com/1000/600/nature/2/", MediaType: "IMAGE" },
					{ Url: "assets/images/posts/video1.mp4", MediaType: "VIDEO" }
				],
				UserInfo : {
					Id : "2a5a1dc2-f5ad-4eb8-ae69-f4ea160422ff",
					Username : "TestUsername",
					ImageURL : "assets/images/profiles/profile-1.jpg"
				},
				Comments: [
					{ Id : "1", Text: "123"},
					{ Id : "2", Text: "123"},
				],
				LikedBy: [
					{ Id : "1", Text: "123"},
					{ Id : "2", Text: "123"},
				] 
			},
			{
				Description : "Deskripcija drugog posta",
    			Location : "Kraljevo",
    			PostType : "REGULAR",
				Media: [
					{ Url :"assets/images/posts/post-1.jpg", MediaType: "IMAGE" },
					{ Url: "images/image2", MediaType: "IMAGE" },
					{ Url: "videos/video1", MediaType: "VIDEO" }
				],
				UserInfo : {
					Id : "2a5a1dc2-f5ad-4eb8-ae69-f4ea160422ff",
					Username : "TestUsernameDrugi",
					ImageURL : "assets/images/profiles/profile-5.jpg"
				},
				Comments: [
					{ Id : "1", Text: "123"},
					{ Id : "2", Text: "123"},
				],
				LikedBy: [
					{ Id : "1", Text: "123"},
					{ Id : "2", Text: "123"},
					{ Id : "2", Text: "123"},
					{ Id : "2", Text: "123"},
				] 
			}
		],
		},
	});

	return <PostContext.Provider value={{ postState, dispatch }}>{props.children}</PostContext.Provider>;
};

export default PostContextProvider;
