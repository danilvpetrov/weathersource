import React, { Component } from 'react';
import Moment from 'react-moment';
import 'moment-timezone';
import { Heading, Box, Table } from 'react-bulma-components';
import WeatherIcon from 'react-icons-weather';



class WeatherData extends Component {
    render() {
        const { data } = this.props
        return (
            <div>
                <Box>
                    <span style={{ fontSize: 72 }}><WeatherIcon name="darksky" iconId={data.icon} /></span>
                    <Heading size={4}>{data.summary}</Heading>
                    <Table>
                        <tbody>
                            <DataItem label={"Temperature"} val={data.temperature} unit={"°C"} />
                            <DataItem label={"Apparent Temperature"} val={data.apparent_temperature} unit={"°C"} />
                            <DataItem label={"Precipitation type"} val={data.precip_type} />
                            <DataItem label={"Liquid precipitation rate"} val={data.precip_intensity} unit={"mm"} />
                            <DataAsPercentage label={"Precipitation Probability"} val={data.precip_probability} />
                            <DataItem label={"Max. Liquid precipitation rate"} val={data.precip_intensity_max} unit={"mm"} />
                            <DataAsTime label={"Max. Liquid precipitation rate time"} val={data.precip_intensity_max_time} />

                            <DataItem label={"Atmospheric Pressure"} val={data.pressure} unit={"millibars"} />
                            <DataAsPercentage label={"Humidity"} val={data.humidity} />
                            <DataItem label={"Dew Point"} val={data.temperature} unit={"°C"} />
                            <DataAsPercentage label={"Cloud Cover"} val={data.cloud_cover} />
                            <DataItem label={"Wind Speed"} val={data.wind_speed} unit={"km/h"} />
                            <DataAsWindDirection label={"Wind Direction"} val={data.wind_bearing} />
                            <DataItem label={"Wind Gust"} val={data.wind_gust} unit={"km/h"} />
                            <DataAsTime label={"Wind Gust Time"} val={data.wind_gust_time} />
                            <DataAsPercentage label={"Cloud Cover"} val={data.cloud_cover} />
                            <DataAsPercentage label={"Moon Phase"} val={data.moon_phase} />
                            <DataItem label={"UV index"} val={data.uv_index} />
                            <DataAsTime label={"UV index Time"} val={data.uv_index_time} />
                            <DataItem label={"Ozone"} val={data.ozone} />
                            <DataItem label={"Visibility"} val={data.visibility} unit={"km"} />

                            <DataItem label={"Apparent Temperature High"} val={data.apparent_temperature_high} unit={"°C"} />
                            <DataAsTime label={"Apparent Temperature High Time"} val={data.apparent_temperature_high_time} />
                            <DataItem label={"Apparent Temperature Low"} val={data.apparent_temperature_low} unit={"°C"} />
                            <DataAsTime label={"Apparent Temperature Low Time"} val={data.apparent_temperature_low_time} />
                            <DataItem label={"Apparent Temperature Max"} val={data.apparent_temperature_max} unit={"°C"} />
                            <DataAsTime label={"Apparent Temperature Max Time"} val={data.apparent_temperature_max_time} />
                            <DataItem label={"Apparent Temperature Min"} val={data.apparent_temperature_min} unit={"°C"} />
                            <DataAsTime label={"Apparent Temperature Min Time"} val={data.apparent_temperature_min_time} />
                            <DataItem label={"Temperature High"} val={data.temperature_high} unit={"°C"} />
                            <DataAsTime label={"Temperature High Time"} val={data.temperature_high_time} />
                            <DataItem label={"Temperature Low"} val={data.temperature_low} unit={"°C"} />
                            <DataAsTime label={"Temperature Low Time"} val={data.temperature_low_time} />
                            <DataItem label={"Temperature Max"} val={data.temperature_max} unit={"°C"} />
                            <DataAsTime label={"Temperature Max Time"} val={data.temperature_max_time} />
                            <DataItem label={"Temperature Min"} val={data.temperature_min} unit={"°C"} />
                            <DataAsTime label={"Temperature Min Time"} val={data.temperature_min_time} />

                            <DataAsTime label={"Sunrise Time"} val={data.sunrise_time} />
                            <DataAsTime label={"Sunset Time"} val={data.sunset_time} />

                        </tbody>
                    </Table>
                </Box>
            </div>
        );
    }
}

class DataItem extends Component {
    render() {
        if (!this.props.val) {
            return null;
        }

        const label = this.props.label
        const val = this.props.val
        const unit = this.props.unit || ""

        return (
            <tr>
                <td><strong>{label}:</strong></td>
                <td>{val}{unit}</td>
            </tr>
        );
    }
}

class DataAsPercentage extends Component {
    render() {
        if (!this.props.val) {
            return null;
        }

        const label = this.props.label
        const val = parseFloat(this.props.val * 100).toFixed(0);

        return (
            <tr>
                <td><strong>{label}:</strong></td>
                <td>{val}%</td>
            </tr>
        );
    }
}

class DataAsTime extends Component {
    render() {
        if (!this.props.val) {
            return null;
        }

        const label = this.props.label
        const val = this.props.val
        const tz = this.props.tz
        const format = this.props.format || "HH:mm"

        return (
            <tr>
                <td><strong>{label}:</strong></td>
                <td><Moment unix tz={tz} format={format}>{val}</Moment></td>
            </tr>
        );
    }
}

class DataAsWindDirection extends Component {
    degToCompass(num) {
        var val = Math.floor((num / 22.5) + 0.5);
        var arr = [
            "N", "NNE", "NE", "ENE", "E", "ESE",
            "SE", "SSE", "S", "SSW", "SW", "WSW",
            "W", "WNW", "NW", "NNW"
        ];
        return arr[(val % 16)];
    }

    render() {
        if (!this.props.val) {
            return null;
        }

        const label = this.props.label
        const val = this.props.val

        return (
            <tr>
                <td><strong>{label}:</strong></td>
                <td>{this.degToCompass(val)}</td>
            </tr>
        );
    }
}

export default WeatherData;
