import 'package:flame/components.dart';
import 'package:flame/sprite.dart';

/// 몬스터 컴포넌트
/// 게임 내 몬스터를 나타냅니다
class Monster extends SpriteComponent with HasGameRef {
  double speed = 50.0; // 초당 이동 속도
  int health = 50;
  int maxHealth = 50;
  int attackPower = 10;
  int goldReward = 10;
  int experienceReward = 5;

  @override
  Future<void> onLoad() async {
    super.onLoad();
    
    // TODO: 몬스터 스프라이트 로드
    // sprite = await gameRef.loadSprite('monster.png');
    // size = Vector2(48, 48);
    
    // 임시로 사각형으로 표시
    size = Vector2(48, 48);
  }

  @override
  void update(double dt) {
    super.update(dt);
    
    // 왼쪽으로 이동 (플레이어를 향해)
    position.x -= speed * dt;
    
    // 화면 밖으로 나가면 제거
    if (position.x + size.x < 0) {
      removeFromParent();
    }
  }

  void takeDamage(int damage) {
    health -= damage;
    if (health <= 0) {
      // 몬스터 처치
      removeFromParent();
    }
  }

  bool get isDead => health <= 0;
}
