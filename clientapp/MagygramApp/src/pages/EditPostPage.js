import React from "react";
import EditPostForm from "../components/EditPostForm";
import Header from "../components/Header";
import PostContextProvider from "../contexts/PostContext";

const EditPostPage = () => {
	return (
		<React.Fragment>
			<div className="container-scroller">
				<div className="container-fluid ">
					<Header />
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
