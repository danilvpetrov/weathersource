import React, { Component } from 'react';
import { Section, Columns, Loader, Heading } from 'react-bulma-components';

class LoadingBlock extends Component {
    render() {
        const { text } = this.props
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
                        <Section>
                            <Heading subtitle size={4} renderAs="h6" textAlignment="centered">
                                {text}
                            </Heading>
                        </Section>
                    </Columns.Column>
                </Columns>
            </Section>
        )
    }
}

export default LoadingBlock;
