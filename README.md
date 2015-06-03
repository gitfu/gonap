# ```go nap ```
#### ```A rethink of my REST-ful ways in Go.``` 


## PACKAGE DOCUMENTATION

#### package gonap
```
    import "github.com/gitfu/gonap"
```

#### VARIABLES
```
var Depth = "5"
    Depth sets the level of detail returned from the REST server .

var Endpoint = "https://private-anon-4354b0b6a-profitbricksrestapi.apiary-mock.com"
    Endpoint is the base url for REST requests .

var Passwd = "my_password"
    Password for authentication .

var Username = "my_username"
    Username for authentication .
```
#### FUNCTIONS
```
func SetAuth(u, p string)
    SetAuth is used to set Username and Passwd. Username and Passwd are
    declared in config.go

func SetDepth(newdepth string) string
    SetDepth is used to set Depth

func SetEndpoint(newendpoint string) string
    SetEnpoint is used to set the REST Endpoint. Endpoint is declared in
    config.go
```

#### ```TYPES```

```go
type Datacenter struct {
    Id         string       `json:"id,omitempty"`
    Type       string       `json:"type,omitempty"`
    Href       string       `json:"href,omitempty"`
    MetaData   MetaData     `json:"metadata,omitempty"`
    Properties DcProperties `json:"properties,omitempty"`
    Entities   DcEntities   `json:"entities,omitempty"`
}
```


    Datacenter is struct to hold data for a datacenter
    
