import fetch from 'cross-fetch'

export const FETCH_FORECAST = 'FETCH_FORECAST'
export const FETCH_FORECAST_SUCCEEDED = 'FETCH_FORECAST_SUCCEEDED'
export const FETCH_FORECAST_FAILED = 'FETCH_FORECAST_FAILED'


const doFetch = () => ({ type: FETCH_FORECAST })
const success = (forecast) => ({ type: FETCH_FORECAST_SUCCEEDED, forecast: forecast })
const failure = (error) => ({ type: FETCH_FORECAST_FAILED, error: error })

export function fetchForecast() {
  return (dispatch) => {
    dispatch(doFetch())
    return fetch("/api/forecast")
      .then(
        response => {
          if (!response.ok) {
            throw new Error(response.status + " " + response.statusText)
          }
          return response.json()
        })
      .then(json => dispatch(success(json)))
      .catch(
        error => {
          dispatch(failure(error.message))
        })
  }
}
