import { useContext } from "react";
import { Button, Modal } from "react-bootstrap";
import { modalConstants } from "../../constants/ModalConstants";
import { productConstants } from "../../constants/ProductConstants";
import { ProductContext } from "../../contexts/ProductContext";
import CreateProductForm from "../CreateProductForm";
import FailureAlert from "../FailureAlert";

const CreateProductModal = () => {
	const { productState, dispatch } = useContext(ProductContext);

	const handleModalClose = () => {
		dispatch({ type: modalConstants.HIDE_CREATE_PRODUCT_MODAL });
	};

	return (
		<Modal show={productState.createProduct.showModal} aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Header closeButton>
				<Modal.Title id="contained-modal-title-vcenter">
					<big>Create product</big>
				</Modal.Title>
			</Modal.Header>
			<Modal.Body>
				<FailureAlert
					hidden={!productState.createProduct.showErrorMessage}
					header="Error"
					message={productState.createProduct.errorMessage}
					handleCloseAlert={() => dispatch({ type: productConstants.CREATE_PRODUCT_MODAL_HIDE_ERROR })}
				/>
				<CreateProductForm />
			</Modal.Body>
			<Modal.Footer>
				<Button onClick={handleModalClose}>Close</Button>
			</Modal.Footer>
		</Modal>
	);
};

export default CreateProductModal;
