package models

type Organization struct {
	Model
	Name string `gorm:"type:string;size:255" json:"name"`
}

type OrganizationModel struct{}

func (m OrganizationModel) Show(id string) Organization {
	organization := Organization{}
	db.First(&organization, id)

	return organization
}

func (m OrganizationModel) Get() []Organization {
	var organizations []Organization

	db.Find(&organizations)

	return organizations
}

func (m OrganizationModel) Create(organization Organization) Organization {
	db.Create(&organization)

	return organization
}
