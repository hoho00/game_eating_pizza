/// 게임 상수 정의
class GameConstants {
  // 화면 크기
  static const double gameWidth = 800;
  static const double gameHeight = 600;

  // 플레이어 설정
  static const double playerSpeed = 100.0;
  static const int playerMaxHealth = 100;
  static const int playerStartingGold = 0;

  // 몬스터 설정
  static const double monsterSpeed = 50.0;
  static const int monsterBaseHealth = 50;
  static const int monsterBaseGoldReward = 10;
  static const int monsterBaseExpReward = 5;

  // 스폰 설정
  static const double spawnInterval = 2.0; // 초

  // 레벨업 설정
  static const int expPerLevel = 100; // 레벨당 필요한 경험치

  // 무기 설정
  static const int weaponUpgradeCostBase = 100; // 기본 강화 비용
}
