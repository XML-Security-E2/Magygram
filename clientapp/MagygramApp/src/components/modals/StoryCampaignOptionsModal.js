import { useContext, useEffect, useRef, useState } from "react";
import { Button, Modal } from "react-bootstrap";
import { modalConstants } from "../../constants/ModalConstants";
import { hasRoles } from "../../helpers/auth-header";
import SuccessAlert from "../SuccessAlert";
import FailureAlert from "../FailureAlert";
import { StoryContext } from "../../contexts/StoryContext";
import { storyService } from "../../services/StoryService";
import { storyConstants } from "../../constants/StoryConstants";
import EditStoryCampaignForm from "../agent-campaigns/EditStoryCampaignForm";

const StoryCampaignOptionsModal = ({ storyId }) => {
	const { storyState, dispatch } = useContext(StoryContext);

	const [hiddenForm, setHiddenForm] = useState(true);
	const [hiddenCampaignAlert, setHiddenCampaignAlert] = useState(true);
	const campaignRef = useRef();

	const handleEditCampaign = () => {
		let campaign = {
			campaignId: storyState.agentCampaignStoryOptionModal.campaign.id,
			minDisplaysForRepeatedly: parseInt(campaignRef.current.campaignState.minDisplaysForRepeatedly),
			targetGroup: {
				minAge: parseInt(campaignRef.current.campaignState.minAge),
				maxAge: parseInt(campaignRef.current.campaignState.maxAge),
				gender: campaignRef.current.campaignState.gender,
			},
			dateFrom: campaignRef.current.campaignState.startDate.getTime(),
			dateTo: campaignRef.current.campaignState.endDate.getTime(),
		};

		storyService.updateStoryCampaign(campaign, dispatch);
	};

	const handleModalClose = () => {
		dispatch({ type: modalConstants.HIDE_STORY_AGENT_OPTIONS_MODAL });
		setHiddenForm(true);
		setHiddenCampaignAlert(true);
	};

	const handleDelete = () => {
		storyService.deleteStoryCampaign(storyId, dispatch);
	};

	const handleOpenCampaignStoryEditModal = () => {
		storyService.getCampaignByStoryId(storyId, dispatch);
	};

	useEffect(() => {
		setHiddenCampaignAlert(storyState.agentCampaignStoryOptionModal.campaign.frequency === "REPEATEDLY" || storyState.agentCampaignStoryOptionModal.campaign.frequency === "");
		setHiddenForm(storyState.agentCampaignStoryOptionModal.campaign.frequency === "ONCE" || storyState.agentCampaignStoryOptionModal.campaign.frequency === "");
	}, [storyState.agentCampaignStoryOptionModal.campaign]);

	return (
		<Modal show={storyState.agentCampaignStoryOptionModal.showModal} aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Header closeButton>
				<Modal.Title id="contained-modal-title-vcenter">{hiddenForm ? "Options" : "Report"}</Modal.Title>
			</Modal.Header>
			<Modal.Body>
				<SuccessAlert
					hidden={!storyState.agentCampaignStoryOptionModal.showSuccessMessage}
					header="Success"
					message={storyState.agentCampaignStoryOptionModal.successMessage}
					handleCloseAlert={() => dispatch({ type: storyConstants.HIDE_STORY_CAMPAIGN_OPTION_ALERTS })}
				/>
				<FailureAlert
					hidden={!storyState.agentCampaignStoryOptionModal.showError}
					header="Error"
					message={storyState.agentCampaignStoryOptionModal.errorMessage}
					handleCloseAlert={() => dispatch({ type: storyConstants.HIDE_STORY_CAMPAIGN_OPTION_ALERTS })}
				/>
				<FailureAlert hidden={hiddenCampaignAlert} header="Error" message="Camapigns that lasts only one day cannot be edited" handleCloseAlert={() => setHiddenCampaignAlert(true)} />

				<div hidden={!hiddenForm}>
					<div className="row">
						<button type="button" className="btn btn-link btn-fw text-secondary w-100 border-0" onClick={handleOpenCampaignStoryEditModal}>
							Edit
						</button>
					</div>
					<div className="row">
						<button hidden={!hasRoles(["agent"])} type="button" className="btn btn-link btn-fw text-danger w-100 border-0" onClick={handleDelete}>
							Delete
						</button>
					</div>
				</div>
				<div hidden={hiddenForm}>
					<EditStoryCampaignForm ref={campaignRef} fontColor="black" />
					<Button onClick={handleEditCampaign} className="btn float-right">
						Edit
					</Button>
				</div>
			</Modal.Body>
		</Modal>
	);
};

export default StoryCampaignOptionsModal;
