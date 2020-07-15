var redis = require('redis');
RDS_PORT = 6379;        //端口号
RDS_HOST = '127.0.1.1';    //服务器IP
RDS_PWD = '';    //密码
RDS_OPTS = {};            //设置项
/*
client = redis.createClient(RDS_PORT, RDS_HOST, RDS_OPTS);
client.auth(RDS_PWD, function () {
    console.log('通过认证');
});
client.on('connect', function () {
    client.hmset('short', { 'js': 'javascript', 'C#': 'C Sharp' }, redis.print);
    client.hmset('short', 'SQL', 'Structured Query Language', 'HTML', 'HyperText Mark-up Language', redis.print);
    client.hgetall("short", function (err, res) {
        if (err) {
            console.log('Error:' + err);
            return;
        }
        console.dir(res);
    });
});
client.on('ready', function (err) {
    console.log('ready');
});
*/

function getRedisClient() {
    client = redis.createClient(RDS_PORT, RDS_HOST, RDS_OPTS);
    // client.auth(RDS_PWD, function () {
    //     console.log('通过认证');
    //     return client;
    // });
    return null;
}

module.exports.getRedisClient = getRedisClient;