const curd = require_plugin("curd");
exports.post = curd.list("device", {
    name: 'string',
    type: 'string',
});