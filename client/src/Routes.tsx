import * as React from "react";
import { Switch } from "react-router";
import { Route } from "react-router-dom";
import { Container, Menu } from "semantic-ui-react";
import "./App.css";
import Home from "./features/home/Components/home";

export class Routes extends React.Component<{}, {}> {
  public render() {
    return (
      <div className="App">
        <Menu size="large">
          <Container>
            <Menu.Item as="span" active={true}>
              reqhack
            </Menu.Item>
            <Menu.Item as="a" active={true}>
              Home
            </Menu.Item>
            <Menu.Item as="a">Work</Menu.Item>
            <Menu.Item as="a">Company</Menu.Item>
            <Menu.Item as="a">Careers</Menu.Item>
            <Menu.Item position="right">Login</Menu.Item>
          </Container>
        </Menu>
        <Switch>
          <Route path="/" component={Home} />
        </Switch>
      </div>
    );
  }
}
