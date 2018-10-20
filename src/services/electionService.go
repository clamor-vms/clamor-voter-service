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

type IElectionService interface {
    CreateElection(models.Election) models.Election
    UpdateElection(models.Election) models.Election
    GetElection(uint64) (models.Election, error)
    GetElections(clamor.QueryRequest) ([]models.Election, uint64, error)
    EnsureElectionTable()
    EnsureElection(models.Election)
}

type ElectionService struct {
    db *gorm.DB
}
func NewElectionService(db *gorm.DB) *ElectionService {
    return &ElectionService{db: db}
}
func (p *ElectionService) CreateElection(election models.Election) models.Election {
    p.db.Create(&election)
    return election
}
func (p *ElectionService) UpdateElection(election models.Election) models.Election {
    p.db.Save(&election)
    return election
}
func (p *ElectionService) GetElection(code uint64) (models.Election, error) {
    var election models.Election
    err := p.db.Where(&models.Election{Code: code}).First(&election).Error
    return election, err
}
func (p *ElectionService) GetElections(query clamor.QueryRequest) ([]models.Election, uint64, error) {
    var count uint64
    var elections []models.Election
    election := models.Election{}

    clamor.BuildQueryWithoutPagination(p.db, query, &models.Election{}).Count(&count)
    err := clamor.BuildQuery(p.db, query, &election).Find(&elections).Error
    return elections, count, err
}
func (p *ElectionService) EnsureElectionTable() {
    p.db.AutoMigrate(&models.Election{})
    p.db.Model(&models.Election{}).AddUniqueIndex("idx_election_code", "code")
}
func (p *ElectionService) EnsureElection(election models.Election) {
    existing, err := p.GetElection(election.Code)

    if err != nil {
        p.CreateElection(election)
    } else {
        existing.Name = election.Name
        p.UpdateElection(existing)
    }
}

