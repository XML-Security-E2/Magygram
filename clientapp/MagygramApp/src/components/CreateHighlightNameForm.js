import { useContext } from "react";
import { Button } from "react-bootstrap";
import { storyConstants } from "../constants/StoryConstants";
import { StoryContext } from "../contexts/StoryContext";

const CreateHighlightNameForm = ({ setHighlightName, handleSubmit }) => {
	const { storyState, dispatch } = useContext(StoryContext);

	const handleHideNameInput = () => {
		dispatch({ type: storyConstants.HIDE_HIGHLIGHTS_NAME_INPUT });
	};

	return (
		<form className="forms-sample w-100" method="post" hidden={!storyState.highlights.showHighlightsName} onSubmit={handleSubmit}>
			<input required type="text" className="form-control" placeholder="Highlight name" onChange={(e) => setHighlightName(e.target.value)} />

			<Button type="submit" className="float-right mt-2">
				Create highlights
			</Button>

			<Button className="float-left mt-2" hidden={!storyState.highlights.showHighlightsName} onClick={handleHideNameInput}>
				Cancel
			</Button>
		</form>
	);
};

export default CreateHighlightNameForm;
