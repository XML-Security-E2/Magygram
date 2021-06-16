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
import NotificationContextProvider from "../contexts/NotificationContext";

const HomePage = () => {
	return (
		<React.Fragment>
			{hasRoles(["user"]) ? (
				<div>
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
				</div>
			) : (
				<div>
					{hasRoles(["admin"]) ? (
						<PostContextProvider>
							<div>
								<div class="mt-4">
									<div class="container d-flex justify-content-center">
										<div class="col-9">
											<div class="row">
												<div class="col-8">
													<h1>Admin zona</h1>
												</div>
											</div>
										</div>
									</div>
								</div>
							</div>
						</PostContextProvider>
					) : (
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
						</div>
					)}
				</div>
			)}
		</React.Fragment>
	);
};

export default HomePage;
