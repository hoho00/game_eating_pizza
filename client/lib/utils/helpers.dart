/// 유틸리티 함수들
class Helpers {
  /// 숫자를 천 단위로 포맷팅 (예: 1000 -> "1,000")
  static String formatNumber(int number) {
    return number.toString().replaceAllMapped(
      RegExp(r'(\d{1,3})(?=(\d{3})+(?!\d))'),
      (Match m) => '${m[1]},',
    );
  }

  /// 경험치를 퍼센트로 변환
  static double getExperiencePercentage(int currentExp, int level) {
    final requiredExp = level * 100;
    if (requiredExp == 0) return 0.0;
    return (currentExp / requiredExp).clamp(0.0, 1.0);
  }

  /// 체력을 퍼센트로 변환
  static double getHealthPercentage(int currentHealth, int maxHealth) {
    if (maxHealth == 0) return 0.0;
    return (currentHealth / maxHealth).clamp(0.0, 1.0);
  }
}
