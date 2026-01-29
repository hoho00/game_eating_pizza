import 'package:flutter/material.dart';

/// 체력 바 위젯
class HealthBar extends StatelessWidget {
  final int currentHealth;
  final int maxHealth;
  final double width;
  final double height;

  const HealthBar({
    super.key,
    required this.currentHealth,
    required this.maxHealth,
    this.width = 200,
    this.height = 20,
  });

  @override
  Widget build(BuildContext context) {
    final percentage = maxHealth > 0 ? (currentHealth / maxHealth).clamp(0.0, 1.0) : 0.0;
    final color = percentage > 0.5
        ? Colors.green
        : percentage > 0.25
            ? Colors.orange
            : Colors.red;

    return Container(
      width: width,
      height: height,
      decoration: BoxDecoration(
        border: Border.all(color: Colors.white, width: 2),
        borderRadius: BorderRadius.circular(10),
        color: Colors.black54,
      ),
      child: Stack(
        children: [
          FractionallySizedBox(
            widthFactor: percentage,
            child: Container(
              decoration: BoxDecoration(
                color: color,
                borderRadius: BorderRadius.circular(8),
              ),
            ),
          ),
          Center(
            child: Text(
              '$currentHealth / $maxHealth',
              style: const TextStyle(
                color: Colors.white,
                fontSize: 12,
                fontWeight: FontWeight.bold,
              ),
            ),
          ),
        ],
      ),
    );
  }
}