```go
func AsDatacenter(body []byte) Datacenter
```
```go
    AsDatacenter unmarshalls a []byte into a Datacenter struct
```
```go
func (dc *Datacenter) Getserver(srvid string) Server
```
```go
func (dc *Datacenter) Listlans() Lans
```
```go
func (dc *Datacenter) Listloadbalancers() Loadbalancers
```
```go
func (dc *Datacenter) Listservers() Servers
    * func (dc *Datacenter) Asjson() []byte {

    } *
```
```go
func (dc *Datacenter) Listvolumes() Volumes
```
```go
type Datacenters struct {
    Id    string       `json:"id,omitempty"`
    Type  string       `json:"type,omitempty"`
    Href  string       `json:"href,omitempty"`
    Items []Datacenter `json:"items,omitempty"`
}
    Datacenters is a struct for Datacenter collections
```
```go
func AsDatacenters(body []byte) Datacenters
```
```go
type DcEntities struct {
    Servers       Servers       `json:"servers,omitempty"`
    Loadbalancers Loadbalancers `json:"loadbalancers,omitempty"`
    Lans          Lans          `json:"lans,omitempty"`
    Volumes       Volumes       `json:"volumes,omitempty"`
}
```
```go
type DcProperties struct {
    Name        string `json:"name"`
    Location    string `json:"location"`
    Description string `json:"description,omitempty"`
}
```
```go
type FwRule struct {
    Id         string           `json:"id,omitempty"`
    Type       string           `json:"type,omitempty"`
    Href       string           `json:"href,omitempty"`
    MetaData   MetaData         `json:"metadata,omitempty"`
    Properties FwRuleProperties `json:"properties,omitempty"`
}
```
```go
func AsFwRule(body []byte) FwRule
```
```go
type FwRuleProperties struct {
    Name           string `json:"name,omitempty"`
    Protocol       string `json:"protocol,omitempty"`
    SourceMac      string `json:"sourceMac,omitempty"`
    SourceIp       string `json:"sourceIp,omitempty"`
    TargetIp       string `json:"targetIp,omitempty"`
    PortRangeStart string `json:"portRangeStart,omitempty"`
    PortRangeEnd   string `json:"portRangeEnd,omitempty"`
    IcmpType       string `json:"icmpType,omitempty"`
    IcmpCode       string `json:"icmpCode,omitempty"`
}
```
```go
type FwRules struct {
    Id    string   `json:"id,omitempty"`
    Type  string   `json:"type,omitempty"`
    Href  string   `json:"href,omitempty"`
    Items []FwRule `json:"items,omitempty"`
}
```
```go
func AsFwRules(body []byte) FwRules
```
```go
type Image struct {
    Id         string          `json:"id,omitempty"`
    Type       string          `json:"type,omitempty"`
    Href       string          `json:"href,omitempty"`
    MetaData   MetaData        `json:"metadata,omitempty"`
    Properties ImageProperties `json:"properties,omitempty"`
}
```
```go
func AsCdRom(body []byte) Image
```
```go
func AsImage(body []byte) Image
```
```go
type ImageProperties struct {
    Name                string `json:"name,omitempty"`
    Description         string `json:"description,omitempty"`
    Location            string `json:"location"`
    Size                int    `json:"size"`
    Public              bool   `json:"public,omitempty"`
    ImageType           string `json:"imageType,omitempty"`
    CpuHotPlug          bool   `json:"cpuHotPlug,omitempty"`
    CpuHotUnplug        bool   `json:"cpuHotUnplug,omitempty"`
    RamHotPlug          bool   `json:"ramHotPlug,omitempty"`
    RamHotUnplug        bool   `json:"ramHotUnplug,omitempty"`
    NicHotPlug          bool   `json:"nicHotPlug,omitempty"`
    NicHotUnplug        bool   `json:"nicHotUnplug,omitempty"`
    DiscVirtioHotPlug   bool   `json:"discVirtioHotPlug,omitempty"`
    DiscVirtioHotUnplug bool   `json:"discVirtioHotUnplug,omitempty"`
    DiscScsiHotPlug     bool   `json:"discScsiHotPlug,omitempty"`
    DiscScsiHotUnplug   bool   `json:"discScsiHotUnplug,omitempty"`
    LicenceType         string `json:"licenceType,omitempty"`
}
```
```go
type Images struct {
    Id    string  `json:"id,omitempty"`
    Type  string  `json:"type,omitempty"`
    Href  string  `json:"href,omitempty"`
    Items []Image `json:"items,omitempty"`
}
```
```go
func AsCdRoms(body []byte) Images
```
```go
func AsImages(body []byte) Images
```
```go
type Ipblock struct {
    Id         string            `json:"id,omitempty"`
    Type       string            `json:"type,omitempty"`
    Href       string            `json:"href,omitempty"`
    MetaData   MetaData          `json:"metadata,omitempty"`
    Properties IpblockProperties `json:"properties"`
}
```
```go
func AsIpblock(body []byte) Ipblock
```
```go
type IpblockProperties struct {
    Location string `json:"location"`
    Size     int    `json:"size"`
}
```
```go
type Ipblocks struct {
    Id    string    `json:"id,omitempty"`
    Type  string    `json:"type,omitempty"`
    Href  string    `json:"href,omitempty"`
    Items []Ipblock `json:"items,omitempty"`
}
```
```go
func AsIpblocks(body []byte) Ipblocks
```
```go
type Lan struct {
    Id         string        `json:"id,omitempty"`
    Type       string        `json:"type,omitempty"`
    Href       string        `json:"href,omitempty"`
    MetaData   MetaData      `json:"metadata,omitempty"`
    Properties LanProperties `json:"properties"`
    Entities   LanEntities   `json:"entities,omitempty"`
}
```
```go
func AsLan(body []byte) Lan
```
```go
type LanEntities struct {
    Nics Nics `json:"nics,omitempty"`
}
```
```go
type LanProperties struct {
    Name   string `json:"name,omitempty"`
    Public bool   `json:"public,omitempty"`
}
```
```go
type Lans struct {
    Id    string `json:"id,omitempty"`
    Type  string `json:"type,omitempty"`
    Href  string `json:"href,omitempty"`
    Items []Lan  `json:"items,omitempty"`
}
```
```go
func AsLans(body []byte) Lans
```
```go
type LbalEnts struct {
    BalancedNics Nics `json:"balancednics,omitempty"`
}
```
```go
type LbalProps struct {
    Name string `json:"name,omitempty"`
    Ip   string `json:"ip,omitempty"`
    Dhcp bool   `json:"dhcp,omitempty"`
}
```
```go
type Loadbalancer struct {
    Id         string    `json:"id,omitempty"`
    Type       string    `json:"type,omitempty"`
    Href       string    `json:"href,omitempty"`
    MetaData   MetaData  `json:"metadata,omitempty"`
    Properties LbalProps `json:"properties,omitempty"`
    Entities   LbalEnts  `json:"entities,omitempty"`
}
```
```go
func AsLoadbalancer(body []byte) Loadbalancer
```
```go
type Loadbalancers struct {
    Id    string         `json:"id,omitempty"`
    Type  string         `json:"type,omitempty"`
    Href  string         `json:"href,omitempty"`
    Items []Loadbalancer `json:"items,omitempty"`
}
```
```go
func AsLoadbalancers(body []byte) Loadbalancers
```
```go
type Location struct {
    Id         string             `json:"id,omitempty"`
    Type       string             `json:"type,omitempty"`
    Href       string             `json:"href,omitempty"`
    MetaData   MetaData           `json:"metadata,omitempty"`
    Properties LocationProperties `json:"properties"`
}
    Location is the struct for a Location
```
```go
func AsLocation(body []byte) Location
    AsLocation Unmarshals json into a Location struct
```
```go
type LocationProperties struct {
    Name string `json:"name"`
}
    LocationProperties is aa struct for Locaation Properties
```
```go
type Locations struct {
    Id    string     `json:"id,omitempty"`
    Type  string     `json:"type,omitempty"`
    Href  string     `json:"href,omitempty"`
    Items []Location `json:"items,omitempty"`
}
    Locations is the struct for a Location Collection
```
```go
func AsLocations(body []byte) Locations
    AsLocations Unmarshals json into a Locations struct
```
```go
type MetaData struct {
    LastModified   string `json:"lastModifiedDate,omitempty"`
    LastModifiedBy string `json:"lastModifiedBy,omitempty"`
    Created        string `json:"createdDate,omitempty"` //"2014-02-01T11:12:12Z",
    CreatedBy      string `json:"createdBy,omitempty"`
    State          string `json:"state,omitempty"`
    Etag           string `json:"etag,omitempty"`
}
```
```go
MetaData is a struct for metadata returned in a PBResp.Body
```
```go
type Nic struct {
    Id         string        `json:"id,omitempty"`
    Type       string        `json:"type,omitempty"`
    Href       string        `json:"href,omitempty"`
    MetaData   MetaData      `json:"metadata,omitempty"`
    Properties NicProperties `json:"properties,omitempty"`
}
```
```go
func AsNic(body []byte) Nic
```
```go
type NicProperties struct {
    Name           string   `json:"name,omitempty"`
    Ips            []string `json:"ips,omitempty"`
    Dhcp           bool     `json:"dhcp,omitempty"`
    Lan            int      `json:"lan"`
    FirewallActive bool     `json:"firewallActive,omitempty"`
    Firewallrules  []FwRule `json:"firewallrules,omitempty"`
}
```
```go
type Nics struct {
    Id    string `json:"id,omitempty"`
    Type  string `json:"type,omitempty"`
    Href  string `json:"href,omitempty"`
    Items []Nic  `json:"items,omitempty"`
}
```
```go
func AsNics(body []byte) Nics
```
```go
type PBResp struct {
    StatusCode int
    Headers    http.Header
    Body       []byte
}
```

