import React, { Component } from 'react';
import { connect } from 'react-redux';
import {
    withRouter, Switch, Route, NavLink, Redirect,
    useRouteMatch, useLocation
} from 'react-router-dom'
import Chart from './Chart'
import { Section, Container, Heading, Columns, Menu } from 'react-bulma-components';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faExclamationTriangle, faChartArea } from '@fortawesome/free-solid-svg-icons'

const mapStateToProps = state => {
    const { daily, hourly } = state.charts;
    return { daily, hourly }
}


class Charts extends Component {
    render() {
        const { daily, hourly } = this.props;
        return (
            <ChartsStructure daily={daily} hourly={hourly} />
        )
    }
}

function ChartsStructure(data) {
    const { daily, hourly } = data;
    const { path, url, isExact } = useRouteMatch();

    const dailyConfigs = [
        {
            pathname: `${url}/daily-max-temp`,
            chartConfig: {
                series: [{
                    name: "Daily Maximum Temperature (°C)",
                    data: daily.temperature
                }],
                options: {
                    title: {
                        text: 'Daily Maximum Temperature (°C)',
                        align: 'left'
                    },
                    labels: daily.time,
                    xaxis: {
                        type: 'datetime',
                    },
                    colors: [
                        '#66DA26'
                    ]
                }
            }
        },
        {
            pathname: `${url}/daily-precip-probability`,
            chartConfig: {
                series: [{
                    name: 'Daily Precipitation Probability (%)',
                    data: daily.precip_probability
                }],
                options: {
                    title: {
                        text: 'Daily Precipitation Probability (%)',
                        align: 'left'
                    },
                    labels: daily.time,
                    xaxis: {
                        type: 'datetime',
                    },
                    colors: [
                        '#ffdd56'
                    ]
                }
            }
        },
        {
            pathname: `${url}/daily-precip-rate`,
            chartConfig: {
                series: [{
                    name: 'Daily liquid precipitation rate (mm/hour)',
                    data: daily.precip_intensity
                }],
                options: {
                    title: {
                        text: 'Daily liquid precipitation rate (mm/hour)',
                        align: 'left'
                    },
                    labels: daily.time,
                    xaxis: {
                        type: 'datetime',
                    },
                    colors: [
                        '#3298dc'
                    ]
                }
            }
        }
    ]

    const hourlyConfigs = [
        {
            pathname: `${url}/hourly-max-temp`,
            chartConfig: {
                series: [{
                    name: 'Hourly Maximum Temperature (°C)',
                    data: hourly.temperature
                }],
                options: {
                    title: {
                        text: 'Hourly Maximum Temperature (°C)',
                        align: 'left'
                    },
                    labels: hourly.time,
                    xaxis: {
                        type: 'datetime',
                    },
                    colors: [
                        '#66DA26'
                    ]
                }
            }
        },
        {
            pathname: `${url}/hourly-precip-probability`,
            chartConfig: {
                series: [{
                    name: 'Hourly Precipitation Probability (%)',
                    data: hourly.precip_probability
                }],
                options: {
                    title: {
                        text: 'Hourly Precipitation Probability (%)',
                        align: 'left'
                    },
                    labels: hourly.time,
                    xaxis: {
                        type: 'datetime',
                    },
                    colors: [
                        '#ffdd56'
                    ]
                }
            }
        },
        {
            pathname: `${url}/hourly-precip-rate`,
            chartConfig: {
                series: [{
                    name: 'Hourly liquid precipitation rate (mm/hour)',
                    data: hourly.precip_intensity
                }],
                options: {
                    title: {
                        text: 'Hourly liquid precipitation rate (mm/hour)',
                        align: 'left'
                    },
                    labels: hourly.time,
                    xaxis: {
                        type: 'datetime',
                    },
                    colors: [
                        '#3298dc'
                    ]
                }
            }
        }
    ]

    if (isExact === true) {
        return <Redirect to={{
            pathname: `${path}/daily-max-temp`,
            state: dailyConfigs[0].chartConfig
        }} />
    }

    const renderLinksFromConfig = (configs) => {
        return configs.map((config, i) => {
            return (
                <NavLink
                    key={i}
                    to={{
                        pathname: config.pathname,
                        state: config.chartConfig
                    }}
                    activeClassName="is-active"
                >
                    <span style={{ fontSize: 12, marginRight: 5 }}>
                        <FontAwesomeIcon icon={faChartArea} />
                    </span>
                    {config.chartConfig.options.title.text}
                </NavLink>
            )
        })
    }

    return (
        <Columns keys={12}>
            <Columns.Column size={3}>
                <Section>
                    <Menu>
                        <Menu.List title="Daily Forecast Chars">
                            {renderLinksFromConfig(dailyConfigs)}
                        </Menu.List>
                        <Menu.List title="Hourly Forecast Chars">
                            {renderLinksFromConfig(hourlyConfigs)}
                        </Menu.List>
                    </Menu>
                </Section>
            </Columns.Column>
            <Columns.Column size={9}>
                <Switch>
                    <Route path={`${path}/:chartId`} >
                        <ChartBody />
                    </Route>
                </Switch>
            </Columns.Column>
        </Columns>
    )
}


function ChartBody() {
    const { state } = useLocation();

    if (!state) {
        return (
            <Section>
                <Heading size={3}>
                    <span style={{ fontSize: 28, marginRight: 5 }}>
                        <FontAwesomeIcon icon={faExclamationTriangle} />
                    </span>
                    Not Available
                </Heading>
            </Section>
        )
    }

    return (
        <Section>
            <Container>
                <Chart config={state} />
            </Container>
        </Section>
    )
}

export default withRouter(connect(mapStateToProps)(Charts));
