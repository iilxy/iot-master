const curd = require_plugin("curd");
exports.post = curd.list("device", {
    name: 'string',
    type: 'string',
}, ctx => {
    //强制修改为
    const filter = {key: "tunnel_id", value: ctx.params._id};
    ctx.body.filter = ctx.body.filter || [];
    ctx.body.filter.push(filter);
});