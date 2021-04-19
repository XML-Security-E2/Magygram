import React, {Component} from 'react';
import RegistrationForm from '../components/RegistrationForm';

const RegistrationPage = () => {
    return ( 
        <React.Fragment>
                <section className="register-photo">
                    <div className="form-container">
                        <div className="image-holder"/>
                        <RegistrationForm/>
                    </div>
                </section>
            </React.Fragment>
     );
}
 
export default RegistrationPage;