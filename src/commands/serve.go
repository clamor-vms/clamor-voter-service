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
    "net/http"

    "github.com/spf13/cobra"
    "github.com/gorilla/mux"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    skaioskit "github.com/nathanmentley/skaioskit-go-core"

    "skaioskit/core"
    "skaioskit/services"
    "skaioskit/controllers"
)

var serveCmd = &cobra.Command{
    Use:   "serve",
    Short: "runs the rest api",
    Long:  `runs the rest api that exposes the voter / voter history data for other microservices to consume.`,
    Run: func(cmd *cobra.Command, args []string) {
        //setup db connection
        conStr := skaioskit.BuildMySqlConnectionString(core.DATABASE_USER, os.Getenv("MYSQL_PASSWORD"), core.DATABASE_HOST, core.DATABASE_NAME)
        db, err := gorm.Open("mysql", conStr)
        if err != nil {
            panic(err)
        }
        defer db.Close()

        //setup services
        schoolService := services.NewSchoolDistrictService(db)
        countyService := services.NewCountyService(db)
        jurisdictionService := services.NewJurisdictionService(db)
        electionService := services.NewElectionService(db)
        villageService := services.NewVillageService(db)
        voterService := services.NewVoterService(db)
        voterHistoryService := services.NewVoterHistoryService(db)

        //build controllers
        aboutController := skaioskit.NewControllerProcessor(controllers.NewAboutController())
        countyController := skaioskit.NewControllerProcessor(controllers.NewCountyController(countyService))
        electionController := skaioskit.NewControllerProcessor(controllers.NewElectionController(electionService))
        jurisdictionController := skaioskit.NewControllerProcessor(controllers.NewJurisdictionController(jurisdictionService))
        schoolDistrictController := skaioskit.NewControllerProcessor(controllers.NewSchoolDistrictController(schoolService))
        villageController := skaioskit.NewControllerProcessor(controllers.NewVillageController(villageService))
        voterController := skaioskit.NewControllerProcessor(controllers.NewVoterController(voterService))
        voterHistoryController := skaioskit.NewControllerProcessor(controllers.NewVoterHistoryController(voterHistoryService))

        //setup routing to controllers
        r := mux.NewRouter()
        r.HandleFunc("/about", aboutController.Logic)
        r.HandleFunc("/county", countyController.Logic)
        r.HandleFunc("/election", electionController.Logic)
        r.HandleFunc("/jurisdiction", jurisdictionController.Logic)
        r.HandleFunc("/schoolDistrict", schoolDistrictController.Logic)
        r.HandleFunc("/village", villageController.Logic)
        r.HandleFunc("/voter", voterController.Logic)
        r.HandleFunc("/voterHistory", voterHistoryController.Logic)

        //wrap everything behind a jwt middleware
        jwtMiddleware := skaioskit.JWTEnforceMiddleware([]byte(os.Getenv("JWT_SECRET")))
        http.Handle("/", skaioskit.PanicHandler(jwtMiddleware(r)))

        //server up app
        if err := http.ListenAndServe(":" + core.PORT_NUMBER, nil); err != nil {
            panic(err)
        }
    },
}

//Entry
func init() {
    RootCmd.AddCommand(serveCmd)
}
