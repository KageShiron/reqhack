import * as React from "react";
import { Button, Menu } from "semantic-ui-react";
import Container from "semantic-ui-react/dist/commonjs/elements/Container/Container";
import "./App.css";
import BinNameForm from "./features/home/Components/binNameForm";

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
        <Container>
          <BinNameForm />
        </Container>
        <br />
        <Container>
          <div>
            <Button primary={true} size="massive">
              {" "}
              Create New Bin
            </Button>
          </div>
        </Container>
      </div>
    );
  }
}

export default App;
