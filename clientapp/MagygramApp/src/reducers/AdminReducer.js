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
            return{
                ...state,
                verificationRequests:{
                    requests:[]
                }
            }
        }
        case adminConstants.GET_PENDING_VERIFICATION_REQUEST_SUCCESS:{
            return{
                ...state,
                verificationRequests:{
                    requests:action.requests
                }
            }
        }
        case adminConstants.GET_PENDING_VERIFICATION_REQUEST_FAILURE:{
            return{
                ...state,
                verificationRequests:{
                    requests:[]
                },
            }
        }
        case adminConstants.APPROVE_VERIFICATION_REQUEST_REQUEST:{
            return{
                ...state,
            }
        }
        case adminConstants.APPROVE_VERIFICATION_REQUEST_SUCCESS:{
            return{
                ...state,
            }
        }
        case adminConstants.APPROVE_VERIFICATION_REQUEST_FAILURE:{
            return{
                ...state,
            }
        }
        case adminConstants.REJECT_VERIFICATION_REQUEST_REQUEST:{
            return{
                ...state,
            }
        }
        case adminConstants.REJECT_VERIFICATION_REQUEST_SUCCESS:{
            return{
                ...state,
            }
        }
        case adminConstants.REJECT_VERIFICATION_REQUEST_FAILURE:{
            return{
                ...state,
            }
        }
		default:
			return state;
	}
};
