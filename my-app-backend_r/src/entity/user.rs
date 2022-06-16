use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
pub struct UserInfo {
    avatar: String,
    nickname: String,
    email: String,
    phone: String,
    we_chat: String,
    qq: String,
    github: String,
    location: String,
    signature: String,
}

impl Default for UserInfo {
    fn default() -> Self {
        Self {
            avatar: "".to_string(),
            nickname: "".to_string(),
            email: "".to_string(),
            phone: "".to_string(),
            we_chat: "".to_string(),
            qq: "".to_string(),
            github: "".to_string(),
            location: "".to_string(),
            signature: "".to_string()
        }
    }
}