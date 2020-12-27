package app

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/DanielTitkov/teams-bot-example/internal/domain"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func (a *App) CreateUser(u *domain.User) error {

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	u.PasswordHash = string(hash)
	_, err = a.repo.CreateUser(u)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) GetUser(u *domain.User) (*domain.User, error) {
	user, err := a.repo.GetUserByUsername(u.Username)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (a *App) ValidateUserPassword(u *domain.User) (bool, error) {
	user, err := a.repo.GetUserByUsername(u.Username)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(u.Password))
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (a *App) GetUserToken(u *domain.User) (string, error) {
	user, err := a.repo.GetUserByUsername(u.Username)
	if err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(a.cfg.Auth.Exp)).Unix()

	t, err := token.SignedString([]byte(a.cfg.Auth.Secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func generateUserLogin(name string) string {
	reg := regexp.MustCompile(`[^\w]+`)
	login := reg.ReplaceAllString(name, "")
	return strings.ToLower(login) + fmt.Sprint(time.Now().Nanosecond())
}

func (a *App) GetOrCreateTeamsUser(turn domain.Turn) (*domain.User, error) {
	user, err := a.repo.GetUserByTeamsID(*turn.User.Meta.Teams.ID)
	if err != nil {
		password := "123" // FIXME
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			return nil, err
		}

		user, err = a.repo.CreateUser(&domain.User{
			DisplayName:  *turn.User.Meta.Teams.Username,
			Username:     generateUserLogin(*turn.User.Meta.Teams.Username),
			PasswordHash: string(hash),
			Meta: domain.UserMeta{
				Teams: domain.UserMessagerData{
					ID: turn.User.Meta.Teams.ID,
				},
			},
		})
		if err != nil {
			return nil, err
		}
		a.logger.Info("user created", fmt.Sprint(user))

		err = a.SendTeamsProactive(&domain.Turn{ // FIXME
			Message: domain.Message{
				Text:      buildUserCreatedMessage(user.DisplayName, user.Username),
				System:    TeamsSystemCode, // TODO: maybe put to turn
				Direction: OutputMessageCode,
				Proactive: true,
			},
			Dialog: domain.TurnDialog{
				Meta: domain.DialogMeta{
					Teams: turn.Dialog.Meta.Teams,
				},
			},
		})
		if err != nil {
			a.logger.Error("failed to send user created notification", err)
		} else {
			a.logger.Info("user created notification sent", fmt.Sprint())
		}

	} else {
		a.logger.Info("user fetched", fmt.Sprint(user))
	}
	return user, nil
}
