package postgres

import (
	"film-adviser/settings"
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresRepo struct {
	ip      string
	port    int
	db_name string
	pass    string
	user    string
	dsn     string

	db *gorm.DB
}

// MustInit implements repository.Repository.

func New() *PostgresRepo {
	return &PostgresRepo{}
}

func (pg_rep *PostgresRepo) MustInit() {
	pg_rep.ip = settings.GetSettings().PgIp
	pg_rep.user = settings.GetSettings().PgUser
	pg_rep.db_name = settings.GetSettings().PgDb
	pg_rep.pass = settings.GetSettings().PgPass
	pg_rep.port = settings.GetSettings().PgPort

	pg_rep.formdsn()

	var err error
	pg_rep.db, err = gorm.Open(postgres.Open(pg_rep.dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect to postgres !")
	}
	pg_rep.db.AutoMigrate(&UsersFilms{})
}

func (pg_rep PostgresRepo) Write(chatid int64, film string) error {
	new_film := UsersFilms{Chatid: chatid, Film: film}
	err := pg_rep.db.Create(&new_film).Error
	if err != nil {
		fmt.Println("Error while writing new film from postgres !")
	}
	return err
}

func (pg_rep PostgresRepo) PickRandom(chatid int64) (error, string) {
	var films []UsersFilms
	if err := pg_rep.db.Where("Chatid = ?", chatid).Find(&films).Error; err != nil {
		fmt.Println("Error while reading films from postgres !")
		return err, ""
	}
	ind := rand.Intn(len(films))

	return nil, films[ind].Film
}

func (pg_rep *PostgresRepo) formdsn() {

	pg_rep.dsn = fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=%s", pg_rep.ip, pg_rep.user, pg_rep.db_name, pg_rep.pass, strconv.Itoa(pg_rep.port), "disable")
}
