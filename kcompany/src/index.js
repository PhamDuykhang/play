import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import { BrowserRouter } from 'react-router-dom'
import App from './App';
import * as serviceWorker from './serviceWorker';
import { createStore,applyMiddleware,compose,combineReducers } from 'redux';
import EmployeeReducer from './reducer/EReducer';
import { Provider } from 'react-redux';
import thunk from 'redux-thunk'

const RootReducer = combineReducers(
    {
        employee:EmployeeReducer,
    }
)

const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;
const store = createStore (RootReducer ,composeEnhancers(applyMiddleware(thunk)))  


ReactDOM.render(<Provider store={store}>
      <BrowserRouter>
    <App />
    </BrowserRouter>
</Provider>, document.getElementById('root'));

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();