import React from "react";
import CreatePostForm from "../components/CreatePostForm";
import Header from "../components/Header";

const CreatePostPage = () => {
	return (
		<React.Fragment>
			<div className="container-scroller">
				<div className="container-fluid ">
					<Header />
					<div className="main-panel">
						<div className="container">
							<CreatePostForm />
						</div>
					</div>
				</div>
			</div>
		</React.Fragment>
	);
};

export default CreatePostPage;
