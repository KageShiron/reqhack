import { combineReducers } from "redux";
import { ActionType } from "typesafe-actions";
import * as actions from "./actions";
import { IBin } from "./models";

export type ReqhackAction = ActionType<typeof actions>;
export type ReqhackState = Readonly<{
  bin: IBin;
}>;

export default combineReducers<ReqhackState, ReqhackAction>({
  bin: (state = { name: "" }, action) => {
    switch (action.type) {
      default:
        return state;
    }
  }
});
