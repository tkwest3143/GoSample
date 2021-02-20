import React from "react";
import UserProfile from "../data/user";
import { BrowserRouter, Link, Route, Switch } from "react-router-dom";
import UserStore from "../data/user";

export default class Login extends React.Component {
  constructor() {
    super();
    this.state = {
      user: UserStore.getUsername(),
    };
  }

  handleChange(e) {
    var xhr = new XMLHttpRequest();

    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;
    var param = "username=" + username + "&password=" + password;
    xhr.open("POST", "/doLogin");
    xhr.send(param);
    console.log(this.props);
    this.props.changeName(username);
    xhr.onreadystatechange = function () {
      if (xhr.readyState === 4 && xhr.status === 200) {
        UserStore.setUsername(username);
        UserStore.on("change", () => {
          this.setState({
            user: UserStore.getUsername(),
          });
        });
      }
    };
  }
  render() {
    return (
      <div>
        <input type="text" id="username" />
        <input type="text" id="password" />
        <button id="login" onClick={this.handleChange.bind(this)}>
          Sign In
        </button>
      </div>
    );
  }
}
