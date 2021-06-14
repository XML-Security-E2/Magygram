import React from "react";
import Header from "../components/Header";
import CreateProductModal from "../components/modals/CreateProductModal";
import EditProductModal from "../components/modals/EditProductModal";
import ProductList from "../components/ProductList";
import ProductContextProvider from "../contexts/ProductContext";

const HomePage = () => {
	return (
		<div>
			<ProductContextProvider>
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
			</ProductContextProvider>
		</div>
	);
};

export default HomePage;
