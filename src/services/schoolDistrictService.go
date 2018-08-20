/*
    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero General Public License as
    published by the Free Software Foundation, either version 3 of the
    License, or (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

    You should have received a copy of the GNU Affero General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package services

import (
    "github.com/jinzhu/gorm"

    skaioskit "github.com/nathanmentley/skaioskit-go-core"

    "skaioskit/models"
)

type ISchoolDistrictService interface {
    CreateSchoolDistrict(models.SchoolDistrict) models.SchoolDistrict
    UpdateSchoolDistrict(models.SchoolDistrict) models.SchoolDistrict
    GetSchoolDistrict(uint) (models.SchoolDistrict, error)
    GetSchoolDistricts(skaioskit.QueryRequest) ([]models.SchoolDistrict, uint64, error)
    EnsureSchoolDistrictTable()
    EnsureSchoolDistrict(models.SchoolDistrict)
}

type SchoolDistrictService struct {
    db *gorm.DB
}
func NewSchoolDistrictService(db *gorm.DB) *SchoolDistrictService {
    return &SchoolDistrictService{db: db}
}
func (p *SchoolDistrictService) CreateSchoolDistrict(school models.SchoolDistrict) models.SchoolDistrict {
    p.db.Create(&school)
    return school
}
func (p *SchoolDistrictService) UpdateSchoolDistrict(school models.SchoolDistrict) models.SchoolDistrict {
    p.db.Save(&school)
    return school
}
func (p *SchoolDistrictService) GetSchoolDistrict(code uint) (models.SchoolDistrict, error) {
    var school models.SchoolDistrict
    err := p.db.Where(&models.SchoolDistrict{Code: code}).First(&school).Error
    return school, err
}
func (p *SchoolDistrictService) GetSchoolDistricts(query skaioskit.QueryRequest) ([]models.SchoolDistrict, uint64, error) {
    var count uint64
    var schoolDistricts []models.SchoolDistrict
    schoolDistrict := models.SchoolDistrict{}

    skaioskit.BuildQueryWithoutPagination(p.db, query, &models.SchoolDistrict{}).Count(&count)
    err := skaioskit.BuildQuery(p.db, query, &schoolDistrict).Find(&schoolDistricts).Error
    return schoolDistricts, count, err
}
func (p *SchoolDistrictService) EnsureSchoolDistrictTable() {
    p.db.AutoMigrate(&models.SchoolDistrict{})
    p.db.Model(&models.SchoolDistrict{}).AddUniqueIndex("idx_school_district_code", "code")
    p.db.Model(&models.SchoolDistrict{}).AddForeignKey("county_code", "counties(code)", "RESTRICT", "RESTRICT")
    p.db.Model(&models.SchoolDistrict{}).AddForeignKey("jurisdiction_code", "jurisdictions(code)", "RESTRICT", "RESTRICT")
}
func (p *SchoolDistrictService) EnsureSchoolDistrict(school models.SchoolDistrict) {
    existing, err := p.GetSchoolDistrict(school.Code)

    if err != nil {
        p.CreateSchoolDistrict(school)
    } else {
        existing.Name = school.Name
        p.UpdateSchoolDistrict(existing)
    }
}

