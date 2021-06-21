import { useContext, useState } from "react";
import { Modal } from "react-bootstrap";
import { modalConstants } from "../../constants/ModalConstants";
import { PostContext } from "../../contexts/PostContext";
import Axios from "axios";
import { authHeader } from "../../helpers/auth-header";
import { hasRoles } from "../../helpers/auth-header";
import { postService } from "../../services/PostService";
import { postConstants } from "../../constants/PostConstants";
import SuccessAlert from "../SuccessAlert";
import FailureAlert from "../FailureAlert";

const OptionsModal = () => {
	const { postState, dispatch } = useContext(PostContext);

	const [reportReasons, setReportReasons] = useState([]);
	const [hiddenForm, setHiddenForm] = useState(true);

	const handleShowForm = () => {
		setHiddenForm(false);
	};

	const handleSubmit = (e) => {
		e.preventDefault();

		let reportDTO = {
			contentId: postState.editPost.post.id,
			contentType: "POST",
			reportReasons,
		};

		postService.reportPost(reportDTO, dispatch);
	};

	const toggleReportReason = (reason) => {
		let a = [...reportReasons];

		if (a.find((col) => col === reason) === undefined) {
			a.push(reason);
		} else {
			a = reportReasons.filter((reas) => reas !== reason);
		}
		setReportReasons(a);
		console.log(a);
	};

	const handleModalClose = () => {
		dispatch({ type: modalConstants.HIDE_POST_OPTIONS_MODAL });
		setHiddenForm(true);
		setReportReasons([]);
	};

	const handleOpenPostEditModal = () => {
		dispatch({ type: modalConstants.SHOW_POST_EDIT_MODAL });
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

	return (
		<Modal show={postState.postOptions.showModal} aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Header closeButton>
				<Modal.Title id="contained-modal-title-vcenter">{hiddenForm ? "Options" : "Report"}</Modal.Title>
			</Modal.Header>
			<Modal.Body>
				<SuccessAlert
					hidden={!postState.postReport.showSuccessMessage}
					header="Success"
					message={postState.postReport.successMessage}
					handleCloseAlert={() => dispatch({ type: postConstants.REPORT_POST_REQUEST })}
				/>
				<FailureAlert
					hidden={!postState.postReport.showError}
					header="Error"
					message={postState.postReport.errorMessage}
					handleCloseAlert={() => dispatch({ type: postConstants.REPORT_POST_REQUEST })}
				/>
				<div hidden={!hiddenForm}>
					<div className="row">
						<button
							hidden={localStorage.getItem("userId") !== postState.editPost.post.userId}
							type="button"
							className="btn btn-link btn-fw text-secondary w-100 border-0"
							onClick={handleOpenPostEditModal}
						>
							Edit
						</button>
					</div>
					<div className="row">
						<button
							hidden={localStorage.getItem("userId") === null || localStorage.getItem("userId") === postState.editPost.post.userId}
							type="button"
							className="btn btn-link btn-fw text-secondary w-100 border-0"
							onClick={handleShowForm}
						>
							Report
						</button>
					</div>
					<div className="row">
						<button hidden={!hasRoles(["admin"])} type="button" className="btn btn-link btn-fw text-danger w-100 border-0" onClick={handleDelete}>
							Delete?
						</button>
					</div>
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

export default OptionsModal;
