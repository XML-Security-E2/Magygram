import React, { useContext, useRef } from "react";
import { productConstants } from "../constants/ProductConstants";
import { ProductContext } from "../contexts/ProductContext";
import { productService } from "../services/ProductService";
import CreateCampaignForm from "./CreateCampaignForm";
import FailureAlert from "./FailureAlert";
import SuccessAlert from "./SuccessAlert";

const CreateCampaignWrapper = ({ productId }) => {
	const { productState, dispatch } = useContext(ProductContext);

	const campaignRef = useRef();

	const handleSubmit = () => {
		let post = {};
		post.productId = productId;

		post.displayTime = parseInt(campaignRef.current.campaignState.displayTime);
		if (campaignRef.current.campaignState.checkedOnce) {
			post.frequency = "ONCE";
		} else {
			post.frequency = "REPEATEDLY";
		}

		if (campaignRef.current.campaignState.checkedPost) {
			post.campaignType = "POST";
		} else {
			post.campaignType = "STORY";
		}
		post.targetGroup = {
			gender: campaignRef.current.campaignState.gender,
			minAge: parseInt(campaignRef.current.campaignState.minAge),
			maxAge: parseInt(campaignRef.current.campaignState.maxAge),
		};
		post.startDate = campaignRef.current.campaignState.startDate.getTime();
		post.endDate = campaignRef.current.campaignState.endDate.getTime();
		post.exposeOnceDate = campaignRef.current.campaignState.exposeOnceDate.getTime();
		post.minDisplays = parseInt(campaignRef.current.campaignState.minDisplaysForRepeatedly);
		console.log(campaignRef.current.campaignState);

		productService.createCampaign(post, dispatch);
	};

	return (
		<React.Fragment>
			<FailureAlert
				hidden={!productState.createCampaign.showErrorMessage}
				header="Error"
				message={productState.createCampaign.errorMessage}
				handleCloseAlert={() => dispatch({ type: productConstants.CREATE_CAMPAIGN_REQUEST })}
			/>
			<SuccessAlert
				hidden={!productState.createCampaign.showSuccessMessage}
				header="Success"
				message={productState.createCampaign.successMessage}
				handleCloseAlert={() => dispatch({ type: productConstants.CREATE_CAMPAIGN_REQUEST })}
			/>
			<h3 className="row">Create campaign</h3>
			<CreateCampaignForm ref={campaignRef} handleSubmit={handleSubmit} />
		</React.Fragment>
	);
};

export default CreateCampaignWrapper;
