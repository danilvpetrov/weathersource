import React, { Component } from 'react';
import { useLocation } from 'react-router-dom'
import { Section, Notification } from 'react-bulma-components';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faExclamationCircle } from '@fortawesome/free-solid-svg-icons'

class NotFound extends Component {
    render() {
        return <NotFoundBody />
    }
}

function NotFoundBody() {
    let location = useLocation();

    return (
        <Section>
            <Notification color="danger" style={{ marginTop: 150 }}>
                <span style={{ fontSize: 28, marginRight: 5 }}><FontAwesomeIcon icon={faExclamationCircle} /></span>Nothing found for the path <code>{location.pathname}</code>.
            </Notification>
        </Section>
    )
}

export default NotFound;
