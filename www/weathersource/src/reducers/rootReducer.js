import { combineReducers } from 'redux';
import forecastReducer from './forecast';
import chartsReducer from './charts';

export default combineReducers({
     forecast: forecastReducer,
     charts: chartsReducer,
});
