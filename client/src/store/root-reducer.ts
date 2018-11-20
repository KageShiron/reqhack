import { routerReducer } from "react-router-redux";
import { combineReducers } from "redux";
import { StateType } from "typesafe-actions";
import { reqhackReducer } from "../features/reqhack";

const rootReducer = combineReducers({
  reqhack: reqhackReducer,
  router: routerReducer
});
export type RootState = StateType<typeof rootReducer>;

export default rootReducer;
