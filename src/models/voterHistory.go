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
    "strconv"
)

type VoterHistory struct {
    ID uint `gorm:"primary_key"`
    VoterId uint64
    CountyCode uint
    JurisdictionCode uint
    SchoolCode uint
    ElectionCode uint64
    AbsenteeInd string  `gorm:"size:1"`
}
func GetVoterHistoryCSVHeader() []string {
    var ret []string

    ret = append(ret, "voter_id")
    ret = append(ret, "county_code")
    ret = append(ret, "jurisdiction_code")
    ret = append(ret, "school_code")
    ret = append(ret, "election_code")
    ret = append(ret, "absentee_ind")

    return ret
}
func (v *VoterHistory) ToSlice() []string {
    var ret []string

    ret = append(ret, strconv.FormatUint(v.VoterId, 10))
    ret = append(ret, strconv.FormatUint(uint64(v.CountyCode), 10))
    ret = append(ret, strconv.FormatUint(uint64(v.JurisdictionCode), 10))
    ret = append(ret, strconv.FormatUint(uint64(v.SchoolCode), 10))
    ret = append(ret, strconv.FormatUint(uint64(v.ElectionCode), 10))
    ret = append(ret, v.AbsenteeInd)

    return ret
}
