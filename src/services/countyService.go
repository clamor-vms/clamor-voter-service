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

    clamor "github.com/clamor-vms/clamor-go-core"

    "clamor/models"
)

type ICountyService interface {
    CreateCounty(models.County) models.County
    UpdateCounty(models.County) models.County
    GetCounties(clamor.QueryRequest) ([]models.County, uint64, error)
    GetCounty(uint) (models.County, error)
    EnsureCountyTable()
    EnsureCounty(models.County)
}

type CountyService struct {
    db *gorm.DB
}
func NewCountyService(db *gorm.DB) *CountyService {
    return &CountyService{db: db}
}
func (p *CountyService) CreateCounty(county models.County) models.County {
    p.db.Create(&county)
    return county
}
func (p *CountyService) UpdateCounty(county models.County) models.County {
    p.db.Save(&county)
    return county
}
func (p *CountyService) GetCounty(code uint) (models.County, error) {
    var county models.County
    err := p.db.Where(&models.County{Code: code}).First(&county).Error
    return county, err
}
func (p *CountyService) GetCounties(query clamor.QueryRequest) ([]models.County, uint64, error) {
    var count uint64
    var counties []models.County

    clamor.BuildQueryWithoutPagination(p.db, query, &models.County{}).Count(&count)
    err := clamor.BuildQuery(p.db, query, &models.County{}).Find(&counties).Error
    return counties, count, err
}
func (p *CountyService) EnsureCountyTable() {
    p.db.AutoMigrate(&models.County{})
    p.db.Model(&models.County{}).AddUniqueIndex("idx_county_code", "code")
}
func (p *CountyService) EnsureCounty(county models.County) {
    existing, err := p.GetCounty(county.Code)
    if err != nil {
        p.CreateCounty(county)
    } else {
        existing.Name = county.Name
        p.UpdateCounty(existing)
    }
}
