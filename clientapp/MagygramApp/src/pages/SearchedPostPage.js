import React from "react";
import PostContextProvider from "../contexts/PostContext";
import StoryContextProvider from "../contexts/StoryContext";
import CreateStoryModal from "../components/modals/CreateStoryModal";
import AddPostToFavouritesModal from "../components/modals/AddPostToFavouritesModal";
import Header from "../components/Header";
import UserContextProvider from "../contexts/UserContext";
import SearchedPostTimeline from "../components/SearchedPostsTimeline";
import NotificationContextProvider from "../contexts/NotificationContext";

const SearchPostPage = (props) => {

	return (
		<React.Fragment>
			<UserContextProvider>
				<StoryContextProvider>
					<PostContextProvider>
						<NotificationContextProvider>
							<Header />
						</NotificationContextProvider>
						<CreateStoryModal />
						<AddPostToFavouritesModal />
						<div>
							<div class="mt-4">
								<div class="container d-flex justify-content-center">
									<div class="col-9">
										<div class="row">
											<div class="col-8">
												<SearchedPostTimeline  id={props.match.params.id}/>
											</div>
										</div>
									</div>
								</div>
							</div>
		    			</div>
					</PostContextProvider>
				</StoryContextProvider>
			</UserContextProvider>
		</React.Fragment>
	);
};

export default SearchPostPage;
