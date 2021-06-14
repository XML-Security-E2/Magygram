const ProductItem = ({ id, imagePath, name, price, handleEditProducts, handleDeleteProducts }) => {
	return (
		<div className="col-md-4 col-lg-3 rounded border img-fluid container-img mt-3">
			<div>
				<img src={imagePath} className="img-fluid rounded-lg w-100" alt="" />
				<h5 className="mt-2">{name}</h5>
				<h5 className="float-right" style={{ color: "#198ae3" }}>
					RSD {Number(price).toFixed(2)}
				</h5>
			</div>
			<div className="overlay-img">
				<button className="btn icon-img" data-toggle="tooltip" title="Edit product">
					<i className="mdi mdi-pencil" onClick={() => handleEditProducts(id)}></i>
				</button>
				<button className="btn btn-danger float-right" data-toggle="tooltip" title="Delete product" onClick={() => handleDeleteProducts(id)}>
					<i className="mdi mdi-close"></i>
				</button>
			</div>
		</div>
	);
};

export default ProductItem;
