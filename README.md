PACKAGE DOCUMENTATION

package gonap
    import "github.com/gitfu/gonap"


CONSTANTS

const CommandHeader = "application/x-www-form-urlencoded"
    CommandHeader is used with is_command

const FullHeader = "application/vnd.profitbricks.resource+json"
    FullHeader is the standard header to include with all http requests
    except is_patch and is_command

const PatchHeader = "application/vnd.profitbricks.partial-properties+json"
    PatchHeader is used with is_patch .

VARIABLES

var Depth = "5"
    Depth sets the level of detail returned from the REST server .

var Endpoint = "https://private-anon-4354b0b6a-profitbricksrestapi.apiary-mock.com"
    Endpoint is the base url for REST requests .

var Passwd = "my_password"
    Password for authentication .

var Username = "my_username"
    Username for authentication .

FUNCTIONS

func MkJson(i interface{}) string

func SetAuth(u, p string)

func SetDepth(newdepth string) string
    SetDepth is used to set Depth

func SetEndpoint(newendpoint string) string
    SetEnpoint is used to set the REST Endpoint. Endpoint is declared in
    config.go

TYPES

type Collection struct {
    Id_Type_Href
    Items []Instance `json:"items,omitempty"`
    Resp  PBResp     `json:"-"`
}

type Datacenter Instance
    Datacenter is struct to hold data for a datacenter

func CreateDatacenter(jason []byte) Datacenter
    CreateDatacenter creates a datacenter and returns a Datacenter struct

func DeleteDatacenter(dcid string) Datacenter
    Deletes a Datacenter where id==dcid

func GetDatacenter(dcid string) Datacenter
    GetDatacenter returns a Datacenter struct where id == dcid

func PatchDatacenter(dcid string, jason []byte) Datacenter
    PatchDatacenter replaces any Datacenter properties with the values in
    jason returns an Datacenter struct where id ==dcid

func UpdateDatacenter(dcid string, jason []byte) Datacenter
    UpdateDatacenter updates all Datacenter properties from values in jason
    returns an Datacenter struct where id ==dcid

type Datacenters Collection
    Datacenters is a struct for Datacenter collections

func ListDatacenters() Datacenters
    ListDatacenters returns a Datacenter collection struct

type FwRule Instance
    FwRule is struct for Firewall rule instance data

func CreateFwRule(dcid string, srvid string, nicid string, jason []byte) FwRule
    CreateFwRule uses the json values in jason to create a new firewall rule
    and returns a FwRule struct

func DeleteFWRule(dcid, srvid, nicid, fwruleid string) FwRule
    DeleteFwRule removes firewall rule

func GetFwRule(dcid, srvid, nicid, fwruleid string) FwRule
    GetFwRule Retrieve a firewall rule and returns fwRule struct

func PatchFWRule(dcid string, srvid string, nicid string, fwruleid string, jason []byte) FwRule
    PatchFwRule Partially updates a firewall rule with data from []byte
    jason returns FwRule struct

func UpdateFwRule(dcid string, srvid string, nicid string, fwruleid string, jason []byte) FwRule
    UpdateFwRule Replaces all the properties of firewall rule, returns a
    FwRule struct

type FwRules Collection
    FwRules is the struct for firewall rule collections

func ListFwRules(dcid, srvid, nicid string) FwRules
    ListFwRules returns a collection of firewall rules

type Id_Type_Href struct {
    Id   string `json:"id"`
    Type string `json:"type"`
    Href string `json:"href"`
}

type Image Instance

func AttachCdrom(dcid string, srvid string, cdid string) Image

func DeleteImage(imageid string) Image
    Deletes an image where id==imageid

func DetachCdrom(dcid, srvid, cdid string) Image

func GetAttachedCdrom(dcid, srvid, cdid string) Image

func GetImage(imageid string) Image
    GetImage returns an Image struct where id ==imageid

