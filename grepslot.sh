#!/bin/bash
slotname=""
for file in `find /opt/flinkx-es/ -type f  -name "*.log"`
do
        slotname=`cat $file|grep "Cannot obtain valid replication slot"|awk -F"'" '{print $2}'|uniq`
        echo $slotname
        if [ "$slotname" != "" ]
        then
                prosname=`echo $file|awk -F"/" '{print $5}'|awk -F'.' '{print $1}'`
                echo $prosname
                kill -9 `ps -ef | grep $prosname|awk '{print $1}'`
                echo "send slotname $slotname to delete..."
                curl postgres.middleware:9990/slot/$slotname
                slotname=""
                echo -n "$prosname slotname is delete, waiting restart..." >> /opt/flinkx-es/slotdelete.log
                date >> /opt/flinkx-es/slotdelete.log
        fi
done