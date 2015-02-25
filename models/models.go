// Packate models defines types and methods used to interact with user
// input data and data associated with a table in the DB
package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// User struct used to read and write to DB table "users"
// see const category in home.go for Interest codes
type User struct {
	Uid       int       // Database primary key. AutoIncrement value
	Username string    `orm:"size(30);unique"`
	Fullname string    `orm:"size(60)"`
	Email    string    `orm:"size(256)"` // encoded email address
	Password string    `orm:"size(128)"` // password hash value
	AutoLog  bool      // true if user selected "Remember Me"
	Interest uint8     // User's primary interest as a category code
	RegKey   string    `orm:"size(60)"` // used to confirm registration email
	ResetKey string    `orm:"size(60)"` // used to confirm password reset
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"auto_now;type(datetime)"`
	UserStatus bool    // true if user is activity
}

// User database CRUD methods include Insert, Read, Update and Delete
func (usr *User) Insert() error {
	if _, err := orm.NewOrm().Insert(usr); err != nil {
		return err
	}
	return nil
}

func (usr *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(usr, fields...); err != nil {
		return err
	}
	return nil
}

func (usr *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(usr, fields...); err != nil {
		return err
	}
	return nil
}

func (usr *User) Delete() error {
	if _, err := orm.NewOrm().Delete(usr); err != nil {
		return err
	}
	return nil
}

// RegFrm struct to hold Registration page and Profile page form values
type RegFrm struct {
	Email    string `form:"email"`
	Username string `form:"username"`
	Fullname string `form:"fullname"`
	Password string `form:"pw1"`
	Confirm  string `form:"pw2"`
	AutoLog  string `form:"autolog"`
}

// Comment struct used to read and write to DB table "comment"
// see const category in home.go for Category codes
type Comment struct {
	Cid       int    // Database primary key. AutoIncrement value
	User     *User  `orm:"rel(fk);index"` // Indexed Foreign Key -> User.Id
	Txt      string `orm:"size(4096)"`
	Category uint8  // Indicates page used to post the comment
	Archived bool
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
}

// Comment database access methods include Insert and Read
// TODO Add Archive and review need for Delete
func (cmnt *Comment) Insert() error {
	if _, err := orm.NewOrm().Insert(cmnt); err != nil {
		return err
	}
	return nil
}

func (cmnt *Comment) Read(fields ...string) error {
	if err := orm.NewOrm().Read(cmnt, fields...); err != nil {
		return err
	}
	return nil
}

// Role struct used for user.
type Role struct {
	Rid       	int    // Database primary key. AutoIncrement value
	RoleName    string `orm:"size(4096)"`
}

// Role database CRUD methods include Insert, Read, Update and Delete
func (usr *Role) Insert() error {
	if _, err := orm.NewOrm().Insert(usr); err != nil {
		return err
	}
	return nil
}

func (usr *Role) Read(fields ...string) error {
	if err := orm.NewOrm().Read(usr, fields...); err != nil {
		return err
	}
	return nil
}

func (usr *Role) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(usr, fields...); err != nil {
		return err
	}
	return nil
}

func (usr *Role) Delete() error {
	if _, err := orm.NewOrm().Delete(usr); err != nil {
		return err
	}
	return nil
}

type RoleFrm struct {
	RoleName	string `form:"rolename"`
}


type PatientInformation struct {
	Pid       			int     // Database primary key. AutoIncrement value
	UserId				int		//`orm:"size(60)"`
	SampleCode    		string    `orm:"size(256)"` // sample code
	PatientCode    		string    `orm:"size(256)"` // patient code
    PatientCity			string    `orm:"size(256)"` //  code
    PatientAge			string    `orm:"size(256)"` //  code
    PatientGender		string    `orm:"size(256)"` //  code
    PatientMHistory		string    `orm:"size(256)"` //  code
    PatientMedication	string    `orm:"size(256)"` //  code
    PatientFHistory		string    `orm:"size(256)"` //  code
    PatientFHabits		string    `orm:"size(256)"` //  code
    PatientSmoking		string    `orm:"size(256)"` //  code
    PatientAlcohol		string    `orm:"size(256)"` //  code
    PatientHeight		string    `orm:"size(256)"` //  code
    PatientWeight		string    `orm:"size(256)"` //  code
    PatientBMI			string    `orm:"size(256)"` //  code
    PatientWaistCircum	string    `orm:"size(256)"` //  code
    PatientDiabp		string    `orm:"size(256)"` //  code
    PatientSysbp		string    `orm:"size(256)"` //  code
    Comments			string    `orm:"size(4096)"` //  code
    CreatedBy			string    `orm:"size(256)"` //  code
    UpdateBy			string    `orm:"size(256)"` //  code
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"auto_now;type(datetime)"`
}

