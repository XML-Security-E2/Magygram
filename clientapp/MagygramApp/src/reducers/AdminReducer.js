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
        case adminConstants.GET_PENDING_AGENT_REGISTRATION_REQUEST_REQUEST:{
            let strCpy = {
				...state,
			};
            strCpy.agentRegistrationRequests.requests=[]
			return strCpy;
        }
        case adminConstants.GET_PENDING_AGENT_REGISTRATION_REQUEST_SUCCESS:{
            let strCpy = {
				...state,
			};
            strCpy.agentRegistrationRequests.requests=action.requests
			return strCpy;
        }
        case adminConstants.GET_PENDING_AGENT_REGISTRATION_REQUEST_FAILURE:{
            let strCpy = {
				...state,
			};
            strCpy.agentRegistrationRequests.requests=[]
			return strCpy;
        }
        case adminConstants.APPROVE_AGENT_REGISTRATION_REQUEST_REQUEST:{
            let strCpy = {
				...state,
			};
            strCpy.agentRegistrationRequests.showSuccessMessage=false
            strCpy.agentRegistrationRequests.successMessage=""
            strCpy.agentRegistrationRequests.showError=false
            strCpy.agentRegistrationRequests.errorMessage=""
			return strCpy;
        }
        case adminConstants.APPROVE_AGENT_REGISTRATION_REQUEST_SUCCESS:{
            let strCpy = {
				...state,
			};
			var newListOfVerificationRequests = strCpy.agentRegistrationRequests.requests.filter((request) => request.Id !== action.requestId);
			strCpy.agentRegistrationRequests.requests = newListOfVerificationRequests;
            strCpy.agentRegistrationRequests.showSuccessMessage=true
            strCpy.agentRegistrationRequests.successMessage=action.successMessage
            strCpy.agentRegistrationRequests.showError=false
            strCpy.agentRegistrationRequests.errorMessage=""
			return strCpy;
        }
        case adminConstants.APPROVE_AGENT_REGISTRATION_REQUEST_FAILURE:{
            let strCpy = {
				...state,
			};
            strCpy.agentRegistrationRequests.showSuccessMessage=false
            strCpy.agentRegistrationRequests.successMessage=""
            strCpy.agentRegistrationRequests.showError=true
            strCpy.agentRegistrationRequests.errorMessage=action.errorMessage
			return strCpy;
        }
        case adminConstants.REJECT_AGENT_REGISTRATION_REQUEST_REQUEST:{
            let strCpy = {
				...state,
			};
            strCpy.agentRegistrationRequests.showSuccessMessage=false
            strCpy.agentRegistrationRequests.successMessage=""
            strCpy.agentRegistrationRequests.showError=false
            strCpy.agentRegistrationRequests.errorMessage=""
			return strCpy;
        }
        case adminConstants.REJECT_AGENT_REGISTRATION_REQUEST_SUCCESS:{
            let strCpy = {
				...state,
			};
			var newListOfVerificationRequests = strCpy.agentRegistrationRequests.requests.filter((request) => request.Id !== action.requestId);
			strCpy.agentRegistrationRequests.requests = newListOfVerificationRequests;
            strCpy.agentRegistrationRequests.showSuccessMessage=true
            strCpy.agentRegistrationRequests.successMessage=action.successMessage
            strCpy.agentRegistrationRequests.showError=false
            strCpy.agentRegistrationRequests.errorMessage=""
			return strCpy;
        }
        case adminConstants.REJECT_AGENT_REGISTRATION_REQUEST_FAILURE:{
            let strCpy = {
				...state,
			};
            strCpy.agentRegistrationRequests.showSuccessMessage=false
            strCpy.agentRegistrationRequests.successMessage=""
            strCpy.agentRegistrationRequests.showError=true
            strCpy.agentRegistrationRequests.errorMessage=action.errorMessage
			return strCpy;
        }
        case modalConstants.SHOW_REGISTER_AGENT_MODAL:{
            return{
                ...state,
                registerNewAgent:{
                    showModal: true,
                    showError:false,
                    errorMessage:"",
                    showSuccessMessage: false,
			        successMessage: "",
                },
            }
        }
        case modalConstants.HIDE_REGISTER_AGENT_MODAL:{
            return{
                ...state,
                registerNewAgent:{
                    showModal: false,
                    showError:false,
                    errorMessage:"",
                    showSuccessMessage: false,
			        successMessage: "",
                },
            }
        }
        case adminConstants.AGENT_REGISTRATION_VALIDATION_FAILURE:{
            return{
                ...state,
                registerNewAgent:{
                    showModal: true,
                    showError:true,
                    errorMessage:action.errorMessage,
                    showSuccessMessage: false,
			        successMessage: "",
                },
            }
        }
        case adminConstants.AGENT_REGISTRATION_REQUEST:{
            return{
                ...state,
                registerNewAgent:{
                    showModal: true,
                    showError:false,
                    errorMessage:"",
                    showSuccessMessage: false,
			        successMessage: "",
                },
            }
        }
        case adminConstants.AGENT_REGISTRATION_SUCCESS:{
            return{
                ...state,
                registerNewAgent:{
                    showModal: false,
                    showError:false,
                    errorMessage:"",
                    showSuccessMessage: true,
			        successMessage: action.message,
                },
            }
        }
        case adminConstants.AGENT_REGISTRATION_FAILURE:{
            return{
                ...state,
                registerNewAgent:{
                    showModal: true,
                    showError:true,
                    errorMessage:action.errorMessage,
                    showSuccessMessage: false,
			        successMessage: "",
                },
            }
        }
		default:
			return state;
	}
};
