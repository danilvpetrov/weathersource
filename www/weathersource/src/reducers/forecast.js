import {
    FETCH_FORECAST,
    FETCH_FORECAST_SUCCEEDED,
    FETCH_FORECAST_FAILED
} from '../actions/forecast';

const initState = {
    isLoading: false,
    isLoaded: false,
    error: null,
    data: null
}

export default (state = initState, action) => {
    switch (action.type) {
        case FETCH_FORECAST:
            return Object.assign({}, state, {
                isLoading: true,
                isLoaded: false,
                error: null,
                data: null
            })
        case FETCH_FORECAST_FAILED:
            return Object.assign({}, state, {
                isLoading: false,
                isLoaded: false,
                error: action.error,
                data: null
            })
        case FETCH_FORECAST_SUCCEEDED:
            return Object.assign({}, state, {
                isLoading: false,
                isLoaded: true,
                error: null,
                data: action.forecast,
            })
        default:
            return state
    }
}
