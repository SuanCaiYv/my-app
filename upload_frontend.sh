cd my-app-frontend
tar -zcvf frontend.tar.gz dist
rsync frontend.tar.gz root@1.117.182.211:/home/my_app/
