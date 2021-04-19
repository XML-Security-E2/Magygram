import React, {useState} from 'react';
import Axios from 'axios'

const RegistrationForm = () => {

    const [name, setName] = useState('')
    const [surname, setSurname] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [repeatedPassword, setRepeatedPassword] = useState('');

    const handleSubmit = (e) => {
        let registrationDTO = {
            name : name,
            surname: surname, 
            email: email,
            password: password,
            repeatedPassword: repeatedPassword, 
        };

        Axios
        .post("https://localhost:443" + "/api/registration", registrationDTO).then((res) =>{
            console.log(res)
        }).catch((err) => {
            console.log(err)
        });
        
        e.preventDefault();
        console.log(name)
        console.log(surname)
        console.log(email)
        console.log(password)
        console.log(repeatedPassword)
    }

    return ( 
        <form method="post" onSubmit={handleSubmit}>
            <h2 className="text-center"><strong>Create</strong> an account.</h2>
            <div className="form-group"><input className="form-control" required type="text" name="nameInput" placeholder="Name" value={name} onChange={(e) => setName(e.target.value)}/></div>
            <div className="form-group"><input className="form-control" required type="text" name="surnameInput" placeholder="Surname" value={surname} onChange={(e) => setSurname(e.target.value)}/></div>
            <div className="form-group"><input className="form-control" required type="email" name="email" placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)}/></div>
            <div className="form-group"><input className="form-control" required type="password" name="passwordInput" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)}/></div>
            <div className="form-group"><input className="form-control" required type="password" name="password-repeat" placeholder="Password (repeat)" value={repeatedPassword} onChange={(e) => setRepeatedPassword(e.target.value)}/></div>
            <div className="form-group">
                <div className="form-check"><label className="form-check-label"><input className="form-check-input" type="checkbox"/>I agree to the license terms</label></div>
            </div>
            <div className="form-group"><input className="btn btn-primary btn-block" type="submit" value="Sign In"/></div><a className="already" href="#/login">You already have an account? Login here.</a>

        </form>
     );
}
 
export default RegistrationForm;