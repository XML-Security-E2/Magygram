import React from "react";
import Timeline from "../components/Timeline";
import PostContextProvider from "../contexts/PostContext";
import StoryContextProvider from "../contexts/StoryContext";
import CreateStoryModal from "../components/modals/CreateStoryModal";
import AddPostToFavouritesModal from "../components/modals/AddPostToFavouritesModal";
import Header from "../components/Header";
import Storyline from "../components/Storyline"

const HomePage = () => {
	return (
		<React.Fragment>
			<Header/>
			<StoryContextProvider>
				<PostContextProvider>
					<CreateStoryModal />
					<AddPostToFavouritesModal />
					<div>
						<div class="mt-4">
							<div class="container d-flex justify-content-center">
								<div class="col-9">
									<div class="row">
										<div class="col-8">
											<Storyline/>
											<Timeline />
										</div>
									</div>
								</div>
							</div>
						</div>
					</div>
				</PostContextProvider>
			</StoryContextProvider>
		</React.Fragment>
	);
};

export default HomePage;
