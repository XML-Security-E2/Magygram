import React from "react";
import CreateAgentPostForm from "../components/agent-post-create/CreateAgentPostForm";
import HeaderWrapper from "../components/HeaderWrapper";
import PostContextProvider from "../contexts/PostContext";

const CreateAgentPostPage = () => {
	return (
		<React.Fragment>
			<div className="container-scroller">
				<div className="container-fluid ">
					<HeaderWrapper />
					<PostContextProvider>
						<div className="main-panel">
							<div className="container">
								<CreateAgentPostForm />
							</div>
						</div>
					</PostContextProvider>
				</div>
			</div>
		</React.Fragment>
	);
};

export default CreateAgentPostPage;
