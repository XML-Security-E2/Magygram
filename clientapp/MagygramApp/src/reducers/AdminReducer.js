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
		default:
			return state;
	}
};
