/// 무기 데이터 모델
class WeaponModel {
  final int id;
  final int playerId;
  final String name;
  final String type; // sword, bow, staff
  int attackPower;
  double attackSpeed;
  final String rarity; // common, rare, epic, legendary
  int level;

  WeaponModel({
    required this.id,
    required this.playerId,
    required this.name,
    required this.type,
    this.attackPower = 10,
    this.attackSpeed = 1.0,
    this.rarity = 'common',
    this.level = 1,
  });

  factory WeaponModel.fromJson(Map<String, dynamic> json) {
    return WeaponModel(
      id: json['id'] as int,
      playerId: json['player_id'] as int,
      name: json['name'] as String,
      type: json['type'] as String,
      attackPower: json['attack_power'] as int? ?? 10,
      attackSpeed: (json['attack_speed'] as num?)?.toDouble() ?? 1.0,
      rarity: json['rarity'] as String? ?? 'common',
      level: json['level'] as int? ?? 1,
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'player_id': playerId,
      'name': name,
      'type': type,
      'attack_power': attackPower,
      'attack_speed': attackSpeed,
      'rarity': rarity,
      'level': level,
    };
  }
}
