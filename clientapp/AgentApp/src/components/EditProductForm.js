import React, { useContext, useEffect, useState } from "react";
import { productConstants } from "../constants/ProductConstants";
import { ProductContext } from "../contexts/ProductContext";
import { productService } from "../services/ProductService";

const EditProductForm = () => {
	const { productState, dispatch } = useContext(ProductContext);

	const [image, setImage] = useState("");
	const imgRef = React.createRef();

	const [id, setId] = useState("");
	const [name, setName] = useState("");
	const [price, setPrice] = useState("");
	const [quantity, setQuantity] = useState("");

	const onImageChange = (e) => {
		setImage(e.target.files[0]);

		if (e.target.files && e.target.files[0]) {
			let img = e.target.files[0];
			dispatch({ type: productConstants.EDIT_PRODUCT_IMAGE_SELECTED, showedImage: URL.createObjectURL(img) });
		}
	};

	const handleImageDeselect = () => {
		dispatch({ type: productConstants.EDIT_PRODUCT_IMAGE_DESELECTED });
	};

	const handleSubmit = (e) => {
		e.preventDefault();

		let product = { name, price: parseFloat(price), quantity: parseInt(quantity) };
		console.log(product);
		productService.updateProductInfo(product, id, dispatch);
	};

	const handleImageChange = () => {
		productService.updateProductImage(image, id, dispatch);
	};

	useEffect(() => {
		setId(productState.updateProduct.product.id);
		setName(productState.updateProduct.product.name);
		setPrice(productState.updateProduct.product.price);
		setQuantity(productState.updateProduct.product.quantity);
	}, [productState.updateProduct.product]);

	return (
		<React.Fragment>
			<form className="forms-sample" method="post" onSubmit={handleSubmit}>
				<div className="form-group">
					{productState.updateProduct.showedImage !== "" && <img src={productState.updateProduct.showedImage} alt="product" className="img-fluid" />}
					<input type="file" ref={imgRef} style={{ display: "none" }} name="image" accept="image/png, image/jpeg" onChange={onImageChange} />
					<div className="row d-flex flex-row-reverse mt-1">
						<button hidden={!productState.updateProduct.imageSelected} type="button" onClick={handleImageChange} className="btn btn-outline-secondary btn-icon-text border-0 mr-3">
							<i className="mdi mdi-file-check btn-icon-prepend"></i> Submit
						</button>
						<button hidden={!productState.updateProduct.imageSelected} type="button" onClick={handleImageDeselect} className="btn btn-outline-danger btn-icon-text border-0">
							Remove<i className="mdi mdi-close ml-1 align-middle"></i>
						</button>
						<button hidden={productState.updateProduct.imageSelected} type="button" onClick={() => imgRef.current.click()} className="btn btn-outline-primary btn-icon-text border-0">
							<i className="mdi mdi-upload btn-icon-prepend"></i> Upload image
						</button>
					</div>
				</div>
				<div className="form-group">
					<label for="name">Name</label>
					<input type="text" required className="form-control" value={name} id="name" placeholder="Name" onChange={(e) => setName(e.target.value)} />
				</div>
				<div className="form-group">
					<label for="price">Price</label>
					<input type="number" step={0.01} required value={price} className="form-control" id="price" min="1" placeholder="Price" onChange={(e) => setPrice(e.target.value)} />
				</div>
				<div className="form-group">
					<label for="quantity">Quantity</label>
					<input type="number" required className="form-control" value={quantity} id="quantity" min="1" placeholder="Quantity" onChange={(e) => setQuantity(e.target.value)} />
				</div>
				<button type="submit" className="btn btn-primary mr-2">
					Submit
				</button>
			</form>
		</React.Fragment>
	);
};

export default EditProductForm;
