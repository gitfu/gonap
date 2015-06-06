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

func SetAuth(u, p string)
    SetAuth is used to set Username and Passwd. Username and Passwd are
    declared in config.go

func SetDepth(newdepth string) string
    SetDepth is used to set Depth

func SetEndpoint(newendpoint string) string
    SetEnpoint is used to set the REST Endpoint. Endpoint is declared in
    config.go

TYPES

type Datacenter struct {
    Id_Type_Href
    MetaData   MetaData              `json:"metadata,omitempty"`
    Properties Datacenter_Properties `json:"properties"`
    Entities   Datacenter_Entities   `json:"entities,omitempty"`
    Resp       PBResp                `json:"-"`
}
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

func (dc *Datacenter) Save()
    Save converts the datacenter struct's properties to json and "patch"es
    them to the rest server

func (dc *Datacenter) ToJson() string
    Datacenter.ToJson marshals the Datacenter struct into json

type Datacenter_Entities struct {
    Servers       Servers       `json:"servers,omitempty"`
    Loadbalancers Loadbalancers `json:"loadbalancers,omitempty"`
    Lans          Lans          `json:"lans,omitempty"`
    Volumes       Volumes       `json:"volumes,omitempty"`
}
    Datacenter_Entities is a struct inside a Datacenter struct to hold
    collections of other structs

type Datacenter_Properties struct {
    Name        string `json:"name"`
    Location    string `json:"location"`
    Description string `json:"description,omitempty"`
}

type Datacenters struct {
    Id_Type_Href
    Items []Datacenter `json:"items,omitempty"`
    Resp  PBResp       `json:"-"`
}
    Datacenters is a struct for Datacenter collections

func ListDatacenters() Datacenters
    ListDatacenters returns a Datacenter collection struct

func (dcs *Datacenters) ToJson() string
    Datacenter.ToJson marshals the Datacenter struct into json

type FwRule struct {
    Id_Type_Href
    MetaData   MetaData          `json:"metadata,omitempty"`
    Properties FwRule_Properties `json:"properties,omitempty"`
    Resp       PBResp            `json:"-"`
}
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

