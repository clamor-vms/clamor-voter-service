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

type IVillageService interface {
    CreateVillage(models.Village) models.Village
    UpdateVillage(models.Village) models.Village
    GetVillage(uint) (models.Village, error)
    GetVillages(clamor.QueryRequest) ([]models.Village, uint64, error)
    EnsureVillageTable()
    EnsureVillage(models.Village)
}

type VillageService struct {
    db *gorm.DB
}
func NewVillageService(db *gorm.DB) *VillageService {
    return &VillageService{db: db}
}
func (p *VillageService) CreateVillage(village models.Village) models.Village {
    p.db.Create(&village)
    return village
}
func (p *VillageService) UpdateVillage(village models.Village) models.Village {
    p.db.Save(&village)
    return village
}
func (p *VillageService) GetVillage(code uint) (models.Village, error) {
    var village models.Village
    err := p.db.Where(&models.Village{Code: code}).First(&village).Error
    return village, err
}
func (p *VillageService) GetVillages(query clamor.QueryRequest) ([]models.Village, uint64, error) {
    var count uint64
    var villages []models.Village
    village := models.Village{}

    clamor.BuildQueryWithoutPagination(p.db, query, &models.Village{}).Count(&count)
    err := clamor.BuildQuery(p.db, query, &village).Find(&villages).Error
    return villages, count, err
}
func (p *VillageService) EnsureVillageTable() {
    p.db.AutoMigrate(&models.Village{})
    p.db.Model(&models.Village{}).AddUniqueIndex("idx_village_code", "code")
    p.db.Model(&models.Village{}).AddForeignKey("county_code", "counties(code)", "RESTRICT", "RESTRICT")
    p.db.Model(&models.Village{}).AddForeignKey("jurisdiction_code", "jurisdictions(code)", "RESTRICT", "RESTRICT")
}
func (p *VillageService) EnsureVillage(village models.Village) {
    existing, err := p.GetVillage(village.Code)

    if err != nil {
        p.CreateVillage(village)
    } else {
        existing.Name = village.Name
        p.UpdateVillage(existing)
    }
}
