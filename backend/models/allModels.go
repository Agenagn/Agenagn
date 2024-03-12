package models

import (
	"time"

	"gorm.io/gorm"
)

const (
    StatusAccepted   = "Accepted"
    StatusInProgress = "In Progress"
    StatusRejected   = "Rejected"
    EmploymentFullTime = "Full time"
    EmploymentPartTime = "Part time"
    EmploymentRemote   = "Remote"
)

type Company struct {
    gorm.Model
    Name              string
    RegistrationEmail string `gorm:"unique"`
    Password          string 
    Location          string
    Sector            string
    NoOfEmployees     int
    Logo              string
    City              string
    Region            string
    SubCityZone       string `gorm:"column:sub_city_zone"`
    WebsiteLink       string
    PhoneNumber       string
    EmailAddress      string `gorm:"unique"`
    Jobs              []Job `gorm:"foreignKey:CompanyID"`
}

type JobSeeker struct {
    gorm.Model
    Name              string `gorm:"column:name" binding:"required" json:"name"`
    Password          string `gorm:"column:password" binding:"required" json:"password"`
    Gender            string `gorm:"column:gender"  json:"gender"`
    DoB               time.Time `gorm:"column:dob"  json:"dob"`
    City              string `gorm:"column:city" json:"city"`
    Region            string `gorm:"column:region"  json:"region"`
    SubCityZone       string `gorm:"column:sub_city_zone"  json:"sub_city_zone"`
    PortfolioWebsite  string `gorm:"column:portfolio_website"`
    LinkedIn          string `gorm:"column:linked_in"  json:"linked_in"`
    // GitHub            string `gorm:"column:github"  json:"github"`
    PhoneNumber       string `gorm:"column:phone_number" binding:"required" json:"phone_number"`
    EmailAddress      string `gorm:"unique" binding:"required" json:"email_address"`
    ProfessionalTitle string `gorm:"column:professional_title"`
    Bio               string `gorm:"column:bio"  json:"bio"`
    Applications      []Application `gorm:"foreignKey:UserID"`
    Bookmarks         []Bookmark    `gorm:"foreignKey:UserID"`
    EducationInfos    []EducationInfo `gorm:"foreignKey:UserID"`
    WorkHistories     []WorkHistory `gorm:"foreignKey:UserID"`
}

type Job struct {
    gorm.Model
    Title            string
    Description      string
    ExperienceLevel  string `gorm:"column:experience_level"`
    DatePosted       time.Time `gorm:"column:date_posted"`
    DateExpiry       time.Time `gorm:"column:date_expiry"`
    Vacancies        int
    CompanyID        uint `gorm:"column:company_id"`
    Location         string
    EmploymentType   string `gorm:"column:employment_type;type:varchar(255);default:'Full time'"`
    Applications     []Application `gorm:"foreignKey:JobID"`
    Bookmarks        []Bookmark    `gorm:"foreignKey:JobID"`
    Sectors          []Sector      `gorm:"many2many:job_sectors;"`
}

type Application struct {
    gorm.Model
    JobID         uint `gorm:"column:job_id"`
    UserID        uint `gorm:"column:user_id"`
    Resume        string
    CoverLetter   string `gorm:"column:cover_letter"`
    Status        string `gorm:"column:status;type:varchar(255);default:'In Progress'"`
}

type EducationInfo struct {
    gorm.Model
    LearningTitle    string `gorm:"column:learning_title"`
    SchoolUniversity string `gorm:"column:school_university"`
    StartDate        time.Time `gorm:"column:start_date"`
    EndDate          time.Time `gorm:"column:end_date"`
    Description      string
    UserID           uint `gorm:"column:user_id"`
}

type WorkHistory struct {
    gorm.Model
    Position    string
    Company     string
    StartDate   time.Time `gorm:"column:start_date"`
    EndDate     time.Time `gorm:"column:end_date"`
    Description string
    UserID      uint `gorm:"column:user_id"`
}


type Bookmark struct {
    gorm.Model
    JobID  uint `gorm:"column:job_id"`
    UserID uint `gorm:"column:user_id"`
    UniqueConstraint string `gorm:"uniqueIndex:idx_job_user"`
}

type Sector struct {
    gorm.Model
    Title string
}

type JobSector struct {
    gorm.Model
    SectorID uint `gorm:"column:sector_id"`
    JobID    uint `gorm:"column:job_id"`
    UniqueConstraint string `gorm:"uniqueIndex:idx_sector_job"`
}


