import 'package:flame/components.dart';
import 'package:flame/sprite.dart';

/// 플레이어 컴포넌트
/// 게임 내 플레이어 캐릭터를 나타냅니다
class Player extends SpriteComponent with HasGameRef {
  double speed = 100.0; // 초당 이동 속도 (픽셀)
  int health = 100;
  int maxHealth = 100;
  int gold = 0;
  int level = 1;
  int experience = 0;

  @override
  Future<void> onLoad() async {
    super.onLoad();
    
    // TODO: 플레이어 스프라이트 로드
    // sprite = await gameRef.loadSprite('player.png');
    // size = Vector2(64, 64);
    
    // 임시로 사각형으로 표시
    position = Vector2(100, gameRef.size.y / 2);
    size = Vector2(64, 64);
  }

  @override
  void update(double dt) {
    super.update(dt);
    
    // 자동으로 오른쪽으로 이동
    position.x += speed * dt;
  }

  void takeDamage(int damage) {
    health -= damage;
    if (health <= 0) {
      health = 0;
      // 게임 오버 처리
    }
  }

  void addGold(int amount) {
    gold += amount;
  }

  void addExperience(int amount) {
    experience += amount;
    // 레벨업 체크 (레벨당 100 경험치 필요)
    while (experience >= level * 100) {
      experience -= level * 100;
      level++;
      maxHealth += 10;
      health = maxHealth; // 레벨업 시 체력 회복
    }
  }
}
