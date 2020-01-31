import React, { Component, Suspense } from 'react';
import { connect } from 'react-redux';
import { Switch, Route } from 'react-router-dom'
import LoadingBlock from './LoadingBlock';

const mapStateToProps = state => {
    const { isLoaded } = state.forecast
    return { isLoaded };
}

// Load all components lazily.
const Current = React.lazy(() => import('./Current'));
const Hourly = React.lazy(() => import('./Hourly'));
const Daily = React.lazy(() => import('./Daily'));
const NotFound = React.lazy(() => import('./NotFound'));
const Charts = React.lazy(() => import('./Charts'));

class AppBody extends Component {
    render() {
        const { isLoaded } = this.props;

        if (isLoaded) {
            return (
                <Suspense fallback={LoadingBlock}>
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
                </Suspense>
            )
        }

        return null;
    }
}

export default connect(mapStateToProps)(AppBody);
