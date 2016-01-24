const React = require("react");

export default class Main extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            selectedFile: "",
            connected: false,
            logs: []
        };
    }

    render() {
        return (
            <div className="T-Window">
                <div className="T-Header">
                    <h1 class="T-Title">Tailor</h1>
                    <Indicator connected={this.state.connected} />
                </div>
                <div class="T-Main">
                    <LogFileList selectedFile={this.state.selectedFile} />
                    <LogContent logs={this.state.logs} />
                </div>
            </div>
        );
    }
}
