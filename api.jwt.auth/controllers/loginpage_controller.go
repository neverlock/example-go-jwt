package controllers

import (
	"net/http"
)

func LoginPageController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Write([]byte("<html><head>"))
	w.Write([]byte("<script type='text/javascript'>\n"))
	//w.Write([]byte("function sendjs(){\n"))
	w.Write([]byte("console.log(\"test\")\n"))
	w.Write([]byte("var url = \"http://localhost:5000/token-auth\";\n"))
	w.Write([]byte("var data = {};\n"))
	w.Write([]byte("data.Username = \"haku\";\n"))
	w.Write([]byte("data.Password  = \"testing\";\n"))
	w.Write([]byte("var json = JSON.stringify(data);\n"))
	w.Write([]byte("var xhr = new XMLHttpRequest();\n"))
	w.Write([]byte("xhr.open(\"POST\", url, true);\n"))
	w.Write([]byte("xhr.setRequestHeader('Content-type','application/json; charset=utf-8');\n"))
	w.Write([]byte("xhr.onload = function () {\n"))
	w.Write([]byte("var users = JSON.parse(xhr.responseText);\n"))
	w.Write([]byte("if (xhr.readyState == 4 && xhr.status == \"200\") {\n"))
	w.Write([]byte("console.table(users);\n"))
	w.Write([]byte("} else {\n"))
	//w.Write([]byte("console.error(users);\n"))
	w.Write([]byte("console.log('error');\n"))
	w.Write([]byte("}\n"))
	w.Write([]byte("}\n"))

	w.Write([]byte("xhr.send(json);\n"))
	//w.Write([]byte("}\n")) //end function
	w.Write([]byte("</script></head><body>"))
	w.Write([]byte("<form method='POST' enctype='application/json' action='/token-auth'>"))
	w.Write([]byte("<input name='Username' value='haku'>"))
	w.Write([]byte("<input name='Password' value='testing'>"))
	w.Write([]byte("<button type='submit'>Submit</button>"))
	w.Write([]byte("</body></html>"))
}

/*
function sendjs() {
    var url = "http://localhost:5000/token-auth";
    var data = {};
    data.Username = "haku";
    data.Password = "testing";
    var json = JSON.stringify(data);
    var xhr = new XMLHttpRequest();
    xhr.open("POST", url, true);
    xhr.setRequestHeader('Content-type', 'application/json; charset=utf-8');
    xhr.onload = function() {
        var users = JSON.parse(xhr.responseText);
        if (xhr.readyState == 4 && xhr.status == "201") {
            console.table(users);
        } else {
            console.error(users);
        }
    }
    xhr.send(json);
}
*/
