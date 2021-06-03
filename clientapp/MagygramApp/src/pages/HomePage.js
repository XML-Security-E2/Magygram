import React from "react";
import Timeline from "../components/Timeline";
import PostContextProvider from "../contexts/PostContext";
import StoryContextProvider from "../contexts/StoryContext";
import CreateStoryModal from "../components/modals/CreateStoryModal";
import AddPostToFavouritesModal from "../components/modals/AddPostToFavouritesModal";
import Header from "../components/Header";
import Storyline from "../components/Storyline"
import UserContextProvider from "../contexts/UserContext";
import { useHistory } from "react-router-dom";
import { authHeader } from "../helpers/auth-header";
import OptionsModal from "../components/modals/OptionsModal";
import EditPostModal from "../components/modals/EditPostModal";

const HomePage = () => {
	return (
		<React.Fragment>
			<UserContextProvider>
			<Header/>
			<StoryContextProvider>
				<PostContextProvider>
					<CreateStoryModal />
					<AddPostToFavouritesModal />
					<OptionsModal />
					<EditPostModal />
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
			</UserContextProvider>
		</React.Fragment>
	);
};

export default HomePage;
