use mongodb::{Client, Collection};
use mongodb::bson::{bson, DateTime};
use mongodb::bson::oid::ObjectId;
use mongodb::error::Error;
use mongodb::options::{ClientOptions, InsertOneOptions};
use mongodb::results::InsertOneResult;
use crate::entity::article::Article;

pub struct ArticleDaoStruct {
    collection: Option<Collection<Article>>,
}

impl ArticleDaoStruct {
    fn new() -> Self {
        ArticleDaoStruct {collection: None}
    }

    async fn connection(&mut self, address: String, port: i32, db: String) -> Self {
        let url = format!("mongodb://{}:{}", address, port);
        let options = ClientOptions::parse(url).await.unwrap();
        let client = Client::with_options(options).unwrap();
        let db = client.database(db.as_str());
        let collection = db.collection::<Article>("article");
        Self {collection: Some(collection)}
    }

    async fn insert(&mut self, article: &mut Article) -> Result<(), Error> {
        let collection = self.collection.as_mut().unwrap();
        article.id = ObjectId::new().to_hex();
        article.created_time = DateTime::now();
        article.updated_time = DateTime::now();
        collection.insert_one(article, Some(InsertOneOptions::default())).await?;
        Ok(())
    }

    async fn select(&mut self, id: String) -> Result<Article, Error> {
        todo!()
    }

    async fn list_by_author(&mut self, author: String, visibility: i32, equally: bool, page_num: usize, page_size: usize, sort: String, desc: bool, tag_list: Vec<String>, search_key: String) -> Result<Vec<Article>, Error> {
        let mut desc_int = -1;
        if desc {
            desc_int = 1;
        }
        let skip = (page_num - 1) + page_size;
        let search_key = search_key.to_lowercase();
        let v;
        if equally {
            v = bson!({"$eq": visibility});
        } else {
            v = bson!({"$ne": visibility});
        }
        if !tag_list.is_empty()
    }
}

#[cfg(test)]
mod tests {
    use futures::TryStreamExt;
    use mongodb::bson::doc;
    use mongodb::{Client, Cursor};
    use mongodb::options::{ClientOptions, FindOptions};
    use crate::entity::sys::SysUser;

    #[tokio::test]
    async fn test() {
        let options = ClientOptions::parse("mongodb://127.0.0.1:27017").await.unwrap();
        let client = Client::with_options(options).unwrap();
        let database = client.database("my_app");
        let collection = database.collection::<SysUser>("sys_user");
        let filter = doc! {};
        let find_options= FindOptions::default();
        let mut cursor: Cursor<SysUser> = collection.find(filter, find_options).await.unwrap();
        while let Some(val) = cursor.try_next().await.unwrap() {
            println!("{:?}", val)
        }
    }
}