// PatientInformation database CRUD methods include Insert, Read, Update and Delete
func (pinfo *PatientInformation) Insert() error {
	if _, err := orm.NewOrm().Insert(pinfo); pinfo != nil {
		return err
	}
	return nil
}

func (pinfo *PatientInformation) Read(fields ...string) error {
	if err := orm.NewOrm().Read(pinfo, fields...); err != nil {
		return err
	}
	return nil
}

func (pinfo *PatientInformation) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(pinfo, fields...); err != nil {
		return err
	}
	return nil
}

func (pinfo *PatientInformation) Delete() error {
	if _, err := orm.NewOrm().Delete(pinfo); err != nil {
		return err
	}
	return nil
}

// RegFrm struct to hold Registration page and Profile page form values
type PatientInformationFrm struct {
	SampleCode    		string    `form:"samplecode"` // sample code
	PatientCode    		string    `form:"patientcode"` // patient code
    PatientCity			string    `form:"patientcity"` //  code
    PatientAge			string    `form:"patientage"` //  code
    PatientGender		string    `form:"patientgender"` //  code
    PatientMHistory		string    `form:"patientmhistory"` //  code
    PatientMedication	string    `form:"patientmedication"` //  code
    PatientFHistory		string    `form:"patientfhistory"` //  code
    PatientFHabits		string    `form:"patientfhabits"` //  code
    PatientSmoking		string    `form:"patientsmoking"` //  code
    PatientAlcohol		string    `form:"patientalcohol"` //  code
    PatientHeight		string    `form:"patientheight"` //  code
    PatientWeight		string    `form:"patientweight"` //  code
    PatientBMI			string    `form:"patientbmi"` //  code
    PatientWaistCircum	string    `form:"patientwaistcircum"` //  code
    PatientDiabp		string    `form:"patientdiabp"` //  code
    PatientSysbp		string    `form:"patientsysbp"` //  code	
}

//Cardiovascular Information

type CardioInformation struct {
	Cvid       			int     // Database primary key. AutoIncrement value
	UserId				int		// User id
	PatientId     		*PatientInformation	`orm:"rel(fk);index"` // Indexed Foreign Key -> User.Id
	SampleCode    		string    `orm:"size(256)"` // sample code
	Ecg    				string    `orm:"size(256)"` // patient code
    TroponinI			string    `orm:"size(256)"` //  code
    CkMb				string    `orm:"size(256)"` //  code
    LdH					string    `orm:"size(256)"` //  code
    Comments			string    `orm:"size(4096)"` //  code
    CreatedBy			string    `orm:"size(256)"` //  code
    UpdateBy			string    `orm:"size(256)"` //  code
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"auto_now;type(datetime)"`
}

// PatientInformation database CRUD methods include Insert, Read, Update and Delete
func (cvinfo *CardioInformation) Insert() error {
	if _, err := orm.NewOrm().Insert(cvinfo); err != nil {
		return err
	}
	return nil
}

func (cvinfo *CardioInformation) Read(fields ...string) error {
	if err := orm.NewOrm().Read(cvinfo, fields...); err != nil {
		return err
	}
	return nil
}

func (cvinfo *CardioInformation) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(cvinfo, fields...); err != nil {
		return err
	}
	return nil
}

func (cvinfo *CardioInformation) Delete() error {
	if _, err := orm.NewOrm().Delete(cvinfo); err != nil {
		return err
	}
	return nil
}

// RegFrm struct to hold Registration page and Profile page form values
type CardioInformationFrm struct {
	SampleCode    		string    `form:"samplecode"` // sample code
	PatientCode    		string    `form:"patientcode"` // patient code
	Ecg    				string    `form:"ecg"` // patient code
    TroponinI			string    `form:"troponin1"` //  code
    CkMb				string    `form:"ckmb"` //  code
    LdH					string    `form:"ldh"` //  code
}



//Cytokines Information

type CytokinesInformation struct {
	Cyid       			int     // Database primary key. AutoIncrement value
	UserId				int		// User id
	PatientId     		*PatientInformation	`orm:"rel(fk);index"` // Indexed Foreign Key -> User.Id
	SampleCode    		string    `orm:"size(256)"` // sample code
	il_1b    			string    `orm:"size(256)"` // patient code
    il_1ra				string    `orm:"size(256)"` //  code
    il_2				string    `orm:"size(256)"` //  code
    il_4				string    `orm:"size(256)"` //  code
    il_5				string    `orm:"size(256)"` //  code
    il_6				string    `orm:"size(256)"` //  code
    il_7				string    `orm:"size(256)"` //  code
    il_8				string    `orm:"size(256)"` //  code
    il_9				string    `orm:"size(256)"` //  code
    il_10				string    `orm:"size(256)"` //  code
    il_12_p70			string    `orm:"size(256)"` //  code
    il_13				string    `orm:"size(256)"` //  code
    il_15				string    `orm:"size(256)"` //  code
    il_17				string    `orm:"size(256)"` //  code
    basic_fgf			string    `orm:"size(256)"` //  code
    eotaxin				string    `orm:"size(256)"` //  code
    g_csf				string    `orm:"size(256)"` //  code
    gm_csf				string    `orm:"size(256)"` //  code
    ifn_g				string    `orm:"size(256)"` //  code
    ip_10				string    `orm:"size(256)"` //  code
    mcp_1_mcaf			string    `orm:"size(256)"` //  code
    mip_1a				string    `orm:"size(256)"` //  code
    mip_1b				string    `orm:"size(256)"` //  code
    pdgf_bb				string    `orm:"size(256)"` //  code
    rantes   			string    `orm:"size(256)"` //  code
    tnf_a				string    `orm:"size(256)"` //  code
    vegf				string    `orm:"size(256)"` //  code
    Comments			string    `orm:"size(4096)"` //  code
    CreatedBy			string    `orm:"size(256)"` //  code
    UpdateBy			string    `orm:"size(256)"` //  code
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"auto_now;type(datetime)"`
}

