import * as React from "react";
import { Button, Input, Menu } from "semantic-ui-react";
import Container from "semantic-ui-react/dist/commonjs/elements/Container/Container";
import "./App.css";

class App extends React.Component {
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
        <div>
          https://<Input icon={{ name: 'redo', circular: true, link: true }} placeholder='bin name' />.reqhack00.esora.xyz
        </div>
        <div>
          <Button> Create New Bin</Button>
        </div>
      </div>
    );
  }
}

export default App;
