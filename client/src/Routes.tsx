import * as React from "react";
import { Switch } from "react-router";
import { Route } from "react-router-dom";
import { Menu } from "semantic-ui-react";
import "./App.css";
import Bin from "./features/bin/Components/bin";
import Home from "./features/home/Components/home";

export class Routes extends React.Component<{}, {}> {
  public render() {
    return (
      <div className="App">
        <Menu size="large">
          <Menu.Item as="a" active={true} href="/" id="logo">
            reqhack
          </Menu.Item>
          <Menu.Item position="right">Login</Menu.Item>
        </Menu>
        <Switch>
          <Route exact={true} path="/" component={Home} />
          <Route path="/bin/:name" component={Bin} />
        </Switch>
      </div>
    );
  }
}