// PatientInformation database CRUD methods include Insert, Read, Update and Delete
func (cyinfo *CytokinesInformation) Insert() error {
	if _, err := orm.NewOrm().Insert(cyinfo); err != nil {
		return err
	}
	return nil
}

func (cyinfo *CytokinesInformation) Read(fields ...string) error {
	if err := orm.NewOrm().Read(cyinfo, fields...); err != nil {
		return err
	}
	return nil
}

func (cyinfo *CytokinesInformation) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(cyinfo, fields...); err != nil {
		return err
	}
	return nil
}

func (cyinfo *CytokinesInformation) Delete() error {
	if _, err := orm.NewOrm().Delete(cyinfo); err != nil {
		return err
	}
	return nil
}

// RegFrm struct to hold Registration page and Profile page form values
type CytokinesInformation struct {
	SampleCode    		string    `form:"samplecode"` // sample code
	PatientCode    		string    `form:"patientcode"` // patient code
	il_1b    			string    `form:"il_1b"` // patient code
    il_1ra				string    `form:"il_1ra"` //  code
    il_2				string    `form:"il_2"` //  code
    il_4				string    `form:"il_4"` //  code
    il_5				string    `form:"il_5"` //  code
    il_6				string    `form:"size(256)"` //  code
    il_7				string    `form:"il_7"` //  code
    il_8				string    `form:"il_8"` //  code
    il_9				string    `form:"il_9"` //  code
    il_10				string    `form:"il_10"` //  code
    il_12_p70			string    `form:"il_12_p70"` //  code
    il_13				string    `form:"il_13"` //  code
    il_15				string    `form:"il_15"` //  code
    il_17				string    `form:"il_17"` //  code
    basic_fgf			string    `form:"basic_fgf"` //  code
    eotaxin				string    `form:"eotaxin"` //  code
    g_csf				string    `form:"g_csf"` //  code
    gm_csf				string    `form:"gm_csf"` //  code
    ifn_g				string    `form:"ifn_g"` //  code
    ip_10				string    `form:"ip_10"` //  code
    mcp_1_mcaf			string    `form:"mcp_1_mcaf"` //  code
    mip_1a				string    `form:"mip_1a"` //  code
    mip_1b				string    `form:"mip_1b"` //  code
    pdgf_bb				string    `form:"pdgf_bb"` //  code
    rantes   			string    `form:"rantes"` //  code
    tnf_a				string    `form:"tnf_a"` //  code
    vegf				string    `form:"vegf"` //  code
}

// Register User and Comment models with the Beego ORM
func init() {
	orm.RegisterModel(new(User), new(Comment), new(Role), new(PatientInformation), new(CardioInformationFrm), new(CytokinesInformation))
}
