import * as React from 'react';
import * as ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import store from './store';

import App from './App';
import './index.css';
import registerServiceWorker from './registerServiceWorker';

ReactDOM.render(
    <Provider store={store}>
        <div>abcde {(store as any).hello} bbb</div>
        <App />
    </Provider>
  ,
  document.getElementById('root') as HTMLElement
);

registerServiceWorker();
