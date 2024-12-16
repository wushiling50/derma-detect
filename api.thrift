namespace go api

struct UserRegisterRequest {
    1: required string username,
    2: required string password,
}

struct UserRegisterResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required i64 user_id,
    4: required string token,
}

struct UserLoginRequest {
    1: required string username,
    2: required string password,
}

struct UserLoginResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required i64 user_id,
    4: required string token,
}

service BasicService {
    UserRegisterResponse UserRegister(1: UserRegisterRequest req) (api.post="/derma/detect/user/register/")
    UserLoginResponse UserLogin(1: UserLoginRequest req) (api.post="/derma/detect/user/login/")
}