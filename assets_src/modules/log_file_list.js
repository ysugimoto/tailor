import React from "react";

export default class LogFileList extends React.Component {
    constructor(props) {
        super(props);
    }

    selectHost(evt) {
        const em = evt.currentTarget.querySelector("em");
        em.normalize();
        this.props.onSelectHost(em.textContent);
    }

    render() {
        return (
            <div className="T-Sidebar">
                <ul className="T-LogList">
                    { this.props.hosts.map((h) => {
                        if ( h === this.props.selectedHost ) {
                            return (
                                <li onClick={this.selectHost.bind(this)}>
                                    <a href="#" className="active">
                                        <em>{h}</em>
                                    </a>
                                </li>
                            );
                        } else {
                            return (
                                <li onClick={this.selectHost.bind(this)}>
                                    <a href="#">
                                        <em>{h}</em>
                                    </a>
                                </li>
                            );
                        }
                    }) }
                </ul>
            </div>
        );
    }
}
