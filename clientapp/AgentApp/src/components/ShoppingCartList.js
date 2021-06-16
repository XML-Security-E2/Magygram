import React, { useContext } from "react";
import { orderConstants } from "../constants/OrderConstants";
import { OrderContext } from "../contexts/OrderContext";
import ShoppingCartItem from "./ShoppingCartItem";

const ShoppingCartList = () => {
	const { orderState, dispatch } = useContext(OrderContext);

	const setProductCount = (id, count) => {
		dispatch({ type: orderConstants.SET_PRODUCT_COUNT_TO_ORDER, id, count });
	};

	const deleteFromShoppingCart = (id) => {
		dispatch({ type: orderConstants.REMOVE_PRODUCT_FROM_ORDER, id });
	};

	return (
		<React.Fragment>
			{orderState.shoppingCart.items.map((item) => {
				console.log(item);
				return (
					<React.Fragment>
						<ShoppingCartItem
							id={item.id}
							key={item.id}
							name={item.name}
							count={item.count}
							price={item.price}
							imageUrl={item.imagePath}
							setProductCount={setProductCount}
							deleteFromShoppingCart={deleteFromShoppingCart}
						/>
						<hr />
					</React.Fragment>
				);
			})}
		</React.Fragment>
	);
};

export default ShoppingCartList;
