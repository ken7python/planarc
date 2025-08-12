cd /home/ec2-user/planarc/app
npm run build

DEPLOY_DIR="/var/www/planarc"

cp -r dist/* $DEPLOY_DIR
