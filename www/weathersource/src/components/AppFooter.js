import React, { Component } from 'react';
import { Footer, Content, Image, } from 'react-bulma-components';

class AppFooter extends Component {
    render() {
        return (
            <Footer backgroundColor={"info"}>
                <Content>
                    <Image
                        src="https://darksky.net/images/darkskylogo.png"
                        alt="Dark Sky Logo"
                        size={48}
                        pull={"left"}
                        marginless={true}
                    />
                    <a href="https://darksky.net/poweredby/" pull={"left"} className="has-text-light">
                        <strong>Powered by Dark Sky</strong>
                    </a>
                </Content>
            </Footer >
        );
    }
}

export default AppFooter;
