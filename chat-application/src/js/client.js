import React from "react";
import ReactDOM from "react-dom";
import Layout from "./components/Layout";
import Login from "./components/login";
import Chat from "./components/chat";
import { BrowserRouter as Router, Route } from "react-router-dom";

const app = document.getElementById("app");
ReactDOM.render(
  <Router>
    <Layout>
      <Route exact path="/login" component={Login}></Route>
      <Route path="/chat" component={Chat}></Route>
    </Layout>
  </Router>,
  app
);
