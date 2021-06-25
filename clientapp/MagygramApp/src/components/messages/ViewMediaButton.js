const ViewMediaButton = ({ disabled, text, openMedia, media, messageId }) => {
	return (
		<button disabled={disabled} type="button" className="btn btn-outline-secondary btn-icon-text" onClick={() => openMedia(messageId, media)}>
			{text} <i className="mdi mdi-play btn-icon-append"></i>
		</button>
	);
};

export default ViewMediaButton;
