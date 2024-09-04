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
	_ = pg_rep.db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s;", "users_ids"))
	_ = pg_rep.db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s;", "users_films"))

	pg_rep.db.AutoMigrate(&UsersFilms{})
	pg_rep.db.AutoMigrate(&UsersId{})

}

func (pg_rep PostgresRepo) Write(userid int64, film string) error {
	new_film := UsersFilms{UserId: userid, Film: film}
	err := pg_rep.db.Create(&new_film).Error
	if err != nil {
		fmt.Println("Error while writing new film from postgres !")
	}
	return err
}

func (pg_rep PostgresRepo) PickRandom(userid int64) (error, string) {
	var films []UsersFilms
	if err := pg_rep.db.Where("UserId = ?", userid).Find(&films).Error; err != nil {
		fmt.Println("Error while reading films from postgres !")
		return err, ""
	}
	ind := rand.Intn(len(films))

	return nil, films[ind].Film
}

func (pg_rep *PostgresRepo) formdsn() {

	pg_rep.dsn = fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=%s", pg_rep.ip, pg_rep.user, pg_rep.db_name, pg_rep.pass, strconv.Itoa(pg_rep.port), "disable")
}

func (pg_rep PostgresRepo) AddChatid(receiver_id, sender_id, user_id int64) error {

	var user UsersId
	if err := pg_rep.db.Where("UserId = ?", user_id).Find(&user).Error; err != nil {
		fmt.Println("Error while reading user from postgres !")
		return err
	}
	if user.ID == 0 {
		fmt.Println("No user found with the specified ID.")
		new_user := UsersId{UserId: user_id, SenderId: sender_id, ReceiverId: receiver_id}
		err := pg_rep.db.Create(&new_user).Error
		if err != nil {
			fmt.Println("Error while writing new film from postgres !")
		}
		return err
	}
	if receiver_id != 0 {
		user.ReceiverId = receiver_id
	}
	if sender_id != 0 {
		user.SenderId = sender_id
	}

	if err := pg_rep.db.Save(&user).Error; err != nil {
		fmt.Println("Error while updating Id in Postgres:", err)
		return err
	}

	return nil
}
func (pg_rep PostgresRepo) GetUserChat(user_id int64) (error, int64, int64) {
	var user UsersId
	if err := pg_rep.db.Where("UserId = ?", user_id).Find(&user).Error; err != nil {
		fmt.Println("Error while reading user from postgres !")
		return err, 0, 0
	}
	return nil, user.ReceiverId, user.SenderId
}
