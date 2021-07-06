import React, { useContext, useEffect } from "react";
import { ProductContext } from "../contexts/ProductContext";
import { productService } from "../services/ProductService";
import CampaignStatisticReportItem from "./CampaignStatisticReportItem";

const CampaignStatisticReportList = () => {
	const { productState, dispatch } = useContext(ProductContext);

	useEffect(() => {
		const getReportsHandler = async () => {
			await productService.getGeneratedCampaignStatisticsReports(dispatch);
		};
		getReportsHandler();
	}, [dispatch]);

	return (
		<React.Fragment>
			<div className="row">
				<div className="col-12">
					<h3 className="text-dark">Reports</h3>
					<div className="row mt-3">
						<div className="col-6">
							<b>File name</b>
						</div>
						<div className="col-2">
							<b>Creation date</b>
						</div>
						<div className="col-4"></div>
					</div>
				</div>

				{productState.reports.map((report) => {
					return <CampaignStatisticReportItem report={report} />;
				})}
			</div>
		</React.Fragment>
	);
};

export default CampaignStatisticReportList;
