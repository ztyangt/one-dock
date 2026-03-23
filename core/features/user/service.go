package user

import (
	"context"
	"errors"
	"one-dock/app/config"
	"one-dock/core/comm"
	"one-dock/pkgs/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	UserInfo(ctx context.Context, id int64) (*userInfoResponse, error)
	Login(ctx context.Context, req *loginRequest) (resp loginResp, err error)
}

type service struct {
	comm.BaseService
	repo Repository
	cfg  *config.Cfg
}

// newService 私有构造函数
func newService(repo Repository, cfg *config.Cfg) Service {
	return &service{repo: repo, cfg: cfg}
}

// login 登录用户
func (s *service) Login(ctx context.Context, req *loginRequest) (resp loginResp, err error) {
	// 1. 判断account是否为邮箱
	conditions := make(map[string]any)
	if utils.Is.Email(req.Account) {
		conditions["email"] = req.Account
	} else {
		conditions["account"] = req.Account
	}

	// 2. 查询用户
	user, err := s.repo.GetByCondition(ctx, conditions)
	if err != nil {
		return
	}

	// 3. 检验用户状态
	if user.Status != comm.UserStatusActive {
		err = errors.New("账号已封禁，请联系管理员！")
		return
	}

	// 4. 验证密码
	if !utils.Password.Verify(user.Password, req.Password) {
		err = errors.New("密码错误！")
		return
	}

	// 5. 生成 Token
	secret := s.cfg.JWT.Secret
	if secret == "" {
		secret = utils.Random.String(32)
	}

	nowTime := time.Now()
	reqClaims := jwt.MapClaims{
		"iss": s.cfg.JWT.Issuer,
		"iat": nowTime.UnixMilli(),
		"sub": s.cfg.JWT.Subject,
		"exp": nowTime.Add(time.Second * time.Duration(s.cfg.JWT.Expire)).UnixMilli(),
		"data": map[string]any{
			"id":      user.Id,
			"email":   user.Email,
			"account": user.Account,
			"role":    user.Role,
			"status":  comm.UserStatus(user.Status),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, reqClaims).SignedString([]byte(secret))
	if err != nil {
		return
	}

	type Resp struct {
		Token string     `json:"token"`
		User  *UserModel `json:"user"`
	}

	s.ModelToDTO(&resp, &Resp{
		Token: token,
		User:  user,
	})

	return
}

// userInfo 获取用户信息
func (s *service) UserInfo(ctx context.Context, id int64) (*userInfoResponse, error) {
	u, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	dto := &userInfoResponse{}
	if err := s.ModelToDTO(dto, u); err != nil {
		return nil, err
	}
	return dto, nil
}
