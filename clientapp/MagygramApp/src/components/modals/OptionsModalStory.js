import { useContext, useState } from "react";
import { Modal } from "react-bootstrap";
import { modalConstants } from "../../constants/ModalConstants";
import { StoryContext } from "../../contexts/StoryContext";
import { storyConstants } from "../../constants/StoryConstants";
import { storyService } from "../../services/StoryService";
import SuccessAlert from "../SuccessAlert";
import FailureAlert from "../FailureAlert";

const OptionsModalStory = () => {
	const { storyState, dispatch } = useContext(StoryContext);

	const [reportReasons, setReportReasons] = useState([]);
	const [hiddenForm, setHiddenForm] = useState(true);

	const handleModalClose = () => {
		dispatch({ type: modalConstants.HIDE_STORY_OPTIONS_MODAL });
		setHiddenForm(true);
	};

	const handleShowForm = () => {
		setHiddenForm(false);
	};

	const handleSubmit = (e) => {
		e.preventDefault();

		let reportDTO = {
			contentId: storyState.storySliderModal.stories[storyState.storySliderModal.firstUnvisitedStory].header.storyId,
			contentType: "STORY",
			reportReasons,
		};

		storyService.reportStory(reportDTO, dispatch);
	};

	const toggleReportReason = (reason) => {
		let a = [...reportReasons];

		if (a.find((col) => col === reason) === undefined) {
			a.push(reason);
		} else {
			a = a.filter((reas) => reas !== reason);
		}
		setReportReasons(a);
		console.log(a);
	};

	return (
		<Modal show={storyState.postOptions.showModal} aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Header closeButton>
				<Modal.Title id="contained-modal-title-vcenter">{hiddenForm ? "Options" : "Report"}</Modal.Title>
			</Modal.Header>
			<Modal.Body>
				<SuccessAlert
					hidden={!storyState.storyReport.showSuccessMessage}
					header="Success"
					message={storyState.storyReport.successMessage}
					handleCloseAlert={() => dispatch({ type: storyConstants.REPORT_STORY_REQUEST })}
				/>
				<FailureAlert
					hidden={!storyState.storyReport.showError}
					header="Error"
					message={storyState.storyReport.errorMessage}
					handleCloseAlert={() => dispatch({ type: storyConstants.REPORT_STORY_REQUEST })}
				/>
				<div hidden={!hiddenForm} className="row">
					<button hidden={localStorage.getItem("userId") === null} type="button" className="btn btn-link btn-fw text-secondary w-100 border-0" onClick={handleShowForm}>
						Report
					</button>
				</div>
				<form hidden={hiddenForm} method="post" onSubmit={handleSubmit}>
					<div>
						<div className="form-group row d-flex align-items-center">
							<div className="col-sm-12 text-center">
								<label className="mr-3">I find it offensive</label>

								<input type="checkbox" onChange={() => toggleReportReason("I find it offensive")} />
							</div>
						</div>
						<hr />
						<div className="form-group row d-flex align-items-center">
							<div className="col-sm-12 text-center">
								<label className="mr-3">It's spam</label>

								<input type="checkbox" onChange={() => toggleReportReason("It's spam")} />
							</div>
						</div>
						<hr />

						<div className="form-group row d-flex align-items-center">
							<div className="col-sm-12 text-center">
								<label className="mr-3">It's sexualy inappropriate</label>

								<input type="checkbox" onChange={() => toggleReportReason("It's sexualy inappropriate")} />
							</div>
						</div>
						<hr />

						<div className="form-group row d-flex align-items-center">
							<div className="col-sm-12 text-center">
								<label className="mr-3">It's a scam or it's misleading</label>

								<input type="checkbox" onChange={() => toggleReportReason("It's a scam or it's misleading")} />
							</div>
						</div>
						<hr />

						<div className="form-group row d-flex align-items-center">
							<div className="col-sm-12 text-center">
								<label className="mr-3">It's violent or prohibited content</label>

								<input type="checkbox" onChange={() => toggleReportReason("It's violent or prohibited content")} />
							</div>
						</div>
						<hr />

						<div className="form-group row d-flex align-items-center">
							<div className="col-sm-12 text-center">
								<label className="mr-3">It refers to a political candidate or issue</label>

								<input type="checkbox" onChange={() => toggleReportReason("It refers to a political candidate or issue")} />
							</div>
						</div>
						<hr />

						<div className="form-group row d-flex align-items-center">
							<div className="col-sm-12 text-center">
								<label className="mr-3">It violates my intellectual property rights</label>

								<input type="checkbox" onChange={() => toggleReportReason("It violates my intellectual property rights")} />
							</div>
						</div>
						<div className="form-group">
							<button className="btn btn-primary float-right  mb-2" type="submit">
								Report
							</button>
						</div>
					</div>
				</form>
			</Modal.Body>
		</Modal>
	);
};

export default OptionsModalStory;
