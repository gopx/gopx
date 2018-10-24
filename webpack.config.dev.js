const path = require('path');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');

const viewsDir = path.resolve(__dirname, './web/views');
const publicDir = path.resolve(__dirname, './web/public');

module.exports = {
  entry: {
    home: path.resolve(viewsDir, 'home/index.js'),
    login: path.resolve(viewsDir, 'login/index.js'),
    signup: path.resolve(viewsDir, 'signup/index.js'),
    search: path.resolve(viewsDir, 'search/index.js'),
    profile: path.resolve(viewsDir, 'profile/index.js'),
    package: path.resolve(viewsDir, 'package/index.js'),
    error: path.resolve(viewsDir, 'error/index.js'),
  },
  output: {
    path: publicDir,
    filename: 'js/[name].js',
  },
  devtool: 'source-map',
  module: {
    rules: [
      {
        test: /\.(s*)css$/,
        use: [
          MiniCssExtractPlugin.loader,
          'css-loader',
          'sass-loader',
        ],
      },
    ],
  },
  plugins: [
    new MiniCssExtractPlugin({
      filename: 'css/[name].css',
    }),
  ],
};
