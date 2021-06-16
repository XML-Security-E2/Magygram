const ShoppingCartItem = ({ id, name, count, price, imageUrl, setProductCount, deleteFromShoppingCart }) => {
	return (
		<div className="row align-items-center" style={{ cursor: "pointer" }}>
			<div className="col-12 col-md-8">
				<div className="ml-2">
					<h4>{name}</h4>

					<div className="row align-items-center">
						<input type="number" className="form-control col-sm-3 ml-3" id="quantity" min="1" value={count} onChange={(e) => setProductCount(id, e.target.value)} />x{" "}
						<span style={{ color: "#198ae3" }}>{Number(price).toFixed(2)} RSD</span>
					</div>
				</div>
			</div>
			<div className="col-12 col-md-4">
				<div className="mr-2">
					<div className=" container-img">
						<img src={imageUrl} className="img-fluid rounded" alt="" />
						<div className="overlay-img">
							<button className="btn  shopp-icon-img" data-toggle="tooltip" title="Delete product" onClick={() => deleteFromShoppingCart(id)}>
								<i className="mdi mdi-close"></i>
							</button>
						</div>
					</div>
				</div>
			</div>
		</div>
	);
};

export default ShoppingCartItem;
