package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

// APIResponse is the standard response format
type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Project model
type Project struct {
	ID               int            `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Title            string         `json:"title" gorm:"not null;column:title"`
	Description      string         `json:"description" gorm:"type:text;column:description"`
	TechStack        pq.StringArray `json:"techStack" gorm:"type:text[];column:tech_stack"`
	ImageTitle       string         `json:"imageTitle" gorm:"column:image_title"`
	ImageDescription string         `json:"imageDescription" gorm:"type:text;column:image_description"`
	ImageURL         string         `json:"imageUrl" gorm:"type:varchar(500);column:image_url"`
	ButtonText       string         `json:"buttonText" gorm:"column:button_text"`
	DetailURL        string         `json:"detailUrl" gorm:"type:varchar(500);column:detail_url"`
	Category         string         `json:"category" gorm:"type:varchar(50);column:category"`
	CreatedAt        time.Time      `json:"createdAt" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt        time.Time      `json:"updatedAt" gorm:"autoUpdateTime;column:updated_at"`
}

// TableName specifies the table name for Project model
func (Project) TableName() string {
	return "projects"
}

// ProjectImage model
type ProjectImage struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	ProjectID int       `json:"projectId" gorm:"not null;index;column:project_id"`
	ImageURL  string    `json:"imageUrl" gorm:"type:varchar(500);not null;column:image_url"`
	Order     int       `json:"order" gorm:"default:0;column:order"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime;column:created_at"`
}

// TableName specifies the table name for ProjectImage model
func (ProjectImage) TableName() string {
	return "project_images"
}

// Paragraph for About content
type Paragraph struct {
	Text string `json:"text"`
	Type string `json:"type"` // "highlight" or "normal"
}

// Hobby for About content
type Hobby struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

// AboutContent represents the JSONB content in AboutCard
type AboutContent struct {
	Paragraphs []Paragraph `json:"paragraphs,omitempty"`
	Hobbies    []Hobby     `json:"hobbies,omitempty"`
	Quote      string      `json:"quote,omitempty"`
}

// Implement driver.Valuer and sql.Scanner for JSONB
func (ac AboutContent) Value() (driver.Value, error) {
	return json.Marshal(ac)
}

func (ac *AboutContent) Scan(value interface{}) error {
	if value == nil {
		*ac = AboutContent{}
		return nil
	}

	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return gorm.ErrInvalidData
	}

	if len(bytes) == 0 {
		*ac = AboutContent{}
		return nil
	}

	return json.Unmarshal(bytes, ac)
}

// AboutCard model
type AboutCard struct {
	ID        string       `json:"id" gorm:"primaryKey;type:varchar(50);column:id"`
	Title     string       `json:"title" gorm:"not null;column:title"`
	Icon      string       `json:"icon" gorm:"type:varchar(50);column:icon"`
	Content   AboutContent `json:"content" gorm:"type:jsonb;column:content"`
	CreatedAt time.Time    `json:"createdAt" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt time.Time    `json:"updatedAt" gorm:"autoUpdateTime;column:updated_at"`
}

// TableName specifies the table name for AboutCard model
func (AboutCard) TableName() string {
	return "about_cards"
}

// Skill model
type Skill struct {
	Name       string `json:"name"`
	Percentage int    `json:"percentage"` // 0-100
	Icon       string `json:"icon"`
}

// Skills array for JSONB
type Skills []Skill

func (s Skills) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *Skills) Scan(value interface{}) error {
	if value == nil {
		*s = []Skill{}
		return nil
	}

	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return gorm.ErrInvalidData
	}

	if len(bytes) == 0 {
		*s = []Skill{}
		return nil
	}

	return json.Unmarshal(bytes, s)
}

// SkillCategory model
type SkillCategory struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Title       string    `json:"title" gorm:"not null;column:title"`
	Description string    `json:"description" gorm:"type:text;column:description"`
	Icon        string    `json:"icon" gorm:"type:varchar(50);column:icon"`
	Skills      Skills    `json:"skills" gorm:"type:jsonb;column:skills"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime;column:updated_at"`
}

// TableName specifies the table name for SkillCategory model
func (SkillCategory) TableName() string {
	return "skill_categories"
}

// ContactInfo model
type ContactInfo struct {
	Icon  string `json:"icon"`
	Label string `json:"label"`
	Value string `json:"value"`
	Link  string `json:"link"`
}

// ContactInfoArray for JSONB
type ContactInfoArray []ContactInfo

func (c ContactInfoArray) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *ContactInfoArray) Scan(value interface{}) error {
	if value == nil {
		*c = []ContactInfo{}
		return nil
	}

	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return gorm.ErrInvalidData
	}

	if len(bytes) == 0 {
		*c = []ContactInfo{}
		return nil
	}

	return json.Unmarshal(bytes, c)
}

// SocialLink model
type SocialLink struct {
	Name  string `json:"name"`
	Icon  string `json:"icon"`
	URL   string `json:"url"`
	Color string `json:"color"`
}

// SocialLinkArray for JSONB
type SocialLinkArray []SocialLink

func (s SocialLinkArray) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *SocialLinkArray) Scan(value interface{}) error {
	if value == nil {
		*s = []SocialLink{}
		return nil
	}

	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return gorm.ErrInvalidData
	}

	if len(bytes) == 0 {
		*s = []SocialLink{}
		return nil
	}

	return json.Unmarshal(bytes, s)
}

// ContactPageData model
type ContactPageData struct {
	ID          int              `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Title       string           `json:"title" gorm:"column:title"`
	Description string           `json:"description" gorm:"type:text;column:description"`
	ContactInfo ContactInfoArray `json:"contactInfo" gorm:"type:jsonb;column:contact_info"`
	SocialLinks SocialLinkArray  `json:"socialLinks" gorm:"type:jsonb;column:social_links"`
	UpdatedAt   time.Time        `json:"updatedAt" gorm:"autoUpdateTime;column:updated_at"`
}

// TableName specifies the table name for ContactPageData model
func (ContactPageData) TableName() string {
	return "contact_data"
}

// ContactMessage model (for future use, not exposed in API)
type ContactMessage struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null"`
	Message   string    `json:"message" gorm:"type:text;not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
}

// Category model for projects categories endpoint
type Category struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

// SidebarButton model for about sidebar endpoint
type SidebarButton struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Icon  string `json:"icon"`
}

// SkillsPageData model for skills page metadata
type SkillsPageData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
