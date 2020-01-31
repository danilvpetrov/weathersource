import React, { Component } from 'react';
import { connect } from 'react-redux';
import LoadingBlock from './LoadingBlock';


const mapStateToProps = state => {
    const { isLoading } = state.forecast
    return { isLoading };
}

class AppLoader extends Component {
    render() {
        const isLoading = this.props.isLoading;

        if (isLoading) {
            return (
                <LoadingBlock />
            )
        }

        return null;
    }
}

export default connect(mapStateToProps)(AppLoader);
