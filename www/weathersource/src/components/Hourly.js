import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import {
    withRouter, Switch, Route, Redirect,
    NavLink, useRouteMatch, useLocation
} from 'react-router-dom'
import Moment from 'react-moment';
import WeatherData from './Weatherdata'
import { Section, Container, Heading, Columns, Menu } from 'react-bulma-components';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faExclamationTriangle, faHourglass, faStopwatch } from '@fortawesome/free-solid-svg-icons'


const calendarStrings = {
    sameDay: '[Today at] LT',
    nextDay: '[Tomorrow at] LT',
    nextWeek: 'dddd [at] LT',
    sameElse: 'L'
};

const mapStateToProps = state => {
    const { data } = state.forecast
    const { hourly } = data;
    return { hourly }
}


class Hourly extends Component {
    render() {
        const { hourly } = this.props;
        return (
            <HourlyStructure items={hourly} />
        )
    }
}

class MenuListCollapsible extends Component {
    constructor(props) {
        super(props);
        this.state = {
            collapsed: props.collapsed
        };
        this.handleClick = this.handleClick.bind(this);
    }

    handleClick(e) {
        e.preventDefault();
        this.setState(state => ({
            collapsed: !state.collapsed
        }));
    }

    render() {
        const { state, props } = this;
        return (
            <>
                <p className="menu-label" onClick={this.handleClick} style={{ cursor: 'pointer' }}>
                    <span style={{ fontSize: 12, marginRight: 5 }}><FontAwesomeIcon icon={faHourglass} /></span>{props.title}
                </p>
                <ul className="menu-list" hidden={!state.collapsed}>
                    {props.children}
                </ul>
            </>
        )
    }
}

MenuListCollapsible.propTypes = {
    collapsed: PropTypes.bool
}

MenuListCollapsible.defaultProps = {
    collapsed: true
};


function HourlyStructure(data) {
    const { items } = data;
    const { path, url, isExact } = useRouteMatch();
    let redirect = false;

    if (isExact === true && items.length > 0) {
        redirect = true;
    }

    items.sort((a, b) => {
        if (a.time > b.time) { return 1; }
        if (a.time < b.time) { return -1; }
        return 0;
    })

    const start = items[0].time
    const grouped = []

    items.forEach(item => {
        // time is in seconds, NOT in ms
        let hourDiff = Math.abs(item.time - start) / 36000;
        let i = Math.floor(hourDiff % 12);

        if (!grouped[i]) {
            grouped[i] = {
                key: (i + 1),
                title: `Next ${(i + 1) * 12} hours`,
                items: []
            };
        }

        grouped[i].items.push(
            <NavLink
                key={item.time}
                to={{
                    pathname: `${url}/${item.time}`,
                    state: item
                }}
                activeClassName="is-active"
            >
                <span style={{ fontSize: 12, marginRight: 5 }}><FontAwesomeIcon icon={faStopwatch} /></span>
                <Moment unix calendar={calendarStrings}>
                    {item.time}
                </Moment>
            </NavLink>
        )
    })

    const subMenus = grouped.map((group, i) => {
        let collapsed = redirect ? i === 0 : true
        return (
            <MenuListCollapsible title={group.title} key={i} collapsed={collapsed}>
                {group.items}
            </MenuListCollapsible>
        )
    })

    return (
        <>
            {redirect && (<Redirect to={{
                pathname: `${path}/${items[0].time}`,
                state: items[0]
            }} />)}
            <Columns keys={12}>
                <Columns.Column size={3}>
                    <Section>
                        <Menu>
                            <Menu.List title="Hourly Forecasts">
                                {subMenus}
                            </Menu.List>
                        </Menu>
                    </Section>
                </Columns.Column>
                <Columns.Column size={9}>
                    <Switch>
                        <Route path={`${path}/:hourlyId`} >
                            <HourlyBody />
                        </Route>
                    </Switch>
                </Columns.Column>
            </Columns>
        </>
    )
}


function HourlyBody(data) {
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
                Forecast for <Moment unix calendar={calendarStrings}>
                    {state.time}
                </Moment>
            </Heading>
            <Container>
                <WeatherData data={state} />
            </Container>
        </Section>
    )
}

export default withRouter(connect(mapStateToProps)(Hourly));
