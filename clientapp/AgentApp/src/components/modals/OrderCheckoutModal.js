import { useContext } from "react";
import { Button, Modal } from "react-bootstrap";
import { modalConstants } from "../../constants/ModalConstants";
import { orderConstants } from "../../constants/OrderConstants";
import { OrderContext } from "../../contexts/OrderContext";
import FailureAlert from "../FailureAlert";
import SuccessAlert from "../SuccessAlert";
import ShoppingCartList from "../ShoppingCartList";
import CreateOrderForm from "../CreateOrderForm";

const OrderCheckoutModal = () => {
	const { orderState, dispatch } = useContext(OrderContext);

	const handleModalClose = () => {
		dispatch({ type: modalConstants.HIDE_ORDER_CHECKOUT_MODAL });
	};

	const getOrderSum = () => {
		let sum = 0;
		orderState.shoppingCart.items.forEach((item) => {
			sum += item.count * item.price;
		});
		return sum;
	};

	return (
		<Modal show={orderState.orderCheckout.showModal} aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Header closeButton>
				<Modal.Title id="contained-modal-title-vcenter">
					<big>Checkout</big>
				</Modal.Title>
			</Modal.Header>
			<Modal.Body>
				<FailureAlert
					hidden={!orderState.orderCheckout.showErrorMessage}
					header="Error"
					message={orderState.orderCheckout.errorMessage}
					handleCloseAlert={() => dispatch({ type: orderConstants.CHECKOUT_MODAL_HIDE_MESSAGE })}
				/>
				<SuccessAlert
					hidden={!orderState.orderCheckout.showSuccessMessage}
					header="Success"
					message={orderState.orderCheckout.successMessage}
					handleCloseAlert={() => dispatch({ type: orderConstants.CHECKOUT_MODAL_HIDE_MESSAGE })}
				/>
				<ShoppingCartList />
				<div className="row d-flex justify-content-end">
					<div className="mr-4">
						<b>Total:</b>
						<span style={{ color: "#198ae3" }} className="ml-2">
							<b>{Number(getOrderSum()).toFixed(2)} RSD</b>
						</span>
					</div>
				</div>
				<hr />
				<CreateOrderForm />
			</Modal.Body>
			<Modal.Footer>
				<Button onClick={handleModalClose}>Close</Button>
			</Modal.Footer>
		</Modal>
	);
};

export default OrderCheckoutModal;
