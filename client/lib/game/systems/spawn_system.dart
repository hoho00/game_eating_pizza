import 'package:flame/components.dart';
import 'package:flame/game.dart';
import '../components/monster.dart';

/// 몬스터 스폰 시스템
/// 일정 간격으로 몬스터를 생성합니다
class SpawnSystem {
  double _spawnTimer = 0.0;
  double _spawnInterval = 2.0; // 2초마다 스폰
  bool _isActive = false;
  FlameGame? gameRef;

  void setGameRef(FlameGame gameRef) {
    this.gameRef = gameRef;
  }

  void start() {
    _isActive = true;
    _spawnTimer = 0.0;
  }

  void stop() {
    _isActive = false;
  }

  void reset() {
    _spawnTimer = 0.0;
    _isActive = false;
  }

  void update(double dt) {
    if (!_isActive || gameRef == null) return;

    _spawnTimer += dt;
    
    if (_spawnTimer >= _spawnInterval) {
      _spawnTimer = 0.0;
      spawnMonster();
    }
  }

  void spawnMonster() {
    if (gameRef == null) return;
    
    final monster = Monster();
    // 화면 오른쪽 끝에서 스폰
    monster.position = Vector2(
      gameRef!.size.x,
      gameRef!.size.y / 2,
    );
    
    gameRef!.add(monster);
  }
}
