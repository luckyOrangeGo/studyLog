#!/bin/bash  
   
STTY_RESTORE=$(stty -g)  
   
echo -n "Username: "  
read username  
   
echo -n "Password: "  
stty -echo cbreak  
while true  
do  
        character=$(dd if=/dev/tty bs=1 count=1 2> /dev/null)  
        case $character in  
        $(echo -e "\n"))  
                break  
                ;;  
        $(echo -e "\b"))  
                if [ -n "$password" ]; then  
                        echo -n -e "\b \b"  
                        password=$(echo "$password" | sed 's/.$//g')  
                fi  
                ;;  
        *)  
                password=$password$character  
                echo -n '*'  
                ;;  
        esac  
done  
   
stty $STTY_RESTORE                      #stty -cbreak echo  
  
echo -e "\n\nUsername is $username"  
echo "Password is $password"  