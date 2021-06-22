import { useContext, useState } from "react";
import { Modal } from "react-bootstrap";
import { modalConstants } from "../../constants/ModalConstants";
import { userConstants } from "../../constants/UserConstants";
import { UserContext } from "../../contexts/UserContext";
import { userService } from "../../services/UserService";
import FailureAlert from "../FailureAlert";
import SuccessAlert from "../SuccessAlert";

const ReportUserModal = ({ userId }) => {
	const { userState, dispatch } = useContext(UserContext);

	const [reportReasons, setReportReasons] = useState([]);

	const handleModalClose = () => {
		dispatch({ type: modalConstants.HIDE_USER_REPORT_MODAL });
		setReportReasons([]);
	};

	const handleSubmit = (e) => {
		e.preventDefault();

		let reportDTO = {
			contentId: userId,
			contentType: "USER",
			reportReasons,
		};

		userService.reportUser(reportDTO, dispatch);
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

	return (
		<Modal show={userState.userReport.showModal} aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Header closeButton>
				<Modal.Title id="contained-modal-title-vcenter">Report</Modal.Title>
			</Modal.Header>
			<Modal.Body>
				<SuccessAlert
					hidden={!userState.userReport.showSuccessMessage}
					header="Success"
					message={userState.userReport.successMessage}
					handleCloseAlert={() => dispatch({ type: userConstants.REPORT_USER_REQUEST })}
				/>
				<FailureAlert
					hidden={!userState.userReport.showError}
					header="Error"
					message={userState.userReport.errorMessage}
					handleCloseAlert={() => dispatch({ type: userConstants.REPORT_USER_REQUEST })}
				/>
				<form method="post" onSubmit={handleSubmit}>
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

export default ReportUserModal;
