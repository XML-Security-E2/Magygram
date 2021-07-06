import React, { useContext } from "react";
import { ProductContext } from "../contexts/ProductContext";
import { productService } from "../services/ProductService";
import CampaignStatisticItem from "./CampaignStatisticItem";

const CampaignStatisticList = () => {
	const { productState, dispatch } = useContext(ProductContext);

	const getCampaignStatisticHandler = async () => {
		await productService.getCampaignStatistics(dispatch);
	};

	return (
		<React.Fragment>
			<div className="row mt-5">
				<div className="col-12">
					<button type="button" className="btn btn-outline-secondary" onClick={getCampaignStatisticHandler}>
						Generate report
					</button>
				</div>
				<div className="col-12">{productState.showedCampaigns.lenght > 0 && <h3 className="text-dark">Top campaigns</h3>}</div>
				{productState.showedCampaigns.map((campaign) => {
					return <CampaignStatisticItem statistics={campaign} />;
				})}
			</div>
		</React.Fragment>
	);
};

export default CampaignStatisticList;
