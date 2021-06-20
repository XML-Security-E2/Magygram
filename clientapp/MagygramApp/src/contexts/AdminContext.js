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
            requests:[],
            showError: false,
			errorMessage: "",
			showSuccessMessage: false,
			successMessage: "",
        },
        reportRequests:{
            requests:[],
            showError: false,
			errorMessage: "",
			showSuccessMessage: false,
			successMessage: "",
        },
        agentRegistrationRequests:{
            requests:[],
            showError: false,
			errorMessage: "",
			showSuccessMessage: false,
			successMessage: "",
        }
	});

	return <AdminContext.Provider value={{ state, dispatch }}>{props.children}</AdminContext.Provider>;
};

export default AdminContextProvider;
