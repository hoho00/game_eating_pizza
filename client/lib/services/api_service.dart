import 'package:dio/dio.dart';

/// API 서비스
/// 서버와의 통신을 담당합니다
class ApiService {
  late Dio _dio;
  final String baseUrl;

  ApiService({String? baseUrl})
      : baseUrl = baseUrl ?? 'http://localhost:8080/api/v1' {
    _dio = Dio(
      BaseOptions(
        baseUrl: this.baseUrl,
        connectTimeout: const Duration(seconds: 5),
        receiveTimeout: const Duration(seconds: 5),
        headers: {
          'Content-Type': 'application/json',
        },
      ),
    );

    // 인터셉터 추가 (에러 처리, 로깅 등)
    _dio.interceptors.add(
      InterceptorsWrapper(
        onRequest: (options, handler) {
          // TODO: 토큰 추가
          // final token = storageService.getToken();
          // if (token != null) {
          //   options.headers['Authorization'] = 'Bearer $token';
          // }
          return handler.next(options);
        },
        onError: (error, handler) {
          // 에러 처리
          return handler.next(error);
        },
      ),
    );
  }

  // 인증 관련
  Future<Response> register(String username, String password) async {
    return await _dio.post('/auth/register', data: {
      'username': username,
      'password': password,
    });
  }

  Future<Response> login(String username, String password) async {
    return await _dio.post('/auth/login', data: {
      'username': username,
      'password': password,
    });
  }

  // 플레이어 관련
  Future<Response> getPlayerInfo() async {
    return await _dio.get('/players/me');
  }

  Future<Response> getLeaderboard({int limit = 10}) async {
    return await _dio.get('/players/leaderboard', queryParameters: {
      'limit': limit,
    });
  }

  // 무기 관련
  Future<Response> getWeapons() async {
    return await _dio.get('/weapons');
  }

  Future<Response> upgradeWeapon(int weaponId) async {
    return await _dio.put('/weapons/$weaponId/upgrade');
  }

  Future<Response> equipWeapon(int weaponId) async {
    return await _dio.put('/weapons/$weaponId/equip');
  }

  // 던전 관련
  Future<Response> getAllDungeons() async {
    return await _dio.get('/dungeons/all');
  }

  Future<Response> getActiveDungeons() async {
    return await _dio.get('/dungeons/active');
  }

  Future<Response> getDungeon(int dungeonId) async {
    return await _dio.get('/dungeons/$dungeonId');
  }
}
