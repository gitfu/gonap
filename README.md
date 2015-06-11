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

func AssociateNics(dcid string, lbalid string, jason []byte) Collection

func ListAttachedCdroms(dcid, srvid string) Collection

func ListAttachedVolumes(dcid, srvid string) Collection

func ListBalancedNics(dcid, lbalid string) Collection

func ListDatacenters() Collection
    ListDatacenters returns a Collection struct

func ListFwRules(dcid, srvid, nicid string) Collection
    ListFwRules returns a collection of firewall rules

func ListImages() Collection
    ListImages returns an Collection struct

func ListIpBlocks() Collection
    ListIpBlocks

func ListLanMembers(dcid, lanid string) Collection
    ListLanMembers returns a Nic struct collection for the Lan

func ListLans(dcid string) Collection
    ListLan returns a Collection for lans in the Datacenter

func ListLoadbalancers(dcid string) Collection
    Listloadbalancers returns a Collection struct for loadbalancers in the
    Datacenter

func ListLocations() Collection
    ListLocations returns location collection data

func ListNics(dcid, srvid string) Collection
    ListNics returns a Nics struct collection

func ListRequests() Collection

func ListServers(dcid string) Collection
    ListServers returns a server struct collection

func ListSnapshots() Collection
    ListSnapshots retrieves a collection of snapshot data returns a
    Collection struct

func ListVolumes(dcid string) Collection
    ListVolumes returns a Collection struct for volumes in the Datacenter

type Id_Type_Href struct {
    Id   string `json:"id"`
    Type string `json:"type"`
    Href string `json:"href"`
}

type Instance struct {
    Id_Type_Href
    MetaData   StringMap      `json:"metaData"`
    Properties StringIfaceMap `json:"properties"`
    Entities   StringIfaceMap `json:"entities"`
    Resp       PBResp         `json:"-"`
}

func AttachCdrom(dcid string, srvid string, cdid string) Instance

func AttachVolume(dcid string, srvid string, volid string) Instance

func CreateDatacenter(jason []byte) Instance
    CreateDatacenter creates a datacenter and returns a Instance struct

func CreateFwRule(dcid string, srvid string, nicid string, jason []byte) Instance
    CreateFwRule uses the json values in jason to create a new firewall rule
    and returns an Instance struct

func CreateLan(dcid string, jason []byte) Instance
    CreateLan creates a lan in the datacenter from a jason []byte and
    returns a Instance struct

func CreateLoadbalancer(dcid string, jason []byte) Instance
    Createloadbalancer creates a loadbalancer in the datacenter from a jason
    []byte and returns a Instance struct

func CreateNic(dcid string, srvid string, jason []byte) Instance
    CreateNic creates a nic on a server from a jason []byte and returns a
    Instance struct

func CreateServer(dcid string, jason []byte) Instance
    CreateServer creates a server from a jason []byte and returns a Instance
    struct

func GetAttachedCdrom(dcid, srvid, cdid string) Instance

func GetAttachedVolume(dcid, srvid, volid string) Instance

func GetBalancedNic(dcid, lbalid, balnicid string) Instance

func GetDatacenter(dcid string) Instance
    GetDatacenter returns a Instance struct where id == dcid

func GetFwRule(dcid, srvid, nicid, fwruleid string) Instance
    GetFwRule Retrieve a firewall rule and returns Instance struct

func GetImage(imageid string) Instance
    GetImage returns an Instance struct where id ==imageid

func GetIpBlock(ipblockid string) Instance

func GetLan(dcid, lanid string) Instance
    GetLan pulls data for the lan where id = lanid returns an Instance
    struct

func GetLoadbalancer(dcid, lbalid string) Instance
    GetLoadbalancer pulls data for the Loadbalancer where id = lbalid
    returns a Instance struct

func GetLocation(locid string) Instance
    GetLocation returns location data

func GetNic(dcid, srvid, nicid string) Instance
    GetNic pulls data for the nic where id = srvid returns a Instance struct

func GetRequest(requestid string) Instance

func GetServer(dcid, srvid string) Instance
    GetServer pulls data for the server where id = srvid returns a Instance
    struct

