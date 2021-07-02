import { useContext, useState } from "react";
import { Modal } from "react-bootstrap";
import { modalConstants } from "../../constants/ModalConstants";
import { PostContext } from "../../contexts/PostContext";
import Axios from "axios";
import { authHeader } from "../../helpers/auth-header";
import { hasRoles } from "../../helpers/auth-header";
import { searchService } from "../../services/SearchService";
import { postService } from "../../services/PostService";
import AsyncSelect from "react-select/async";

const SearchInfluencerModal = () => {
	const { postState, dispatch } = useContext(PostContext);

	const handleModalClose = () => {
		dispatch({ type: modalConstants.HIDE_SEARCH_INFLUENCER_MODAL });
	};


	const [location, setLocation] = useState("");
	const [description, setDescription] = useState("");
	const [showedMedia, setShowedMedia] = useState([]);
	const [tags, setTags] = useState([]);
	const [tagInput, setTagInput] = useState("");
	const [search, setSearch] = useState("");
	const [usernameSearch, setUsername] = useState("");

	const loadOptions = (value, callback) => {
		setTimeout(() => {
			searchService.influencerSearchTags(value, callback);
		}, 1000);
	};

	const onInputChange = (inputValue, { action }) => {
		
		switch (action) {
			case "set-value":
				return;
			case "menu-close":
				setSearch("");
				return;
			case "input-change":
				setSearch(inputValue);
				return;
			default:
				return;
		}
	};

	const onChange = (option) => {
		setTags(option);
		setUsername(option.value)
		return false;
	};
	const handleSend = () => {
		
		let username = search
		console.log(search)
		let requestDTO = {
			contentId: postState.viewPostModal.post.Id,
			username: usernameSearch,
			status: "PENDING",
		};
		console.log(requestDTO)
		postService.sendCampaign(requestDTO,dispatch)
	};

	return (
		<Modal show={postState.campaignOptions.showModal} aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Header closeButton>
				<Modal.Title id="contained-modal-title-vcenter">{ "Search" }</Modal.Title>
			</Modal.Header>
			<Modal.Body>
				
				    <div>
                    <div className="form-group">
                        <label for="tags">Search influencer</label>

                        <AsyncSelect defaultOptions loadOptions={loadOptions} onInputChange={onInputChange} onChange={onChange} placeholder="search" inputValue={search} />
                    </div>
                     <hr />
					<div className="row">
						<button
							type="button"
							className="btn btn-link btn-fw text-secondary w-100 border-0"
							onClick={handleSend}
						>
							Send request
						</button>
					</div>
				</div>
			</Modal.Body>
		</Modal>
	);
};

export default SearchInfluencerModal;