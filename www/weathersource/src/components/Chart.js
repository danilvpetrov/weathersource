import React, { Component } from 'react';
import { Container } from 'react-bulma-components';
import ReactApexChart from "react-apexcharts";


class Chart extends Component {

    render() {
        const { config } = this.props
        const { options, series } = config

        const defaultOptions = {
            chart: {
                type: 'area',
                height: 350,
                zoom: {
                    enabled: true
                },
            },
            fill: {
                type: 'gradient'
            },
            dataLabels: {
                enabled: false
            },
            stroke: {
                curve: 'smooth',
                width: 4
            },
            title: {
                align: 'left'
            },
            yaxis: {
                opposite: true
            },
            legend: {
                horizontalAlign: 'left'
            }
        };

        const mergedOption = Object.assign({}, defaultOptions, options);

        return (
            <Container>
                <div id="chart">
                    <ReactApexChart
                        type={mergedOption.chart.type}
                        options={mergedOption}
                        series={series}
                        height={mergedOption.chart.height}
                    />
                </div>
            </Container>
        );
    }
}

export default Chart;