type FwRule_Properties struct {
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
    FwRule_Properties is a struct for firewall rule properties

type FwRules struct {
    Id_Type_Href
    Items []FwRule `json:"items,omitempty"`
    Resp  PBResp   `json:"-"`
}
    FwRules is the struct for firewall rule collections

func ListFwRules(dcid, srvid, nicid string) FwRules
    ListFwRules returns a collection of firewall rules

type Id_Type_Href struct {
    Id   string `json:"id"`
    Type string `json:"type"`
    Href string `json:"href"`
}

type Image struct {
    Id_Type_Href
    MetaData   MetaData         `json:"metadata,omitempty"`
    Properties Image_Properties `json:"properties,omitempty"`
    Resp       PBResp           `json:"-"`
}
    Image is the struct for image and cdrom data

func AttachCdrom(dcid string, srvid string, cdid string) Image

func CreateImage(jason []byte) Image
    CreateImage creates an Image and returns an Image struct

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

type Image_Properties struct {
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
    Image_Properties for image and cdrom data

type Images struct {
    Id_Type_Href
    Items []Image `json:"items,omitempty"`
    Resp  PBResp  `json:"-"`
}
    Images is a struct for Image and cdrom collections

func ListAttachedCdroms(dcid, srvid string) Images

func ListImages() Images
    ListImages returns an Images struct

type Ipblock struct {
    Id_Type_Href
    MetaData   MetaData           `json:"metadata,omitempty"`
    Properties Ipblock_Properties `json:"properties"`
    Resp       PBResp             `json:"-"`
}
    Ipblock is the struct for Ipblock data

func GetIpBlock(ipblockid string) Ipblock

func ReleaseIpBlock(ipblockid string) Ipblock

func ReserveIpBlock(jason []byte) Ipblock

type Ipblock_Properties struct {
    Location string `json:"location"`
    Size     int    `json:"size"`
}
    Ipblock_Properties is just that

type Ipblocks struct {
    Id_Type_Href
    Items []Ipblock `json:"items,omitempty"`
    Resp  PBResp    `json:"-"`
}
    Ipblocks is the struct for an Ipblock collection

func ListIpBlocks() Ipblocks
    ListIpBlocks

type Lan struct {
    Id_Type_Href
    MetaData   MetaData       `json:"metadata,omitempty"`
    Properties Lan_Properties `json:"properties"`
    Entities   Lan_Entities   `json:"entities,omitempty"`
    Resp       PBResp         `json:"-"`
}

func CreateLan(dcid string, jason []byte) Lan
    CreateLan creates a lan in the datacenter from a jason []byte and
    returns a Lan struct

func DeleteLan(dcid, lanid string) Lan
    DeleteLan deletes a lan where id == lanid

func GetLan(dcid, lanid string) Lan
    GetLan pulls data for the lan where id = lanid returns a Lan struct

func PatchLan(dcid string, lanid string, jason []byte) Lan

func UpdateLan(dcid string, lanid string, jason []byte) Lan

type Lan_Entities struct {
    Nics Nics `json:"nics,omitempty"`
}

type Lan_Properties struct {
    Name   string `json:"name,omitempty"`
    Public bool   `json:"public,omitempty"`
}

type Lans struct {
    Id_Type_Href
    Items []Lan  `json:"items,omitempty"`
    Resp  PBResp `json:"-"`
}

func ListLans(dcid string) Lans
    ListLan returns a Lans struct collection for lans in the Datacenter

type Loadbalancer struct {
    Id_Type_Href
    MetaData   MetaData                `json:"metadata,omitempty"`
    Properties Loadbalancer_Properties `json:"properties,omitempty"`
    Entities   Loadbalancer_Entities   `json:"entities,omitempty"`
    Resp       PBResp                  `json:"-"`
}

func CreateLoadbalancer(dcid string, jason []byte) Loadbalancer
    Createloadbalancer creates a loadbalancer in the datacenter from a jason
    []byte and returns a Loadbalancer struct

func DeleteLoadbalancer(dcid, lbalid string) Loadbalancer

func GetLoadbalancer(dcid, lbalid string) Loadbalancer
    GetLoadbalancer pulls data for the Loadbalancer where id = lbalid
    returns a Loadbalancer struct

func PatchLoadBalancer(dcid string, lbalid string, jason []byte) Loadbalancer

func UpdateLoadbalancer(dcid string, lbalid string, jason []byte) Loadbalancer

type Loadbalancer_Entities struct {
    BalancedNics Nics `json:"balancednics,omitempty"`
}

type Loadbalancer_Properties struct {
    Name string `json:"name,omitempty"`
    Ip   string `json:"ip,omitempty"`
    Dhcp bool   `json:"dhcp,omitempty"`
}

type Loadbalancers struct {
    Id_Type_Href
    Items []Loadbalancer `json:"items,omitempty"`
    Resp  PBResp         `json:"-"`
}

func ListLoadbalancers(dcid string) Loadbalancers
    Listloadbalancers returns a Loadbalancers struct for loadbalancers in
    the Datacenter

type Location struct {
    Id_Type_Href
    MetaData   MetaData            `json:"metadata,omitempty"`
    Properties Location_Properties `json:"properties"`
    Resp       PBResp              `json:"-"`
}
    Location is the struct for a Location

func GetLocation(locid string) Location
    GetLocation returns a PBResp with data for a location in the Body

type Location_Properties struct {
    Name string `json:"name"`
}
    Location_Properties is aa struct for Locaation Properties

type Locations struct {
    Id_Type_Href
    Items []Location `json:"items,omitempty"`
    Resp  PBResp     `json:"-"`
}
    Locations is the struct for a Location Collection

func ListLocations() Locations
    ListLocations returns a PBResp with location collection data in the Body

type MetaData struct {
    LastModified   string `json:"lastModifiedDate,omitempty"`
    LastModifiedBy string `json:"lastModifiedBy,omitempty"`
    Created        string `json:"createdDate,omitempty"` //"2014-02-01T11:12:12Z",
    CreatedBy      string `json:"createdBy,omitempty"`
    State          string `json:"state,omitempty"`
    Etag           string `json:"etag,omitempty"`
}
    MetaData is a struct for metadata returned in a PBResp.Body

type Nic struct {
    Id_Type_Href
    MetaData   MetaData       `json:"metadata,omitempty"`
    Properties Nic_Properties `json:"properties,omitempty"`
    Resp       PBResp         `json:"-"`
}

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

type Nic_Properties struct {
    Name           string   `json:"name,omitempty"`
    Ips            []string `json:"ips,omitempty"`
    Dhcp           bool     `json:"dhcp,omitempty"`
    Lan            int      `json:"lan"`
    FirewallActive bool     `json:"firewallActive,omitempty"`
    Firewallrules  []FwRule `json:"firewallrules,omitempty"`
}

type Nics struct {
    Id_Type_Href
    Items []Nic  `json:"items,omitempty"`
    Resp  PBResp `json:"-"`
}

func AssociateNics(dcid string, lbalid string, jason []byte) Nics

func ListBalancedNics(dcid, lbalid string) Nics

func ListLanMembers(dcid, lanid string) Nics
    ListLanMembers returns a Nic struct collection for the Lan

func ListNics(dcid, srvid string) Nics
    ListNics returns a Nics struct collection

type PBResp struct {
    StatusCode int
    Headers    http.Header
    Body       []byte
    // contains filtered or unexported fields
}
    PBResp is the struct returned by all Rest request functions

func (r *PBResp) PrintHeaders()
    PrintHeaders prints the http headers as k,v pairs

type RestRequest struct {
    Id_Type_Href
    Metadata               MetaData `json:"metadata"`
    RestRequest_Properties `json:"properties"`
    Resp                   PBResp `json:"-"`
}

func GetRequest(requestid string) RestRequest

func StatusRequest(requestid string) RestRequest

type RestRequest_Properties struct {
    Method  string            `json:"method"`
    Headers map[string]string `json:"headers"`
    Body    string            `json:"body"`
    Url     string            `json:"url"`
}

type RestRequests struct {
    Id_Type_Href
    Metadata MetaData      `json:"metadata"`
    Items    []RestRequest `json:"items"`
    Resp     PBResp        `json:"-"`
}

func ListRequests() RestRequests

type Server struct {
    Id_Type_Href
    MetaData   MetaData          `json:"metadata,omitempty"`
    Properties Server_Properties `json:"properties,omitempty"`
    Entities   Server_Entities   `json:"entities,omitempty"`
    Resp       PBResp            `json:"-"`
}
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

type Server_Entities struct {
    Cdroms  Images  `json:"cdroms,omitempty"`
    Nics    Nics    `json:"nics,omitempty"`
    Volumes Volumes `json:"volumes,omitempty"`
}

type Server_Properties struct {
    Name             string `json:"name,omitempty"`
    Cores            int    `json:"cores,omitempty"`
    Ram              int    `json:"ram,omitempty"`
    Availabilityzone string `json:"availabilityzone,omitempty"`
    Licencetype      string `json:"licencetype,omitempty"`
    Bootvolume       string `json:"bootvolume,omitempty"`
    Bootcdrom        string `json:"bootcdrom,omitempty"`
}

type Servers struct {
    Id_Type_Href
    Items []Server `json:"items,omitempty"`
    Resp  PBResp   `json:"-"`
}
    Servers is a struct for Server struct collections

func ListServers(dcid string) Servers
    ListServers returns a server struct collection

type Snapshot struct {
    Id_Type_Href
    MetaData   MetaData            `json:"metadata,omitempty"`
    Properties Snapshot_Properties `json:"properties,omitempty"`
    Resp       PBResp              `json:"-"`
}
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

type Snapshot_Properties struct {
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
    Snapshot_Properties struct

type Snapshots struct {
    Id_Type_Href
    Items []Snapshot `json:"items,omitempty"`
    Resp  PBResp     `json:"-"`
}
    Snapshots struct for a Snapshot collection

func ListSnapshots() Snapshots
    ListSnapshots retrieves a collection of snapshot data returns a
    Snapshots struct

type Volume struct {
    Id_Type_Href
    MetaData   MetaData          `json:"metadata,omitempty"`
    Properties Volume_Properties `json:"properties,omitempty"`
    Resp       PBResp            `json:"-"`
}

func AttachVolume(dcid string, srvid string, volid string) Volume

func DeleteVolume(dcid, volid string) Volume

func DetachVolume(dcid, srvid, volid string) Volume

func GetAttachedVolume(dcid, srvid, volid string) Volume

func PatchVolume(dcid string, volid string, jason []byte) Volume

func UpdateVolume(dcid string, volid string, jason []byte) Volume

func (vol *Volume) Save()

type Volume_Properties struct {
    Name                string `json:"name,omitempty"`
    Size                int    `json:"size"`
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

type Volumes struct {
    Id_Type_Href
    Items []Volume `json:"items,omitempty"`
    Resp  PBResp   `json:"-"`
}

func ListAttachedVolumes(dcid, srvid string) Volumes

func ListVolumes(dcid string) Volumes
    ListVolumes returns a Volumes struct for volumes in the Datacenter

SUBDIRECTORIES

	p
	tests

