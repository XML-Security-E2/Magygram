import React, { createContext, useReducer } from "react";
import { userReducer } from "../reducers/UserReducer";

export const UserContext = createContext();

const UserContextProvider = (props) => {
	const [users, dispatch] = useReducer(userReducer, []);

	return <UserContext.Provider value={{ users, dispatch }}>{props.children}</UserContext.Provider>;
};

export default UserContextProvider;
