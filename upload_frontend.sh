cd my-app-frontend
npm run build
tar -zcvf frontend.tar.gz dist
rsync frontend.tar.gz -e 'ssh -p 2022' root@1.117.182.211:/home/my_app/
