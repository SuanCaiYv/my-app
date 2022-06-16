use mongodb::bson::DateTime;
use crate::entity::user::UserInfo;
use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
pub struct SysUser {
    #[serde(rename="_id")]
    id: String,
    username: String,
    credential: String,
    salt: String,
    role: String,
    info: UserInfo,
    available: bool,
    created_time: DateTime,
    updated_time: DateTime,
}

impl Default for SysUser {
    fn default() -> Self {
        Self {
            id: "".to_string(),
            username: "cwb".to_string(),
            credential: "cwb".to_string(),
            salt: "cwb".to_string(),
            role: "reader".to_string(),
            info: UserInfo::default(),
            available: false,
            created_time: DateTime::now(),
            updated_time: DateTime::now(),
        }
    }
}