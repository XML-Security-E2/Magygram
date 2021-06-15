const SuccessAlert = (props) => {
	return (
		<div className="alert alert-success alert-dismissible fade show" hidden={props.hidden} role="alert">
			<h4 className="alert-heading">{props.header}</h4>
			<hr />
			<p className="mb-0">{props.message}</p>
			<button type="button" className="close" onClick={props.handleCloseAlert}>
				<span aria-hidden="true">&times;</span>
			</button>
		</div>
	);
};

export default SuccessAlert;
