const CracoLessPlugin = require("craco-less");

module.exports = {
  plugins: [
    {
      plugin: CracoLessPlugin,
      options: {
        lessLoaderOptions: {
          lessOptions: {
            modifyVars: {"@primary-color": "rgb(128,128,128)"},
            javascriptEnabled: true,
          },
        },
      },
    },
  ],
  module: {
    rules: [
      {
        test: /\.less$/,
        use: [
          "style-loader",
          "css-loader",
          "less-loader",
        ],
      },
    ],
  },
  webpack: {
    configure: {
      resolve: {
        fallback: {
          "path": false,
        },
      },
    },
  },
};
