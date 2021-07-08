import React, { useContext } from "react";
import { productConstants } from "../constants/ProductConstants";
import { ProductContext } from "../contexts/ProductContext";
import { productService } from "../services/ProductService";

const CampaignStatisticReportItem = ({ report }) => {
	const { dispatch } = useContext(ProductContext);

	const setShowedCampaigns = () => {
		dispatch({ type: productConstants.CHANGE_SHOWED_CAMPAIGNS, campaigns: report.campaigns });
	};

	const downloadFIle = () => {
		productService.downloadReport(report.fileId + ".xml");
	};

	return (
		<React.Fragment>
			<div className="col-12 mt-2">
				<div className="row rounded-lg border p-2" style={{ cursor: "pointer" }} onClick={setShowedCampaigns}>
					<div className="col-6">
						<b>{report.fileId + ".xml"}</b>
					</div>
					<div className="col-2">
						{new Date(report.dateCreating).toLocaleDateString("en-US", {
							day: "2-digit",
							month: "2-digit",
							year: "numeric",
						})}
					</div>
					<div className="col-4">
						<button type="button" className="btn btn-outline-secondary" onClick={downloadFIle}>
							Generate PDF
						</button>
					</div>
				</div>
			</div>
		</React.Fragment>
	);
};

export default CampaignStatisticReportItem;
