// Code generated by go-bindata.
// sources:
// web/templates/index.html
// web/templates/navigation_bar.html
// web/templates/second_view.html
// web/templates/third_view.html
// web/static/navigation_bar.css
// web/static/style.css
// web/static/third_view.css
// web/static/todolist.css
// DO NOT EDIT!

package assets

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _webTemplatesIndexHtml = []byte(`<!DOCTYPE html>
<html lang="en">
    <head>
        <!-- Required meta tags -->
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
        <title>ToDoList | index</title>

        
        <!-- Bootstrap CSS -->
        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/css/bootstrap.min.css" integrity="sha384-/Y6pD6FV/Vv2HJnA6t+vslU6fwYXjCFtcEpHbNJ0lyAFsXTsjBbfaDjzALeQsN6M" crossorigin="anonymous">

        <link rel="stylesheet" href="web/static/style.css">
        <link rel="stylesheet" href="web/static/navigation_bar.css">
        <link rel="stylesheet" href="web/static/todolist.css">
    </head>


    <body>
        <div class="container">
            {{.NavigationBar}}
        </div>

        



        <div class="container">
            <div id="myDIV" class="header">
                <h2>My To Do List</h2>
                <input type="text" id="myInput" placeholder="Write your task...">
                <span onclick="newElement()" class="addBtn">Add</span>
              </div>
          
              <ul id="myUL">
                <li onclick="checkTask(this)" class="nametask">Hit the gym</li>
                <li  onclick="checkTask(this)" class="nametask">Pay bills</li>
              </ul>
        </div>

        
    </body>



    <!-- Optional JavaScript -->
    <!-- jQuery first, then Popper.js, then Bootstrap JS -->
    <script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.11.0/umd/popper.min.js" integrity="sha384-b/U6ypiBEHpOf/4+1nzFpr53nxSS+GLCkfwBdFNTxtclqqenISfwAzpKaMNFNmj4" crossorigin="anonymous"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/js/bootstrap.min.js" integrity="sha384-h0AbiXch4ZDo7tp9hKZ4TsHbi047NrKGLO3SEJAg45jXxnGIfYzk4Si90RDIqNm1" crossorigin="anonymous"></script>
    <script>
     
        // Create a "close" button and append it to each list item
        var myNodelist = document.getElementsByClassName("nametask");
        var i;
        for (i = 0; i < myNodelist.length; i++) {
          var span = document.createElement("SPAN");
          var txt = document.createTextNode("\u00D7");
          span.className = "close";
          span.appendChild(txt);
          myNodelist[i].appendChild(span);
        }
        
        // Click on a close button to hide the current list item
        var close = document.getElementsByClassName("close");
        var i;
        for (i = 0; i < close.length; i++) {
          close[i].onclick = function() {
            var div = this.parentElement;
            div.style.display = "none";

            deleteinbackend(this.parentElement.id);  //se llama a la función para borrar el elemento en el backend
          }
        }
        
        ///////////////////////////////////////////////////////
		//////////////////////////////////////////////////////
        
		
		//Add a "checked" symbol when clicking on a list item and send the PUT request to the backend
		function checkTask(element) {
			element.className="checked";
			
			// solicitud a la API REST vía método PUT
            fetch(url + "/tasks/"+ element.id, {
            method: 'PUT',
            
            })
            // promesas encadenadas
            .then(res => res.text())
            // procesamiento de la respuesta
            .then(text => {
                // depuración en consola del navegador
                console.log(text);
            })
            // procesamiento de errores
            .catch(err => {
                // depuración en consola del navegador
                console.error(err);
            });
			 
		}
        
        /////////////////////////////////////////////////////////////
        /////////////////////////////////////////////////////////////
        
        const url = "http://localhost:8000";
        window.onload = handleGETtasksButton();

        function handleGETtasksButton() {   //muestra todas las tareas al cargar la pagina
            
            // solicitud a la API REST vía método GET
            fetch(url + "/tasks")
            // promesas encadenadas
            .then(res => res.json())
            // procesamiento de la respuesta
            .then(json => {
                // depuración en consola del navegador
                console.log(json);
                
                
                for(i=0; i<json.length; i++){
                    
                    var li = document.createElement("li");
                    var nombre = json[i].name;
                    
                    var t = document.createTextNode(nombre);
                    li.appendChild(t);
                    li.id = json[i].ID; //le asigna a cada etiqueta HTML <li> el ID de la task
					
					if(json[i].check_valid == true){
						li.className="checked";
					}
					li.setAttribute("onclick","checkTask(this)");
					
					
                    var span = document.createElement("SPAN");
                    var txt = document.createTextNode("\u00D7");
                    span.className = "close";
                    span.appendChild(txt);
                    li.appendChild(span);

                    document.getElementById("myUL").appendChild(li);
                    
                }	
                for (i = 0; i < close.length; i++) {
                  close[i].onclick = function() {
                    var div = this.parentElement;
                    div.style.display = "none";
                    deleteinbackend(this.parentElement.id);
                  }
                }
            
            })
			
            
        }	 
        
        ///////////////////////////////////////
        ///////////////////////////////////////
        

        // Create a new list item when clicking on the "Add" button
        function newElement() {
          var li = document.createElement("li");
          var taskname = document.getElementById("myInput").value;
          var t = document.createTextNode(taskname);
          li.appendChild(t);

          
          if (taskname === '') {
            alert("You must write something!");
          } else {
            document.getElementById("myUL").appendChild(li);


            // solicitud a la API REST vía método POST
            fetch(url + "/tasks", {
            method: 'POST',
            
            // se incluyen en el cuerpo del mensaje los datos del nuevo libro a registrar 
            body: JSON.stringify({            
                name: taskname
            })
            })
            // promesas encadenadas
            .then(res => res.text())
            // procesamiento de la respuesta
            .then(text => {
                // depuración en consola del navegador
                console.log(text);
                li.id=text; //le asigno el ID de la task devuelto por el backend a la etiqueta HTML <li>

                // conversión de respuesta de texto (html) a objeto json
                 //json = JSON.parse(text);
                 //volcado a pantalla de la respuesta
                 //movieId.innerHTML = json.id;
            })
            // procesamiento de errores
            .catch(err => {
                // depuración en consola del navegador
                console.error(err);
            });



          }
          document.getElementById("myInput").value = "";  //se vuelve a poner en blanco el textbox
        
          var span = document.createElement("SPAN");
          var txt = document.createTextNode("\u00D7");
          span.className = "close";
          span.appendChild(txt);
          li.appendChild(span);
        
          for (i = 0; i < close.length; i++) {
            close[i].onclick = function() {
              var div = this.parentElement;
              div.style.display = "none";
              deleteinbackend(this.parentElement.id);
            }
          }

        }


        /////////////////////////////
        ////////////////////////////

        function deleteinbackend(id){   //para borrar la task en el backend
            
            // solicitud a la API REST vía método POST
            fetch(url + "/tasks/"+id, {
            method: 'DELETE',
            
            })
            // promesas encadenadas
            .then(res => res.text())
            // procesamiento de la respuesta
            .then(text => {
                // depuración en consola del navegador
                console.log(text);
                
            })
            // procesamiento de errores
            .catch(err => {
                // depuración en consola del navegador
                console.error(err);
            });

        }
        </script>
        
    </html>`)

