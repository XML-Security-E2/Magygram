import { useContext } from "react";
import { modalConstants } from "../constants/ModalConstants";
import { OrderContext } from "../contexts/OrderContext";
import { hasRoles } from "../helpers/auth-header";
import { productService } from "../services/ProductService";
import { userService } from "../services/UserService";
import ShoppingCartList from "./ShoppingCartList";

const Header = () => {
	const { orderState, dispatch } = useContext(OrderContext);

	const navStyle = { height: "50px", borderBottom: "1px solid rgb(200,200,200)" };
	const iconStyle = { fontSize: "30px", margin: "0px", marginLeft: "13px" };

	const openCheckoutModal = () => {
		dispatch({ type: modalConstants.SHOW_ORDER_CHECKOUT_MODAL });
	};

	const getOrderSum = () => {
		let sum = 0;
		orderState.shoppingCart.items.forEach((item) => {
			sum += item.count * item.price;
		});
		return sum;
	};

	const backToHome = () => {
		window.location = "#/";
	};

	const getStats = () => {
		window.location = "#/campaign-stats";
	};

	return (
		<nav className="navbar navbar-light navbar-expand-md navigation-clean" style={navStyle}>
			<div className="container">
				<div>
					<img onClick={() => backToHome()} src="assets/img/logotest.png" alt="NistagramLogo" />
				</div>
				<button className="navbar-toggler" data-toggle="collapse">
					<span className="sr-only">Toggle navigation</span>
					<span className="navbar-toggler-icon"></span>
				</button>
				<div className="d-flex align-items-center">
					<div className="dropdown" hidden={hasRoles(["admin", "agent"])}>
						<i className="fa fa-shopping-cart" style={iconStyle} id="dropdownMenu2" data-toggle="dropdown" />
						<span className="ml-1 bg-primary rounded text-white pl-1 pr-1">{orderState.shoppingCart.items.length}</span>
						<ul style={{ width: "300px", marginLeft: "15px", minWidth: "350px" }} className="dropdown-menu dropdown-menu-right" aria-labelledby="dropdownMenu2">
							<li className="mb-3">
								<h4 className="ml-2">Shopping cart</h4>
							</li>
							<hr />
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
							<div className="row d-flex justify-content-end">
								<button disabled={orderState.shoppingCart.items.length === 0} type="button" className="btn btn-primary mr-4" onClick={openCheckoutModal}>
									Checkout
								</button>
							</div>
						</ul>
					</div>

					<i className="fa fa-home" style={iconStyle} />
					<div className="dropdown" hidden={!hasRoles(["*"])}>
						<i className="fa fa-user" style={iconStyle} id="dropdownMenu1" data-toggle="dropdown" aria-haspopup="true" />

						<ul style={{ width: "200px", marginLeft: "15px" }} className="dropdown-menu" aria-labelledby="dropdownMenu1">
							<li>
								<button className=" btn shadow-none" onClick={getStats}>
									Stats
								</button>
							</li>
							<li>
								<button className=" btn shadow-none" onClick={() => userService.logout()}>
									Logout
								</button>
							</li>
						</ul>
					</div>
				</div>
			</div>
		</nav>
	);
};

export default Header;
