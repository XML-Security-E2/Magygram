import React, { useContext, useEffect } from "react";
import { modalConstants } from "../constants/ModalConstants";
import { ProductContext } from "../contexts/ProductContext";
import { productService } from "../services/ProductService";
import ProductItem from "./ProductItem";

const ProductList = () => {
	const { productState, dispatch } = useContext(ProductContext);

	useEffect(() => {
		const getProductsHandler = async () => {
			await productService.findAllProducts(dispatch);
		};
		getProductsHandler();
	}, [dispatch]);

	const handleOpenCreateProductsModal = () => {
		dispatch({ type: modalConstants.SHOW_CREATE_PRODUCT_MODAL });
	};

	const getUserHandler = async (id) => {
		await productService.findById(id, dispatch);
	};

	const handleEditProducts = (id) => {
		getUserHandler(id);
		dispatch({ type: modalConstants.SHOW_EDIT_PRODUCT_MODAL });
	};

	const handleDeleteProducts = (id) => {
		productService.deleteProduct(id, dispatch);
	};

	return (
		<React.Fragment>
			<button type="button" className="btn btn-primary row" onClick={handleOpenCreateProductsModal}>
				+ Create product
			</button>
			<div className="content-wrapper mt-4">
				<div className="row">
					{productState.listProducts.products.map((product) => {
						return (
							<ProductItem
								id={product.id}
								imagePath={product.imageUrl}
								key={product.id}
								name={product.name}
								price={product.price}
								handleEditProducts={handleEditProducts}
								handleDeleteProducts={handleDeleteProducts}
							/>
						);
					})}
				</div>
			</div>
		</React.Fragment>
	);
};

export default ProductList;
