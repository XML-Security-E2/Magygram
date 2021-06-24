import { useContext, useState } from "react";
import { Modal } from "react-bootstrap";
import { modalConstants } from "../../constants/ModalConstants";
import { MessageContext } from "../../contexts/MessageContext";
import { searchService } from "../../services/SearchService";
import AsyncSelect from "react-select/async";
import { messageService } from "../../services/MessageService";

const MessageUserSearchModal = () => {
	const { messageState, dispatch } = useContext(MessageContext);
	const [search, setSearch] = useState("");
	const [selectedUserId, setSelectedUserId] = useState("");

	const handleModalClose = () => {
		dispatch({ type: modalConstants.HIDE_MESSAGE_SELECT_USER_MODAL });
	};

	const loadOptions = (value, callback) => {
		setTimeout(() => {
			searchService.userSearchUsers(value, callback);
		}, 1000);
	};

	const onInputChange = (inputValue, { action }) => {
		switch (action) {
			case "set-value":
				return;
			case "menu-close":
				setSearch("");
				return;
			case "input-change":
				setSearch(inputValue);
				return;
			default:
				return;
		}
	};

	const onChange = (option) => {
		console.log(option);
		setSelectedUserId(option.id);
		return false;
	};

	const handleStartConversation = () => {
		messageService.getMessagesFromUser(selectedUserId, dispatch);
	};

	return (
		<Modal show={messageState.selectUserModal.showModal} aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Header closeButton>
				<button type="button" className="btn btn-link text-primary" onClick={handleStartConversation}>
					Next
				</button>
			</Modal.Header>
			<Modal.Body>
				<div className="row" style={{ overflowX: "hidden", minHeight: "500px", maxHeight: "500px" }}>
					<div className="w-100 mr-3 ml-3">
						<AsyncSelect defaultOptions loadOptions={loadOptions} onInputChange={onInputChange} onChange={onChange} placeholder="search" inputValue={search} />
					</div>
				</div>
			</Modal.Body>
		</Modal>
	);
};

export default MessageUserSearchModal;
