use mongodb::bson::DateTime;
use serde::{Serialize, Deserialize};

#[derive(Debug, Serialize, Deserialize)]
pub struct Article {
    #[serde(rename="_id")]
    pub id: String,
    pub name: String,
    pub author: String,
    pub summary: String,
    pub cover_img: String,
    pub catalog: Catalog,
    pub content: String,
    pub kind: Kind,
    pub tag_list: Vec<Tag>,
    pub release_time: DateTime,
    pub visibility: i32,
    pub fulltext_title: String,
    pub fulltext_content: String,
    pub available: bool,
    pub created_time: DateTime,
    pub updated_time: DateTime,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct Catalog {
    name: String,
    children: Vec<Catalog>,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct Kind {
    #[serde(rename="_id")]
    id: String,
    name: String,
    available: bool,
    created_time: DateTime,
    updated_time: DateTime,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct Tag {
    #[serde(rename="_id")]
    id: String,
    name: String,
    available: bool,
    created_time: DateTime,
    updated_time: DateTime,
}