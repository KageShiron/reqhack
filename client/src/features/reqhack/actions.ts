import { createStandardAction } from "typesafe-actions";
import { IBin } from "./models";

const CREATE_BIN = "reqhack/CREATE_BIN";
const GENERATE_RANDOM_NAME = "reqhack/GENERATE_RANDOM_NAME";

export const createBin = createStandardAction(CREATE_BIN)<IBin>();
export const generateRandomName = createStandardAction(GENERATE_RANDOM_NAME)();
