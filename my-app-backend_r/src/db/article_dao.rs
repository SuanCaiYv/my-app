use futures::TryStreamExt;
use mongodb::{Client, Collection};
use mongodb::bson::{bson, DateTime, doc};
use mongodb::bson::oid::ObjectId;
use mongodb::error::Error;
use mongodb::options::{ClientOptions, CountOptions, FindOptions, InsertOneOptions};
use mongodb::results::InsertOneResult;
use crate::entity::article::Article;

pub struct ArticleDaoStruct {
    collection: Option<Collection<Article>>,
}

impl ArticleDaoStruct {
    fn new() -> Self {
        ArticleDaoStruct { collection: None }
    }

    async fn connection(&self, address: String, port: i32, db: String) -> Self {
        let url = format!("mongodb://{}:{}", address, port);
        let options = ClientOptions::parse(url).await.unwrap();
        let client = Client::with_options(options).unwrap();
        let db = client.database(db.as_str());
        let collection = db.collection::<Article>("article");
        Self { collection: Some(collection) }
    }

    async fn insert(&mut self, article: &mut Article) -> Result<(), Error> {
        let collection = self.collection.as_mut().unwrap();
        article.id = ObjectId::new().to_hex();
        article.created_time = DateTime::now();
        article.updated_time = DateTime::now();
        collection.insert_one(article, Some(InsertOneOptions::default())).await?;
        Ok(())
    }

    async fn select(&mut self, id: String) -> Result<Option<Article>, Error> {
        todo!()
    }

    async fn select_by_author_name(&mut self, author: String, name: String) -> Result<Option<Article>, Error> {
        todo!()
    }

    async fn list_by_author0(&mut self, author: String, visibility: i32, equally: bool) -> Result<Vec<Article>, Error> {
        todo!()
    }

    async fn list_by_author(&mut self, author: String, visibility: i32, equally: bool, page_num: usize, page_size: usize, sort: String, desc: bool, tag_list: Vec<String>, search_key: String) -> Result<(Vec<Article>, i64), Error> {
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
        let mut filter = doc! {};
        let mut options = FindOptions::default();
        options.limit = Some(page_size as i64);
        options.skip = Some(skip as u64);
        options.sort = Some(doc! {
            sort: desc_int
        });
        if !tag_list.is_empty() && !search_key.is_empty() {
            filter = doc! {
				"author": author,
				"available":  true,
				"visibility": v,
				"tag_list._id": 1,
				"$text": doc! {
					"$search": search_key,
				}
            };
        } else if !tag_list.is_empty() {
            filter = doc! {
				"author":     author,
				"available":  true,
				"visibility": v,
				"tag_list._id": doc! {
					"$all": tag_list,
				}
            };
        } else if !search_key.is_empty() {
            filter = doc! {
				"author":     author,
				"available":  true,
				"visibility": v,
				"$text": doc! {
					"$search": search_key,
				}
            };
        } else {
            filter = doc! {
				"author":     author,
				"available":  true,
				"visibility": v,
            }
        }
        let mut ans = Vec::new();
        let mut cursor = self.collection.as_mut().unwrap().find(filter.clone(), options).await?;
        while let Some(val) = cursor.try_next().await? {
            ans.push(val);
        }
        self.collection.as_mut().unwrap().count_documents(filter, CountOptions::default()).await?;
        Ok((ans, 0))
    }

    async fn list_all0(&self, author: String) -> Result<Article, Error> {
        todo!()
    }

    async fn update(&mut self, article: &Article) -> Result<(), Error> {
        todo!()
    }

    async fn delete(&mut self, id: String) -> Result<(), Error> {
        todo!()
    }
}

#[cfg(test)]
mod tests {
    use futures::TryStreamExt;
    use mongodb::bson::doc;
    use mongodb::{Client, Cursor};
    use mongodb::options::{ClientOptions, FindOptions, InsertOneOptions};
    use crate::entity::sys::SysUser;

    #[tokio::test]
    async fn test() {
        let options = ClientOptions::parse("mongodb://127.0.0.1:27017").await.unwrap();
        let client = Client::with_options(options).unwrap();
        let database = client.database("my_app");
        let collection = database.collection::<SysUser>("sys_user");
        let user = SysUser::default();
        collection.insert_one(&user, InsertOneOptions::default()).await.unwrap();
        let filter = doc! {};
        let find_options = FindOptions::default();
        let mut cursor: Cursor<SysUser> = collection.find(filter.clone(), find_options.clone()).await.unwrap();
        while let Some(val) = cursor.try_next().await.unwrap() {
            println!("{:?}", val)
        }
    }
}