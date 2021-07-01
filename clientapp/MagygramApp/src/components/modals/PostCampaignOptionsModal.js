import { useContext, useRef, useState } from "react";
import { Button, Modal } from "react-bootstrap";
import { modalConstants } from "../../constants/ModalConstants";
import { PostContext } from "../../contexts/PostContext";
import Axios from "axios";
import { authHeader } from "../../helpers/auth-header";
import { hasRoles } from "../../helpers/auth-header";
import { postConstants } from "../../constants/PostConstants";
import SuccessAlert from "../SuccessAlert";
import FailureAlert from "../FailureAlert";
import EditCampaignForm from "../agent-campaigns/EditCampaignForm";
import { postService } from "../../services/PostService";

const PostCampaignOptionsModal = ({ postId }) => {
	const { postState, dispatch } = useContext(PostContext);

	const [hiddenForm, setHiddenForm] = useState(true);
	const campaignRef = useRef();

	const handleEditCampaign = () => {
		// 	CampaignId string `bson:"campaign_id,omitempty" json:"campaignId"`
		// MinDisplaysForRepeatedly int `bson:"min_displays_for_repeatedly" json:"minDisplaysForRepeatedly"`
		// Frequency CampaignFrequency `bson:"frequency" json:"frequency"`
		// TargetGroup TargetGroup `bson:"target_group" json:"targetGroup"`
		// DateFrom time.Time `bson:"date_from" json:"dateFrom"`
		// DateTo time.Time `bson:"date_to" json:"dateTo"`

		// type TargetGroup struct {
		// 	MinAge int `bson:"min_age" json:"minAge" validate:"required,numeric,min=12,max=70"`
		// 	MaxAge int `bson:"max_age" json:"maxAge" validate:"required,numeric,min=12,max=120"`
		// 	Gender GenderType `bson:"gender" json:"gender"`
		// }
		let campaign = {
			campaignId: postState.agentCampaignPostOptionModal.campaign.id,
			minDisplaysForRepeatedly: 0, //TOOOOODOOOOO
			frequency: campaignRef.current.campaignState.checkedOnce ? "ONCE" : "REPEATEDLY",
			targetGroup: {
				minAge: campaignRef.current.campaignState.minAge,
				maxAge: campaignRef.current.campaignState.maxAge,
				gender: campaignRef.current.campaignState.gender,
			},
			dateFrom: campaignRef.current.campaignState.startDate.getTime(),
			dateTo: campaignRef.current.campaignState.endDate.getTime(),
		};
	};

	const handleModalClose = () => {
		dispatch({ type: modalConstants.HIDE_POST_AGENT_OPTIONS_MODAL });
		setHiddenForm(true);
	};

	const handleDelete = () => {
		let requestId = postState.editPost.post.id;
		Axios.put(`/api/posts/${requestId}/delete`, {}, { validateStatus: () => true, headers: authHeader() })
			.then((res) => {
				console.log(res);
				if (res.status === 200) {
					console.log("Post has been deleted");
					alert("You have successfully deleted this post!");
					dispatch({ type: modalConstants.HIDE_POST_OPTIONS_MODAL });
				} else {
					console.log("ERROR");
				}
			})
			.catch((err) => {
				console.log("ERROR");
			});
	};

	const handleOpenCampaignPostEditModal = () => {
		postService.getCampaignByPostId(postId, dispatch);
		setHiddenForm(false);
	};

	return (
		<Modal show={postState.agentCampaignPostOptionModal.showModal} aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Header closeButton>
				<Modal.Title id="contained-modal-title-vcenter">{hiddenForm ? "Options" : "Report"}</Modal.Title>
			</Modal.Header>
			<Modal.Body>
				<SuccessAlert
					hidden={!postState.agentCampaignPostOptionModal.showSuccessMessage}
					header="Success"
					message={postState.agentCampaignPostOptionModal.successMessage}
					handleCloseAlert={() => dispatch({ type: postConstants.HIDE_POST_CAMPAIGN_OPTION_ALERTS })}
				/>
				<FailureAlert
					hidden={!postState.agentCampaignPostOptionModal.showError}
					header="Error"
					message={postState.agentCampaignPostOptionModal.errorMessage}
					handleCloseAlert={() => dispatch({ type: postConstants.HIDE_POST_CAMPAIGN_OPTION_ALERTS })}
				/>
				<div hidden={!hiddenForm}>
					<div className="row">
						<button type="button" className="btn btn-link btn-fw text-secondary w-100 border-0" onClick={handleOpenCampaignPostEditModal}>
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
					<EditCampaignForm ref={campaignRef} fontColor="black" />
					<Button onClick={handleEditCampaign} className="btn float-right">
						Edit
					</Button>
				</div>
			</Modal.Body>
		</Modal>
	);
};

export default PostCampaignOptionsModal;
