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

    skaioskit "github.com/nathanmentley/skaioskit-go-core"

    "skaioskit/services"
)

type VillageController struct {
    villageService services.IVillageService
}
func NewVillageController(villageService services.IVillageService) *VillageController {
    return &VillageController{
        villageService: villageService,
    }
}
func (p *VillageController) Get(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    queryStr := r.URL.Query().Get("query")
    query := skaioskit.QueryRequest{}
    err := json.Unmarshal([]byte(queryStr), &query)

    if err != nil {
        return skaioskit.ControllerResponse{Status: http.StatusBadRequest, Body: skaioskit.EmptyResponse{}}
    }

    villages, count, err := p.villageService.GetVillages(query)

    if err == nil {
        return skaioskit.ControllerResponse{Status: http.StatusOK, Body: GetVillagesResponse{Villages: villages, Total: count}}
    } else {
        panic(err)
    }
}
func (p *VillageController) Post(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
}
func (p *VillageController) Put(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
}
func (p *VillageController) Delete(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
}
