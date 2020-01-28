import {
    FETCH_FORECAST_SUCCEEDED,
} from '../actions/forecast';

const initState = {
    forecast: {
        hourly: {},
        daily: {}
    }
}

export default (state = initState, action) => {
    switch (action.type) {
        case FETCH_FORECAST_SUCCEEDED:
            const { hourly, daily } = action.forecast

            return Object.assign(
                {},
                state,
                {
                    hourly: {
                        time: hourly.map(item => item.time * 1000),
                        temperature: hourly.map(item => item.temperature || 0),
                        precip_probability: hourly.map(item => item.precip_probability * 100 || 0),
                        precip_intensity: hourly.map(item => item.precip_intensity || 0)
                    },
                    daily: {
                        time: daily.map(item => item.time * 1000),
                        temperature: daily.map(item => item.temperature_high || 0),
                        precip_probability: daily.map(item => item.precip_probability * 100 || 0),
                        precip_intensity: daily.map(item => item.precip_intensity || 0)
                    }
                })
        default:
            return state
    }
}
