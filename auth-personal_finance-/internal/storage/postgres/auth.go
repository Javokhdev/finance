package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	pb "auth/genproto/auth"
	"golang.org/x/crypto/bcrypt"
	"auth/api/token"
)

type AuthRepo struct {
	db 		 *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (r *AuthRepo) Register(req *pb.RegisterReq) (*pb.RegisterRes, error) {
	res := &pb.RegisterRes{Message: "User registered successfully"}

	tr, err := r.db.Begin()
	if err!= nil {
        return nil, err
    }

	var id string
	query := `INSERT INTO users (username, email, password, full_name, date_of_birth) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err = tr.QueryRow(query, req.Username, req.Email, req.Password, req.FullName, req.DateOfBirth).Scan(&id)
	if err != nil {
		tr.Rollback()
		return nil, err
	}

	query = `INSERT INTO settings (user_id) VALUES ($1)`
	_, err = tr.Exec(query, id)
	if err!= nil {
        tr.Rollback()
        return nil, err
    }

	err = tr.Commit()
	if err!= nil {
		tr.Rollback()
        return nil, err
    }
	
	return res, nil
}

func (r *AuthRepo) Login(req *pb.LoginReq) (*pb.User, error) {
	res := &pb.User{}

	var passwordHash string
	query := `SELECT id, username, email, role, password FROM users WHERE username = $1`
	err := r.db.QueryRow(query, req.Username).Scan(
		&res.Id,
		&res.Username,
		&res.Email,
		&res.Role,
		&passwordHash,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	} else if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("invalid password for username: %s", req.Username)
	}

	return res, nil
}
func (r *AuthRepo) ForgotPassword(req *pb.GetByEmail) (*pb.ForgotPassRes, error) {
	res := &pb.ForgotPassRes{Message: "Email sent"}

	query := `SELECT email FROM users WHERE email = $1`

	var email string
	err := r.db.QueryRow(query, req.Email).Scan(&email)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%s Email not found", req.Email)
		}
		return nil, err
	}

	return res, nil
}

func (r *AuthRepo) ResetPassword(req *pb.ResetPassReq) (*pb.ResetPasswordRes, error) {
	res := &pb.ResetPasswordRes{Message: "Password reset successfully"}

	query := `UPDATE users SET password = $1, updated_at=now() WHERE email = $2`

	_, err := r.db.Exec(query, req.NewPassword, req.Email)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *AuthRepo) SaveRefreshToken(req *pb.RefToken) (*pb.SaveRefereshTokenRes, error) {
	res := &pb.SaveRefereshTokenRes{Message: "Token saved successfully"}

    query := `INSERT INTO tokens (user_id, token) VALUES ($1, $2)`

    _, err := r.db.Exec(query, req.UserId, req.Token)
    if err != nil {
        return nil, err
    }

    return res, nil
}

func (r *AuthRepo) RefreshToken(req *pb.GetByEmail) (*pb.LoginRes, error) {
	res := &pb.LoginRes{}

	query := `SELECT token FROM tokens WHERE user_id = (SELECT id FROM users WHERE email = $1)`
	var tokenString string
	err := r.db.QueryRow(query, req.Email).Scan(&tokenString)
	if err != nil {
		if err == sql.ErrNoRows {
            return nil, fmt.Errorf("token not found for user with email: %s", req.Email)
        }
		return nil, err
	}

    claims, err :=token.ExtractClaim(tokenString)
	if err!= nil {
        return nil, err
    }

	id := claims["user_id"].(string)
	username, _ := claims["username"].(string)
	email, _ := claims["email"].(string)
	role, _ := claims["role"].(string)

	res, _ = token.GenerateJWTToken(&pb.User{
		Id:       id,
        Username: username,
        Email:    email,
        Role:     role,
	})

    return res, nil
}

func (r *AuthRepo) ChangeRole(req *pb.Role) (*pb.ChangeRoleRes, error) {
	res := &pb.ChangeRoleRes{Message: "Role changed successfully"}

    query := `UPDATE users SET role = $1 WHERE id = $2 AND deleted_at = 0`

    _, err := r.db.Exec(query, req.Role, req.Id)
    if err!= nil {
        return nil, err
    }

    return res, nil
}