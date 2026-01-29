import 'package:flame/game.dart';
import '../components/player.dart';
import '../systems/spawn_system.dart';
import '../systems/combat_system.dart';

/// 게임 상태를 관리하는 매니저
class GameManager {
  late FlameGame gameRef;
  Player? player;
  late SpawnSystem spawnSystem;
  late CombatSystem combatSystem;
  
  bool _isPaused = false;
  bool _isGameOver = false;

  /// 게임 초기화
  Future<void> initialize(FlameGame game) async {
    gameRef = game;
    
    // 시스템 초기화
    spawnSystem = SpawnSystem();
    spawnSystem.setGameRef(game);
    combatSystem = CombatSystem();
    
    // 플레이어 생성 (나중에 구현)
    // player = Player();
    // gameRef.add(player);
  }

  /// 게임 시작
  void startGame() {
    _isPaused = false;
    _isGameOver = false;
    spawnSystem.start();
  }

  /// 게임 업데이트
  void update(double dt) {
    if (_isPaused || _isGameOver) return;
    
    spawnSystem.update(dt);
    combatSystem.update(dt);
  }

  /// 게임 일시정지
  void pauseGame() {
    _isPaused = true;
  }

  /// 게임 재개
  void resumeGame() {
    _isPaused = false;
  }

  /// 게임 리셋
  void resetGame() {
    _isGameOver = false;
    _isPaused = false;
    spawnSystem.reset();
    combatSystem.reset();
  }

  /// 게임 오버
  void gameOver() {
    _isGameOver = true;
    spawnSystem.stop();
  }

  bool get isPaused => _isPaused;
  bool get isGameOver => _isGameOver;
}
