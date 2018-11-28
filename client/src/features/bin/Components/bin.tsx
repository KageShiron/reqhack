import * as React from "react";
import { connect } from "react-redux";
import { RouteComponentProps, withRouter } from "react-router";
import { RootState } from "../../../store";

interface IProps extends RouteComponentProps<{ name: string }> {}

const Bin = (props: IProps) => (
  <div>
    <h1>Inspect {props.match.params.name}</h1>
  </div>
);

const mapStateToProps = (state: RootState, props: IProps) => ({});

export default withRouter(connect(mapStateToProps)(Bin));
