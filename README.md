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
```
go get "github.com/gitfu/goprofitbricks"
```
###  Test stuff
```
cd $GOPATH/src/github.com/gitfu/goprofitbricks
```
```
vi config_test.
(Set Username, Password, and Endpoint for testing)
```
```
go test -v 
```


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
	
func main() {
	/**
	 List Datacenter returns a collection of (Datacenter) Instances
	See file resp.go
	
	**/
	
	resp:=goprofitbricks.ListDatacenters()

	//get the Id of the first Item in the collection

	dc := goprofitbricks.GetDatacenter(resp.Items[0].Id)

	// Instances have methods like ShowProps()	

	dc.ShowProps()
	
	// Instances have a  SetProp method.

	dc.SetProp("name","fu")
	dc.ShowProps()
	// Instances have a  ShowEnts method.

	dc.ShowEnts()
	
	// Calling Save on an Instance patches the properties to the current values. 
	
	dc.Save()
	
	```
	
	build 
	
	```go 
	    go build testrest.go
	    go install testrest.go
	```
	
	```go 
	  testrest
	  ```
```go	  
version  :  4
location  :  de/fkb
name  :  datacenter1
description  :  Description of my DC
name  :  fu
description  :  Description of my DC
version  :  4
location  :  de/fkb
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
lans  :  {
    "id": "700e1cab-99b2-4c30-ba8c-1d273ddba022/lans",
    "type": "collection",
    "href": "https://api.profitbricks.com/rest/datacenters/700e1cab-99b2-4c30-ba8c-1d273ddba022/lans"
}
save status code is  202
```
  
	
	
