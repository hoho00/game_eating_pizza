/// 플레이어 데이터 모델
class PlayerModel {
  final int id;
  final String username;
  int level;
  int experience;
  int gold;
  double maxDistance;
  int totalKills;

  PlayerModel({
    required this.id,
    required this.username,
    this.level = 1,
    this.experience = 0,
    this.gold = 0,
    this.maxDistance = 0.0,
    this.totalKills = 0,
  });

  factory PlayerModel.fromJson(Map<String, dynamic> json) {
    return PlayerModel(
      id: json['id'] as int,
      username: json['username'] as String,
      level: json['level'] as int? ?? 1,
      experience: json['experience'] as int? ?? 0,
      gold: json['gold'] as int? ?? 0,
      maxDistance: (json['max_distance'] as num?)?.toDouble() ?? 0.0,
      totalKills: json['total_kills'] as int? ?? 0,
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'username': username,
      'level': level,
      'experience': experience,
      'gold': gold,
      'max_distance': maxDistance,
      'total_kills': totalKills,
    };
  }
}
