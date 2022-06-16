use std::time::Duration;
use redis::{aio, AsyncCommands, RedisResult, ToRedisArgs};
use tokio::runtime::Builder;

pub struct RedisOps {
    connection: Option<aio::MultiplexedConnection>,
}

impl RedisOps {
    pub fn new() -> Self {
        RedisOps{ connection: None }
    }

    pub async fn connection(self, address: String, port: i32) -> Self {
        let url = format!("redis://{}:{}", address, port);
        let connection = redis::Client::open(url).unwrap().get_multiplexed_async_connection().await.unwrap();
        RedisOps{ connection: Some(connection) }
    }

    pub async fn set<T: ToRedisArgs>(&mut self, key: String, value: T) -> RedisResult<()> {
        redis::cmd("SET")
            .arg(&key)
            .arg(&value)
            .query_async(self.connection.as_mut().unwrap())
            .await
    }

    pub async fn set_ref<T: ToRedisArgs>(&mut self, key: &'static str, value: T) -> RedisResult<()> {
        redis::cmd("SET")
            .arg(key)
            .arg(value)
            .query_async(self.connection.as_mut().unwrap())
            .await
    }

    pub async fn set_exp<T: ToRedisArgs>(&mut self, key: String, value: T, exp: Duration) -> RedisResult<()> {
        redis::cmd("PSETEX")
            .arg(&key)
            .arg(exp.as_millis() as u64)
            .arg(&value)
            .query_async(self.connection.as_mut().unwrap())
            .await
    }

    pub async fn set_exp_ref<T: ToRedisArgs>(&mut self, key: &'static str, value: T, exp: Duration) -> RedisResult<()> {
        redis::cmd("PSETEX")
            .arg(key)
            .arg(exp.as_millis() as u64)
            .arg(&value)
            .query_async(self.connection.as_mut().unwrap())
            .await
    }

    pub async fn get(&mut self, key: String) -> RedisResult<String> {
        redis::cmd("GET")
            .arg(&key)
            .query_async(self.connection.as_mut().unwrap())
            .await
    }

    pub async fn get_ref(&mut self, key: &'static str) -> RedisResult<String> {
        redis::cmd("GET")
            .arg(key)
            .query_async(self.connection.as_mut().unwrap())
            .await
    }

    pub async fn del(&mut self, key: String) -> RedisResult<()> {
        redis::cmd("DEL")
            .arg(&key)
            .query_async(self.connection.as_mut().unwrap())
            .await
    }

    pub async fn del_ref(&mut self, key: &'static str) -> RedisResult<()> {
        redis::cmd("DEL")
            .arg(key)
            .query_async(self.connection.as_mut().unwrap())
            .await
    }

    pub async fn push_sort_queue<T: ToRedisArgs>(&mut self, key: String, val: T, score: f64) -> RedisResult<()> {
        redis::cmd("ZADD")
            .arg(&key)
            .arg(score)
            .arg(&val)
            .query_async(self.connection.as_mut().unwrap())
            .await
    }

    pub async fn peek_sort_queue<T: ToRedisArgs>(&mut self, key: String) -> RedisResult<T> {
        todo!()
    }

    pub async fn pop_sort_queue<T: ToRedisArgs>(&mut self, key: String) -> RedisResult<T> {
        todo!()
    }
}

#[cfg(test)]
mod tests {
    use std::thread;
    use std::time::Duration;
    use crate::nosql::redis_ops::RedisOps;

    #[tokio::test]
    async fn test() {
        let mut redis_ops = RedisOps::new().connection("127.0.0.1".to_string(), 6379).await;
        let val = redis_ops.get("key1".to_string()).await.unwrap();
        println!("{}", val);
    }
}