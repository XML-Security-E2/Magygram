import React from "react";
import Timeline from "../components/Timeline";
import PostContextProvider from "../contexts/PostContext";
import StoryContextProvider from "../contexts/StoryContext";
import CreateStoryModal from "../components/modals/CreateStoryModal";
import AddPostToFavouritesModal from "../components/modals/AddPostToFavouritesModal";
import Header from "../components/Header";
import Storyline from "../components/Storyline";
import UserContextProvider from "../contexts/UserContext";
import { hasRoles } from "../helpers/auth-header";
import GuestHeader from "../components/GuestHeader";
import GuestTimeline from "../components/GuestTimeline";
import AdminHomePageTabs from "../components/AdminHomePageTabs";
import AdminVerificationRequestTabContent from "../components/AdminVerificationRequestTabContent";
import AdminContextProvider from "../contexts/AdminContext";

const HomePage = () => {

	return (
		<React.Fragment>
			{hasRoles(["user"]) ? 
			<div>
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
												<Storyline />
												<Timeline search={false} />
											</div>
										</div>
									</div>
								</div>
							</div>
						</div>
					</PostContextProvider>
				</StoryContextProvider>
			</UserContextProvider>
		</div> : 
			
		<div>
			{hasRoles(["admin"]) ? 
			<UserContextProvider>
				<StoryContextProvider>
					<PostContextProvider>
						<AdminContextProvider>
							<Header />
							<div>
								<div className="mt-4">
									<div className="container d-flex justify-content-center">
										<div className="col-12">
											<div className="row">
												<div className="col-12">
													<AdminHomePageTabs/>
													<AdminVerificationRequestTabContent/>
												</div>
											</div>
										</div>
									</div>
								</div>
							</div>
						</AdminContextProvider>
					</PostContextProvider>
				</StoryContextProvider>
			</UserContextProvider>
			:
			<div>
				<PostContextProvider>
				<GuestHeader />
					<div>
						<div class="mt-4">
							<div class="container d-flex justify-content-center">
								<div class="col-9">
									<div class="row">
										<div class="col-8">
											<GuestTimeline />
										</div>
									</div>
								</div>
							</div>
						</div>
					</div>
				</PostContextProvider>
			</div>}
	</div>}
			
		</React.Fragment>
	);
};

export default HomePage;
