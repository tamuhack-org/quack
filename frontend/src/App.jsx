import React, { Component } from 'react';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      response: '',
    };
  }

  componentDidMount() {
    fetch('http://localhost:8080')
      .then(response => response.json())
      .then(res => this.setState({ response: res.express }))
      .catch(err => console.log(err));
  }

  // callApi() {
  //   fetch("api/hello
  // }
  // callApi = async () => {
  //   const response = await fetch("/api/hello");
  //   const body = await response.json();
  //
  //   if (response.status !== 200) throw Error(body.message);
  //
  //   return body;
  // };

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <h1 className="App-title">Welcome to React</h1>
        </header>
        <p className="App-intro">{this.state.response}</p>
      </div>
    );
  }
}

export default App;
