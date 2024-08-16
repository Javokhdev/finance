package kafka

import (
	"context"
	"encoding/json"
	"log"

	pb "auth/genproto/auth"
	"auth/service"
)

func UserEditProfileHandler(u *service.UserService) func(message []byte) {
	return func(message []byte) {
		var user pb.UserRes
		if err := json.Unmarshal(message, &user); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		_, err := u.EditProfile(context.Background(), &user)
		if err != nil {
			log.Printf("failed to edit user via Kafka: %v", err)
			return
		}
		log.Printf("User updated")
	}
}

func UserEditPasswordHandler(u *service.UserService) func(message []byte) {
	return func(message []byte) {
		var user pb.ChangePasswordReq
		if err := json.Unmarshal(message, &user); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		_, err := u.ChangePassword(context.Background(), &user)
		if err != nil {
			log.Printf("failed to edit password via Kafka: %v", err)
			return
		}
		log.Printf("Password updated")
	}
}

func UserEditSettingHandler(u *service.UserService) func(message []byte) {
	return func(message []byte) {
        var user pb.SettingReq
        if err := json.Unmarshal(message, &user); err != nil {
            log.Printf("Cannot unmarshal JSON: %v", err)
            return
        }

        _, err := u.EditSetting(context.Background(), &user)
        if err != nil {
            log.Printf("failed to edit user settings via Kafka: %v", err)
            return
        }
        log.Printf("User settings updated")
    }
}