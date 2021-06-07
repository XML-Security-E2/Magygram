import React from "react";
import CreatePostForm from "../components/CreatePostForm";
import HeaderWrapper from "../components/HeaderWrapper";
import PostContextProvider from "../contexts/PostContext";

const CreatePostPage = () => {
	return (
		<React.Fragment>
			<div className="container-scroller">
				<div className="container-fluid ">
					<HeaderWrapper />
					<PostContextProvider>
						<div className="main-panel">
							<div className="container">
								<CreatePostForm />
							</div>
						</div>
					</PostContextProvider>
				</div>
			</div>
		</React.Fragment>
	);
};

export default CreatePostPage;
