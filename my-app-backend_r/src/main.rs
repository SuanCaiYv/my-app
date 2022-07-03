mod nosql;
mod entity;
mod db;
mod config;

#[tokio::main]
async fn main() {
}

#[cfg(test)]
mod tests {

    trait A {}

    struct B {}

    impl A for &B {}

    #[test]
    fn test() {
        let b = &B{};
        f(b);
    }

    fn f(t: impl A) {}
}
