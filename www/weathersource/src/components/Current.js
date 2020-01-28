import React, { Component } from 'react';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom'
import WeatherData from './Weatherdata'
import { Section, Container, Heading } from 'react-bulma-components';

const mapStateToProps = state => {
    const { data } = state.forecast
    const { current } = data;
    return { current }
}

class Current extends Component {
    render() {
        const { current } = this.props;
        return (
            <Section>
                <Heading size={3}>Current Weather</Heading>
                <Container>
                    <WeatherData data={current} />
                </Container>
            </Section>
        )
    }
}

export default withRouter(connect(mapStateToProps)(Current));
