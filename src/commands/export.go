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

package commands

import (
    "os"
    "encoding/csv"

    "github.com/spf13/cobra"
    "github.com/jinzhu/copier"

    "skaioskit/models"
    "skaioskit/providers"
)

var exportCmd = &cobra.Command{
    Use:   "export",
    Short: "Generates csv files from data provider",
    Long:  `Generates csv files from the configured data provider / data set.`,
    Run: func(cmd *cobra.Command, args []string) {
        provider := providers.NewMichiganByteWidthDataProvider()

        writeCounties(provider)
        writeElections(provider)
        writeJurisdictions(provider)
        writeSchoolDistricts(provider)
        writeVillages(provider)
        writeVoters(provider)
        writeVoterHistories(provider)
    },
}

func writeCounties(provider providers.IVoterDataProvider) {
    chnl := make(chan models.IExportable)
    go func() {
        for county := range provider.ParseCounties() {
            obj := models.County{}
            copier.Copy(&obj, &county)
            chnl <- &obj
        }
        close(chnl)
    }()
    writeCsv("/working/export/counties.csv", models.GetCountyCSVHeader(), chnl)
}

func writeElections(provider providers.IVoterDataProvider) {
    chnl := make(chan models.IExportable)
    go func() {
        for election := range provider.ParseElections() {
            obj := models.Election{}
            copier.Copy(&obj, &election)
            chnl <- &obj
        }
        close(chnl)
    }()
    writeCsv("/working/export/elections.csv", models.GetElectionCSVHeader(), chnl)
}

func writeJurisdictions(provider providers.IVoterDataProvider) {
    chnl := make(chan models.IExportable)
    go func() {
        for jurisdiction := range provider.ParseJurisdictions() {
            obj := models.Jurisdiction{}
            copier.Copy(&obj, &jurisdiction)
            chnl <- &obj
        }
        close(chnl)
    }()
    writeCsv("/working/export/jurisdictions.csv", models.GetJurisdictionCSVHeader(), chnl)
}

func writeSchoolDistricts(provider providers.IVoterDataProvider) {
    chnl := make(chan models.IExportable)
    go func() {
        for schoolDistrict := range provider.ParseSchools() {
            obj := models.SchoolDistrict{}
            copier.Copy(&obj, &schoolDistrict)
            chnl <- &obj
        }
        close(chnl)
    }()
    writeCsv("/working/export/schoolDistricts.csv", models.GetSchoolDistrictCSVHeader(), chnl)
}

func writeVillages(provider providers.IVoterDataProvider) {
    chnl := make(chan models.IExportable)
    go func() {
        for village := range provider.ParseVillages() {
            obj := models.Village{}
            copier.Copy(&obj, &village)
            chnl <- &obj
        }
        close(chnl)
    }()
    writeCsv("/working/export/villages.csv", models.GetVillageCSVHeader(), chnl)
}

func writeVoters(provider providers.IVoterDataProvider) {
    chnl := make(chan models.IExportable)
    go func() {
        for voter := range provider.ParseVoters() {
            obj := models.Voter{}
            copier.Copy(&obj, &voter)
            chnl <- &obj
        }
        close(chnl)
    }()
    writeCsv("/working/export/voters.csv", models.GetVoterCSVHeader(), chnl)
}

func writeVoterHistories(provider providers.IVoterDataProvider) {
    chnl := make(chan models.IExportable)
    go func() {
        for voterHistory := range provider.ParseVoterHistories() {
            obj := models.VoterHistory{}
            copier.Copy(&obj, &voterHistory)
            chnl <- &obj
        }
        close(chnl)
    }()
    writeCsv("/working/export/voterHistories.csv", models.GetVoterHistoryCSVHeader(), chnl)
}

func writeCsv(filename string, header []string, chnl <-chan models.IExportable) {
    file, err := os.Create(filename)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    w := csv.NewWriter(file)
    w.Write(header)
    for record := range chnl {
        if err := w.Write(record.ToSlice()); err != nil {
            panic(err)
        }
    }
    w.Flush()
}

//Entry
func init() {
    RootCmd.AddCommand(exportCmd)
}
