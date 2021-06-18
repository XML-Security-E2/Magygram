import { useContext } from "react";
import { Modal } from "react-bootstrap";
import { modalConstants } from "../../constants/ModalConstants";
import { notificationConstants } from "../../constants/NotificationConstants";
import { NotificationContext } from "../../contexts/NotificationContext";
import NotificationSettingsForm from "../NotificationSettingsForm";
import SuccessAlert from "../SuccessAlert";

const NotificationSettingsModal = () => {
	const { notificationState, dispatch } = useContext(NotificationContext);

	const handleModalClose = () => {
		dispatch({ type: modalConstants.HIDE_NOTIFICATION_SETTINGS_MODAL });
	};

	return (
		<Modal show={notificationState.notificationSettingsModal.showModal} aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Header closeButton>
				<Modal.Title id="contained-modal-title-vcenter">Notification settings</Modal.Title>
			</Modal.Header>
			<Modal.Body>
				<SuccessAlert
					hidden={!notificationState.notificationSettingsModal.showSuccessMessage}
					header="Success"
					message={notificationState.notificationSettingsModal.successMessage}
					handleCloseAlert={() => dispatch({ type: notificationConstants.HIDE_NOTIFICATION_SETTINGS_SUCCESS_MESSAGE })}
				/>
				<NotificationSettingsForm />
			</Modal.Body>
		</Modal>
	);
};

export default NotificationSettingsModal;
