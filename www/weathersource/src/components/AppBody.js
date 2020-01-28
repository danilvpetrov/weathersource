import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Switch, Route } from 'react-router-dom'
import Current from './Current'
import Charts from './Charts'
import Hourly from './Hourly';
import Daily from './Daily';
import NotFound from './NotFound';

const mapStateToProps = state => {
    const { isLoaded } = state.forecast
    return { isLoaded };
}

class AppBody extends Component {
    render() {
        const { isLoaded } = this.props;

        if (isLoaded) {
            return (
                <Switch>
                    <Route path="/hourly">
                        <Hourly />
                    </Route>
                    <Route path="/daily">
                        <Daily />
                    </Route>
                    <Route path="/current">
                        <Current />
                    </Route>
                    <Route path="/charts">
                        <Charts />
                    </Route>
                    <Route exact path="/">
                        <Current />
                    </Route>
                    <Route path="*">
                        <NotFound />
                    </Route>
                </Switch>
            )
        }

        return null;
    }
}

export default connect(mapStateToProps)(AppBody);
