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

package controllers

import (
    "encoding/json"
    "net/http"

    clamor "github.com/clamor-vms/clamor-go-core"

    "clamor/services"
)

type SchoolDistrictController struct {
    schoolDistrictService services.ISchoolDistrictService
}
func NewSchoolDistrictController(schoolDistrictService services.ISchoolDistrictService) *SchoolDistrictController {
    return &SchoolDistrictController{
        schoolDistrictService: schoolDistrictService,
    }
}
func (p *SchoolDistrictController) Get(w http.ResponseWriter, r *http.Request) clamor.ControllerResponse {
    queryStr := r.URL.Query().Get("query")
    query := clamor.QueryRequest{}
    err := json.Unmarshal([]byte(queryStr), &query)

    if err != nil {
        return clamor.ControllerResponse{Status: http.StatusBadRequest, Body: clamor.EmptyResponse{}}
    }

    schoolDistricts, count, err := p.schoolDistrictService.GetSchoolDistricts(query)

    if err == nil {
        return clamor.ControllerResponse{Status: http.StatusOK, Body: GetSchoolDistrictsResponse{SchoolDistricts: schoolDistricts, Total: count}}
    } else {
        panic(err)
    }
}
func (p *SchoolDistrictController) Post(w http.ResponseWriter, r *http.Request) clamor.ControllerResponse {
    return clamor.ControllerResponse{Status: http.StatusNotFound, Body: clamor.EmptyResponse{}}
}
func (p *SchoolDistrictController) Put(w http.ResponseWriter, r *http.Request) clamor.ControllerResponse {
    return clamor.ControllerResponse{Status: http.StatusNotFound, Body: clamor.EmptyResponse{}}
}
func (p *SchoolDistrictController) Delete(w http.ResponseWriter, r *http.Request) clamor.ControllerResponse {
    return clamor.ControllerResponse{Status: http.StatusNotFound, Body: clamor.EmptyResponse{}}
}
