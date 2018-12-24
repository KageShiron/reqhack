#/bin/bash

mkdir -p /var/log/backup
h databasemysqldump -h database -uroot -pmysql reqhack | bzip2 -9 > /var/log/backup/`date "+%Y%m%d_%H%M%S"`
                                                                                                                                          < autodelete.sqlmysql -uroot -pmysql -