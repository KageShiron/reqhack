import { LocationChangeAction, RouterAction } from "react-router-redux";

import { ReqhackAction } from "../features/reqhack";

type ReactRouterAction = RouterAction | LocationChangeAction;
export type RootAction = ReactRouterAction | ReqhackAction;
