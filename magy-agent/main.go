package main

import (
	"flag"
	"fmt"
	"github.com/labstack/echo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"magyAgent/conf"
	"magyAgent/controller/middleware"
	"magyAgent/controller/router"
	"magyAgent/domain/model"
	"magyAgent/interactor"
)

var Db *gorm.DB
var runServer = flag.Bool("server", false, "production is -server option require")

func main() {
	conf.NewConfig(*runServer)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Current.Database.Host, conf.Current.Database.Port, conf.Current.Database.User, conf.Current.Database.Password, conf.Current.Database.Database)

	Db, _ = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	// Auto Migrate
	Db.AutoMigrate(&model.User{}, &model.AccountActivation{}, &model.Permission{}, &model.Role{}, &model.LoginEvent{}, &model.AccountResetPassword{}, &model.Product{}, &model.Order{}, &model.OrderItem{})
	initialInsert()
	e := echo.New()

	i := interactor.NewInteractor(Db)
	h := i.NewAppHandler()

	router.NewRouter(e, h)
	middleware.NewMiddleware(e)

	e.Logger.Fatal(e.StartTLS(":" + conf.Current.Server.Port, "certificate.pem", "certificate-key.pem"))

}

func initialInsert() {
	role1 := &model.Role{ Id: "0a98c96e-4474-4f82-a1b2-e5ea92c4d392", Name: "admin"}
	role2 := &model.Role{ Id: "7a753a24-5a20-4021-a3e0-0afdf3744675", Name: "user"}
	role3 := &model.Role{ Id: "53a23483-fa51-4f83-ba53-6979d87c77af", Name: "agent"}

	permission1 := &model.Permission{ Id : "2d166d22-3bce-433d-b0a9-0e9786715498", Name: "execute_agent_check", Roles: []model.Role{*role3}}
	permission2 := &model.Permission{ Id : "a4f07b70-00dd-4c8c-8e7f-07f9c08d9796", Name: "execute_admin_check", Roles: []model.Role{*role1}}
	permission3 := &model.Permission{ Id : "28ecaf42-bd26-4fa6-94e3-22cedefd2237", Name: "execute_admin_agent_check", Roles: []model.Role{*role3, *role1}}

	adminPassword, _ := model.HashAndSaltPasswordIfStrongAndMatching("Adminadmin1*", "Adminadmin1*")
	agentPassword, _ := model.HashAndSaltPasswordIfStrongAndMatching("Agentagent1*", "Agentagent1*")
	admin := &model.User{Id: "c8381fb9-6e7a-4f4c-88c0-cafdfa11a4d7", Name: "Admin", Surname: "Adminic", Email: "admin@admin.com", Active: true, Password: adminPassword, Roles: []model.Role{*role1}}
	agent := &model.User{Id: "1c033f28-af44-4394-be2c-71d0c93890e1", Name: "Agent", Surname: "Agentic", Email: "agent@agent.com", Active: true, Password: agentPassword, Roles: []model.Role{*role3}}

	Db.Create(role1)
	Db.Create(role2)
	Db.Create(role3)

	Db.Create(permission1)
	Db.Create(permission2)
	Db.Create(permission3)

	Db.Create(admin)
	Db.Create(agent)
}
