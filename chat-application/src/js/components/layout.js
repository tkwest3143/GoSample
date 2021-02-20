import React from "react";
import { Link, withRouter } from "react-router-dom";
import Header from "./header";
import Footer from "./footer";
import Login from "./login";

class Layout extends React.Component {
  navigate() {
    console.log(this.props.history);
    this.props.history.push("/");
  }
  constructor() {
    super();

    this.state = {
      title: "",
      name: "not login",
      loginFlg: 0,
    };
  }
  changeTitle(title) {
    this.setState({ title });
  }

  changeName(name) {
    this.setState({ name });
    this.setState({ loginFlg: 1 });
    this.props.history.push("/chat");
  }
  render() {
    var component = "";
    if (this.state.loginFlg == 1) {
      component = (
        <div class="btn-area">
          <Link to="/chat">
            <button class="btn btn-info">chat</button>
          </Link>
          <button class="btn btn-info" onClick={this.navigate.bind(this)}>
            featured
          </button>
        </div>
      );
    } else {
      component = <Login changeName={this.changeName.bind(this)} />;
    }
    setTimeout(() => {
      if (this.state.title == "Tsutomu") {
      } else {
      }
    }, 1000);
    return (
      <div>
        <div class="loginuser">{this.state.name}</div>
        <Header
          changeTitle={this.changeTitle.bind(this)}
          title={this.state.title}
        />
        {component}
        {this.props.children}
        <Footer />
      </div>
    );
  }
}
export default withRouter(Layout);
