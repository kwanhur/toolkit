#!/usr/bin/env bash

#author kwanhur.huang
#date 2017-07-14
#version 1.0.1
#description stat nginx access log for pv uv
# nginx access log format must be:
#log_format log_access "$remote_addr" "\t$remote_user" "\t$time_local" "\t$request" "\t$request_time" "\t$upstream_response_time" "\t$status" "\t$body_bytes_sent" "\t$http_referer" "\t$http_user_agent" "\t$http_x_forwarded_for" "\t$host" "\t$hostname" "\tCustomName1" "\t$http_Cdn_Src_Ip" "\t$http_Cdn_Src_Port" "\tCustomName4" "\tCustomName5" "\tCustomName6" "\tCustomName7" "\tCustomName8" "\t-";

SCRIPT_NAME=$0
COMMAND=$1

function usage()
{
    echo "Usage: $SCRIPT_NAME {version|help|uv|pv|ip_pv|ip_top20|pv_top20|4xx|5xx}"
    exit 1
}

function stat()
{
    domain=$1
    case $COMMAND in
        "uv")
            echo "$domain $COMMAND :"
            awk '{print $1}' $ACCESS_LOG|sort|uniq|wc -l
        ;;
        "pv")
            echo "$domain $COMMAND :"
            awk '{print $1}' $ACCESS_LOG|wc -l
        ;;
        "ip_pv")
            echo "$domain $COMMAND :"
            awk '{++S[$1]} END {for (a in S) print a,S[a]}' $ACCESS_LOG|sort -nr -t ' ' -k 2
        ;;
        "ip_top20")
            echo "$domain $COMMAND :"
            awk '{print $1}' $ACCESS_LOG |sort|uniq -c|sort -nr |head -20 |awk '{print $2,$1}'
        ;;
        "pv_top20")
            echo "$domain $COMMAND :"
            awk -F '[\t]' '{print $4}' $ACCESS_LOG|awk '{counts[$2]+=1};END {for(url in counts) print counts[url],url}'|sort -nr -k 1|head -n 20
        ;;
        "4xx")
            echo "$domain $COMMAND :"
            awk -F '[\t]' '{if($7 ~ /^40?/) print $7}' $ACCESS_LOG |uniq -c |sort -nr|awk '{print $2,$1}'
        ;;
        "5xx")
            echo "$domain $COMMAND :"
            awk -F '[\t]' '{if($7 ~ /^50?/) print $7}' $ACCESS_LOG |uniq -c |sort -nr|awk '{print $2,$1}'
        ;;
        *)
            usage
        ;;
    esac

}

function domain_stat()
{
    NGX_LOG=/home/log/nginx
    if [ -d $NGX_LOG ]; then
        ALL_DOMAIN_ACCESS=`ls -l $NGX_LOG|grep -i .access.log$|awk '{print $9}'`
        for access in $ALL_DOMAIN_ACCESS
        do
                ACCESS_LOG=$NGX_LOG/$access

                domain=`echo $access|awk -F '.access.log' '{print $1}'`
                stat $domain
        done
    else
        echo "$NGX_LOG directory not exited!!!"
        exit 2
    fi

}


case $COMMAND in
    "version")
        echo "1.0.1"
    ;;
    "h"|"help")
        usage
    ;;
    *)
    domain_stat
    ;;
esac

exit 0
