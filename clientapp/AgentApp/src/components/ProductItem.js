import { useContext, useState } from "react";
import { modalConstants } from "../constants/ModalConstants";
import { ProductContext } from "../contexts/ProductContext";
import { hasRoles } from "../helpers/auth-header";

const ProductItem = ({ id, imagePath, name, price, handleEditProducts, handleDeleteProducts, handleAddToCart }) => {
	const { dispatch } = useContext(ProductContext);

	const [count, setCount] = useState(1);

	const handleOpenOptionsModal = () => {
		dispatch({ type: modalConstants.SHOW_OPTIONS_MODAL, productId: id });
	};
	return (
		<div className="col-md-4 col-lg-3 rounded border img-fluid container-img mt-3">
			<div>
				{hasRoles(["agent"]) && (
					<div className="row d-flex justify-content-end">
						<button className="btn float-right p-0 mr-3" onClick={handleOpenOptionsModal}>
							<i className="fa fa-ellipsis-h" aria-hidden="true"></i>
						</button>
					</div>
				)}
				<img src={imagePath} className="img-fluid rounded-lg w-100 mt-1" alt="" />
				<div className="row">
					<h5 className="mt-2 ml-3">{name}</h5>
				</div>
				<div className="row d-flex justify-content-end">
					<h5 className="mr-3" style={{ color: "#198ae3" }}>
						RSD {Number(price).toFixed(2)}
					</h5>
				</div>

				<hr className="mt-2" hidden={hasRoles(["admin"])} />

				<div hidden={hasRoles(["admin", "agent"])} className={hasRoles(["admin", "agent"]) ? "" : "ml-1 mb-2 row w-100 d-flex justify-content-between"}>
					<input type="number" required className="form-control col-sm-4" id="quantity" min="1" value={count} onChange={(e) => setCount(e.target.value)} />
					<button
						type="button"
						className="btn btn-primary btn-rounded btn-icon col-sm-3"
						data-toggle="tooltip"
						title="Add to cart"
						onClick={() => {
							handleAddToCart({ id, imagePath, name, price, count });
							setCount(1);
						}}
					>
						<i className="mdi mdi-plus"></i>
					</button>
				</div>
			</div>
			<div hidden={!hasRoles(["admin"])} className="overlay-img">
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
