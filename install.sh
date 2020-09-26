mkdir ~/s5_server
service firewalld stop
ip=$(curl -s https://ipinfo.io/ |grep 'ip"'|awk -F '"' '{print$4}')
cd ~/s5_server
wget https://raw.githubusercontent.com/code-scan/s5_server/master/bin/s5 -O s5 
chmod +x s5
./s5 s5 a536e254fa6a6dcad3f42bfd6f343150 65520
echo socks5://$ip 65520 s5 a536e254fa6a6dcad3f42bfd6f343150