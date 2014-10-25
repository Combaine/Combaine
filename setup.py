#!/usr/bin/env python
# -*- coding: utf-8 -*-

import setuptools


__author__ = 'Anton Tiurin'
__copyright__ = 'Copyright 2014'
__credits__ = []

__license__ = 'Apache 2.0'
__version__ = '0.6.5.3'
__maintainer__ = 'Anton Tiurin'
__email__ = 'noxiouz@yandex.ru'
__status__ = 'Production'

__title__ = 'Combaine'
__build__ = 0x000000

__url__ = 'https://github.com/noxiouz/Combaine'
__description__ = 'Bleeding edge'
d = 'https://github.com/noxiouz/docker-registry-driver-elliptics/archive/\
master.zip'

setuptools.setup(
    name=__title__,
    version=__version__,
    author=__author__,
    author_email=__email__,
    maintainer=__maintainer__,
    maintainer_email=__email__,
    url=__url__,
    description=__description__,
    long_description=open('./README.md').read(),
    download_url=d,
    classifiers=['Development Status :: 4 - Beta',
                 'Intended Audience :: Developers',
                 'Programming Language :: Python',
                 # 'Programming Language :: Python :: 2.6',
                 'Programming Language :: Python :: 2.7',
                 # 'Programming Language :: Python :: 3.2',
                 # 'Programming Language :: Python :: 3.3',
                 # 'Programming Language :: Python :: 3.4',
                 'Programming Language :: Python :: Implementation :: CPython',
                 'Operating System :: OS Independent',
                 'Topic :: Utilities',
                 'License :: OSI Approved :: Apache Software License'],
    platforms=['Independent'],
    license=open('./LICENSE').read(),
    namespace_packages=['combaine'],
    packages=['combaine',
              'combaine.cloud',
              'combaine.common',
              'combaine.utils'],
    install_requires=open('./requirements.txt').read(),
    zip_safe=False,
    tests_require=open('./tests/requirements.txt').read(),
    test_suite='nose.collector'
)
