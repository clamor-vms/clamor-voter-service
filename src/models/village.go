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

    "github.com/jinzhu/gorm"
)

type Village struct {
    gorm.Model

    VillageId uint64
    Name string  `gorm:"size:255"`
    Code uint
    JurisdictionCode uint
    CountyCode uint
}
func GetVillageCSVHeader() []string {
    var ret []string

    ret = append(ret, "village_id")
    ret = append(ret, "name")
    ret = append(ret, "school_code")
    ret = append(ret, "jurisdiction_code")
    ret = append(ret, "county_code")

    return ret
}
func (v *Village) ToSlice() []string {
    var ret []string

    ret = append(ret, strconv.FormatUint(v.VillageId, 10))
    ret = append(ret, v.Name)
    ret = append(ret, strconv.FormatUint(uint64(v.Code), 10))
    ret = append(ret, strconv.FormatUint(uint64(v.JurisdictionCode), 10))
    ret = append(ret, strconv.FormatUint(uint64(v.CountyCode), 10))

    return ret
}
