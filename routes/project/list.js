const curd = require_plugin("curd");
exports.post = curd.list("project", {
    name: 'string',
    type: 'string',
    address: 'string',
    port: 'number',
});