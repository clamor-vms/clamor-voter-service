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

type JurisdictionController struct {
    jurisdictionService services.IJurisdictionService
}
func NewJurisdictionController(jurisdictionService services.IJurisdictionService) *JurisdictionController {
    return &JurisdictionController{
        jurisdictionService: jurisdictionService,
    }
}
func (p *JurisdictionController) Get(w http.ResponseWriter, r *http.Request) clamor.ControllerResponse {
    queryStr := r.URL.Query().Get("query")
    query := clamor.QueryRequest{}
    err := json.Unmarshal([]byte(queryStr), &query)

    if err != nil {
        return clamor.ControllerResponse{Status: http.StatusBadRequest, Body: clamor.EmptyResponse{}}
    }

    jurisdictions, count, err := p.jurisdictionService.GetJurisdictions(query)

    if err == nil {
        return clamor.ControllerResponse{Status: http.StatusOK, Body: GetJurisdictionsResponse{Jurisdictions: jurisdictions, Total: count}}
    } else {
        panic(err)
    }
}
func (p *JurisdictionController) Post(w http.ResponseWriter, r *http.Request) clamor.ControllerResponse {
    return clamor.ControllerResponse{Status: http.StatusNotFound, Body: clamor.EmptyResponse{}}
}
func (p *JurisdictionController) Put(w http.ResponseWriter, r *http.Request) clamor.ControllerResponse {
    return clamor.ControllerResponse{Status: http.StatusNotFound, Body: clamor.EmptyResponse{}}
}
func (p *JurisdictionController) Delete(w http.ResponseWriter, r *http.Request) clamor.ControllerResponse {
    return clamor.ControllerResponse{Status: http.StatusNotFound, Body: clamor.EmptyResponse{}}
}
