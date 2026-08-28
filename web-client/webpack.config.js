var path = require("path");

module.exports = {
    mode: 'development',
    entry: './src/index.js',
    output: {
		path: path.resolve(__dirname, 'public'),
		filename: "index.js"
	},
    // Generated proto JS stubs require 'annotations_pb.js' due to HTTP transcoding
    // imports in the proto definitions. Since annotations are not used in the web client,
    // we alias the module to `false` so Webpack ignores it and stubs it with an empty object.
    resolve: {
        alias: {
            '../google/api/annotations_pb.js': false,
        }
    }
};
