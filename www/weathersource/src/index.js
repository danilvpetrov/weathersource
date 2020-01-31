import React from 'react';
import ReactDOM from 'react-dom';
import configureStore from './store';
import { fetchForecast } from './actions/forecast'
import 'react-bulma-components/dist/react-bulma-components.min.css';
import './index.css';
import Root from './Root';
import * as serviceWorker from './serviceWorker';

const store = configureStore()

ReactDOM.render(
    <Root store={store} />,
    document.getElementById('root')
);

store.dispatch(fetchForecast())

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
