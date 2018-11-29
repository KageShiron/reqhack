import { createBrowserHistory } from "history";
import { Provider } from "mobx-react";
import * as React from "react";
import { render } from "react-dom";
import { Router } from "react-router";
import "semantic-ui-css/semantic.min.css";
import "./index.css";
import registerServiceWorker from "./registerServiceWorker";
import { Routes } from "./Routes";
import BinStore from "./store/BinStore";

const binStore = new BinStore();
const stores = {
  bin: binStore
};

const history = createBrowserHistory();
const Root = () => (
  <div>
    <Provider store={stores}>
      <Router history={history}>
        <Routes />
      </Router>
    </Provider>
  </div>
);

render(<Root />, document.getElementById("root"));
registerServiceWorker();
