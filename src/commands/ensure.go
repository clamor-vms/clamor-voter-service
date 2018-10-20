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
    "github.com/spf13/cobra"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    clamor "github.com/clamor-vms/clamor-go-core"

    "clamor/core"
    "clamor/services"
    "clamor/providers"
)

//The ensure command will load data from an IVoterDataProvider and ensure each record exists and is updated in the database.
var ensureCmd = &cobra.Command{
    Use:   "ensure",
    Short: "imports the database",
    Long:  `ensures the database schema exists and has imported the voter data.`,
    Run: func(cmd *cobra.Command, args []string) {
        //setup db connection
        conStr := clamor.BuildMySqlConnectionString(core.DATABASE_USER, os.Getenv("MYSQL_PASSWORD"), core.DATABASE_HOST, core.DATABASE_NAME)
        db, err := gorm.Open("mysql", conStr)
        if err != nil {
            panic(err)
        }
        defer db.Close()

        //setup services
        // TODO: Issue#12 For sanity we should setup a DI pattern or something to setup instances of all of this.
        //                  That would make unit testing a bit more sane and we'd be able to probably call common code
        //                  to get all these services instead of setting up each one at the start of a command.
        schoolService := services.NewSchoolDistrictService(db)
        countyService := services.NewCountyService(db)
        villageService := services.NewVillageService(db)
        jurisdictionService := services.NewJurisdictionService(db)
        electionService := services.NewElectionService(db)
        voterService := services.NewVoterService(db)
        voterHistoryService := services.NewVoterHistoryService(db)

        // TODO: Issue#8 - We should get the provider dynamically based on the args.
        //                  This will be needed to support multiple states worth of voter data.
        //                  We could also use this to support pulling data from multiple sources for the same state.
        provider := providers.NewMichiganByteWidthDataProvider()

        //ensure db

        //setup records for all the counties, jurisdictions, schools, and villages. These will be referenced by the voter records.
        countyService.EnsureCountyTable()
        for county := range provider.ParseCounties() {
            countyService.EnsureCounty(county)
        }

        jurisdictionService.EnsureJurisdictionTable()
        for jurisdiction := range provider.ParseJurisdictions() {
            jurisdictionService.EnsureJurisdiction(jurisdiction)
        }

        schoolService.EnsureSchoolDistrictTable()
        for school := range provider.ParseSchools() {
            schoolService.EnsureSchoolDistrict(school)
        }

        villageService.EnsureVillageTable()
        for village := range provider.ParseVillages() {
            villageService.EnsureVillage(village)
        }

        //ensure every election refereced by the provider has a record.
        electionService.EnsureElectionTable()
        for election := range provider.ParseElections() {
            electionService.EnsureElection(election)
        }

        //ensure we have a record for each voter and records for each election they voted in.
        voterService.EnsureVoterTable()
        for voter := range provider.ParseVoters() {
            voterService.EnsureVoter(voter)
        }
        voterHistoryService.EnsureVoterHistoryTable()
        for voterHistory := range provider.ParseVoterHistories() {
            voterHistoryService.EnsureVoterHistory(voterHistory)
        }
    },
}

//Entry
func init() {
    RootCmd.AddCommand(ensureCmd)
}
