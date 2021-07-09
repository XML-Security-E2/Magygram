import React, { useContext, useEffect } from "react";
import { AdsContext } from "../../contexts/AdsContext";
import { adsService } from "../../services/AdsService";
import CampaignStatisticItem from "./CampaignStatisticItem";

const CampaignStatisticList = () => {
	const { adsState, dispatch } = useContext(AdsContext);

	useEffect(() => {
		const getPostStatisticHandler = async () => {
			await adsService.getPostCampaignStatistic(dispatch);
		};
		const getStoryStatisticHandler = async () => {
			await adsService.getStoryCampaignStatistic(dispatch);
		};
		getPostStatisticHandler();
		getStoryStatisticHandler();
	}, [dispatch]);

	return (
		<React.Fragment>
			<div className="row">
				<div className="col-12">
					<h3 className="text-dark">Post campaigns</h3>
				</div>
				{adsState.postCampaigns !== null &&
					adsState.postCampaigns.map((campaign) => {
						return <CampaignStatisticItem statistics={campaign} />;
					})}
			</div>
			<div className="row mb-3">
				<div className="col-12">
					<h3 className="text-dark">Story campaigns</h3>
				</div>

				{adsState.storyCampaigns !== null &&
					adsState.storyCampaigns.map((campaign) => {
						return <CampaignStatisticItem statistics={campaign} />;
					})}
			</div>
		</React.Fragment>
	);
};

export default CampaignStatisticList;
