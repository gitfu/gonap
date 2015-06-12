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
```
cd $GOPATH/src/github.com/gitfu/goprofitbricks
```
```
vi config_test.go
```
* Set Username, Password, and Endpoint for testing

```go
go test -v 
```
* runs all the tests and reports pass/fail 

### Use
```
cd ~/
```

```
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

#### CONSTANTS
```go
const CommandHeader = "application/x-www-form-urlencoded"
```
    CommandHeader is used with is_command
```go
const FullHeader = "application/vnd.profitbricks.resource+json"
```
    FullHeader is the standard header to include with all http requests
    except is_patch and is_command
```go
const PatchHeader = "application/vnd.profitbricks.partial-properties+json"
```
    PatchHeader is used with is_patch .

	
	
