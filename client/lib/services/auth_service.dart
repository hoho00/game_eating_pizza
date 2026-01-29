import 'package:flutter/foundation.dart';
import 'api_service.dart';
import 'storage_service.dart';

/// 인증 서비스
/// 사용자 인증 및 세션 관리를 담당합니다
class AuthService extends ChangeNotifier {
  final ApiService apiService;
  final StorageService storageService;

  String? _token;
  bool _isAuthenticated = false;

  AuthService({
    required this.apiService,
    required this.storageService,
  }) {
    _loadToken();
  }

  bool get isAuthenticated => _isAuthenticated;
  String? get token => _token;

  /// 저장된 토큰 로드
  Future<void> _loadToken() async {
    _token = await storageService.getToken();
    _isAuthenticated = _token != null;
    notifyListeners();
  }

  /// 회원가입
  Future<bool> register(String username, String password) async {
    try {
      final response = await apiService.register(username, password);
      if (response.statusCode == 201) {
        // 회원가입 성공 후 자동 로그인
        return await login(username, password);
      }
      return false;
    } catch (e) {
      debugPrint('Register error: $e');
      return false;
    }
  }

  /// 로그인
  Future<bool> login(String username, String password) async {
    try {
      final response = await apiService.login(username, password);
      if (response.statusCode == 200) {
        _token = response.data['token'];
        _isAuthenticated = true;
        await storageService.saveToken(_token!);
        notifyListeners();
        return true;
      }
      return false;
    } catch (e) {
      debugPrint('Login error: $e');
      return false;
    }
  }

  /// 로그아웃
  Future<void> logout() async {
    _token = null;
    _isAuthenticated = false;
    await storageService.clearToken();
    notifyListeners();
  }
}
