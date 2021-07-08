import { useContext } from "react";
import { Modal } from "react-bootstrap";
import { modalConstants } from "../../constants/ModalConstants";
import { ProductContext } from "../../contexts/ProductContext";
import { Link } from "react-router-dom";

const OptionsModal = () => {
	const { productState, dispatch } = useContext(ProductContext);

	const handleModalClose = () => {
		dispatch({ type: modalConstants.HIDE_OPTIONS_MODAL });
	};

	return (
		<Modal show={productState.optionsModal.showModal} aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Body>
				<div>
					<div className="row">
						<Link type="button" className="btn btn-link btn-fw text-secondary w-100 border-0" to={"/create-campaign?productId=" + productState.optionsModal.productId}>
							Create Campaign
						</Link>
					</div>
				</div>
			</Modal.Body>
		</Modal>
	);
};

export default OptionsModal;
