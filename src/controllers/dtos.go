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
    "skaioskit/models"
)

//TODO: These should probably be defined in the same go file as the controllers than use them.

type GetAboutResponse struct {
    CoreVersion string
    Version string
    BuildTime string
}

type GetCountiesResponse struct {
    Counties []models.County
    Total uint64
}

type GetElectionsResponse struct {
    Elections []models.Election
    Total uint64
}

type GetJurisdictionsResponse struct {
    Jurisdictions []models.Jurisdiction
    Total uint64
}

type GetSchoolDistrictsResponse struct {
    SchoolDistricts []models.SchoolDistrict
    Total uint64
}

type GetVillagesResponse struct {
    Villages []models.Village
    Total uint64
}

type GetVotersResponse struct {
    Voters []models.Voter
    Total uint64
}

type GetVoterHistoriesResponse struct {
    VoterHistories []models.VoterHistory
    Total uint64
}
