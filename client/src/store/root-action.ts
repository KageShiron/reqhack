import { LocationChangeAction, RouterAction } from "react-router-redux";

import { ReqhackAction } from "../features/home";

type ReactRouterAction = RouterAction | LocationChangeAction;
export type RootAction = ReactRouterAction | ReqhackAction;
