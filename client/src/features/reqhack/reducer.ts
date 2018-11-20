import { combineReducers } from "redux";
import { ActionType, getType } from "typesafe-actions";
import * as actions from "./actions";
import { IBin } from "./models";

export type ReqhackAction = ActionType<typeof actions>;
export type ReqhackState = Readonly<{
  bin: IBin;
}>;

export default combineReducers<ReqhackState, ReqhackAction>({
  bin: (state = { name: "" }, action) => {
    switch (action.type) {
      case getType(actions.generateRandomName):
        return { name: randomName() };
      default:
        return state;
    }
  }
});

function randomName() {
  const list = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
  const len = 8;
  let res = "";
  for (let i = 0; i < len; i++) {
    res += list.charAt(Math.floor(Math.random() * list.length));
  }
  return res;
}
