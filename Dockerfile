FROM ubuntu:trusty

# basic configure
RUN ln -vsTf /bin/bash /bin/sh
RUN ln -vsTf /bin/bash /bin/dash

RUN echo -e \
"deb http://mirror.yandex.ru/ubuntu/ trusty main restricted universe multiverse\n"\
"deb http://mirror.yandex.ru/ubuntu/ trusty-updates main restricted universe multiverse\n"\
"deb http://mirror.yandex.ru/ubuntu/ trusty-backports main restricted universe multiverse\n"\
"deb http://mirror.yandex.ru/ubuntu/ trusty-security main restricted universe multiverse\n"\
"\n"\
"deb http://common.dist.yandex.ru/common stable/amd64/\n"\
"deb http://common.dist.yandex.ru/common stable/all/\n"\
"deb http://yandex-trusty.dist.yandex.ru/yandex-trusty/ stable/all/\n"\
"deb http://yandex-trusty.dist.yandex.ru/yandex-trusty/ stable/amd64/\n"\
> /etc/apt/sources.list

RUN echo -e \
"127.0.0.1 localhost.localdomain localhost\n"\
"::1 localhost ip6-localhost ip6-loopback\n"\
"fe00::0 ip6-localnet\n"\
"ff00::0 ip6-mcastprefix\n"\
"ff02::1 ip6-allnodes\n"\
"ff02::2 ip6-allrouters\n"\
"ff02::3 ip6-allhosts\n"\
> /etc/hosts

RUN apt-get -qq update \
    && apt-get install -y --force-yes --no-install-recommends \
    libjemalloc1 unbound psmisc python-yaml jq=1.4-2.1~ubuntu14.04.1 lsof \
    jnettop util-linux strace tcpdump htop curl moreutils \
    \
    libyandex-ubic-shared-perl libfcgi-procmanager-perl libfcgi-perl \
    libsys-hostname-long-perl libconfig-tiny-perl liburi-perl \
    libfcgi-client-perl libbsd-resource-perl \
    \
    salt-minion=2018.3.1-yandex1 zstd mawk yandex-timetail media-graphite-sender \
    cocaine-tools cocaine-runtime cocaine-framework-python=0.11.1.12 libcocaine-core2 \
    && apt-get clean && rm -rf /var/lib/apt/lists/* /var/cache/apt/archives/*
RUN curl -O http://dist.yandex.ru/storage/1028049/common/yandex-3132-fastcgi-loggiver_0.49_all.deb \
    && dpkg-deb --extract yandex-3132-fastcgi-loggiver*deb /

RUN curl -kO https://raw.githubusercontent.com/pixelb/ps_mem/master/ps_mem.py \
    && mv ps_mem.py /usr/bin/ps_mem \
    && chmod +x /usr/bin/ps_mem

RUN locale-gen en_US.utf8 ru_RU.utf8
RUN ln -vsTf /usr/share/zoneinfo/Europe/Moscow /etc/localtime \
    && echo Europe/Moscow > /etc/timezone

COPY build/combainer /usr/bin/
COPY build/worker /usr/bin/combaine-worker

COPY deploy                        /usr/lib/combaine
COPY plugins/aggregators/custom.py /usr/lib/combaine/apps/
COPY build/agave                   /usr/lib/combaine/apps/
COPY build/graphite                /usr/lib/combaine/apps/
COPY build/razladki                /usr/lib/combaine/apps/
COPY build/cbb                     /usr/lib/combaine/apps/
COPY build/solomon                 /usr/lib/combaine/apps/
COPY build/juggler                 /usr/lib/combaine/apps/