import React, {useEffect} from "react";
import Timeline from "../components/Timeline";
import PostContextProvider from "../contexts/PostContext";
import StoryContextProvider from "../contexts/StoryContext";
import CreateStoryModal from "../components/modals/CreateStoryModal";
import AddPostToFavouritesModal from "../components/modals/AddPostToFavouritesModal";
import Header from "../components/Header";
import Storyline from "../components/Storyline"
import UserContextProvider from "../contexts/UserContext";
import { hasRoles } from "../helpers/auth-header";

const HomePage = () => {
	var role = hasRoles(["user"]) ? "user":"guest";

	return (
		<React.Fragment>
			<div hidden={role==="guest"}>
				<UserContextProvider>
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
				</UserContextProvider>
			</div>
			<div hidden={role==="user"}>

			</div>
		</React.Fragment>
	);
};

export default HomePage