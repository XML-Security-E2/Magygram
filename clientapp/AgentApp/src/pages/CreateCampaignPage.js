import React from "react";
import CreateCampaignWrapper from "../components/CreateCampaignWrapper";
import Header from "../components/Header";
import OrderContextProvider from "../contexts/OrderContext";
import ProductContextProvider from "../contexts/ProductContext";

const CreateCampaignPage = (props) => {
	const search = props.location.search;
	const productId = new URLSearchParams(search).get("productId");

	return (
		<React.Fragment>
			<ProductContextProvider>
				<OrderContextProvider>
					<Header />
					<div>
						<div className="mt-4">
							<div className="container d-flex justify-content-center">
								<div className="col-10">
									<CreateCampaignWrapper productId={productId} />
								</div>
							</div>
						</div>
					</div>
				</OrderContextProvider>
			</ProductContextProvider>
		</React.Fragment>
	);
};

export default CreateCampaignPage;
