#/bin/bash

mkdir -p /var/log/backup
mysqldump -uroot -pmysql reqhack | bzip2 -9 > /var/log/backup/`date "+%Y%m%d_%H%M%S"`
mysql -uroot -pmysql < autodelete.sql