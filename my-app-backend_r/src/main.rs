mod nosql;
mod entity;
mod db;
mod config;

macro_rules! zoom_and_enhance {
    (struct $name:ident { $($fname:ident : $ftype:ty),* }) => {
        struct $name {
            $($fname : $ftype),*
        }

        impl $name {
            fn field_names() -> &'static [&'static str] {
                static NAMES: &'static [&'static str] = &[$(stringify!($fname)),*];
                NAMES
            }
        }
    }
}

zoom_and_enhance! {
struct Export {
        first_name: String,
        last_name: String,
        gender: String,
        date_of_birth: String,
        address: String
    }
}

#[tokio::main]
async fn main() {
}

#[cfg(test)]
mod tests {
    use std::io::Error;

    #[test]
    fn test() {
        println!("{}", f().unwrap())
    }

    fn f() -> Result<i32, i64> {
        ff()?;
        Ok(2)
    }

    fn ff() -> Result<i32, i64> {
        Err(0)
    }
}
