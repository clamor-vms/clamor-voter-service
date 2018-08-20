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

package models
import (
    "time"
)

type Voter struct {
    ID uint `gorm:"primary_key"`
    LastName string  `gorm:"size:35"`
    FirstName string  `gorm:"size:20"`
    MiddleName string  `gorm:"size:20"`
    NameSuffix string  `gorm:"size:3"`
    BirthYear string  `gorm:"size:4"`
    Gender string       `gorm:"size:3"`
    DateOfRegistration time.Time `gorm:"type:datetime"`
    HouseNumberCharacter string       `gorm:"size:1"`
    ResidenceStreetNumber string       `gorm:"size:7"`
    HouseSuffix string       `gorm:"size:4"`
    AddressPreDirection string       `gorm:"size:2"`
    StreetName string       `gorm:"size:30"`
    StreetType string       `gorm:"size:6"`
    SuffixDirection string       `gorm:"size:2"`
    ResidenceRxtension string       `gorm:"size:13"`
    City string       `gorm:"size:35"`
    State string       `gorm:"size:2"`
    Zip string       `gorm:"size:5"`
    MailAddress1 string       `gorm:"size:50"`
    MailAddress2 string       `gorm:"size:50"`
    MailAddress3 string       `gorm:"size:50"`
    MailAddress4 string       `gorm:"size:50"`
    MailAddress5 string       `gorm:"size:50"`
    VoterId uint64
    CountyCode uint
    JurisdictionCode uint
    Ward string `gorm:"size:6"`
    SchoolCode uint
    StateHouse uint
    StateSenate uint
    UsCongress uint
    CountyCommissioner uint
    VillageCode *uint
    VillagePrecinct string `gorm:"size:6"`
    SchoolPrecinct string `gorm:"size:6"`
    PermanentAbsenteeInd string  `gorm:"size:1"`
    StatusType string        `gorm:"size:2"`
    UOCAVAStatus string        `gorm:"size:1"`

    VoterHistories []VoterHistory   `gorm:"foreignkey:VoterId;association_foreignkey:VoterId"`
}
func GetVoterCSVHeader() []string {
    var ret []string

    return ret
}
func (v *Voter) ToSlice() []string {
    var ret []string

    return ret
}
