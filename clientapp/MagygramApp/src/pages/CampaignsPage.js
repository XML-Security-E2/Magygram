import React from "react";
import StoryContextProvider from "../contexts/StoryContext";
import PostContextProvider from "../contexts/PostContext";
import HeaderWrapper from "../components/HeaderWrapper";
import PostCampaignList from "../components/agent-campaigns/PostCampaignList";
import ViewPostCampaignModal from "../components/modals/ViewPostCampaignModal";
import StoryCampaignList from "../components/agent-campaigns/StoryCampaignList";
import StoryCampaignModal from "../components/modals/StoryCampaignModal";

const CampaignsPage = () => {
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
											<StoryCampaignList />
											<PostCampaignList />
											<StoryCampaignModal />
											<ViewPostCampaignModal />
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

export default CampaignsPage;
