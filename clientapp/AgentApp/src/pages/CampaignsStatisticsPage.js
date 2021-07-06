import React from "react";
import Header from "../components/Header";
import CampaignStatisticList from "../components/CampaignStatisticList";
import OrderContextProvider from "../contexts/OrderContext";
import ProductContextProvider from "../contexts/ProductContext";

const CampaignsStatisticsPage = () => {
	return (
		<React.Fragment>
			<ProductContextProvider>
				<OrderContextProvider>
					<Header />
					<div>
						<div className="mt-4">
							<div className="container d-flex justify-content-center">
								<div className="col-10">
									<CampaignStatisticList />
								</div>
							</div>
						</div>
					</div>
				</OrderContextProvider>
			</ProductContextProvider>
		</React.Fragment>
	);
};

export default CampaignsStatisticsPage;
