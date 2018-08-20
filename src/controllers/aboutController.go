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
    "os"
    "net/http"

    skaioskit "github.com/nathanmentley/skaioskit-go-core"

    "skaioskit/core"
)

type AboutController struct {
}
func NewAboutController() *AboutController {
    return &AboutController{}
}
func (p *AboutController) Get(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    return skaioskit.ControllerResponse{Status: http.StatusOK, Body: GetAboutResponse{
        CoreVersion: skaioskit.VERSION,
        Version: core.SERVICE_VERSION,
        BuildTime: os.Getenv("BUILD_DATETIME"),
    }}
}
func (p *AboutController) Post(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
}
func (p *AboutController) Put(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
}
func (p *AboutController) Delete(w http.ResponseWriter, r *http.Request) skaioskit.ControllerResponse {
    return skaioskit.ControllerResponse{Status: http.StatusNotFound, Body: skaioskit.EmptyResponse{}}
}
