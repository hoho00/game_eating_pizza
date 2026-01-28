package services

import (
	"errors"
	"game_eating_pizza/internal/models"
	"game_eating_pizza/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

// AuthService는 인증 관련 비즈니스 로직을 담당합니다
type AuthService struct {
	playerRepo repository.PlayerRepositoryInterface
}

// NewAuthService는 새로운 AuthService 인스턴스를 생성합니다
func NewAuthService(playerRepo repository.PlayerRepositoryInterface) *AuthService {
	return &AuthService{
		playerRepo: playerRepo,
	}
}

// Register는 새로운 플레이어를 등록합니다
func (s *AuthService) Register(username, password string) (*models.Player, error) {
	// 사용자명 중복 확인
	exists, err := s.playerRepo.ExistsByUsername(username)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("username already exists")
	}

	// 비밀번호 해시
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 플레이어 생성
	player := &models.Player{
		Username: username,
		Password: string(hashedPassword),
		Level:    1,
		Gold:     0,
	}

	if err := s.playerRepo.Create(player); err != nil {
		return nil, err
	}

	return player, nil
}

// Login은 플레이어 로그인을 처리합니다
func (s *AuthService) Login(username, password string) (string, *models.Player, error) {
	// 플레이어 조회
	player, err := s.playerRepo.FindByUsername(username)
	if err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	// 비밀번호 확인
	err = bcrypt.CompareHashAndPassword([]byte(player.Password), []byte(password))
	if err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	// TODO: JWT 토큰 생성
	// 임시로 사용자 ID를 토큰으로 사용 (실제로는 JWT 생성 필요)
	token := "temp_token_" + string(rune(player.ID))

	return token, player, nil
}
