tar -zcvf backend.tar.gz my-app-backend
rsync backend.tar.gz -e 'ssh -p 2022' root@1.117.182.211:/home/my_app/
