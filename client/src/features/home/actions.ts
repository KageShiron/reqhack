import { createStandardAction } from "typesafe-actions";
import { IBin } from "./models";

const CREATE_BIN = "home/CREATE_BIN";

export const createBin = createStandardAction(CREATE_BIN)<IBin>();
