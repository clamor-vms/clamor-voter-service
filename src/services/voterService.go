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

type IVoterService interface {
    CreateVoter(models.Voter) models.Voter
    UpdateVoter(models.Voter) models.Voter
    GetVoters(clamor.QueryRequest) ([]models.Voter, uint64, error)
    GetVoter(uint64) (models.Voter, error)
    GetVoterCount() uint64
    EnsureVoterTable()
    EnsureVoter(models.Voter)
}

type VoterService struct {
    db *gorm.DB
}
func NewVoterService(db *gorm.DB) *VoterService {
    return &VoterService{db: db}
}
func (p *VoterService) CreateVoter(voter models.Voter) models.Voter {
    p.db.Create(&voter)
    return voter
}
func (p *VoterService) UpdateVoter(voter models.Voter) models.Voter {
    p.db.Save(&voter)
    return voter
}
func (p *VoterService) GetVoters(query clamor.QueryRequest) ([]models.Voter, uint64, error) {
    var count uint64
    var voters []models.Voter
    voter := models.Voter{}

    clamor.BuildQueryWithoutPagination(p.db, query, &models.Voter{}).Count(&count)
    err := clamor.BuildQuery(p.db, query, &voter).Find(&voters).Error
    return voters, count, err
}
func (p *VoterService) GetVoter(voterId uint64) (models.Voter, error) {
    var voter models.Voter
    err := p.db.Where(&models.Voter{VoterId: voterId}).First(&voter).Error
    return voter, err
}
func (p *VoterService) GetVoterCount() uint64 {
    var count uint64
    p.db.Model(&models.Voter{}).Count(&count)
    return count
}
func (p *VoterService) EnsureVoterTable() {
    p.db.AutoMigrate(&models.Voter{})
    p.db.Model(&models.Voter{}).AddUniqueIndex("idx_voter_voter_id", "voter_id")
    p.db.Model(&models.Voter{}).AddForeignKey("county_code", "counties(code)", "RESTRICT", "RESTRICT")
    p.db.Model(&models.Voter{}).AddForeignKey("jurisdiction_code", "jurisdictions(code)", "RESTRICT", "RESTRICT")
    p.db.Model(&models.Voter{}).AddForeignKey("school_code", "school_districts(code)", "RESTRICT", "RESTRICT")
    p.db.Model(&models.Voter{}).AddForeignKey("village_code", "villages(code)", "RESTRICT", "RESTRICT")
}
func (p *VoterService) EnsureVoter(voter models.Voter) {
    existing, err := p.GetVoter(voter.VoterId)

    if err != nil {
        p.CreateVoter(voter)
    } else {
        voter.ID = existing.ID
        p.UpdateVoter(voter)
    }
}
