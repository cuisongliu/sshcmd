# shell

ssh:

shell --user cuisongliu --passwd admin --host 127.0.0.1 --cmd "ls -l"

scp:

shell --user cuisongliu --passwd admin --host 127.0.0.1 --mode "scp" --local-path "/aa.txt" --remote-path "/aa.txt"

mod:
ssh,scp,ssh|scp,scp|ssh

cmd 最后得使用引号包起来

