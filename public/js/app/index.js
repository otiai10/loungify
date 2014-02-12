var Lounge = {};
Lounge.__socket = null;
Lounge.socket = function(opt){
    if (Lounge.__socket === null) {
        Lounge.__socket = new WebSocket('ws://'+window.location.host+'/websocket/socket');
    }
    return Lounge.__socket;
};
$(function(){
    Lounge.socket().onopen = function(e){
        Lounge.socket().send("Hello, server!");
    };
    Lounge.socket().onmessage = function(e){
        console.log("return from server,", e);
    };
});
