package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"prodcat/config"
	"prodcat/ent"
	"prodcat/ent/schema"
	"prodcat/ent/user"
	"prodcat/server"

	"golang.org/x/crypto/bcrypt"

	"github.com/fsnotify/fsnotify"
	_ "github.com/lib/pq"
)

func main() {

	log.Println("Starting Prodcat App")

	log.Println("Initializing configuration")
	configv := config.InitConfig()

	configv.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	tz := configv.GetString("app.TZ")
	loc, err := time.LoadLocation(tz)
	if err != nil {
		log.Fatalf("Timezone %s is invalid", tz)
	}
	time.Local = loc
	log.Printf("Setting timezone to %s", tz)
	os.Setenv("TZ", tz)

	log.Println("Initializing database")
	client := server.InitDatabase(configv)

	if config.State == "DEV" {

		log.Println("CREATE ADMIN ACCOUNT")
		InitAdminAccount(client)
	}

	log.Println("Initializing HTTP server")
	httpServer := server.InitHttpServer(configv, client)
	configv.WatchConfig()
	httpServer.Start()
}

func InitAdminAccount(db *ent.Client) {
	ctx := context.Background()
	_, err := db.User.
		Query().
		Where(user.LoginEQ("superadmin")).
		Only(ctx)
	if err != nil {
		// log.Panic(fmt.Errorf("failed querying user: %w", err))
		hashedPass, _ := HashPassword("123456")
		role := "admin"
		settings := db.UserSettings.Create().SetFrontend(schema.DefaultSetting{ProductView: "cards"}).SaveX(ctx)
		_, err := db.User.Create().
			SetFirstName("John").
			SetLastName("Doe").
			SetRole(role).
			SetEmail("test@local.lan").
			SetLogin("superadmin").
			SetPassword(string(hashedPass)).
			SetSettings(settings).
			Save(ctx)
		if err != nil {
			log.Panic(fmt.Errorf("admin not created: %w", err))
		}
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
