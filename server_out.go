package main 
import (
			"net/http"
			"time"
			"github.com/gorilla/sessions"
			"errors"
			"github.com/cheikhshift/db"
			"github.com/elazarl/go-bindata-assetfs"
			"bytes"
			"encoding/json"
			"fmt"
			"html"
			"html/template"
			"github.com/fatih/color"
			"strings"
			"reflect"
			"unsafe"
		)
				var store = sessions.NewCookieStore([]byte("a very very very very secret key"))

				type NoStruct struct {
					/* emptystruct */
				}

				func net_sessionGet(key string,s *sessions.Session) string {
					return s.Values[key].(string)
				}


				func net_sessionDelete(s *sessions.Session) string {
						//keys := make([]string, len(s.Values))

						//i := 0
						for k := range s.Values {
						   // keys[i] = k.(string)
						    net_sessionRemove(k.(string), s)
						    //i++
						}

					return ""
				}

				func net_sessionRemove(key string,s *sessions.Session) string {
					delete(s.Values, key)
					return ""
				}
				func net_sessionKey(key string,s *sessions.Session) bool {					
				 if _, ok := s.Values[key]; ok {
					    //do something here
				 		return true
					}

					return false
				}

				 

			 func net_add(x,v float64) float64 {
					return v + x				   
			 }

			 func net_subs(x,v float64) float64 {
				   return v - x
			 }

			 func net_multiply(x,v float64) float64 {
				   return v * x
			 }

			 func net_divided(x,v float64) float64 {
				   return v/x
			 }

	

				func net_sessionGetInt(key string,s *sessions.Session) interface{} {
					return s.Values[key]
				}

				func net_sessionSet(key string, value string,s *sessions.Session) string {
					 s.Values[key] = value
					 return ""
				}
				func net_sessionSetInt(key string, value interface{},s *sessions.Session) string {
					 s.Values[key] = value
					 return ""
				}

				func net_importcss(s string) string {
					return "<link rel=\"stylesheet\" href=\"" + s + "\" /> "
				}

				func net_importjs(s string) string {
					return "<script type=\"text/javascript\" src=\"" + s + "\" ></script> "
				}



				func formval(s string, r*http.Request) string {
					return r.FormValue(s)
				}
			
				func renderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, p *Page)  bool {
				     filename :=  tmpl  + ".tmpl"
				    body, err := Asset(filename)
				    session, er := store.Get(r, "session-")

				 	if er != nil {
				           session,er = store.New(r,"session-")
				    }
				    p.Session = session
				    p.R = r
				    if err != nil {
				      // fmt.Print(err)
				    	return false
				    } else {
				    t := template.New("PageWrapper")
				    t = t.Funcs(template.FuncMap{"a":net_add,"s":net_subs,"m":net_multiply,"d":net_divided,"js" : net_importjs,"css" : net_importcss,"sd" : net_sessionDelete,"sr" : net_sessionRemove,"sc": net_sessionKey,"ss" : net_sessionSet,"sso": net_sessionSetInt,"sgo" : net_sessionGetInt,"sg" : net_sessionGet,"form" : formval,"eq": equalz, "neq" : nequalz, "lte" : netlt,"UserSpaceInterface" : net_UserSpaceInterface,"UserSpace" : net_structUserSpace,"isUserSpace" : net_castUserSpace})
				    t, _ = t.Parse(strings.Replace(strings.Replace(strings.Replace(BytesToString(body), "/{", "\"{",-1),"}/", "}\"",-1 ) ,"`", `\"` ,-1) )
				    outp := new(bytes.Buffer)
				    error := t.Execute(outp, p)
				    if error != nil {
				   // fmt.Print(error)
				    	 http.Redirect(w,r,"/your-500-page",301)
				    return false
				    }  else {
				    p.Session.Save(r, w)

				    fmt.Fprintf(w, html.UnescapeString(outp.String()) )
				    return true
					}
				    }
				}

				func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
				  return func(w http.ResponseWriter, r *http.Request) {
				  	if !apiAttempt(w,r) {
				      fn(w, r, "")
				  	}
				  }
				} 

				func mHandler(w http.ResponseWriter, r *http.Request) {
				  	
				  	if !apiAttempt(w,r) {
				      handler(w, r, "")
				  	}
				  
				} 
				func mResponse(v interface{}) string {
					data,_ := json.Marshal(&v)
					return string(data)
				}
				func apiAttempt(w http.ResponseWriter, r *http.Request) bool {
					session, er := store.Get(r, "session-")
					response := ""
					if er != nil {
						session,_ = store.New(r, "session-")
					}
					callmet := false

					 
				if  r.URL.Path == "/service" && r.Method == strings.ToUpper("GET") { 
					
			Increase()

			w.Header().Set("Content-Type",  "text/html")
			w.Write([]byte(fmt.Sprintf("This is the %v request", Counter) ))
		
					callmet = true
				}
				

					if callmet {
						session.Save(r,w)
						if response != "" {
							//Unmarshal json
							w.Header().Set("Access-Control-Allow-Origin", "*")
							w.Header().Set("Content-Type",  "application/json")
							w.Write([]byte(response))
						}
						return true
					}
					return false
				}
				func SetField(obj interface{}, name string, value interface{}) error {
					structValue := reflect.ValueOf(obj).Elem()
					structFieldValue := structValue.FieldByName(name)

					if !structFieldValue.IsValid() {
						return fmt.Errorf("No such field: %s in obj", name)
					}

					if !structFieldValue.CanSet() {
						return fmt.Errorf("Cannot set %s field value", name)
					}

					structFieldType := structFieldValue.Type()
					val := reflect.ValueOf(value)
					if structFieldType != val.Type() {
						invalidTypeError := errors.New("Provided value type didn't match obj field type")
						return invalidTypeError
					}

					structFieldValue.Set(val)
					return nil
				}
			func handler(w http.ResponseWriter, r *http.Request, context string) {
				  // fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
				  p,err := loadPage(r.URL.Path , context,r,w)
				  if err != nil {
				  	fmt.Println(err)
				        http.Redirect(w,r,"/your-404-page",307)
				        return
				  }

				   w.Header().Set("Cache-Control",  "public")
				  if !p.isResource {
				  		w.Header().Set("Content-Type",  "text/html")
				  		 defer func() {
					        if n := recover(); n != nil {
					           	color.Red("Error loading template in path : web" + r.URL.Path + ".tmpl reason :" )
					           	 fmt.Println(n)
					           	 http.Redirect(w,r,"/your-500-page",307)
					        }
					    }()
				      renderTemplate(w, r,  "web" + r.URL.Path, p)
				     
				     // fmt.Println(w)
				  } else {
				  		if strings.Contains(r.URL.Path, ".css") {
				  	  		w.Header().Add("Content-Type",  "text/css")
				  	  	} else if strings.Contains(r.URL.Path, ".js") {
				  	  		w.Header().Add("Content-Type",  "application/javascript")
				  	  	} else {
				  	  	w.Header().Add("Content-Type",  http.DetectContentType(p.Body))
				  	  	}
				  	 
				  	 
				      w.Write(p.Body)
				  }
				}

				func loadPage(title string, servlet string,r *http.Request,w http.ResponseWriter) (*Page,error) {
				    filename :=  "web" + title + ".tmpl"
				    if title == "/" {
				      http.Redirect(w,r,"/index",302)
				    }
				    body, err := Asset(filename)
				    if err != nil {
				      filename = "web" + title + ".html"
				      if title == "/" {
				    	filename = "web/index.html"
				    	}
				      body, err = Asset(filename)
				      if err != nil {
				         filename = "web" + title
				         body, err = Asset(filename)
				         if err != nil {
				            return nil, err
				         } else {
				          if strings.Contains(title, ".tmpl") || title == "/" {
				              return nil,nil
				          }
				          return &Page{Title: title, Body: body,isResource: true,request: nil}, nil
				         }
				      } else {
				         return &Page{Title: title, Body: body,isResource: true,request: nil}, nil
				      }
				    } 
				    //load custom struts
				    return &Page{Title: title, Body: body,isResource:false,request:r}, nil
				}
				func BytesToString(b []byte) string {
				    bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
				    sh := reflect.StringHeader{bh.Data, bh.Len}
				    return *(*string)(unsafe.Pointer(&sh))
				}
				func equalz(args ...interface{}) bool {
		    	    if args[0] == args[1] {
		        	return true;
				    }
				    return false;
				 }
				 func nequalz(args ...interface{}) bool {
				    if args[0] != args[1] {
				        return true;
				    }
				    return false;
				 }

				 func netlt(x,v float64) bool {
				    if x < v {
				        return true;
				    }
				    return false;
				 }
				 func netgt(x,v float64) bool {
				    if x > v {
				        return true;
				    }
				    return false;
				 }
				 func netlte(x,v float64) bool {
				    if x <= v {
				        return true;
				    }
				    return false;
				 }
				 func netgte(x,v float64) bool {
				    if x >= v {
				        return true;
				    }
				    return false;
				 }
				 type Page struct {
					    Title string
					    Body  []byte
					    request *http.Request
					    isResource bool
					    R *http.Request
					    Session *sessions.Session
					}
			type UserSpace struct {
				/* Property Type */
		
			}

			func  net_castUserSpace(args ...interface{}) *UserSpace  {
				
				s := UserSpace{}
				mapp := args[0].(db.O)
				if _, ok := mapp["_id"]; ok {
					mapp["Id"] = mapp["_id"]
				}
				data,_ := json.Marshal(&mapp)
				
				err := json.Unmarshal(data, &s) 
				if err != nil {
					fmt.Println(err.Error())
				}
				
				return &s
			}
			func net_structUserSpace() *UserSpace{ return &UserSpace{} }
			type UserSpaceInterface UserSpace
				func  net_UserSpaceInterface(args ...interface{}) (d UserSpace){
					if len(args) > 0 {
					jso := args[0].(string)
					var jsonBlob = []byte(jso)
					err := json.Unmarshal(jsonBlob, &d)
					if err != nil {
						fmt.Println("error:", err)
						return
					}
					return
					} else {
						d = UserSpace{} 
						return
					}
				}
			func dummy_timer(){
				dg := time.Second *5
				fmt.Println(dg)
			}

			func main() {
				
					 
					 fmt.Printf("Listenning on Port %v\n", "8080")
					 http.HandleFunc( "/",  makeHandler(handler))
					 http.Handle("/dist/",  http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: "web"}))
					 http.ListenAndServe(":8080", nil)
					 }