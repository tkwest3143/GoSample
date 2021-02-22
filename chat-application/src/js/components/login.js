import React from "react";
import UserProfile from "../data/user";
import { BrowserRouter, Link, Route, Switch } from "react-router-dom";
import UserStore from "../data/user";

export default class Login extends React.Component {
  constructor() {
    super();
    this.state = {
      user: UserStore.getUser(),
    };
  }

  handleChange(e) {
    var xhr = new XMLHttpRequest();

    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;
    var param = "?username=" + username + "&password=" + password;
    xhr.open("POST", "/doLogin" + param);
    xhr.send(param);
    console.log(this.props);
    xhr.onreadystatechange = function () {
      if (xhr.readyState === 4 && xhr.status === 200) {
        var data = JSON.parse(xhr.responseText);

        if (data.err != null) {
          console.log("エラー有");
          UserStore.setError(data.err);
          UserStore.on("change", () => {
            this.setState({
              error: UserStore.getError(),
            });
          });
        } else {
          console.log("エラーなし");
          var userid = data.userinfo.UserID;
          console.log(userid);
          UserStore.setUser(username, userid);
          UserStore.on("change", () => {
            this.setState({
              user: UserStore.getUsername(),
            });
          });
        }
      }
    };
    xhr.onreadystatechange;
    this.props.changeName(username);
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
