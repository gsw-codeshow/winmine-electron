const { execSync } = require('child_process');

var newSquare = function(x, y, num, syncStd) {
    var l = []
    var len = x * y
    console.log(syncStd)
    for(var i = 0; i < x; i++) {
        var a = []
        for(var j = 0; j < y; j++){
            num = syncStd[i*x + j] - 48
            a.push(num)
        }
        l.push(a)
    }
    return l
}

var newArea = function(x, y, num) {
    syncStd = execSync('"winminermap.exe" '+"-n "+num);
    console.log(syncStd)
    return newSquare(x, y, num, syncStd)
}
