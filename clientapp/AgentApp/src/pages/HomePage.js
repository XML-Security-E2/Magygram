import React from "react";
import Header from "../components/Header";
import CreateProductModal from "../components/modals/CreateProductModal";
import OrderCheckoutModal from "../components/modals/OrderCheckoutModal";
import EditProductModal from "../components/modals/EditProductModal";
import ProductList from "../components/ProductList";
import OrderContextProvider from "../contexts/OrderContext";
import ProductContextProvider from "../contexts/ProductContext";

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
