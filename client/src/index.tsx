import * as React from "react";
import { render } from "react-dom";
import { Provider } from "react-redux";
import "semantic-ui-css/semantic.min.css";
import App from "./App";
import "./index.css";
import registerServiceWorker from "./registerServiceWorker";
import store from "./store";

const Root = () => (
  <div>
    <Provider store={store}>
      <App />
    </Provider>
  </div>
);

render(<Root />, document.getElementById("root"));
registerServiceWorker();
