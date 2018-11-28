import * as React from "react";
import { Button } from "semantic-ui-react";
import BinNameForm from "./binNameForm";

export default class Home extends React.Component<{}, {}> {
  public render() {
    return (
      <div>
        <BinNameForm />
        <br />
        <div>
          <Button primary={true} size="massive" onClick={this.handlePost}>
            {" "}
            Create New Bin
          </Button>
        </div>
      </div>
    );
  }

  public handlePost = () => {
    //    fetch("/")
  };
}