func PatchImage(imageid string, jason []byte) Image
    PatchImage replaces any image properties from values in jason returns an
    Image struct where id ==imageid

func UpdateImage(imageid string, jason []byte) Image
    UpdateImage updates all image properties from values in jason returns an
    Image struct where id ==imageid

type Images Collection
    Images is a struct for Image and cdrom collections

func ListAttachedCdroms(dcid, srvid string) Images

func ListImages() Images
    ListImages returns an Images struct

type Instance struct {
    Id_Type_Href
    MetaData   StringMap      `json:"metaData"`
    Properties StringIfaceMap `json:"properties"`
    Entities   StringIfaceMap `json:"entities"`
    Resp       PBResp         `json:"-"`
}

func (ins *Instance) Save()
    Save converts the datacenter struct's properties to json and "patch"es
    them to the rest server

type Ipblock Instance
    Ipblock is the struct for Ipblock data

func GetIpBlock(ipblockid string) Ipblock

func ReleaseIpBlock(ipblockid string) Ipblock

func ReserveIpBlock(jason []byte) Ipblock

type Ipblocks Collection
    Ipblocks is the struct for an Ipblock collection

func ListIpBlocks() Ipblocks
    ListIpBlocks

type Lan Instance

func CreateLan(dcid string, jason []byte) Lan
    CreateLan creates a lan in the datacenter from a jason []byte and
    returns a Lan struct

func DeleteLan(dcid, lanid string) Lan
    DeleteLan deletes a lan where id == lanid

func GetLan(dcid, lanid string) Lan
    GetLan pulls data for the lan where id = lanid returns a Lan struct

func PatchLan(dcid string, lanid string, jason []byte) Lan
    PatchLan does a partial update to a lan using json from []byte jason
    returns a Lan struct

func UpdateLan(dcid string, lanid string, jason []byte) Lan
    UpdateLan does a complete update to a lan using json from []byte jason
    returns a Lan struct

type Lans Collection

func ListLans(dcid string) Lans
    ListLan returns a Lans struct collection for lans in the Datacenter

type Loadbalancer Instance

func CreateLoadbalancer(dcid string, jason []byte) Loadbalancer
    Createloadbalancer creates a loadbalancer in the datacenter from a jason
    []byte and returns a Loadbalancer struct

func DeleteLoadbalancer(dcid, lbalid string) Loadbalancer

func GetLoadbalancer(dcid, lbalid string) Loadbalancer
    GetLoadbalancer pulls data for the Loadbalancer where id = lbalid
    returns a Loadbalancer struct

func PatchLoadbalancer(dcid string, lbalid string, jason []byte) Loadbalancer

func UpdateLoadbalancer(dcid string, lbalid string, jason []byte) Loadbalancer

type Loadbalancers Collection

func ListLoadbalancers(dcid string) Loadbalancers
    Listloadbalancers returns a Loadbalancers struct for loadbalancers in
    the Datacenter

type Location Instance
    Location is the struct for a Location

func GetLocation(locid string) Location
    GetLocation returns a PBResp with data for a location in the Body

type Locations Collection
    Locations is the struct for a Location Collection

func ListLocations() Locations
    ListLocations returns a PBResp with location collection data in the Body

type MetaData StringIfaceMap

type Nic Instance

func CreateNic(dcid string, srvid string, jason []byte) Nic
    CreateNic creates a nic on a server from a jason []byte and returns a
    Nic struct

func DeleteBalancedNic(dcid, lbalid, balnicid string) Nic

func DeleteNic(dcid, srvid, nicid string) Nic
    DeleteNic deletes the nic where id=nicid and returns a PBResp struct

func GetBalancedNic(dcid, lbalid, balnicid string) Nic

func GetNic(dcid, srvid, nicid string) Nic
    GetNic pulls data for the nic where id = srvid returns a Nic struct

func PatchNic(dcid string, srvid string, nicid string, jason []byte) Nic
    PatchNic partial update of nic properties passed in as jason []byte
    Returns Nic struct

