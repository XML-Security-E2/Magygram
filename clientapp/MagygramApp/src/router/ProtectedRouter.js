import React from "react";
import { Route, Redirect } from "react-router-dom";
import { hasRoles } from "../helpers/auth-header";

export const ProtectedRoute = ({ component: Component, ...rest }) => {
	return (
		<Route
			{...rest}
			render={(props) => {
				if(rest.roles==="*"){
					return <Component {...props} />;
				}
				if (hasRoles(rest.roles) === true) {
					return <Component {...props} />;
				} else {
					return (
						<Redirect
							to={{
								pathname: rest.redirectTo,
								state: {
									from: props.location,
								},
							}}
						/>
					);
				}
			}}
		/>
	);
};
