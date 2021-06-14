import React, { useContext, useState } from "react";
import { productConstants } from "../constants/ProductConstants";
import { ProductContext } from "../contexts/ProductContext";
import { productService } from "../services/ProductService";

const CreateProductForm = () => {
	const { productState, dispatch } = useContext(ProductContext);

	const [image, setImage] = useState("");
	const imgRef = React.createRef();

	const [name, setName] = useState("");
	const [price, setPrice] = useState("");

	const onImageChange = (e) => {
		setImage(e.target.files[0]);

		if (e.target.files && e.target.files[0]) {
			let img = e.target.files[0];
			dispatch({ type: productConstants.CREATE_PRODUCT_IMAGE_SELECTED, showedImage: URL.createObjectURL(img) });
		}
	};

	const handleImageDeselect = () => {
		dispatch({ type: productConstants.CREATE_PRODUCT_IMAGE_DESELECTED });
	};

	const handleSubmit = (e) => {
		e.preventDefault();

		let product = { name, price, image };
		productService.createProduct(product, dispatch);
	};

	return (
		<React.Fragment>
			<form className="forms-sample" method="post" onSubmit={handleSubmit}>
				<div className="form-group">
					{productState.createProduct.showedImage !== "" && <img src={productState.createProduct.showedImage} alt="product" className="img-fluid" />}
					<input type="file" ref={imgRef} style={{ display: "none" }} name="image" accept="image/png, image/jpeg" onChange={onImageChange} />
					<div className="row d-flex flex-row-reverse">
						<button hidden={!productState.createProduct.imageSelected} type="button" onClick={handleImageDeselect} className="btn btn-outline-danger btn-icon-text border-0  mt-4">
							Remove<i className="mdi mdi-close ml-1 align-middle"></i>
						</button>
						<button hidden={productState.createProduct.imageSelected} type="button" onClick={() => imgRef.current.click()} className="btn btn-outline-primary btn-icon-text border-0">
							<i className="mdi mdi-upload btn-icon-prepend"></i> Upload image
						</button>
					</div>
				</div>
				<div className="form-group">
					<label for="name">Name</label>
					<input type="text" required className="form-control" id="name" placeholder="Name" onChange={(e) => setName(e.target.value)} />
				</div>
				<div className="form-group">
					<label for="price">Price</label>
					<input type="number" step={0.01} required className="form-control" id="price" min="1" placeholder="Price" onChange={(e) => setPrice(e.target.value)} />
				</div>
				<button type="submit" className="btn btn-primary mr-2">
					Submit
				</button>
			</form>
		</React.Fragment>
	);
};

export default CreateProductForm;
