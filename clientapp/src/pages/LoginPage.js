import React,{Component} from 'react';

class LoginPage extends Component {
    render() { 
        return ( 
            <React.Fragment>
                <section className="login-clean">
                    <form method="post">
                        <h2 className="sr-only">Login Form</h2>
                        <div class="illustration"><i class="icon ion-ios-navigate"></i></div>
                        <div class="form-group"><input class="form-control" type="email" name="email" placeholder="Email"></input></div>
                        <div class="form-group"><input class="form-control" type="password" name="password" placeholder="Password"></input></div>
                        <div class="form-group"><button class="btn btn-primary btn-block" type="submit">Log In</button></div><a class="forgot" href="#">Forgot your email or password?</a>
                    </form>
                </section>
            </React.Fragment>
            
        );
    }
}
 
export default LoginPage;