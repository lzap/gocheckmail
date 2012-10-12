gocheckmail
===========

Ultra lightweight maildir checker using inotify Linux syscall

Features
--------

 * Linux inotify watching
 * Maildir support
 * Configurable command to execute

Todo:

 * Mailbox support
 * D-Bus notifications

Installation
------------

Using Google go runtime:

    git checkout git://github.com/lzap/gocheckmail.git
    cd gocheckmail
    go build
    cp gocheckmail ~/bin

Configuration
-------------

Is very easy:

    $ cat ~/.gocheckmail.conf 
    path = $HOME/mail/INBOX
    command = /usr/bin/notify-send 'New mail'

