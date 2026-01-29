import 'package:flutter/material.dart';
import 'package:flame/game.dart';
import '../game/game.dart';

/// 게임 화면
/// Flame 게임을 표시하는 화면입니다
class GameScreen extends StatefulWidget {
  const GameScreen({super.key});

  @override
  State<GameScreen> createState() => _GameScreenState();
}

class _GameScreenState extends State<GameScreen> {
  late GameEatingPizza game;

  @override
  void initState() {
    super.initState();
    game = GameEatingPizza();
  }

  @override
  void dispose() {
    game.pauseGame();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Stack(
        children: [
          // Flame 게임
          GameWidget<GameEatingPizza>.controlled(
            gameFactory: () => game,
          ),
          // HUD 오버레이 (나중에 구현)
          Positioned(
            top: 40,
            left: 20,
            right: 20,
            child: Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                // 골드 표시 (나중에 구현)
                Container(
                  padding: const EdgeInsets.symmetric(
                    horizontal: 16,
                    vertical: 8,
                  ),
                  decoration: BoxDecoration(
                    color: Colors.black54,
                    borderRadius: BorderRadius.circular(20),
                  ),
                  child: const Text(
                    '골드: 0',
                    style: TextStyle(
                      color: Colors.white,
                      fontSize: 16,
                      fontWeight: FontWeight.bold,
                    ),
                  ),
                ),
                // 일시정지 버튼
                IconButton(
                  icon: const Icon(Icons.pause, color: Colors.white),
                  onPressed: () {
                    game.pauseGame();
                    _showPauseDialog();
                  },
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }

  void _showPauseDialog() {
    showDialog(
      context: context,
      barrierDismissible: false,
      builder: (context) => AlertDialog(
        title: const Text('일시정지'),
        content: const Text('게임이 일시정지되었습니다.'),
        actions: [
          TextButton(
            onPressed: () {
              Navigator.pop(context);
              game.resumeGame();
            },
            child: const Text('계속하기'),
          ),
          TextButton(
            onPressed: () {
              Navigator.pop(context);
              Navigator.pop(context);
            },
            child: const Text('나가기'),
          ),
        ],
      ),
    );
  }
}
