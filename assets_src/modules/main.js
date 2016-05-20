import React from "react";
import Indicator from "./indicator";
import LogFileList from "./log_file_list";
import LogContent from "./log_content";

export default class Main extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            selectedHost: "",
            connected: false
        };
    }

    handleSelectHost(host) {
        if ( this.state.selectedHost === host ) {
            host = "";
        }
        this.setState({selectedHost: host});
    }

    flushLog() {
        if ( this.state.selectedHost && this.state.selectedHost in this.props.logs ) {
            this.props.logs[this.state.selectedHost] = [];
            this.forceUpdate();
        }
    }

    render() {
        let logs = [];
        const hosts = Object.keys(this.props.logs);

        if ( this.state.selectedHost && this.state.selectedHost in this.props.logs ) {
            logs = this.props.logs[this.state.selectedHost];
        } else {
            if ( hosts.length > 0 ) {
                logs = this.props.logs[hosts[0]];
                this.state.selectedHost = hosts[0];
            }
        }

        return (
            <div className="T-Window">
                <div className="T-Header">
                    <h1 className="T-Title">Tailor</h1>
                    <Indicator connected={this.state.connected} />
                </div>
                <div className="T-Main">
                    <LogFileList hosts={hosts} selectedHost={this.state.selectedHost} onSelectHost={this.handleSelectHost.bind(this)} />
                    <LogContent logs={logs} flushLog={this.flushLog.bind(this)}/>
                </div>
            </div>
        );
    }
}
