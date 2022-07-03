use mongodb::bson::{DateTime, Document};
use mongodb::bson::oid::ObjectId;
use mongodb_gridfs::GridFSBucket;
use mongodb::error::Error;
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug, Clone)]
struct GridFSFile {
    #[serde(rename="_id")]
    pub id: ObjectId,
    pub length: i64,
    #[serde(rename="chunkSize")]
    pub chunk_size: i32,
    #[serde(rename="uploadDate")]
    pub upload_date: DateTime,
    pub filename: String,
    pub metadata: Document,
}

impl From<Document> for GridFSFile {
    fn from(document: Document) -> Self {
        let file: GridFSFile = serde_json::from_value(serde_json::to_value(document).unwrap()).unwrap();
        file
    }
}

impl From<&Document> for GridFSFile {
    fn from(document: &Document) -> Self {
        let file: GridFSFile = serde_json::from_value(serde_json::to_value(document).unwrap()).unwrap();
        file
    }
}

pub struct GridFSDaoStruct {
    bucket: Option<GridFSBucket>
}

impl GridFSDaoStruct {
    fn new() -> Self {
        GridFSDaoStruct {
            bucket: None,
        }
    }

    async fn connection(&self, address: String, port: i32, db: String) -> Self {
        todo!()
    }

    async fn upload(&mut self, file_content: Vec<u8>, filename: String, metadata: Document) -> Result<(), Error> {
        todo!()
    }
}

#[cfg(test)]
mod test {
    use futures::StreamExt;
    use mongodb::bson::{DateTime, doc, Document};
    use mongodb::bson::oid::ObjectId;
    use mongodb::Client;
    use mongodb::options::ClientOptions;
    use mongodb_gridfs::GridFSBucket;
    use mongodb_gridfs::options::{GridFSBucketOptions, GridFSFindOptions};
    use serde::{Deserialize, Serialize};

    #[derive(Serialize, Deserialize, Debug, Clone)]
    struct GridFSFile {
        #[serde(rename="_id")]
        pub id: ObjectId,
        pub length: i64,
        #[serde(rename="chunkSize")]
        pub chunk_size: i32,
        #[serde(rename="uploadDate")]
        pub upload_date: DateTime,
        pub filename: String,
        pub metadata: Document,
    }

    impl From<Document> for GridFSFile {
        fn from(document: Document) -> Self {
            let file: GridFSFile = serde_json::from_value(serde_json::to_value(document).unwrap()).unwrap();
            file
        }
    }

    impl From<&Document> for GridFSFile {
        fn from(document: &Document) -> Self {
            let file: GridFSFile = serde_json::from_value(serde_json::to_value(document).unwrap()).unwrap();
            file
        }
    }

    #[tokio::test]
    async fn test() {
        let options = ClientOptions::parse("mongodb://127.0.0.1:27017").await.unwrap();
        let client = Client::with_options(options).unwrap();
        let database = client.database("my_app_grid_fs");
        let mut bucket = GridFSBucket::new(database, Some(GridFSBucketOptions::default()));
        let mut cursor = bucket.find(doc! {"filename": "default-avatar.png"}, GridFSFindOptions::default()).await.unwrap();
        let document = cursor.next().await.unwrap().unwrap();
        let file: GridFSFile = serde_json::from_value(serde_json::to_value(&document).unwrap()).unwrap();
        // 时区问题
        println!("{:?}", file.upload_date.to_string());
    }
}