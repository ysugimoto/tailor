const ReactDOM = require("react-dom");
const Main     = require("./modules/main");

class Controller {
    constructor(element) {
        this.element = element;
        this.elemment.style.height = (window.innerHeight - 40) + "px";
    }

    init() {
        this.main = ReactDOM.render(main, <Main />);
        this.ws   = new WebSocket(WEBSOCKET_URL);
        this.ws.onmessage = (message) => {
            this.main.setState({log: message});
        };
    }
}

const controller = new Controller(
    document.querySelector(".T-Main")
);

