import React, { Component } from 'react';
import { Section, Columns, Loader } from 'react-bulma-components';

class LoadingBlock extends Component {
    render() {
        return (
            <Section>
                <Columns centered={true} >
                    <Columns.Column size={"half"}>
                        <Loader
                            style={{
                                margin: 'auto',
                                width: 200,
                                height: 200,
                                border: '20px solid #209cee',
                                borderTopColor: 'transparent',
                                borderRightColor: 'transparent'
                            }}
                        />
                    </Columns.Column>
                </Columns>
            </Section>
        )
    }
}

export default LoadingBlock;