func webTemplatesIndexHtmlBytes() ([]byte, error) {
	return _webTemplatesIndexHtml, nil
}

func webTemplatesIndexHtml() (*asset, error) {
	bytes, err := webTemplatesIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/templates/index.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webTemplatesNavigation_barHtml = []byte(`<!-- Navigation Bar -->
<nav class="navbar navbar-light bg-light">
    <div class="container">
        <span class="navbar-brand">My ToDoList</span>

        <ul class="nav justify-content-end">
            <li class="nav-item">
                <a id="homeNav" class="nav-link" href="/">Home</a>
            </li>
            <!-- <li class="nav-item">
                <a id="secondNav" class="nav-link" href="/second">Second</a>
            </li>
            <li class="nav-item">
                <a id="thirdNav" class="nav-link" href="/third/1">Third</a>
            </li> -->
        </ul>
    </div>
</nav>`)

func webTemplatesNavigation_barHtmlBytes() ([]byte, error) {
	return _webTemplatesNavigation_barHtml, nil
}

func webTemplatesNavigation_barHtml() (*asset, error) {
	bytes, err := webTemplatesNavigation_barHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/templates/navigation_bar.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webTemplatesSecond_viewHtml = []byte(`<!DOCTYPE html>
<html lang="en">
    <head>
        <!-- Required meta tags -->
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
        <title>Golang HTML Server</title>

        <!-- Bootstrap CSS -->
        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/css/bootstrap.min.css" integrity="sha384-/Y6pD6FV/Vv2HJnA6t+vslU6fwYXjCFtcEpHbNJ0lyAFsXTsjBbfaDjzALeQsN6M" crossorigin="anonymous">

        <link rel="stylesheet" href="/static/style.css">
        <link rel="stylesheet" href="/static/navigation_bar.css">
    </head>
    <body>
        <div class="container">
            {{.NavigationBar}}

            <h1>Another View</h1>
            <h2>- Content Goes Here -</h2>
        </div>
    </body>

    <!-- Optional JavaScript -->
    <!-- jQuery first, then Popper.js, then Bootstrap JS -->
    <script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.11.0/umd/popper.min.js" integrity="sha384-b/U6ypiBEHpOf/4+1nzFpr53nxSS+GLCkfwBdFNTxtclqqenISfwAzpKaMNFNmj4" crossorigin="anonymous"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/js/bootstrap.min.js" integrity="sha384-h0AbiXch4ZDo7tp9hKZ4TsHbi047NrKGLO3SEJAg45jXxnGIfYzk4Si90RDIqNm1" crossorigin="anonymous"></script>
</html>`)

func webTemplatesSecond_viewHtmlBytes() ([]byte, error) {
	return _webTemplatesSecond_viewHtml, nil
}

func webTemplatesSecond_viewHtml() (*asset, error) {
	bytes, err := webTemplatesSecond_viewHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/templates/second_view.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webTemplatesThird_viewHtml = []byte(`<!DOCTYPE html>
<html lang="en">
    <head>
        <!-- Required meta tags -->
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
        <title>Golang HTML Server</title>

        <!-- Bootstrap CSS -->
        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/css/bootstrap.min.css" integrity="sha384-/Y6pD6FV/Vv2HJnA6t+vslU6fwYXjCFtcEpHbNJ0lyAFsXTsjBbfaDjzALeQsN6M" crossorigin="anonymous">

        <link rel="stylesheet" href="/static/style.css">
        <link rel="stylesheet" href="/static/navigation_bar.css">
        <link rel="stylesheet" href="/static/third_view.css">
    </head>
    <body>
        <div class="container">
            {{.NavigationBar}}

            <h1>Rendering Data</h1>
            <h4>This page takes the integer passed in and determines if it is odd or even</h4>
            <div class="result-box">
                {{if .StringQuery}}
                    <h2 class="result-underlined">You didn't enter an integer. This is what was entered:</h2>
                    <h3>{{.StringQuery}}</h3>
                {{else}}
                    <h2 class="result-underlined">The number entered is:</h2>
                    <h3>{{.Number}}</h3>
                    <h2 class="result-underlined">This number is:</h2>
                    <h3>{{.Number | formatOddOrEven}}</h3>
                {{end}}
            </div>
        </div>
    </body>

    <!-- Optional JavaScript -->
    <!-- jQuery first, then Popper.js, then Bootstrap JS -->
    <script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.11.0/umd/popper.min.js" integrity="sha384-b/U6ypiBEHpOf/4+1nzFpr53nxSS+GLCkfwBdFNTxtclqqenISfwAzpKaMNFNmj4" crossorigin="anonymous"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/js/bootstrap.min.js" integrity="sha384-h0AbiXch4ZDo7tp9hKZ4TsHbi047NrKGLO3SEJAg45jXxnGIfYzk4Si90RDIqNm1" crossorigin="anonymous"></script>
</html>`)

func webTemplatesThird_viewHtmlBytes() ([]byte, error) {
	return _webTemplatesThird_viewHtml, nil
}

func webTemplatesThird_viewHtml() (*asset, error) {
	bytes, err := webTemplatesThird_viewHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/templates/third_view.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webStaticNavigation_barCss = []byte(``)

func webStaticNavigation_barCssBytes() ([]byte, error) {
	return _webStaticNavigation_barCss, nil
}

func webStaticNavigation_barCss() (*asset, error) {
	bytes, err := webStaticNavigation_barCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/static/navigation_bar.css", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webStaticStyleCss = []byte(`h1, h2, h3, h4, h5 {
    text-align: center;
}

ul li {
    list-style:none;
}

`)

func webStaticStyleCssBytes() ([]byte, error) {
	return _webStaticStyleCss, nil
}

func webStaticStyleCss() (*asset, error) {
	bytes, err := webStaticStyleCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/static/style.css", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webStaticThird_viewCss = []byte(`.result-box {
    margin-top: 50px;
}

.result-underlined {
    text-decoration: underline;
}`)

func webStaticThird_viewCssBytes() ([]byte, error) {
	return _webStaticThird_viewCss, nil
}

func webStaticThird_viewCss() (*asset, error) {
	bytes, err := webStaticThird_viewCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/static/third_view.css", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webStaticTodolistCss = []byte(`/* Include the padding and border in an element's total width and height */
* {
  box-sizing: border-box;
}

/* Remove margins and padding from the list */
ul {
  margin: 0;
  padding: 0;
}

/* Style the list items */
ul li {
  cursor: pointer;
  position: relative;
  padding: 12px 8px 12px 40px;
  background: #eee;
  font-size: 18px;
  transition: 0.2s;

  /* make the list items unselectable */
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}

/* Set all odd list items to a different color (zebra-stripes) */
ul li:nth-child(odd) {
  background: #f9f9f9;
}

/* Darker background-color on hover */
ul li:hover {
  background: #ddd;
}

/* When clicked on, add a background color and strike out text */
ul li.checked {
  background: #888;
  color: #fff;
  text-decoration: line-through;
}

/* Add a "checked" mark when clicked on */
ul li.checked::before {
  content: '';
  position: absolute;
  border-color: #fff;
  border-style: solid;
  border-width: 0 2px 2px 0;
  top: 10px;
  left: 16px;
  transform: rotate(45deg);
  height: 15px;
  width: 7px;
}


.nametask {
  list-style:none;
}

/* Style the close button */
.close {
  position: absolute;
  right: 0;
  top: 0;
  padding: 12px 16px 12px 16px;
}

.close:hover {
  background-color: #f44336;
  color: white;
}

/* Style the header */
.header {
  background-color: #47309C;
  padding: 30px 40px;
  color: white;
  text-align: center;
}

/* Clear floats after the header */
.header:after {
  content: "";
  display: table;
  clear: both;
}

/* Style the input */
input {
  margin: 0;
  border: none;
  border-radius: 0;
  width: 75%;
  padding: 10px;
  float: left;
  font-size: 16px;
}

/* Style the "Add" button */
.addBtn {
  padding: 10px;
  width: 25%;
  background: #d9d9d9;
  color: #555;
  float: left;
  text-align: center;
  font-size: 16px;
  cursor: pointer;
  transition: 0.3s;
  border-radius: 0;
}

.addBtn:hover {
  background-color: #bbb;
}`)

func webStaticTodolistCssBytes() ([]byte, error) {
	return _webStaticTodolistCss, nil
}

func webStaticTodolistCss() (*asset, error) {
	bytes, err := webStaticTodolistCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/static/todolist.css", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"web/templates/index.html": webTemplatesIndexHtml,
	"web/templates/navigation_bar.html": webTemplatesNavigation_barHtml,
	"web/templates/second_view.html": webTemplatesSecond_viewHtml,
	"web/templates/third_view.html": webTemplatesThird_viewHtml,
	"web/static/navigation_bar.css": webStaticNavigation_barCss,
	"web/static/style.css": webStaticStyleCss,
	"web/static/third_view.css": webStaticThird_viewCss,
	"web/static/todolist.css": webStaticTodolistCss,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"web": &bintree{nil, map[string]*bintree{
		"static": &bintree{nil, map[string]*bintree{
			"navigation_bar.css": &bintree{webStaticNavigation_barCss, map[string]*bintree{}},
			"style.css": &bintree{webStaticStyleCss, map[string]*bintree{}},
			"third_view.css": &bintree{webStaticThird_viewCss, map[string]*bintree{}},
			"todolist.css": &bintree{webStaticTodolistCss, map[string]*bintree{}},
		}},
		"templates": &bintree{nil, map[string]*bintree{
			"index.html": &bintree{webTemplatesIndexHtml, map[string]*bintree{}},
			"navigation_bar.html": &bintree{webTemplatesNavigation_barHtml, map[string]*bintree{}},
			"second_view.html": &bintree{webTemplatesSecond_viewHtml, map[string]*bintree{}},
			"third_view.html": &bintree{webTemplatesThird_viewHtml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

