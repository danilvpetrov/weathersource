import React, { Component } from 'react';
import { connect } from 'react-redux';
import {
    withRouter, Switch, Route, NavLink, Redirect,
    useRouteMatch, useLocation
} from 'react-router-dom'
import Moment from 'react-moment';
import WeatherData from './Weatherdata'
import { Section, Container, Heading, Columns, Menu } from 'react-bulma-components';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faExclamationTriangle, faCalendarDay } from '@fortawesome/free-solid-svg-icons'

const mapStateToProps = state => {
    const { data } = state.forecast
    const { daily } = data;
    return { daily }
}


class Daily extends Component {
    render() {
        const { daily } = this.props;
        return (
            <DailyStructure items={daily} />
        )
    }
}

function DailyStructure(data) {
    const { items } = data;
    const { path, url, isExact } = useRouteMatch();

    if (isExact === true && items.length > 0) {
        return <Redirect to={{
            pathname: `${path}/${items[0].time}`,
            state: items[0]
        }} />
    }

    const renderedItems = items.map(item => {
        return (
            <NavLink
                key={item.time}
                to={{
                    pathname: `${url}/${item.time}`,
                    state: item
                }}
                activeClassName="is-active"
            >
                <span style={{ fontSize: 12, marginRight: 5 }}><FontAwesomeIcon icon={faCalendarDay} /></span>
                <Moment unix format="dddd, MMMM Do">{item.time}</Moment>
            </NavLink>
        )
    })
    return (
        <Columns keys={12}>
            <Columns.Column size={3}>
                <Section>
                    <Menu>
                        <Menu.List title="Daily Forecasts">
                            {renderedItems}
                        </Menu.List>
                    </Menu>
                </Section>
            </Columns.Column>
            <Columns.Column size={9}>
                <Switch>
                    <Route path={`${path}/:dailyId`} >
                        <DailyBody />
                    </Route>
                </Switch>
            </Columns.Column>
        </Columns>
    )
}


function DailyBody(data) {
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
            <Heading size={3}>
                Forecast for <Moment unix tz={state.timezone} format="dddd, MMMM Do">
                    {state.time}
                </Moment>
            </Heading>
            <Container>
                <WeatherData data={state} />
            </Container>
        </Section>
    )
}

export default withRouter(connect(mapStateToProps)(Daily));
