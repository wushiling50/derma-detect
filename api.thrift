namespace go api

struct User {
    1: i64 id,
    2: string username,
    3: string avatar,
    4: string signature,
    5: string email,
    6: i64 phone,
    7: string birth,
}

struct Picture {
    1: i64 id,
    2: string picture_url,
    3: string create_date,
    4: string percent,
    5: string describe,
}

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

struct UserInfoRequest {
    1: required string token,
}

struct UserInfoResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required User user,
}

struct ResetInfoRequest {
    1: required string token,
    2: optional string username,
    3: optional string email,
    4: optional i64 phone,
    5: optional string birth,
}

struct ResetInfoResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
}

struct ResetPasswordRequest {
    1: required string token,
    2: required string old_password,
    3: required string new_password,
    4: required string check_password,
}

struct ResetPasswordResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
}

struct UploadAvatarRequest {
    1: required string token,
    2: required binary avatar,
}

struct UploadAvatarResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required string avatar_url,
}

struct UploadPictureRequest {
    1: required string token,
    2: required binary picture,
}

struct UploadPictureResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required list<Picture> result_list,
}

struct HistoryRequest {
    1: required string token,
}

struct HistoryResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required list<Picture> picture_list,
}

service BasicService {
    UserRegisterResponse UserRegister(1: UserRegisterRequest req) (api.post="/derma/detect/user/register/")
    UserLoginResponse UserLogin(1: UserLoginRequest req) (api.post="/derma/detect/user/login/")

    UserInfoResponse UserInfo(1: UserInfoRequest req) (api.get="/derma/detect/user/info/")
    ResetInfoResponse ResetInfo(1: ResetInfoRequest req) (api.post="/derma/detect/user/reset-info/")
    ResetPasswordResponse ResetPassword(1: ResetPasswordRequest req) (api.post="/derma/detect/user/reset-password/")
    
    UploadAvatarResponse UploadAvatar(1:UploadAvatarRequest req) (api.post="/derma/detect/user/upload-avatar/")
}

service PictureService {
    HistoryResponse UserHistory(1: HistoryRequest req) (api.get="/derma/detect/picture/history/")
    UploadPictureResponse UploadPicture(1: UploadPictureRequest req) (api.post="/derma/detect/picture/upload/")
}