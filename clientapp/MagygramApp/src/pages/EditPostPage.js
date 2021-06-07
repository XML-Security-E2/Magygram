import React from "react";
import EditPostForm from "../components/EditPostForm";
import HeaderWrapper from "../components/HeaderWrapper";
import PostContextProvider from "../contexts/PostContext";

const EditPostPage = () => {
	return (
		<React.Fragment>
			<div className="container-scroller">
				<div className="container-fluid ">
					<HeaderWrapper />
					<div className="main-panel">
						<div className="container">
							<PostContextProvider>
								<EditPostForm />
							</PostContextProvider>
						</div>
					</div>
				</div>
			</div>
		</React.Fragment>
	);
};

export default EditPostPage;
