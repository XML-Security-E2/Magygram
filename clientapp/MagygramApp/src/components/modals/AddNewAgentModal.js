import { useContext, useState } from "react";
import { Button, Modal } from "react-bootstrap";
import { modalConstants } from "../../constants/ModalConstants";
import { AdminContext } from "../../contexts/AdminContext";
import { userService } from "../../services/UserService";

const AddNewAgentModal = () => {
	const { state,dispatch } = useContext(AdminContext);

	const [name, setName] = useState("");
	const [surname, setSurname] = useState("");
	const [email, setEmail] = useState("");
	const [username, setUsername] = useState("");
	const [password, setPassword] = useState("");
	const [repeatedPassword, setRepeatedPassword] = useState("");
	const [webSite, setWebSite] = useState("");

	const handleModalClose = () => {
		setName("")
		setSurname("")
		setEmail("")
		setUsername("")
		setPassword("")
		setRepeatedPassword("")
		setWebSite("")
		dispatch({ type: modalConstants.HIDE_REGISTER_AGENT_MODAL });
	};

	const handleSubmit = (e) => {
		e.preventDefault();

        let agent = {
				name,
				surname,
				email,
				username,
				password,
				repeatedPassword,
				webSite,
			};

        userService.registerAgentByAdmin(agent, dispatch);
	};

	return (
		<Modal show={state.registerNewAgent.showModal} aria-labelledby="contained-modal-title-vcenter" centered onHide={handleModalClose}>
			<Modal.Header closeButton>
				<Modal.Title id="contained-modal-title-vcenter">
					<big>Register new agent</big>
				</Modal.Title>
			</Modal.Header>
			<Modal.Body>
            <form method="post" onSubmit={handleSubmit}>
				<div>
					<div className="form-group">
						<input className="form-control" required type="text" name="nameInput" placeholder="Name" value={name} onChange={(e) => setName(e.target.value)} />
					</div>

					<div className="form-group">
						<input className="form-control" required type="text" name="surnameInput" placeholder="Surname" value={surname} onChange={(e) => setSurname(e.target.value)} />
					</div>

					<div className="form-group">
						<input className="form-control" required type="email" name="email" placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)} />
					</div>

					<div className="form-group">
						<input className="form-control" required type="username" name="username" placeholder="@Username" value={username} onChange={(e) => setUsername(e.target.value)} />
					</div>
					
                    <div className="form-group">
						<input className="form-control" required type="webSite" name="webSite" placeholder="Web site" value={webSite} onChange={(e) => setWebSite(e.target.value)} />
					</div>
					

					<div className="form-group">
						<input className="form-control" required type="password" name="passwordInput" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} />
					</div>

					<div className="form-group">
						<input
							className="form-control"
							required
							type="password"
							name="password-repeat"
							placeholder="Password (repeat)"
							value={repeatedPassword}
							onChange={(e) => setRepeatedPassword(e.target.value)}
						/>
					</div>

                    <div className="form-group text-center" style={{ color: "red", fontSize: "0.8em" }} hidden={!state.registerNewAgent.showError}>
						{state.registerNewAgent.errorMessage}
					</div>

					<div className="form-group">
						<input className="btn btn-primary btn-block" type="submit" value="Sign In" />
					</div>
				</div>



			</form>
			</Modal.Body>
			<Modal.Footer>
				<Button onClick={handleModalClose}>Close</Button>
			</Modal.Footer>
		</Modal>
	);
};

export default AddNewAgentModal;
