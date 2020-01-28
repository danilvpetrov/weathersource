import React, { Component } from 'react';
import { connect } from 'react-redux';
import { NavLink } from 'react-router-dom'
import { fetchForecast } from '../actions/forecast';
import { Navbar } from 'react-bulma-components';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import {
    faSync, faCloudSun, faClock,
    faHourglass, faCalendarDay
} from '@fortawesome/free-solid-svg-icons'


const mapStateToProps = state => {
    const { isLoading, isLoaded, error, data } = state.forecast
    return { isLoading, isLoaded, error, data }
}

const mapDispatchToProps = dispatch => ({
    fetchForecast: () => dispatch(fetchForecast())
})


class AppNavbar extends Component {
    refresh = (event) => {
        this.props.fetchForecast();
    }

    // set active state for hamburger
    state = { active: false }

    handleBurgerClick = () => {
        const { active } = this.state;
        this.setState({ active: !active });
    }

    render() {
        return (
            <Navbar color={"info"} fixed={"top"} active={this.state.active}>
                <Navbar.Brand>
                    <Navbar.Item renderAs="a" href="/">
                        <span style={{ fontSize: 28, marginRight: 5 }}><FontAwesomeIcon icon={faCloudSun} /></span>
                        <strong>Weather Forecast</strong>
                    </Navbar.Item>
                    <Navbar.Burger
                        active={this.state.active.toString()}
                        onClick={this.handleBurgerClick}
                    />
                </Navbar.Brand>
                <Navbar.Menu >
                    <Navbar.Container>
                        <NavLink className="navbar-item" to="/current" activeStyle={{ backgroundColor: '#118fe4', color: '#fff' }}>
                            <span style={{ fontSize: 18, marginRight: 5 }}><FontAwesomeIcon icon={faClock} /></span>
                            Current
                        </NavLink>
                        <NavLink className="navbar-item" to="/hourly" activeStyle={{ backgroundColor: '#118fe4', color: '#fff' }}>
                            <span style={{ fontSize: 18, marginRight: 5 }}><FontAwesomeIcon icon={faHourglass} /></span>
                            Hourly forecast
                        </NavLink>
                        <NavLink className="navbar-item" to="/daily" activeStyle={{ backgroundColor: '#118fe4', color: '#fff' }}>
                            <span style={{ fontSize: 18, marginRight: 5 }}><FontAwesomeIcon icon={faCalendarDay} /></span>
                            Daily forecast
                        </NavLink>
                    </Navbar.Container>
                    <Navbar.Container position="end">
                        <Navbar.Item onClick={this.refresh} renderAs="a">
                            <FontAwesomeIcon icon={faSync} />
                        </Navbar.Item>
                    </Navbar.Container>
                </Navbar.Menu>
            </Navbar>
        );
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(AppNavbar);
