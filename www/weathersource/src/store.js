
import { createStore, applyMiddleware, compose } from 'redux';
import thunk from 'redux-thunk';
import { createLogger } from 'redux-logger';
import rootReducer from './reducers/rootReducer';

export default function configureStore(preloadedState) {

    const composeEnhancers = (
        typeof window !== 'undefined' && window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__) || compose;

    return createStore(
        rootReducer,
        preloadedState,
        composeEnhancers(
            applyMiddleware(
                thunk,
                createLogger(),
            )
        )
    );
}
