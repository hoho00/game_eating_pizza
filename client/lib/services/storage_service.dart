import 'package:shared_preferences/shared_preferences.dart';

/// 저장소 서비스
/// 로컬 데이터 저장을 담당합니다
class StorageService {
  static const String _tokenKey = 'auth_token';
  static const String _playerIdKey = 'player_id';

  /// 토큰 저장
  Future<void> saveToken(String token) async {
    final prefs = await SharedPreferences.getInstance();
    await prefs.setString(_tokenKey, token);
  }

  /// 토큰 불러오기
  Future<String?> getToken() async {
    final prefs = await SharedPreferences.getInstance();
    return prefs.getString(_tokenKey);
  }

  /// 토큰 삭제
  Future<void> clearToken() async {
    final prefs = await SharedPreferences.getInstance();
    await prefs.remove(_tokenKey);
  }

  /// 플레이어 ID 저장
  Future<void> savePlayerId(int playerId) async {
    final prefs = await SharedPreferences.getInstance();
    await prefs.setInt(_playerIdKey, playerId);
  }

  /// 플레이어 ID 불러오기
  Future<int?> getPlayerId() async {
    final prefs = await SharedPreferences.getInstance();
    return prefs.getInt(_playerIdKey);
  }
}
