import React from "react";
import StoryContextProvider from "../contexts/StoryContext";
import PostContextProvider from "../contexts/PostContext";
import HeaderWrapper from "../components/HeaderWrapper";
import ViewPostModal from "../components/modals/ViewPostModal";
import UserDislikedPosts from "../components/UserDislikedPosts";

const DislikedPostPage = () => {
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
                                            <UserDislikedPosts/>
                                            <ViewPostModal />
										</div>
									</div>
								</div>
							</div>
						</div>
					</div>
				</StoryContextProvider>
			</PostContextProvider>
		</React.Fragment>
	);
};

export default DislikedPostPage;
