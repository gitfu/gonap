# goprofitbricks

### ```Install``` 

install go 
* OpenBSD  
``` pkg_add go ```

*  Arch linux 
    ```pacman -S go```
    
* Debian 
     ``` apt-get install golang-go ```
* etc ....
      https://golang.org/doc/install

#### ```Set your Environment```

```
mkdir -p ~/go/bin
export GOPATH=~/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```
### ``` Fetch goprofitbricks```
```go
go get "github.com/gitfu/goprofitbricks"
```
####  ```Test stuff```
```go
cd $GOPATH/src/github.com/gitfu/goprofitbricks
```
```go
vi config_test.go
```
* Set Username, Password, and Endpoint for testing

```go
go test -v 
```
* runs all the tests and reports pass/fail 

### Use
```ada
cd ~/
```

```ada
vi testrest.go
```

```go
package  main
import "github.com/gitfu/goprofitbricks"
import "fmt"	


func main() {
	// auth however you like,  I like Jasmin's technique myself. 
	goprofitbricks.SetAuth("your_username","your_password")
	/**
	 List Datacenter returns a collection of (Datacenter) Instances
	See file resp.go
	
	**/
	
	resp:=goprofitbricks.ListDatacenters()

	//get the Id of the first Item in the collection

	dc := goprofitbricks.GetDatacenter(resp.Items[0].Id)

	// Instances have methods like ShowProps()	
	fmt.Println("Before...")
	dc.ShowProps()
	
	// Instances have a  SetProp method.
	dc.SetProp("name","fu")

	fmt.Println("After.....")
	dc.ShowProps()
	// Instances have a  ShowEnts method.

	dc.ShowEnts()
	
	// Calling Save on an Instance patches the properties to the current values. 
	
	dc.Save()
	}
	
```
###### ```build``` 
```go 
	    go build testrest.go
	    go install testrest.go
```
###### ```Run```
```go 
	  testrest
```


### ```PACKAGE DOCUMENTATION```

```go
package goprofitbricks
    import "github.com/gitfu/goprofitbricks"
```

#### ```Variables```
```go
var Depth string
```
Depth controls the amount of data returned from the rest server ( range 1-5 )
```go
var Endpoint string
```
* Endpoint is the base url for REST requests .
```go 
var Passwd string
```
* Password for authentication .
```go
var Username string
```
* Username for authentication .

#### Functions
```go
func MkJson(i interface{}) string
```
*  Turn just about anything into Json

```go
	func SetAuth(u, p string)
```
```go
	func SetDepth(newdepth string) string
```

```go
	func SetEndpoint(newendpoint string) string
```

*  	SetEndpoint is used to set the REST Endpoint. 
*  	Endpoint is declared in config.go

#### Types

```go
	type StringCollectionMap map[string]Collection
```
```go
	type StringIfaceMap map[string]interface{}
```
```go
	type StringMap map[string]string
```
#### Resp struct
```go
	type Resp struct {
    		Req        *http.Request
    		StatusCode int
    		Headers    http.Header
    		Body       []byte
	}
```
* 		Resp is the struct returned by all Rest request functions

##### Resp methods
```go

			func (r *Resp) PrintHeaders()
```
* 				PrintHeaders prints the http headers as k,v pairs

```go 

	type Id_Type_Href struct {
    		Id   string `json:"id"`
    		Type string `json:"type"`
    		Href string `json:"href"`
	}
```
*  	The Id_Type_Href struct is embedded in Instance structs and Collection structs

#### Instance struct
```go

	type Instance struct {
    		Id_Type_Href
    		MetaData   StringMap           `json:"metaData"`
    		Properties StringIfaceMap      `json:"properties"`
    		Entities   StringCollectionMap `json:"entities"`
    		Resp       Resp                `json:"-"`
	}
```
* 		Get, Create, Update, and Patch functions all return an Instance struct.
*		A Resp struct is embedded in the Instance struct,
*		the raw server response is available as Instance.Resp.Body
		
##### Instance methods
```go
		func (ins *Instance) Save()
```
* 			Patch'es current Instance properties to the rest server 
			
```go
		func (ins *Instance) SetProp(key, val string)
```
* 			Sets an Instance property
```go
		func (ins *Instance) ShowEnts()
```
* 			ShowEnts prints the Instance Entities as k,v pairs
```go
		func (ins *Instance) ShowProps()
```   
* 			ShowProps prints the properties as k,v pairs

##### Get functions that return an instance 

```go
	type Collection struct {
    		Id_Type_Href
    		Items []Instance `json:"items,omitempty"`
    		Resp  Resp       `json:"-"`
	}

```
* 	Collection Structs contain Instance arrays. 
* 	List functions return Collections

### ```Functions by target```

#### ```Datacenter```
```go
 func ListDatacenters() Collection  
```
```go
 func CreateDatacenter(jason []byte) Instance  
```
```go
 func GetDatacenter(dcid string) Instance  
```
```go
 func UpdateDatacenter(dcid string, jason []byte) Instance  
```
```go
 func PatchDatacenter(dcid string, jason []byte) Instance  
```
```go
 func DeleteDatacenter(dcid string) Resp  
```

####  ```Server```
```go
 func ListServers(dcid string) Collection  
```
```go
 func CreateServer(dcid string, jason []byte) Instance  
```
```go
 func GetServer(dcid, srvid string) Instance  
```
```go
 func UpdateServer(dcid string, srvid string, jason []byte) Instance  
```
```go
 func PatchServer(dcid string, srvid string, jason []byte) Instance  
```
```go
 func DeleteServer(dcid, srvid string) Resp  
```
##### ``` Server Attached Cdroms```
```go
 func ListAttachedCdroms(dcid, srvid string) Collection  
```
```go
 func AttachCdrom(dcid string, srvid string, cdid string) Instance  
```
```go
 func GetAttachedCdrom(dcid, srvid, cdid string) Instance  
```
```go
 func DetachCdrom(dcid, srvid, cdid string) Resp  
```
##### ```Server Attached Volumes```
```go
 func ListAttachedVolumes(dcid, srvid string) Collection  
```
```go
 func AttachVolume(dcid string, srvid string, volid string) Instance  
```
```go
 func GetAttachedVolume(dcid, srvid, volid string) Instance  
```
```go
 func DetachVolume(dcid, srvid, volid string) Resp  
```
##### ``` Server Commands ``
```go
 func server_command(dcid, srvid, cmd string) Resp  
```
```go
 func StartServer(dcid, srvid string) Resp  
```
```go
 func StopServer(dcid, srvid string) Resp  
```
```go
 func RebootServer(dcid, srvid string) Resp  
```
#### ```Nics```

```go
 func ListNics(dcid, srvid string) Collection  
```
```go
 func CreateNic(dcid string, srvid string, jason []byte) Instance  
```
```go
 func GetNic(dcid, srvid, nicid string) Instance  
```
```go
 func UpdateNic(dcid string, srvid string, nicid string, jason []byte) Instance  
```
```go
 func PatchNic(dcid string, srvid string, nicid string, jason []byte) Instance  
```
```go
 func DeleteNic(dcid, srvid, nicid string) Resp  
```
