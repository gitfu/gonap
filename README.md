# goprofitbricks

### Install 

install go 
* OpenBSD  
``` pkg_add go ```

*  Arch linux 
    ```pacman -S go```
    
* Debian 
     ``` apt-get install golang-go ```
* etc ....
      https://golang.org/doc/install

### Set your Environment

```
mkdir -p ~/go/bin
export GOPATH=~/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```
### Get goprofitbricks
```go
go get "github.com/gitfu/goprofitbricks"
```
###  Test stuff
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
#### build 
```go 
	    go build testrest.go
	    go install testrest.go
```
#### Run
```go 
	  testrest
```
##### (Output)


```go	  
Before...
description  :  Description of my DC
version  :  4
location  :  de/fkb
name  :  datacenter1
After.....
version  :  4
location  :  de/fkb
name  :  fu
description  :  Description of my DC
lans  :  {
    "id": "700e1cab-99b2-4c30-ba8c-1d273ddba022/lans",
    "type": "collection",
    "href": "https://api.profitbricks.com/rest/datacenters/700e1cab-99b2-4c30-ba8c-1d273ddba022/lans"
}
servers  :  {
    "id": "700e1cab-99b2-4c30-ba8c-1d273ddba022/servers",
    "type": "collection",
    "href": "https://api.profitbricks.com/rest/datacenters/700e1cab-99b2-4c30-ba8c-1d273ddba022/servers"
}
volumes  :  {
    "id": "700e1cab-99b2-4c30-ba8c-1d273ddba022/volumes",
    "type": "collection",
    "href": "https://api.profitbricks.com/rest/datacenters/700e1cab-99b2-4c30-ba8c-1d273ddba022/volumes"
}
loadbalancers  :  {
    "id": "700e1cab-99b2-4c30-ba8c-1d273ddba022/loadbalancers",
    "type": "collection",
    "href": "https://api.profitbricks.com/rest/datacenters/700e1cab-99b2-4c30-ba8c-1d273ddba022/loadbalancers"
}
save status code is  202

```
  


#### PACKAGE DOCUMENTATION
```go
package goprofitbricks
    import "github.com/gitfu/goprofitbricks"
```

#### Constants
```go
const CommandHeader = "application/x-www-form-urlencoded"
```
*    CommandHeader is used with is_command
```go
const FullHeader = "application/vnd.profitbricks.resource+json"
```
*    FullHeader is the standard header to include with all http requests
    except is_patch and is_command
```go
const PatchHeader = "application/vnd.profitbricks.partial-properties+json"
```
*    PatchHeader is used with is_patch .

#### Variables
```go
var Depth = "5"
```

```go
var Endpoint = "https://private-anon-4354b0b6a-profitbricksrestapi.apiary-mock.com"

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
```go
func SetAuth(u, p string)
```

```go
func SetDepth(newdepth string) string
```

*   SetDepth is used to set Depth

```go
func SetEndpoint(newendpoint string) string
```

*  SetEndpoint is used to set the REST Endpoint. 
  Endpoint is declared in config.go

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

```go

type Resp struct {
    Req        *http.Request
    StatusCode int
    Headers    http.Header
    Body       []byte
}
```
* Resp is the struct returned by all Rest request functions

```go

func (r *Resp) PrintHeaders()
```
* PrintHeaders prints the http headers as k,v pairs



```go 

type Id_Type_Href struct {
    Id   string `json:"id"`
    Type string `json:"type"`
    Href string `json:"href"`
}
```
*  The Id_Type_Href struct is embedded in Instance structs and Collection structs

```go

type Instance struct {
    Id_Type_Href
    MetaData   StringMap           `json:"metaData"`
    Properties StringIfaceMap      `json:"properties"`
    Entities   StringCollectionMap `json:"entities"`
    Resp       Resp                `json:"-"`
}

```
* Get, Create, Update, and Patch functions all return an Instance struct.

```go
func (ins *Instance) Save()
```
* Save converts the Instance struct's properties to json and "patch"es
    them to the rest server
```go
func (ins *Instance) SetProp(key, val string)
```
* Sets an Instance property
```go
func (ins *Instance) ShowEnts()
```
* ShowEnts prints the Instance Entities as k,v pairs
```go
func (ins *Instance) ShowProps()
```   
* ShowProps prints the properties as k,v pairs


```go

type Collection struct {
    Id_Type_Href
    Items []Instance `json:"items,omitempty"`
    Resp  Resp       `json:"-"`
}

```
* Collection Structs contain Instance arrays. 
* List functions return Collections

##### List functions that return Collection structs
```go
	ListAttachedCdroms(dcid, srvid string) 
	ListAttachedVolumes(dcid, srvid string) 
	ListBalancedNics(dcid, lbalid string)
	ListDatacenters()
	ListFwRules(dcid, srvid, nicid string)
	ListImages()
	ListIpBlocks()
	ListLanMembers(dcid, lanid string)
	func ListLans(dcid string)
	ListLoadbalancers(dcid string)
	ListLocations()
	ListNics(dcid, srvid string)
	ListRequests()
	ListServers(dcid string)
	ListSnapshots() 
	ListVolumes(dcid string) 
```
