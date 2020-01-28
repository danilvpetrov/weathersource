import React, { Component } from 'react';
import { connect } from 'react-redux';
import Moment from 'react-moment';
import { Section, Content, Notification } from 'react-bulma-components';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faExclamationTriangle } from '@fortawesome/free-solid-svg-icons'


const mapStateToProps = state => {
    const { isLoading, isLoaded, error, data } = state.forecast
    return { isLoading, isLoaded, error, data };
}

class AppNotify extends Component {
    render() {
        const { isLoading, error } = this.props;

        if (!isLoading && error) {
            return (
                <Section>
                    <Notification color="danger" style={{ marginTop: 150 }}>
                        <span style={{ fontSize: 28, marginRight: 5 }}>
                            <FontAwesomeIcon icon={faExclamationTriangle} />
                        </span>
                        Error loading weather data. {error}
                    </Notification>
                </Section>
            )
        }

        const { isLoaded, data } = this.props;

        if (isLoaded) {
            const { current } = data
            return (
                <Content size="small">
                    <Notification color="info">
                        Current weather and forecast are updated
                        for <a target="blank" href={`https://www.google.com/maps?q=${current.latitude},${current.longitude}`}>this
                        location</a> <Moment unix tz={current.timezone} fromNow>{current.time}</Moment>
                    </Notification>
                </Content>
            )
        }

        return null;
    }

}

export default connect(mapStateToProps)(AppNotify);
