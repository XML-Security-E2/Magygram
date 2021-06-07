import React from "react";
import CreatePostForm from "../components/CreatePostForm";
import Header from "../components/Header";
import PostContextProvider from "../contexts/PostContext";

const CreatePostPage = () => {
	return (
		<React.Fragment>
			<div className="container-scroller">
				<div className="container-fluid ">
					<PostContextProvider>
						<Header />
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
