// Proto.enumerate Proto.stringify
var Proto;
(function (Proto) {
    function type(data) {
        var type = typeof (data);
        return type != "object" ? type : data == null ? "null" : (data.constructor).toString().indexOf("Array") != -1 ? "array" : "structure";
    }

    function named(name) {
        return name == "" ? name : name + ": ";
    }

    var enumFlag = "_%$PROTO_ENUM$%";

    function unenumerate(data) {
        return data.length <= enumFlag.length ? '"' + data + '"' : data.substring(data.length - enumFlag.length) == enumFlag ? data.substr(0, data.length - enumFlag.length) : '"' + data + '"';
    }

    // mark enumeration value
    function enumerate(data) {
        return data + enumFlag;
    }
    Proto.enumerate = enumerate;

    function marshal(data, name) {
        switch (type(data)) {
            case "number":
            case "boolean":
                return named(name) + data.toString();

            case "string":
                return named(name) + unenumerate(data);

            case "array":
                var res = "";
                for (var index in data) {
                    res += marshal(data[index], name) + " ";
                }
                return res;

            case "structure":
                var res = "";
                for (var index in data) {
                    res += marshal(data[index], index) + " ";
                }
                return name == "" ? res : name + " { " + res + "}";

            default:
                return "";
        }
    }

    // stringify protocol data
    function stringify(name, data) {
        return { content: marshal(data, name) };
    }
    Proto.stringify = stringify;

    Proto.emptyProto = { content: "" };
})(Proto || (Proto = {}));
