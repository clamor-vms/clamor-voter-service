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

type Jurisdiction struct {
    gorm.Model

    Name string  `gorm:"size:255"`
    Code uint
    CountyCode uint
}
func GetJurisdictionCSVHeader() []string {
    var ret []string

    ret = append(ret, "name")
    ret = append(ret, "jurisdiction_code")
    ret = append(ret, "county_code")

    return ret
}
func (j *Jurisdiction) ToSlice() []string {
    var ret []string

    ret = append(ret, j.Name)
    ret = append(ret, strconv.FormatUint(uint64(j.Code), 10))
    ret = append(ret, strconv.FormatUint(uint64(j.CountyCode), 10))

    return ret
}
