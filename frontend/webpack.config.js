const path = require("path");
const webpack = require("webpack");

const dev = process.env.NODE_ENV !== "production";

module.exports = {
  devServer: {
    host: "localhost",
    port: "3000",
    hot: true,
    headers: {
      "Access-Control-Allow-Origin": "*"
    },
    historyApiFallback: true
  },
  entry: ["react-hot-loader/patch", path.join(__dirname, "/src/index.jsx")],
  module: {
    rules: [
      {
        test: /\.jsx?$/,
        exclude: /node_modules/,
        loaders: ["babel-loader"]
      },
      {
        test: /\.scss$/,
        loader: "style-loader!css-loader!sass-loader"
      },
      {
        test: /\.(jpe?g|png|gif|svg)$/i,
        loader: "url-loader",
        options: {
          limit: 10000
        }
      }
    ]
  },
  resolve: {
    extensions: [".js", ".jsx"]
  },
  output: {
    filename: "index.js",
    path: path.join(__dirname, "/dist")
  },
  mode: dev ? "development" : "production",
  plugins: dev
    ? [
        new webpack.DefinePlugin({
          "process.env.NODE_ENV": JSON.stringify("development"),
          "process.env.API_URL": JSON.stringify(`${process.env.API_URL}`),
        })
      ]
    : [
        new webpack.DefinePlugin({
          "process.env.NODE_ENV": JSON.stringify("production"),
          "process.env.API_URL": JSON.stringify(`${process.env.API_URL}`),
        })
      ]
};
