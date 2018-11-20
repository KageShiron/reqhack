import * as React from 'react';
import { connect } from 'react-redux';
import {Icon, Input, Popup} from "semantic-ui-react";
import {reqhackActions} from "../";
import {RootState} from "../../../store";

interface IProps {
  generateRandomName: (name: string) => any;
}

interface IState{
  name: string;
}

class BinNameForm extends React.Component<IProps, IState> {
  public readonly state: Readonly<IState> = { name: "" };
  public render() {
    return (            <div className="kanit">
        <span>https://</span>
        <Input icon={true} placeholder="BinName">
          <input />
          <Popup trigger={<Icon circular={true} name='redo' link={true } />}>Random Name</Popup>
        </Input>
        <span>.reqhack00.esora.xyz</span>
      </div>
    )
  }
}
const mapStateToProps = (state:RootState) => ({});
export default connect(mapStateToProps,{
  generateRandomName: () => reqhackActions.generateRandomName(),
})(BinNameForm);