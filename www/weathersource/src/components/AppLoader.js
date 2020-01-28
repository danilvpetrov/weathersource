import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Section, Columns, Loader } from 'react-bulma-components';


const mapStateToProps = state => {
    const { isLoading } = state.forecast
    return { isLoading };
}

class AppLoader extends Component {
    render() {
        const isLoading = this.props.isLoading;

        if (isLoading) {
            return (
                <Section>
                    <Columns centered={true} >
                        <Columns.Column size={"half"}>
                            <Loader
                                style={{ margin: 'auto', width: 200, height: 200, border: '20px solid #209cee', borderTopColor: 'transparent', borderRightColor: 'transparent' }}
                            />
                        </Columns.Column>
                    </Columns>
                </Section>
            )
        }

        return null;
    }
}

export default connect(mapStateToProps)(AppLoader);
