/// 던전 데이터 모델
class DungeonModel {
  final int id;
  final String name;
  final String type; // normal, event, boss
  final int difficulty;
  final bool isActive;
  final DateTime? startTime;
  final DateTime? endTime;

  DungeonModel({
    required this.id,
    required this.name,
    required this.type,
    required this.difficulty,
    required this.isActive,
    this.startTime,
    this.endTime,
  });

  factory DungeonModel.fromJson(Map<String, dynamic> json) {
    return DungeonModel(
      id: json['id'] as int,
      name: json['name'] as String,
      type: json['type'] as String,
      difficulty: json['difficulty'] as int? ?? 1,
      isActive: json['is_active'] as bool? ?? true,
      startTime: json['start_time'] != null
          ? DateTime.parse(json['start_time'] as String)
          : null,
      endTime: json['end_time'] != null
          ? DateTime.parse(json['end_time'] as String)
          : null,
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'type': type,
      'difficulty': difficulty,
      'is_active': isActive,
      'start_time': startTime?.toIso8601String(),
      'end_time': endTime?.toIso8601String(),
    };
  }
}
