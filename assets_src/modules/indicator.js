import React from "react";

export default class Indicator extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        if ( this.props.connected === true ) {
            return (
                <span className="T-Indicator c">Connected</span>
            );
        } else {
            return (
                <span className="T-Indicator n">Not Connected</span>
            );
        }
    }
}
