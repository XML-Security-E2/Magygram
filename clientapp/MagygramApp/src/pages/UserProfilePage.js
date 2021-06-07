import React from "react";
import Header from "../components/Header";
import StoryContextProvider from "../contexts/StoryContext";
import UserStoriesModal from "../components/modals/UserStoriesModal";
import UserProfileHeaderInfo from "../components/UserProfileHeaderInfo";
import UserContextProvider from "../contexts/UserContext";
import UserProfileContent from "../components/UserProfileContent";
import PostContextProvider from "../contexts/PostContext";
import GuestHeader from "../components/GuestHeader";

const UserProfilePage = (props) => {
	const search = props.location.search;
	const userId = new URLSearchParams(search).get("userId");

	return (
		<React.Fragment>
			<PostContextProvider>
				{localStorage.getItem("userId") === null ? 
					<GuestHeader />: <Header />
				}
				<StoryContextProvider>
					<div>
						<div className="mt-4">
							<div className="container d-flex justify-content-center">
								<div className="col-12">
									<div className="row">
										<div className="col-12">
											<UserContextProvider>
												<UserProfileHeaderInfo userId={userId} />
											</UserContextProvider>
											<UserProfileContent userId={userId} />
										</div>
									</div>
								</div>
							</div>
						</div>
					</div>
					<UserStoriesModal />
				</StoryContextProvider>
			</PostContextProvider>
		</React.Fragment>
	);
};

export default UserProfilePage;
