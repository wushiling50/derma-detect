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
}

struct Detection {
    1: i64 id,
    2: string name,
    3: string percent,
}

struct Article {
    1: i64 id,
    2: string title,
    3: string content,
    4: string cover,
    5: string publish_time,
    6: string author,
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
    // 2: required binary avatar,
}

struct UploadAvatarResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required string avatar_url,
}

struct UploadPictureRequest {
    1: required string token,
    // 2: required binary picture,
}

struct UploadPictureResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required Picture picture,
    4: required list<Detection> detection_list,
}

struct HistoryRequest {
    1: required string token,
}

struct HistoryResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required list<Picture> picture_list,
}

struct HistoryInfoRequest {
    1: required string token,
    2: required i64 picture_id,
}

struct HistoryInfoResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required Picture picture,
    4: required list<Detection> detection_list,
}

struct PublishRequest {
    1: required string token,
    2: required string title,
    3: required string content,
    4: required string cover,
}

struct PublishResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required i64 article_id,
}

struct UploadCoverRequest {
    1: required string token,
}

struct UploadCoverResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required string cover_url,
}

struct ArticleListRequest {
    1: required string token,
    2: required i64 page_num,
}

struct ArticleListResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required list<Article> article_list,
}

struct ArticleDetailRequest {
    1: required string token,
    2: required i64 article_id,
}

struct ArticleDetailResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required Article article,
}

struct ArticleDeleteRequest {
    1: required string token,
    2: required i64 article_id,
}

struct ArticleDeleteResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
}

struct AiDoctorRequest {
    1: required string token,
    2: required string question,
}

struct AiDoctorResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required string reply,
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
    HistoryInfoResponse HistoryInfo(1: HistoryInfoRequest req) (api.get="/derma/detect/picture/history/info/")
}

service ArticleService {
    PublishResponse Publish(1: PublishRequest req) (api.post="/derma/detect/article/publish/")
    UploadCoverResponse UploadCover(1: UploadCoverRequest req)(api.post="/derma/detect/article/upload-cover/")

    ArticleListResponse ArticleList(1: ArticleListRequest req)(api.get="/derma/detect/article/articles/")

    ArticleDetailResponse ArticleDetail(1: ArticleDetailRequest req)(api.get="/derma/detect/article/article-detail/")
    ArticleDeleteResponse ArticleDelete(1: ArticleDeleteRequest req)(api.delete="/derma/detect/article/delete-article/")
}

service AiDoctorService {
    AiDoctorResponse AIDoctor(1: AiDoctorRequest req) (api.get="/derma/detect/ai-doctor")
}