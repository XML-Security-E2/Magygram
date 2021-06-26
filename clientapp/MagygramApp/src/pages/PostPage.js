import React from "react";
import PostContextProvider from "../contexts/PostContext";
import StoryContextProvider from "../contexts/StoryContext";
import CreateStoryModal from "../components/modals/CreateStoryModal";
import AddPostToFavouritesModal from "../components/modals/AddPostToFavouritesModal";
import Header from "../components/Header";
import UserContextProvider from "../contexts/UserContext";
import NotificationContextProvider from "../contexts/NotificationContext";
import FollowRecommendation from "../components/FollowRecommendation";
import MessageContextProvider from "../contexts/MessageContext";
import PostWrapper from "../components/PostWrapper";

const PostPage = (props) => {
	const search = props.location.search;
	const postId = new URLSearchParams(search).get("postId");

	return (
		<React.Fragment>
			<div>
				<UserContextProvider>
					<StoryContextProvider>
						<PostContextProvider>
							<NotificationContextProvider>
								<Header />
							</NotificationContextProvider>
							<CreateStoryModal />
							<AddPostToFavouritesModal />
							<MessageContextProvider>
								<div>
									<div className="mt-4">
										<div className="container d-flex justify-content-center">
											<div className="col-9">
												<div className="row">
													<div className="col-8">
														<PostWrapper postId={postId} />
													</div>
													<FollowRecommendation />
												</div>
											</div>
										</div>
									</div>
								</div>
							</MessageContextProvider>
						</PostContextProvider>
					</StoryContextProvider>
				</UserContextProvider>
			</div>
		</React.Fragment>
	);
};

export default PostPage;
