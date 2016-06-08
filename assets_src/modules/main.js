import React from "react";
import Indicator from "./indicator";
import LogFileList from "./log_file_list";
import LogContent from "./log_content";
import WindowMode from "./WindowMode";

export default class Main extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            selectedHost: "",
            connected: false,
            windowMode: 1
        };
    }

    handleSelectHost(host) {
        if ( this.state.selectedHost === host ) {
            host = "";
        }
        this.setState({selectedHost: host});
    }

    handleChangeWindowMode(mode) {
        this.setState({windowMode: mode});
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

        let content = "";
        switch ( this.state.windowMode ) {
            case 1:
                content = (
                    <div className="T-Main">
                        <LogFileList hosts={hosts} selectedHost={this.state.selectedHost} onSelectHost={this.handleSelectHost.bind(this)} />
                        <LogContent logs={logs} flushLog={this.flushLog.bind(this)}/>
                    </div>
                );
            break;
            case 2:
                content = (
                    <div className="T-Main T-Main--split2">
                        <SplitWindow hosts={hosts} logs={this.props.logs} />
                        <SplitWindow hosts={hosts} logs={this.props.logs} />
                    </div>
                );
            break;
            case 3:
                content = (
                    <div className="T-Main T-Main--split4">
                        <SplitWindow hosts={hosts} logs={this.props.logs} />
                        <SplitWindow hosts={hosts} logs={this.props.logs} />
                        <SplitWindow hosts={hosts} logs={this.props.logs} />
                        <SplitWindow hosts={hosts} logs={this.props.logs} />
                    </div>
                );
            break;
        }

        return (
            <div className="T-Window">
                <div className="T-Header">
                    <h1 className="T-Title">Tailor</h1>
                    <WindowMode selected={this.state.windowMode} handler={this.handleChangeWindowMode.bind(this)}/>
                    <Indicator connected={this.state.connected} />
                </div>
                {content}
            </div>
        );
    }
}