func GetSnapshot(snapid string) Instance
    GetSnapshot retrieves Instance data where id==snapid returns a` snapshot
    struct

func PatchDatacenter(dcid string, jason []byte) Instance
    PatchDatacenter replaces any Datacenter properties with the values in
    jason returns an Instance struct where id ==dcid

func PatchFWRule(dcid string, srvid string, nicid string, fwruleid string, jason []byte) Instance
    PatchFwRule Partially updates a firewall rule with data from []byte
    jason returns Instance struct

func PatchImage(imageid string, jason []byte) Instance
    PatchImage replaces any image properties from values in jason returns an
    Instance struct where id ==imageid

func PatchLan(dcid string, lanid string, jason []byte) Instance
    PatchLan does a partial update to a lan using json from []byte jason
    returns a Instance struct

func PatchLoadbalancer(dcid string, lbalid string, jason []byte) Instance

func PatchNic(dcid string, srvid string, nicid string, jason []byte) Instance
    PatchNic partial update of nic properties passed in as jason []byte
    Returns Instance struct

func PatchServer(dcid string, srvid string, jason []byte) Instance
    PatchServer partial update of server properties passed in as jason
    []byte Returns Instance struct

func PatchSnapshot(snapid string, jason []byte) Instance
    PatchSnapshot replaces any snapshot properties from values in jason
    returns an Instance struct where id ==snapid

func PatchVolume(dcid string, volid string, jason []byte) Instance

func ReserveIpBlock(jason []byte) Instance

func StatusRequest(requestid string) Instance

func UpdateDatacenter(dcid string, jason []byte) Instance
    UpdateDatacenter updates all Datacenter properties from values in jason
    returns an Instance struct where id ==dcid

func UpdateFwRule(dcid string, srvid string, nicid string, fwruleid string, jason []byte) Instance
    UpdateFwRule Replaces all the properties of firewall rule, returns a
    Instance struct

func UpdateImage(imageid string, jason []byte) Instance
    UpdateImage updates all image properties from values in jason returns an
    Instance struct where id ==imageid

func UpdateLan(dcid string, lanid string, jason []byte) Instance
    UpdateLan does a complete update to a lan using json from []byte jason
    returns a Instance struct

func UpdateLoadbalancer(dcid string, lbalid string, jason []byte) Instance

func UpdateNic(dcid string, srvid string, nicid string, jason []byte) Instance
    UpdateNic is a full update of nic properties passed in as jason []byte
    Returns Instance struct

func UpdateServer(dcid string, srvid string, jason []byte) Instance
    UpdateServer is a full update of server properties passed in as jason
    []byte Returns Instance struct

func UpdateSnapshot(snapid string, jason []byte) Instance
    UpdateSnapshot replaces all snapshot properties from values in jason
    returns an Instance struct where id ==snapid

func UpdateVolume(dcid string, volid string, jason []byte) Instance

func (ins *Instance) Save()
    Save converts the Instance struct's properties to json and "patch"es
    them to the rest server

type Loadbalancer Instance

type Loadbalancers Collection

type MetaData StringIfaceMap

type Nic Instance

type Nics Collection

type PBResp struct {
    Req        *http.Request
    StatusCode int
    Headers    http.Header
    Body       []byte
}
    PBResp is the struct returned by all Rest request functions

func DeleteBalancedNic(dcid, lbalid, balnicid string) PBResp

func DeleteDatacenter(dcid string) PBResp
    Deletes a Datacenter where id==dcid

func DeleteFWRule(dcid, srvid, nicid, fwruleid string) PBResp
    DeleteFwRule removes firewall rule

func DeleteImage(imageid string) PBResp
    Deletes an image where id==imageid

func DeleteLan(dcid, lanid string) PBResp
    DeleteLan deletes a lan where id == lanid

func DeleteLoadbalancer(dcid, lbalid string) PBResp

func DeleteNic(dcid, srvid, nicid string) PBResp
    DeleteNic deletes the nic where id=nicid and returns a PBResp struct

func DeleteServer(dcid, srvid string) PBResp
    DeleteServer deletes the server where id=srvid and returns PBResp struct

func DeleteSnapshot(snapid string) PBResp
    Deletes a Snapshot with id == snapid

func DeleteVolume(dcid, volid string) PBResp

func DetachCdrom(dcid, srvid, cdid string) PBResp

func DetachVolume(dcid, srvid, volid string) PBResp

func RebootServer(dcid, srvid string) PBResp
    RebootServer reboots a server

func ReleaseIpBlock(ipblockid string) PBResp

func StartServer(dcid, srvid string) PBResp
    StartServer starts a server

func StopServer(dcid, srvid string) PBResp
    StopServer stops a server

func (r *PBResp) PrintHeaders()
    PrintHeaders prints the http headers as k,v pairs

type RestRequest_Properties struct {
    Method  string            `json:"method"`
    Headers map[string]string `json:"headers"`
    Body    string            `json:"body"`
    Url     string            `json:"url"`
}

type StringIfaceMap map[string]interface{}

type StringMap map[string]string
    MetaData is a map for metadata returned in a PBResp.Body