func UpdateNic(dcid string, srvid string, nicid string, jason []byte) Nic
    UpdateNic is a full update of nic properties passed in as jason []byte
    Returns Nic struct

type Nics Collection

func AssociateNics(dcid string, lbalid string, jason []byte) Nics

func ListBalancedNics(dcid, lbalid string) Nics

func ListLanMembers(dcid, lanid string) Nics
    ListLanMembers returns a Nic struct collection for the Lan

func ListNics(dcid, srvid string) Nics
    ListNics returns a Nics struct collection

type PBResp struct {
    Req        *http.Request
    StatusCode int
    Headers    http.Header
    Body       []byte
}
    PBResp is the struct returned by all Rest request functions

func (r *PBResp) PrintHeaders()
    PrintHeaders prints the http headers as k,v pairs

type RestRequest Instance

func GetRequest(requestid string) RestRequest

func StatusRequest(requestid string) RestRequest

type RestRequest_Properties struct {
    Method  string            `json:"method"`
    Headers map[string]string `json:"headers"`
    Body    string            `json:"body"`
    Url     string            `json:"url"`
}

type RestRequests Collection

func ListRequests() RestRequests

type Server Instance
    Server is a struct for Server data

func CreateServer(dcid string, jason []byte) Server
    CreateServer creates a server from a jason []byte and returns a Server
    struct

func DeleteServer(dcid, srvid string) Server
    DeleteServer deletes the server where id=srvid and returns Server struct

func GetServer(dcid, srvid string) Server
    GetServer pulls data for the server where id = srvid returns a Server
    struct

func PatchServer(dcid string, srvid string, jason []byte) Server
    PatchServer partial update of server properties passed in as jason
    []byte Returns Server struct

func RebootServer(dcid, srvid string) Server
    RebootServer reboots a server

func StartServer(dcid, srvid string) Server
    StartServer starts a server

func StopServer(dcid, srvid string) Server
    StopServer stops a server

func UpdateServer(dcid string, srvid string, jason []byte) Server
    UpdateServer is a full update of server properties passed in as jason
    []byte Returns Server struct

type Servers Collection
    Servers is a struct for Server struct collections

func ListServers(dcid string) Servers
    ListServers returns a server struct collection

type Snapshot Instance
    Snapshot struct for Snapshot data

func DeleteSnapshot(snapid string) Snapshot
    Deletes a Snapshot with id == snapid

func GetSnapshot(snapid string) Snapshot
    GetSnapshot retrieves Snapshot data where id==snapid returns a` snapshot
    struct

func PatchSnapshot(snapid string, jason []byte) Snapshot
    PatchSnapshot replaces any snapshot properties from values in jason
    returns an Snapshot struct where id ==snapid

func UpdateSnapshot(snapid string, jason []byte) Snapshot
    UpdateSnapshot replaces all snapshot properties from values in jason
    returns an Snapshot struct where id ==snapid

type Snapshots Collection
    Snapshots struct for a Snapshot collection

func ListSnapshots() Snapshots
    ListSnapshots retrieves a collection of snapshot data returns a
    Snapshots struct

type StringIfaceMap map[string]interface{}

type StringMap map[string]string
    MetaData is a map for metadata returned in a PBResp.Body

type Volume Instance

func AttachVolume(dcid string, srvid string, volid string) Volume

func DeleteVolume(dcid, volid string) Volume

func DetachVolume(dcid, srvid, volid string) Volume

func GetAttachedVolume(dcid, srvid, volid string) Volume

func PatchVolume(dcid string, volid string, jason []byte) Volume

func UpdateVolume(dcid string, volid string, jason []byte) Volume

func (vol *Volume) Save()

type Volumes Collection

func ListAttachedVolumes(dcid, srvid string) Volumes

func ListVolumes(dcid string) Volumes
    ListVolumes returns a Volumes struct for volumes in the Datacenter

SUBDIRECTORIES

	p
	tests

