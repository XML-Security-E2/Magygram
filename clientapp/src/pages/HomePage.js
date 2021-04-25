import React, { Component } from "react";

class HomePage extends Component {
	render() {
        const navStyle = {"height":"50px","borderBottom":"1px solid rgb(200,200,200)"};
        const inputStyle= {"border":"1px solid rgb(200,200,200)","color":"rgb(210,210,210)","textAlign":"center"};
        const iconStyle= {"fontSize":"30px","margin":"0px","marginLeft":"13px"};
        const imgStyle= {"width":"30px","height":"30px","marginLeft":"13px","borderWidth":"1px","borderStyle":"solid"};
		return (
            <nav className="navbar navbar-light navbar-expand-md navigation-clean" style={navStyle}>
                <div className="container">
                    <div>
                        <img src="assets/img/logotest.png" alt="NistagramLogo"/>
                    </div>
                    <button className="navbar-toggler" data-toggle="collapse">
                        <span className="sr-only">Toggle navigation</span>
                        <span className="navbar-toggler-icon"></span>
                    </button>
                    <div>
                        <input type="text" style={inputStyle} placeholder="Search" value="Search"/>
                    </div>
                    <div className="d-xl-flex align-items-xl-center">
                        <i className="fa fa-home" style={iconStyle}/>
                        <i className="la la-wechat" style={iconStyle}/>
                        <i className="la la-compass" style={iconStyle}/>
                        <i className="fa fa-heart-o" style={iconStyle}/>
                        <img className="rounded-circle" style={imgStyle} src="assets/img/hitmanImage.jpeg" alt="ProfileImage"/>
                    </div>
                </div>
            </nav>
		);
	}
}

export default HomePage;
