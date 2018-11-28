import { createBrowserHistory } from "history";
import * as React from "react";
import { render } from "react-dom";
import { Provider } from "react-redux";
import { Router } from "react-router";
import "semantic-ui-css/semantic.min.css";
import "./index.css";
import registerServiceWorker from "./registerServiceWorker";
import { Routes } from "./Routes";
import store from "./store";

const history = createBrowserHistory();
const Root = () => (
  <div>
    <Provider store={store}>
      <Router history={history}>
        <Routes />
      </Router>
    </Provider>
  </div>
);

render(<Root />, document.getElementById("root"));
registerServiceWorker();
