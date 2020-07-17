var rw = require("./redis_wrapper.js")

client = rw.getRedisClient();

if (client == null){
    console.log("client == null");
    return;
}

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