##### ```   PBResp is the struct returned by all Rest request functions ```

```go
func AssociateNics(dcid string, lbalid string, jason []byte) PBResp
```
```go
func AttachCdrom(dcid string, srvid string, cdid string) PBResp

func AttachVolume(dcid string, srvid string, volid string) PBResp
```
```go
func CreateDatacenter(jason []byte) PBResp

func CreateFwRule(dcid string, srvid string, nicid string, jason []byte) PBResp

func CreateLan(dcid string, jason []byte) PBResp

func CreateLoadbalancer(dcid string, jason []byte) PBResp

func CreateNic(dcid string, srvid string, jason []byte) PBResp
    *

	CreateNic takes []byte jason to create a nic
	Returns a PBResp struct with Nic data in PBResp.Body .

    *

func CreateServer(dcid string, jason []byte) PBResp
```
```go
func DeleteBalancedNic(dcid, lbalid, balnicid string) PBResp

func DeleteDatacenter(dcid string) PBResp

func DeleteFWRule(dcid, srvid, nicid, fwruleid string) PBResp

func DeleteImage(imageid string) PBResp

func DeleteLan(dcid, lanid string) PBResp

func DeleteLoadbalancer(dcid, lbalid string) PBResp

func DeleteNic(dcid, srvid, nicid string) PBResp

func DeleteServer(dcid, srvid string) PBResp

func DeleteSnapshot(snapid string) PBResp

func DeleteVolume(dcid, volid string) PBResp

func DetachCdrom(dcid, srvid, cdid string) PBResp

func DetachVolume(dcid, srvid, volid string) PBResp
```
```go
func GetAttachedCdrom(dcid, srvid, cdid string) PBResp

func GetAttachedVolume(dcid, srvid, volid string) PBResp

func GetBalancedNic(dcid, lbalid, balnicid string) PBResp

func GetDatacenter(dcid string) PBResp

func GetFwRule(dcid, srvid, nicid, fwruleid string) PBResp
    * Firewall rule

	Retrieve a firewall rule
	Replace properties of firewall rule
	Partially update a firewall rule
	Remove firewall rule

    *
```
```go
func GetImage(imageid string) PBResp

func GetIpBlock(ipblockid string) PBResp

func GetLan(dcid, lanid string) PBResp

func GetLoadbalancer(dcid, lbalid string) PBResp

func GetLocation(locid string) PBResp
    GetLocations returns a PBResp with data for a location in the Body

func GetNic(dcid, srvid, nicid string) PBResp
    GetNic returns a PBResp with nic data in PBResp.Body

func GetRequest(requestid string) PBResp

func GetServer(dcid, srvid string) PBResp

func GetSnapshot(snapid string) PBResp

func GetVolume(dcid, volid string) PBResp
```
```go
func ListAttachedCdroms(dcid, srvid string) PBResp

func ListAttachedVolumes(dcid, srvid string) PBResp

func ListBalancedNics(dcid, lbalid string) PBResp

func ListDatacenters() PBResp

func ListFwRules(dcid, srvid, nicid string) PBResp

func ListImages() PBResp

func ListIpBlocks() PBResp

func ListLanMembers(dcid, lanid string) PBResp

func ListLans(dcid string) PBResp

func ListLoadbalancers(dcid string) PBResp

func ListLocations() PBResp
    ListLocations returns a PBResp with location collection data in the Body

func ListNics(dcid, srvid string) PBResp
    ListNics returns a PBResp with nic collection data in PBResp.Body

func ListRequests() PBResp

func ListServers(dcid string) PBResp
    *

    Server collection

	List servers
	Create a server

    *

func ListSnapshots() PBResp

func ListVolumes(dcid string) PBResp
```
```go
func PatchDatacenter(dcid string, jason []byte) PBResp

func PatchFWRule(dcid string, srvid string, nicid string, fwruleid string, jason []byte) PBResp

func PatchImage(imageid string, jason []byte) PBResp

func PatchLan(dcid string, lanid string, jason []byte) PBResp

func PatchLoadBalancer(dcid string, lbalid string, jason []byte) PBResp

func PatchNic(dcid string, srvid string, nicid string, jason []byte) PBResp

func PatchServer(dcid string, srvid string, jason []byte) PBResp

func PatchSnapshot(snapid string, jason []byte) PBResp

func PatchVolume(dcid string, volid string, jason []byte) PBResp

func RebootServer(dcid, srvid string) PBResp

func ReleaseIpBlock(ipblockid string) PBResp

func ReserveIpBlock(jason []byte) PBResp

func StartServer(dcid, srvid string) PBResp
    * Start server command

	Execute start server

    Stop server command

	Execute Stop Server

    Reboot server command

	Execute reboot server

    *

func StatusRequest(requestid string) PBResp

func StopServer(dcid, srvid string) PBResp

func UpdateDatacenter(dcid string, jason []byte) PBResp

func UpdateFwRule(dcid string, srvid string, nicid string, fwruleid string, jason []byte) PBResp

func UpdateImage(imageid string, jason []byte) PBResp

func UpdateLan(dcid string, lanid string, jason []byte) PBResp
```
```go
func UpdateLoadbalancer(dcid string, lbalid string, jason []byte) PBResp
```
```go
func UpdateNic(dcid string, srvid string, nicid string, jason []byte) PBResp
```
```go
func UpdateServer(dcid string, srvid string, jason []byte) PBResp
```
```go
func UpdateSnapshot(snapid string, jason []byte) PBResp
```
```go
func UpdateVolume(dcid string, volid string, jason []byte) PBResp
```
```go
func (r *PBResp) AsJson()
    AsJson returns PBResp.Body as the raw json
```
```go
func (r *PBResp) PrintHeaders()
    PrintHeaders prints the http headers as k,v pairs
```
```go
type Server struct {
    Id         string   `json:"id,omitempty"`
    Type       string   `json:"type,omitempty"`
    Href       string   `json:"href,omitempty"`
    MetaData   MetaData `json:"metadata,omitempty"`
    Properties SrvProps `json:"properties,omitempty"`
    Entities   SrvEnts  `json:"entities,omitempty"`
}
```
```go
func AsServer(body []byte) Server
```
```go
type Servers struct {
    Id    string   `json:"id,omitempty"`
    Type  string   `json:"type,omitempty"`
    Href  string   `json:"href,omitempty"`
    Items []Server `json:"items,omitempty"`
}
```
```go
func AsServers(body []byte) Servers
```
```go
type SnapProps struct {
    Name                string `json:"name,omitempty"`
    Description         string `json:"description,omitempty"`
    Location            string `json:"location,omitempty"`
    Size                int    `json:"size,omitempty"`
    CpuHotPlug          bool   `json:"cpuHotPlug,omitempty"`
    CpuHotUnplug        bool   `json:"cpuHotUnplug,omitempty"`
    RamHotPlug          bool   `json:"ramHotPlug,omitempty"`
    RamHotUnplug        bool   `json:"ramHotUnplug,omitempty"`
    NicHotPlug          bool   `json:"nicHotPlug,omitempty"`
    NicHotUnplug        bool   `json:"nicHotUnplug,omitempty"`
    DiscVirtioHotPlug   bool   `json:"discVirtioHotPlug,omitempty"`
    DiscVirtioHotUnplug bool   `json:"discVirtioHotUnplug,omitempty"`
    DiscScsiHotPlug     bool   `json:"discScsiHotPlug,omitempty"`
    DiscScsiHotUnplug   bool   `json:"discScsiHotUnplug,omitempty"`
    LicenceType         string `json:"licenceType,omitempty"`
}
```
```go
type Snapshot struct {
    Id         string    `json:"id,omitempty"`
    Type       string    `json:"type,omitempty"`
    Href       string    `json:"href,omitempty"`
    MetaData   MetaData  `json:"metadata,omitempty"`
    Properties SnapProps `json:"properties,omitempty"`
}
```
```go
func AsSnapshot(body []byte) Snapshot
```
```go
type Snapshots struct {
    Id    string     `json:"id,omitempty"`
    Type  string     `json:"type,omitempty"`
    Href  string     `json:"href,omitempty"`
    Items []Snapshot `json:"items,omitempty"`
}
```
```go
func AsSnapshots(body []byte) Snapshots
```
```go
type SrvEnts struct {
    Cdroms  Images  `json:"cdroms,omitempty"`
    Nics    Nics    `json:"nics,omitempty"`
    Volumes Volumes `json:"volumes,omitempty"`
}
```
```go
type SrvProps struct {
    Name             string `json:"name,omitempty"`
    Cores            int    `json:"cores,omitempty"`
    Ram              int    `json:"ram,omitempty"`
    Availabilityzone string `json:"availabilityzone,omitempty"`
    Licencetype      string `json:"licencetype,omitempty"`
    Bootvolume       string `json:"bootvolume,omitempty"`
    Bootcdrom        string `json:"bootcdrom,omitempty"`
}
```
```go
type VolProps struct {
    Name                string `json:"name,omitempty"`
    Size                int    `json:"size,omitempty"`
    Bus                 string `json:"bus,omitempty"`
    Image               string `json:"image,omitempty"`
    ImagePassword       string `json:"imagePassword,omitempty"`
    Type                string `json:"type,omitempty"`
    LicenceType         string `json:"licenceType,omitempty"`
    CpuHotPlug          bool   `json:"cpuHotPlug,omitempty"`
    CpuHotUnplug        bool   `json:"cpuHotUnplug,omitempty"`
    RamHotPlug          bool   `json:"ramHotPlug,omitempty"`
    RamHotUnplug        bool   `json:"ramHotUnplug,omitempty"`
    NicHotPlug          bool   `json:"nicHotPlug,omitempty"`
    NicHotUnplug        bool   `json:"nicHotUnplug,omitempty"`
    DiscVirtioHotPlug   bool   `json:"discVirtioHotPlug,omitempty"`
    DiscVirtioHotUnplug bool   `json:"discVirtioHotUnplug,omitempty"`
    DiscScsiHotPlug     bool   `json:"discScsiHotPlug,omitempty"`
    DiscScsiHotUnplug   bool   `json:"discScsiHotUnplug,omitempty"`
}
```
```go
type Volume struct {
    Id         string   `json:"id,omitempty"`
    Type       string   `json:"type,omitempty"`
    Href       string   `json:"href,omitempty"`
    MetaData   MetaData `json:"metadata,omitempty"`
    Properties VolProps `json:"properties,omitempty"`
}
```
```go
func AsVolume(body []byte) Volume
```
```go
type Volumes struct {
    Id    string   `json:"id,omitempty"`
    Type  string   `json:"type,omitempty"`
    Href  string   `json:"href,omitempty"`
    Items []Volume `json:"items,omitempty"`
}
```
```go
func AsVolumes(body []byte) Volumes

```
