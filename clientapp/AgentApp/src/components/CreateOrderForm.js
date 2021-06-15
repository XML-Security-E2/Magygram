import { useContext, useState } from "react";
import { OrderContext } from "../contexts/OrderContext";
import { orderService } from "../services/OrderService";

const CreateOrderForm = () => {
	const { orderState, dispatch } = useContext(OrderContext);

	const [address, setAddress] = useState("");

	const handleSubmit = (e) => {
		e.preventDefault();

		let items = [];
		orderState.shoppingCart.items.forEach((item) => {
			items.push({ productId: item.id, count: parseInt(item.count) });
		});
		let order = { address, items };

		orderService.createOrder(order, dispatch);
	};
	return (
		<form className="forms-sample" method="post" onSubmit={handleSubmit}>
			<div className="form-group">
				<label for="address">Address</label>
				<input type="text" required className="form-control" id="address" placeholder="Address" onChange={(e) => setAddress(e.target.value)} />
			</div>
			<button type="submit" className="btn btn-primary mr-2">
				Submit
			</button>
		</form>
	);
};

export default CreateOrderForm;
