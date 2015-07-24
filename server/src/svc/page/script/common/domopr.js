// DomOpr.disabled DomOpr.undisabled DomOpr.property DomOpr.select.selected DomOpr.textarea.readLine
var DomOpr;
(function (DomOpr) {
    function jquery(obj) {
        return obj[0] ? obj : $(obj);
    }

    function dom(obj) {
        return obj[0] ? obj[0] : obj;
    }

    // diasbled element
    function disabled(obj) {
        dom(obj).disabled = true;
    }
    DomOpr.disabled = disabled;

    // undisabled element
    function undisabled(obj) {
        dom(obj).disabled = false;
    }
    DomOpr.undisabled = undisabled;

    //
    function property(obj, name) {
        return dom(obj).getAttribute(name);
    }
    DomOpr.property = property;

    (function (select) {
        // get selected element
        function selected(obj) {
            var ele = dom(obj);
            return ele.options[ele.selectedIndex];
        }
        select.selected = selected;
    })(DomOpr.select || (DomOpr.select = {}));
    var select = DomOpr.select;

    (function (textarea) {
        // read line
        function readLine(obj) {
            var value = dom(obj).value;
            var texts = value.split("\n");
            var array = new Array;
            for (var index in texts) {
                var text = texts[index].trim();
                if (text != "") {
                    array.push(text);
                }
            }
            return array;
        }
        textarea.readLine = readLine;
    })(DomOpr.textarea || (DomOpr.textarea = {}));
    var textarea = DomOpr.textarea;
})(DomOpr || (DomOpr = {}));