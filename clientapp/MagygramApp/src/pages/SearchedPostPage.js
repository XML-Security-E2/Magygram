import React from "react";
import Timeline from "../components/Timeline";
import PostContextProvider from "../contexts/PostContext";
import StoryContextProvider from "../contexts/StoryContext";
import CreateStoryModal from "../components/modals/CreateStoryModal";
import AddPostToFavouritesModal from "../components/modals/AddPostToFavouritesModal";
import Header from "../components/Header";
import UserContextProvider from "../contexts/UserContext";

const SearchPostPage = () => {

	return (
		<React.Fragment>
			<UserContextProvider>
				<StoryContextProvider>
					<PostContextProvider>
						<Header />
						<CreateStoryModal />
						<AddPostToFavouritesModal />
						<div>
							<div class="mt-4">
								<div class="container d-flex justify-content-center">
									<div class="col-9">
										<div class="row">
											<div class="col-8">
												<Timeline search={true}/>
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
