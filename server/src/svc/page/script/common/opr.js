var Opr;
(function (Opr) {
    function result(data, status, callback) {
        if (typeof data != "object") {
            alert(data);
            return;
        }

        if (data.ErrorMsg && data.ErrorMsg != "") {
            alert(data.ErrorMsg);
            return;
        }

        if (!callback) {
            alert("Opr.result() error");
            return;
        }

        callback(data);
    }
    Opr.result = result;
})(Opr || (Opr = {}));