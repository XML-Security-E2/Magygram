import React, {Component} from 'react';

class RegistrationPage extends Component {
    render() { 
        return (
            <React.Fragment>
                <section className="register-photo">
                    <div className="form-container">
                    <div className="image-holder"/>
                    <form method="post">
                        <h2 className="text-center"><strong>Create</strong> an account.</h2>
                        <div className="form-group"><input className="form-control" type="text" name="nameInput" placeholder="Name"/></div>
                        <div className="form-group"><input className="form-control" type="text" name="surnameInput" placeholder="Surname"/></div>
                        <div className="form-group"><input className="form-control" type="email" name="email" placeholder="Email"/></div>
                        <div className="form-group"><input className="form-control" type="password" name="passwordInput" placeholder="Password"/></div>
                        <div className="form-group"><input className="form-control" type="password" name="password-repeat" placeholder="Password (repeat)"/></div>
                        <div className="form-group">
                            <div className="form-check"><label class="form-check-label"><input className="form-check-input" type="checkbox"/>I agree to the license terms</label></div>
                        </div>
                        <div className="form-group"><button className="btn btn-primary btn-block" type="submit">Sign Up</button></div><a className="already" href="#">You already have an account? Login here.</a>
                    </form>
                    </div>
                </section>
            </React.Fragment>
        );
    }
}
 
export default RegistrationPage;