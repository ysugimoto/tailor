import ReactDOM from "react-dom";
import React    from "react";
import Main     from "./modules/main";

class Controller {
    constructor(element) {
        this.element = element;
        this.element.style.height = (window.innerHeight - 40) + "px";

        this.logs = {};
    }

    init() {
        this.main = ReactDOM.render(<Main logs={this.logs} hostList={this.hosts}/>, this.element);
        this.ws   = new WebSocket("ws://" + location.host + "/reader");
        this.ws.onopen = () => {
            this.main.setState({connected: true});
            this.ws.onmessage = (evt) => this.handleSocket(evt);
        };
        this.ws.onclose = () => {
            this.main.setState({connected: false});
            this.ws.onmessage = null;
        };
    }

    handleSocket(evt) {
        const message = JSON.parse(evt.data);
        console.log(message);

        if ( message.host in this.logs ) {
            this.logs[message.host].push(message.message);
        } else {
            this.logs[message.host] = [message.message];
        }
        this.main.setState({
            logs: this.logs
        });
    }
}

const c = new Controller(
    document.querySelector(".T-Wrap")
);
c.init();

