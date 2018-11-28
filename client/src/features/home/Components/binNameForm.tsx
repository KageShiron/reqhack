import * as React from 'react';
import { connect } from 'react-redux';
import {Icon, Input, Popup} from "semantic-ui-react";
import {RootState} from "../../../store";

interface IState{
  name: string;
}

class BinNameForm extends React.Component<{}, IState> {

  private static randomName() {
    const list = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
    const len = 8;
    let res = "";
    for (let i = 0; i < len; i++) {
      res += list.charAt(Math.floor(Math.random() * list.length));
    }
    return res;
  }

  public readonly state: Readonly<IState> = { name: "" };

  public render() {
    const { name } = this.state;
    return (
      <div className="kanit">
        <span>https://</span>
        <Input icon={true} placeholder="BinName" value={name} onChange={this.handleChangeText} >
          <input />
          <Popup trigger={<Icon circular={true} name='redo' link={true } onClick={this.handleGenerateRandomName} />}>Random Name</Popup>
        </Input>
        <span>.reqhack00.esora.xyz</span>
      </div>
    )
  }

  public componentDidMount = (): void => {
    this.setState({name:BinNameForm.randomName()});
  };

  private handleChangeText = (e: React.ChangeEvent<HTMLInputElement>) => {
    this.setState({name:e.target.value});
  };

  private handleGenerateRandomName = () => {
    this.setState({name:BinNameForm.randomName()});
  };
}
const mapStateToProps = (state:RootState) => ({});
export default connect(mapStateToProps,{})(BinNameForm);