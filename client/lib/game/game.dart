import 'package:flame/game.dart';
import 'package:flutter/material.dart';
import 'managers/game_manager.dart';

/// 메인 게임 클래스
/// Flame의 Game 클래스를 상속하여 게임 로직을 관리합니다
class GameEatingPizza extends FlameGame {
  late GameManager gameManager;

  @override
  Future<void> onLoad() async {
    super.onLoad();
    
    // 게임 매니저 초기화
    gameManager = GameManager();
    await gameManager.initialize(this);
    
    // 게임 시작
    gameManager.startGame();
  }

  @override
  void update(double dt) {
    super.update(dt);
    gameManager.update(dt);
  }

  @override
  void render(Canvas canvas) {
    super.render(canvas);
    // 게임 렌더링은 컴포넌트들이 자동으로 처리
  }

  void pauseGame() {
    pauseEngine();
    gameManager.pauseGame();
  }

  void resumeGame() {
    resumeEngine();
    gameManager.resumeGame();
  }

  void resetGame() {
    gameManager.resetGame();
  }
}
