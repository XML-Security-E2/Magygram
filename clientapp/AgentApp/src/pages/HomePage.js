import React from "react";
import Header from "../components/Header";
import CreateProductModal from "../components/modals/CreateProductModal";
import OrderCheckoutModal from "../components/modals/OrderCheckoutModal";
import EditProductModal from "../components/modals/EditProductModal";
import ProductList from "../components/ProductList";
import OrderContextProvider from "../contexts/OrderContext";
import ProductContextProvider from "../contexts/ProductContext";
import OptionsModal from "../components/modals/OptionsModal";

const HomePage = () => {
	return (
		<div>
			<ProductContextProvider>
				<OrderContextProvider>
					<Header />
					<div>
						<div className="mt-4">
							<div className="container d-flex justify-content-center">
								<div className="col-10">
									<ProductList />
									<CreateProductModal />
									<EditProductModal />
									<OptionsModal />
								</div>
							</div>
						</div>
					</div>
					<OrderCheckoutModal />
				</OrderContextProvider>
			</ProductContextProvider>
		</div>
	);
};

export default HomePage;
