are-you-sure
============

a program to help you avoid fat fingering those important, potentially career-destroying commands

how to install
============
    go get -u github.com/tobz/are-you-sure
    sudo ln -snf $GOPATH/bin/are-you-sure /usr/local/bin/ays

how to use
============
    toby:~ ays
    [ays] You must specify a command to run!
    toby:~ ays sudo apt-get install htop
    [ays] Command entered: `sudo apt-get install htop`
    Are you sure you want to run this?  Type 'yes' to proceed: yes
    [ays] Running command `sudo apt-get install htop`...
    Reading package lists... Done
    Building dependency tree
    Reading state information... Done
    htop is already the newest version.
    0 upgraded, 0 newly installed, 0 to remove and 21 not upgraded.
