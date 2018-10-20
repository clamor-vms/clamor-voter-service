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

type IJurisdictionService interface {
    CreateJurisdiction(models.Jurisdiction) models.Jurisdiction
    UpdateJurisdiction(models.Jurisdiction) models.Jurisdiction
    GetJurisdiction(uint) (models.Jurisdiction, error)
    GetJurisdictions(clamor.QueryRequest) ([]models.Jurisdiction, uint64, error)
    EnsureJurisdictionTable()
    EnsureJurisdiction(models.Jurisdiction)
}

type JurisdictionService struct {
    db *gorm.DB
}
func NewJurisdictionService(db *gorm.DB) *JurisdictionService {
    return &JurisdictionService{db: db}
}
func (p *JurisdictionService) CreateJurisdiction(jurisdiction models.Jurisdiction) models.Jurisdiction {
    p.db.Create(&jurisdiction)
    return jurisdiction
}
func (p *JurisdictionService) UpdateJurisdiction(jurisdiction models.Jurisdiction) models.Jurisdiction {
    p.db.Save(&jurisdiction)
    return jurisdiction
}
func (p *JurisdictionService) GetJurisdiction(code uint) (models.Jurisdiction, error) {
    var jurisdiction models.Jurisdiction
    err := p.db.Where(&models.Jurisdiction{Code: code}).First(&jurisdiction).Error
    return jurisdiction, err
}
func (p *JurisdictionService) GetJurisdictions(query clamor.QueryRequest) ([]models.Jurisdiction, uint64, error) {
    var count uint64
    var jurisdictions []models.Jurisdiction
    jurisdiction := models.Jurisdiction{}

    clamor.BuildQueryWithoutPagination(p.db, query, &models.Jurisdiction{}).Count(&count)
    err := clamor.BuildQuery(p.db, query, &jurisdiction).Find(&jurisdictions).Error
    return jurisdictions, count, err
}
func (p *JurisdictionService) EnsureJurisdictionTable() {
    p.db.AutoMigrate(&models.Jurisdiction{})
    p.db.Model(&models.Jurisdiction{}).AddUniqueIndex("idx_jurisdiction_code", "code")
    p.db.Model(&models.Jurisdiction{}).AddForeignKey("county_code", "counties(code)", "RESTRICT", "RESTRICT")
}
func (p *JurisdictionService) EnsureJurisdiction(jurisdiction models.Jurisdiction) {
    existing, err := p.GetJurisdiction(jurisdiction.Code)

    if err != nil {
        p.CreateJurisdiction(jurisdiction)
    } else {
        existing.Name = jurisdiction.Name
        p.UpdateJurisdiction(existing)
    }
}

