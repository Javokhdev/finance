package storage

import pb "auth/genproto/auth"

type StorageI interface {
	Auth() AuthI
	User() UserI
}

type AuthI interface {
	Register(*pb.RegisterReq) (*pb.RegisterRes, error)
    Login(*pb.LoginReq) (*pb.User, error)
    ForgotPassword(*pb.GetByEmail) (*pb.ForgotPassRes, error)
    ResetPassword(*pb.ResetPassReq) (*pb.ResetPasswordRes, error)
    SaveRefreshToken(*pb.RefToken) (*pb.SaveRefereshTokenRes, error)
    RefreshToken(*pb.GetByEmail) (*pb.LoginRes, error)
    ChangeRole(*pb.Role) (*pb.ChangeRoleRes, error)
}

type UserI interface {
	GetProfile(*pb.GetById) (*pb.UserRes, error)
    EditProfile(*pb.UserRes) (*pb.UserRes, error)
    ChangePassword(*pb.ChangePasswordReq) (*pb.ChangePasswordRes, error)
    GetSetting(*pb.GetById) (*pb.Setting, error)
    EditSetting(*pb.SettingReq) (*pb.SettingRes, error)
    DeleteUser(*pb.GetById) (*pb.DeleteRes, error)
}