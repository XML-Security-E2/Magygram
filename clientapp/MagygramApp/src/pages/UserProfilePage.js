import React from "react";
import StoryContextProvider from "../contexts/StoryContext";
import UserStoriesModal from "../components/modals/UserStoriesModal";
import UserProfileHeaderInfo from "../components/UserProfileHeaderInfo";
import UserContextProvider from "../contexts/UserContext";
import UserProfileContent from "../components/UserProfileContent";
import PostContextProvider from "../contexts/PostContext";
import HeaderWrapper from "../components/HeaderWrapper";
import EditPostModal from "../components/modals/EditPostModal";
import NotificationContextProvider from "../contexts/NotificationContext";

const UserProfilePage = (props) => {
	const search = props.location.search;
	const userId = new URLSearchParams(search).get("userId");

	return (
		<React.Fragment>
			<HeaderWrapper />
			<PostContextProvider>
				<StoryContextProvider>
					<div>
						<div className="mt-4">
							<div className="container d-flex justify-content-center">
								<div className="col-12">
									<div className="row">
										<div className="col-12">
											<UserContextProvider>
												<NotificationContextProvider>
													<UserProfileHeaderInfo userId={userId} />
												</NotificationContextProvider>
											</UserContextProvider>
											<UserProfileContent userId={userId} />
										</div>
									</div>
								</div>
							</div>
						</div>
					</div>
					<EditPostModal />

					<UserStoriesModal />
				</StoryContextProvider>
			</PostContextProvider>
		</React.Fragment>
	);
};

export default UserProfilePage;
