import React from "react";
import CreatePostForm from "../components/CreatePostForm";
import Header from "../components/Header";
import PostContextProvider from "../contexts/PostContext";

const CreatePostPage = () => {
	return (
		<React.Fragment>
			<div className="container-scroller">
				<div className="container-fluid ">
					<Header />
					<div className="main-panel">
						<div className="container">
							<PostContextProvider>
								<CreatePostForm />
							</PostContextProvider>
						</div>
					</div>
				</div>
			</div>
		</React.Fragment>
	);
};

export default CreatePostPage;
