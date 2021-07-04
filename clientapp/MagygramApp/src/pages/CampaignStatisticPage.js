import React from "react";
import HeaderWrapper from "../components/HeaderWrapper";
import AdsContextProvider from "../contexts/AdsContext";
import CampaignStatisticList from "../components/campaign-statistic/CampaignStatisticList";

const CampaignStatisticPage = () => {
	return (
		<React.Fragment>
			<HeaderWrapper />
			<AdsContextProvider>
				<div>
					<div className="mt-4">
						<div className="container d-flex justify-content-center">
							<div className="col-12">
								<div className="row">
									<div className="col-12">
										<CampaignStatisticList />
									</div>
								</div>
							</div>
						</div>
					</div>
				</div>
			</AdsContextProvider>
		</React.Fragment>
	);
};

export default CampaignStatisticPage;
