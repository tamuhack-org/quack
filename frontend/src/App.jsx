import React, { Component } from "react";

const API_URL = process.env.API_URL;

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      response: "not loaded"
    };
  }

  componentDidMount() {
    const request_url = `${API_URL}` + "/test";
    console.log(request_url);
    fetch(request_url)
      .then(response => response.json())
      .then(this.setState({ response: "I received a message" }))
      .catch(err => console.log(err));
  }

  render() {
    return (
        <div className="App">
          <header className="App-header">
            <h1 className="App-title">Welcome to tamuhack's new registration thingy (here's another update)</h1>
          </header>
          <p className="App-intro">{this.state.response}</p>
        </div>
    );
  }
}

export default App;
