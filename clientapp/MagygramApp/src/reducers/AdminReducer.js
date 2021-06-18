import { modalConstants } from "../constants/ModalConstants";
import { adminConstants } from "../constants/AdminConstants";


export const adminReducer = (state, action) => {
	switch (action.type) {
		case adminConstants.SHOW_VERIFICATION_REQUEST_TAB:{
            return{
                ...state,
                activeTab:{
                    verificationRequestsShow: true,
                    contentReportShow: false,
                    agentRequestsShow: false,
                },
            }
        }
        case adminConstants.SHOW_CONTENT_REPORTS_TAB:{
            return{
                ...state,
                activeTab:{
                    verificationRequestsShow: false,
                    contentReportShow: true,
                    agentRequestsShow: false,
                },
            }
        }
        case adminConstants.SHOW_AGENT_REQUEST_TAB:{
            return{
                ...state,
                activeTab:{
                    verificationRequestsShow: false,
                    contentReportShow: false,
                    agentRequestsShow: true,
                },
            }
        }
        case adminConstants.GET_PENDING_VERIFICATION_REQUEST_REQUEST:{
            let strCpy = {
				...state,
			};
            strCpy.verificationRequests.requests=[]
			return strCpy;
        }
        case adminConstants.GET_PENDING_VERIFICATION_REQUEST_SUCCESS:{
            let strCpy = {
				...state,
			};
            strCpy.verificationRequests.requests=action.requests
			return strCpy;
        }
        case adminConstants.GET_PENDING_VERIFICATION_REQUEST_FAILURE:{
            let strCpy = {
				...state,
			};
            strCpy.verificationRequests.requests=action.requests
			return strCpy;
        }
        case adminConstants.GET_REPORT_REQUEST:{
            let strCpy = {
				...state,
			};
            strCpy.reportRequests.requests=[]
			return strCpy;
        }
        case adminConstants.GET_REPORT_REQUEST_SUCCESS:{
            let strCpy = {
				...state,
			};
            strCpy.reportRequests.requests=action.requests
			return strCpy;
        }
        case adminConstants.GET_REPORT_REQUEST_FAILURE:{
            let strCpy = {
				...state,
			};
            strCpy.reportRequests.requests=action.requests
			return strCpy;
        }
        case adminConstants.APPROVE_VERIFICATION_REQUEST_REQUEST:{
            let strCpy = {
				...state,
			};
            strCpy.verificationRequests.showSuccessMessage=false
            strCpy.verificationRequests.successMessage=""
            strCpy.verificationRequests.showError=false
            strCpy.verificationRequests.errorMessage=""
			return strCpy;
        }
        case adminConstants.APPROVE_VERIFICATION_REQUEST_SUCCESS:
            let strCpy = {
				...state,
			};
			var newListOfVerificationRequests = strCpy.verificationRequests.requests.filter((request) => request.Id !== action.requestId);
			strCpy.verificationRequests.requests = newListOfVerificationRequests;
            strCpy.verificationRequests.showSuccessMessage=true
            strCpy.verificationRequests.successMessage=action.successMessage
            strCpy.verificationRequests.showError=false
            strCpy.verificationRequests.errorMessage=""
			return strCpy;
        case adminConstants.APPROVE_VERIFICATION_REQUEST_FAILURE:{
            let strCpy = {
				...state,
			};
            strCpy.verificationRequests.showSuccessMessage=false
            strCpy.verificationRequests.successMessage=""
            strCpy.verificationRequests.showError=true
            strCpy.verificationRequests.errorMessage=action.errorMessage
			return strCpy;
        }
        case adminConstants.REJECT_VERIFICATION_REQUEST_REQUEST:{
            let strCpy = {
				...state,
			};
            strCpy.verificationRequests.showSuccessMessage=false
            strCpy.verificationRequests.successMessage=""
            strCpy.verificationRequests.showError=false
            strCpy.verificationRequests.errorMessage=""
			return strCpy;
        }
        case adminConstants.REJECT_VERIFICATION_REQUEST_SUCCESS:{
            let strCpy = {
				...state,
			};
			var newListOfVerificationRequests = strCpy.verificationRequests.requests.filter((request) => request.Id !== action.requestId);
			strCpy.verificationRequests.requests = newListOfVerificationRequests;
            strCpy.verificationRequests.showSuccessMessage=true
            strCpy.verificationRequests.successMessage=action.successMessage
            strCpy.verificationRequests.showError=false
            strCpy.verificationRequests.errorMessage=""
			return strCpy;
        }
        case adminConstants.REJECT_VERIFICATION_REQUEST_FAILURE:{
            let strCpy = {
				...state,
			};
            strCpy.verificationRequests.showSuccessMessage=false
            strCpy.verificationRequests.successMessage=""
            strCpy.verificationRequests.showError=true
            strCpy.verificationRequests.errorMessage=action.errorMessage
			return strCpy;
        }
		default:
			return state;
	}
};
