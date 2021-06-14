import { useContext } from "react";
import { Button, Modal } from "react-bootstrap";
import { modalConstants } from "../../constants/ModalConstants";
import { productConstants } from "../../constants/ProductConstants";
import { ProductContext } from "../../contexts/ProductContext";
import EditProductForm from "../EditProductForm";
import FailureAlert from "../FailureAlert";

const EditProductModal = () => {
	const { productState, dispatch } = useContext(ProductContext);

	const handleModalClose = () => {
		dispatch({ type: modalConstants.HIDE_EDIT_PRODUCT_MODAL });
	};

	return (
		<Modal show={productState.updateProduct.showModal} aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Header closeButton>
				<Modal.Title id="contained-modal-title-vcenter">
					<big>Edit product</big>
				</Modal.Title>
			</Modal.Header>
			<Modal.Body>
				<FailureAlert
					hidden={!productState.updateProduct.showErrorMessage}
					header="Error"
					message={productState.updateProduct.errorMessage}
					handleCloseAlert={() => dispatch({ type: productConstants.EDIT_PRODUCT_MODAL_HIDE_ERROR })}
				/>
				<EditProductForm />
			</Modal.Body>
			<Modal.Footer>
				<Button onClick={handleModalClose}>Close</Button>
			</Modal.Footer>
		</Modal>
	);
};

export default EditProductModal;
