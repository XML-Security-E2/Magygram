import React, { createContext, useReducer } from "react";
import { adminReducer } from "../reducers/AdminReducer";

export const AdminContext = createContext();

const AdminContextProvider = (props) => {
	const [state, dispatch] = useReducer(adminReducer, {
		activeTab:{
            verificationRequestsShow: true,
            contentReportShow: false,
            agentRequestsShow: false,
        },
        verificationRequests:{
            requests:[]
        }
	});

	return <AdminContext.Provider value={{ state, dispatch }}>{props.children}</AdminContext.Provider>;
};

export default AdminContextProvider;
