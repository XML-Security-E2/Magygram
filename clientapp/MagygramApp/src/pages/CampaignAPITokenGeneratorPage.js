import React from "react";
import CampaignTokenGenerator from "../components/agent-campaigns/CampaignTokenGenerator";
import HeaderWrapper from "../components/HeaderWrapper";
import UserContextProvider from "../contexts/UserContext";

const CampaignAPITokenGeneratorPage = () => {
	return (
		<React.Fragment>
			<HeaderWrapper />
				<div>
					<div className="mt-4">
						<div className="container d-flex justify-content-center">
								<div className="row">
									<div className="col-12 mt-5">
                                    <UserContextProvider>
                                        <CampaignTokenGenerator/>
                                    </UserContextProvider>
									</div>
								</div>
						</div>
					</div>
				</div>
		</React.Fragment>
	);
};

export default CampaignAPITokenGeneratorPage;
