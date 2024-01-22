package repositories

import (
	"context"
	"errors"
	"fmt"
	"log"
	"prodcat/dto"
	"prodcat/ent"
	"prodcat/ent/predicate"
	"prodcat/ent/schema"
	"prodcat/ent/user"
	"prodcat/services"

	"entgo.io/ent/dialect/sql"
	"github.com/gin-gonic/gin"
)

type UserRepository struct {
	db          *ent.Client
	authService *services.AuthService
}

func NewUserRepository(db *ent.Client, au *services.AuthService) *UserRepository {
	return &UserRepository{
		db:          db,
		authService: au,
	}
}

func (ur *UserRepository) CreateUser(ctx *gin.Context, dto dto.UserDTO) (*ent.User, error) {

	_, err := ur.db.User.
		Query().
		Where(user.LoginEQ(dto.Username)).
		Only(ctx)
	if !ent.IsNotFound(err) {
		return nil, errors.New("user already exists")
	}

	hashedPass, err := ur.authService.HashPassword(dto.Password)
	if err != nil {
		panic(err)
	}

	role := "user"
	settings := ur.db.UserSettings.Create().SaveX(ctx)
	user, err := ur.db.User.Create().
		SetFirstName(dto.FirstName).
		SetLastName(dto.LastName).
		SetEmail(dto.Email).
		SetRole(role).
		SetLogin(dto.Username).
		SetPassword(string(hashedPass)).
		SetSettings(settings).
		Save(ctx)
	if err != nil {
		log.Panic(err)
	}

	log.Println("user returned: ", user)
	return user, nil
}

func (ur *UserRepository) FindUserByLogin(ctx *gin.Context, username string) (u *ent.User, err error) {
	foundUser, err := ur.db.User.
		Query().
		Where(user.LoginEQ(username)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	log.Println("user returned: ", u)
	return foundUser, nil
}

func (ur *UserRepository) GetUsers(c *gin.Context, offset int, limit int) (u []*ent.User, err error) {
	model := ur.db.User
	users, err := model.Query().
		Select(user.FieldID,
			user.FieldFirstName,
			user.FieldLastName,
			user.FieldPhonenumber,
			user.FieldEmail,
			user.FieldCreatedAt,
			user.FieldUpdatedAt,
			user.FieldDeletedAt,
			user.FieldRole,
		).
		Limit(limit).
		Order(user.ByID(sql.OrderDesc())).
		All(c)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) GetUser(c context.Context, id int) (*ent.User, error) {
	u, err := ur.db.User.
		Query().
		Select(user.FieldID,
			user.FieldFirstName,
			user.FieldLastName,
			user.FieldPhonenumber,
			user.FieldEmail,
			user.FieldCreatedAt,
			user.FieldUpdatedAt,
			user.FieldDeletedAt,
			user.FieldRole,
		).
		Where(user.ID(id)).
		Only(c)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	return u, nil
}

func (ur *UserRepository) SetUserSettings(c context.Context, id int) {
	ur.db.UserSettings.Update().SetFrontend(schema.DefaultSetting{ProductView: "biba"}).Where(predicate.UserSettings(user.IDEQ(id)))
}
